package scraper

import (
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/PuerkitoBio/fetchbot"
)

type PropertyScraper interface {
	SetupMux(*fetchbot.Mux)
	GetProperties() <-chan data.Property
}
