package main

import "github.com/PuerkitoBio/fetchbot"
import (
	"fmt"
	"github.com/DennisDenuto/property-price-collector/data/training"
	pphc "github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
	log "github.com/Sirupsen/logrus"
	pachdclient "github.com/pachyderm/pachyderm/src/client"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)

	mux := fetchbot.NewMux()

	//2000 2155
	minPostcode, err := strconv.Atoi(os.Getenv("START_POSTCODE"))
	maxPostcode, err := strconv.Atoi(os.Getenv("END_POSTCODE"))

	pphcFetcher := pphc.NewPropertyPriceHistoryCom("propertypricehistory.com", minPostcode, maxPostcode)
	pphcFetcher.SetupMux(mux)

	fetcher := fetchbot.New(mux)
	fetcher.AutoClose = true

	queue := fetcher.Start()

	for _, seed := range pphcFetcher.SeedUrls {
		queue.SendStringGet(seed)
	}

	client, err := getPachdClient()
	if err != nil {
		log.WithError(err).Error("unable to get pachd client")
		panic(err)
	}
	repo := training.NewTrainingDataRepo(client)

	err = repo.Create()
	if err != nil {
		log.WithError(err).Error("unable to create repo")
		panic(err)
	}
	for {
		select {
		case property := <-pphcFetcher.GetProperties():
			err = retryDuring(10*time.Minute, 10*time.Second, func() error {
				commitId, err := repo.StartTxn()
				if err != nil {
					log.WithError(err).Error("unable to start txn")
					return err
				}

				err = repo.Add(property)
				if err != nil {
					fmt.Errorf("adding property to repo errored: %s", err)
				}

				err = repo.Commit(commitId)
				if err != nil {
					log.WithError(err).Error("unable to commit txn")
					return err
				}

				log.Infof("%+#v", property)

				return nil
			}, func() {
				client, err := getPachdClient()
				if err != nil {
					repo = training.NewTrainingDataRepo(client)
				}
			})

			if err != nil {
				log.WithError(err).Error("unable to write property into datastore")
				panic(err)
			}
		case <-queue.Done():
			log.Info("Finished")
			return
		}
	}

}

func getPachdClient() (training.APIClient, error) {
	host, foundHost := os.LookupEnv("PACHD_SERVICE_HOST")
	port, foundPort := os.LookupEnv("PACHD_SERVICE_PORT")

	if !foundHost || !foundPort {
		return nil, fmt.Errorf("missing required env variable (PACHD_SERVICE_HOST|PACHD_SERVICE_PORT)")
	}

	client, err := pachdclient.NewFromAddress(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.WithError(err).Error("unable to connect to pachyderm")
		return nil, err
	}
	return client, nil
}

func retryDuring(duration time.Duration, sleep time.Duration, callback func() error, attemptRepair func()) (err error) {
	t0 := time.Now()
	i := 0
	for {
		i++

		err = callback()
		if err == nil {
			return
		}

		delta := time.Now().Sub(t0)
		if delta > duration {
			return fmt.Errorf("after %d attempts (during %s), last error: %s", i, delta, err)
		}

		time.Sleep(sleep)

		log.WithError(err).Debug("retrying after error")
		attemptRepair()
	}
}
