package main

import (
	pphc "github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
	"fmt"
	"github.com/PuerkitoBio/fetchbot"
)

func main() {

	mux := fetchbot.NewMux()

	f := pphc.NewPropertyPriceHistoryCom("propertypricehistory.com", 2155, 2155)
	f.SetupMux(mux)

	fetcher := fetchbot.New(mux)

	queue := fetcher.Start()

	for _, seed := range f.SeedUrls {
		queue.SendStringGet(seed)
	}

	for value := range f.GetProperties() {
		println(fmt.Sprintf("%+v", value))
	}

	queue.Block()
}
