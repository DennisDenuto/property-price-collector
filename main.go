package main

import "github.com/PuerkitoBio/fetchbot"
import (
	pphc "github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
	"github.com/DennisDenuto/property-price-collector/data/training"
	"github.com/pachyderm/pachyderm/src/client"
	"fmt"
	"os"
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

	repo := training.NewTrainingDataRepo(client)

	err = repo.Create()
	if err != nil {
		panic(err)
	}

	for property := range pphcFetcher.GetProperties() {
		err = repo.StartTxn()
		if err != nil {
			panic(err)
		}
		println(fmt.Sprintf("%+#v", property))

		err := repo.Add(property)
		if err != nil {
			fmt.Errorf("adding property to repo errored: %s", err)
		}

		err = repo.Commit()
		if err != nil {
			panic(err)
		}
	}

	queue.Block()
}
func getPachdClient() (*client.APIClient, error) {
	host, foundHost := os.LookupEnv("PACHD_SERVICE_HOST")
	port, foundPort := os.LookupEnv("PACHD_SERVICE_PORT")

	if !foundHost || !foundPort {
		panic("missing required env variable (PACHD_SERVICE_HOST|PACHD_SERVICE_PORT")
	}

	client, err := client.NewFromAddress(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}
	return client, err
}
