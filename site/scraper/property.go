package scraper

import (
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/PuerkitoBio/fetchbot"
)

type ListedPropertyScraper interface {
	SetupMux(*fetchbot.Mux)
	GetProperties() <-chan data.Property
}

type HistoricalPropertyScraper interface {
	SetupMux(*fetchbot.Mux)
	GetProperties() <-chan data.PropertyHistoryData
}
