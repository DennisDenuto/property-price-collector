package realestatecomau

import (
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/PuerkitoBio/fetchbot"
	"net/http"
	"fmt"
	"io/ioutil"
)

//https://www.realestate.com.au/auction-results/nsw.html?rsf=edm:auction:nsw

func NewRealEstateComAu(host string) RealEstateComAu {
	return RealEstateComAu{
		Host: host,
		SeedUrls: []string{
			fmt.Sprintf("http://%s/buy/in-nsw/list-1", host),
			fmt.Sprintf("https://%s/auction-results/nsw.html?rsf=edm:auction:nsw", host),
		},
		PropertyChannel: make(chan data.Property, 100),
	}
}

type RealEstateComAu struct {
	Host            string
	SeedUrls        []string
	PropertyChannel <-chan data.Property
}

func (r RealEstateComAu) SetupMux(mux *fetchbot.Mux) {
	mux.Response().Host(r.Host).Path("/buy").Handler(propertyList())
	mux.Response().Host(r.Host).Path("/auction-results").Handler(auction())
}

func (r RealEstateComAu) GetProperties() <-chan data.Property {
	return r.PropertyChannel
}

func propertyList() fetchbot.Handler {
	return fetchbot.HandlerFunc(func(context *fetchbot.Context, response *http.Response, err error) {
		println("hi")
	})
}

func auction() fetchbot.Handler {
	return fetchbot.HandlerFunc(func(context *fetchbot.Context, response *http.Response, err error) {

	})
}
