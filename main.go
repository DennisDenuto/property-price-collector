package main

import "github.com/PuerkitoBio/fetchbot"
import (
	pphc "github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
	"github.com/DennisDenuto/property-price-collector/data/training"
	pachdclient "github.com/pachyderm/pachyderm/src/client"
	"fmt"
	"os"
	"time"
	"log"
)

func main() {

	mux := fetchbot.NewMux()

	pphcFetcher := pphc.NewPropertyPriceHistoryCom("propertypricehistory.com", 2000, 2155)
	pphcFetcher.SetupMux(mux)

	fetcher := fetchbot.New(mux)

	queue := fetcher.Start()

	for _, seed := range pphcFetcher.SeedUrls {
		queue.SendStringGet(seed)
	}

	client, err := getPachdClient()
	if err != nil {
		panic(err)
	}
	repo := training.NewTrainingDataRepo(client)

	err = repo.Create()
	if err != nil {
		panic(err)
	}

	for property := range pphcFetcher.GetProperties() {
		err = retryDuring(10*time.Minute, 10*time.Second, func() error {
			err = repo.StartTxn()
			if err != nil {
				return err
			}

			err := repo.Add(property)
			if err != nil {
				fmt.Errorf("adding property to repo errored: %s", err)
			}

			err = repo.Commit()
			if err != nil {
				return err
			}

			println(fmt.Sprintf("%+#v", property))
			return nil
		}, func() {
			client, err := getPachdClient()
			if err != nil {
				repo = training.NewTrainingDataRepo(client)
			}
		})

		if err != nil {
			panic(err)
		}
	}

	queue.Block()
}

func getPachdClient() (training.APIClient, error) {
	host, foundHost := os.LookupEnv("PACHD_SERVICE_HOST")
	port, foundPort := os.LookupEnv("PACHD_SERVICE_PORT")

	if !foundHost || !foundPort {
		return nil, fmt.Errorf("missing required env variable (PACHD_SERVICE_HOST|PACHD_SERVICE_PORT")
	}

	client, err := pachdclient.NewFromAddress(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
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

		log.Println("retrying after error:", err)
		attemptRepair()
	}
}
