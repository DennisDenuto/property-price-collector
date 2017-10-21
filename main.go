package main

import "github.com/PuerkitoBio/fetchbot"
import (
	"fmt"
	"github.com/DennisDenuto/property-price-collector/data/training/dropbox"
	"github.com/DennisDenuto/property-price-collector/site"
	pphc "github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
	log "github.com/Sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)

	var dropboxToken string
	var found bool

	if dropboxToken, found = os.LookupEnv("DROPBOX_TOKEN"); !found {
		log.Error("missing DROPBOX_TOKEN ENV")
		os.Exit(1)
	}

	mux := fetchbot.NewMux()

	//2000 2155
	minPostcode, err := strconv.Atoi(os.Getenv("START_POSTCODE"))
	if err != nil {
		log.Error("missing START_POSTCODE ENV")
		os.Exit(1)
	}
	maxPostcode, err := strconv.Atoi(os.Getenv("END_POSTCODE"))
	if err != nil {
		log.Error("missing END_POSTCODE ENV")
		os.Exit(1)
	}

	pphcFetcher := pphc.NewPropertyPriceHistoryCom("propertypricehistory.com", minPostcode, maxPostcode, &site.PostcodeSuburbStore{})
	pphcFetcher.SetupMux(mux)

	fetcher := fetchbot.New(mux)
	fetcher.AutoClose = true

	queue := fetcher.Start()

	for _, seed := range pphcFetcher.SeedUrls {
		queue.SendStringGet(seed)
	}

	repo := dropbox.NewPropertyHistoryDataRepo(dropboxToken)

	err = saveProperties(pphcFetcher, repo, queue)
	if err != nil {
		log.WithError(err).Error("saving properties returned an error.")
		os.Exit(1)
	}

	log.Debug("exiting now")
}

func saveProperties(pphcFetcher pphc.PropertyPriceHistoryCom, repo *dropbox.PropertyHistoryDataRepo, queue *fetchbot.Queue) error {
	for {
		select {
		case property, ok := <-pphcFetcher.GetProperties():
			if !ok {
				log.Debug("no more properties to save. exiting")
				return nil
			}
			err := retryDuring(1*time.Minute, 10*time.Second, func() error {
				err := repo.Add(property)
				if err != nil {
					log.WithError(err).Error("adding property to repo errored")
					return err
				}

				log.Infof("%+#v", property)

				return nil
			})

			if err != nil {
				log.WithError(err).Error("unable to write property into datastore (SKIPPING property)")
			}
		case <-queue.Done():
			log.Info("Finished: no more urls to fetch.")
			pphcFetcher.Done()
		}
	}
}

func retryDuring(duration time.Duration, sleep time.Duration, callback func() error) (err error) {
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
	}
}
