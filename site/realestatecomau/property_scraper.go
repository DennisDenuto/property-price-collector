package realestatecomau

import (
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/PuerkitoBio/fetchbot"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"strings"
	"strconv"
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
	PropertyChannel chan data.Property
}

func (r RealEstateComAu) SetupMux(mux *fetchbot.Mux) {
	mux.Response().Host(r.Host).Path("/buy").Handler(propertyList(r.PropertyChannel))
	mux.Response().Host(r.Host).Path("/auction-results").Handler(auction())
}

func (r RealEstateComAu) GetProperties() <-chan data.Property {
	return r.PropertyChannel
}

func propertyList(propertyChannel chan<- data.Property) fetchbot.Handler {
	return fetchbot.HandlerFunc(func(context *fetchbot.Context, response *http.Response, err error) {
		doc, err := goquery.NewDocumentFromResponse(response)
		if err != nil {
			panic(err)
		}

		vm := otto.New()
		doc.Find("script").Each(func(index int, selection *goquery.Selection) {
			if strings.Contains(selection.Text(), "LMI.Data") {
				_, err = vm.Run(selection.Text())
			}
		})
		lmi, err := vm.Get("LMI")
		lmiData, err := lmi.Object().Get("Data")
		lmiListings, err := lmiData.Object().Get("listings")

		exportedListingArray, err := lmiListings.Export()
		listingsArrays := exportedListingArray.([]map[string]interface{})
		for _, val := range listingsArrays {
			id := getString(val["id"])
			propertyPrice := doc.Find(fmt.Sprintf("#t%s .priceText", id))
			propertyFeatures := doc.Find(fmt.Sprintf("#t%s .rui-property-features dd", id))

			numBeds := getPropertyFeature(propertyFeatures, 0)
			numBaths := getPropertyFeature(propertyFeatures, 1)
			numCars := getPropertyFeature(propertyFeatures, 2)


			property := data.Property{
				Type:     getPropertyType(getString(val["prettyDetailsUrl"])),
				Price:    propertyPrice.Text(),
				NumBeds:  numBeds,
				NumBaths: numBaths,
				NumCars:  numCars,
				Address: data.Address{
					AddressLine1: getString(val["streetAddress"]),
					State:        getString(val["state"]),
					PostCode:     getString(val["postalCode"]),
					Suburb:       getString(val["city"]),
					LonLat: data.LonLat{
						Lat: getString(val["latitude"]),
						Lon: getString(val["longitude"]),
					},
				},
			}
			propertyChannel <- property
		}
	})
}

func getPropertyType(prettyDetailsUrl string) string {
	switch true {
	case strings.Contains(prettyDetailsUrl, "house"):
		return "house"
	case strings.Contains(prettyDetailsUrl, "unit"):
		return "apartment"

	default:
		return ""
	}
}

func getPropertyFeature(propertyFeatures *goquery.Selection, idx int) (propertyFeature string) {
	defer func() {
		if r := recover(); r != nil {
			propertyFeature = ""
		}
	}()
	return propertyFeatures.Nodes[idx].FirstChild.Data
}

func getString(obj interface{}) string {
	switch objVal := obj.(type) {
	case string:
		return objVal
	case float64:
		return strconv.FormatFloat(objVal, 'f', 10, 64)
	case float32:
		return strconv.FormatFloat(float64(objVal), 'f', 10, 32)
	case int:
		return strconv.Itoa(objVal)
	default:
		return fmt.Sprintf("%v", objVal)
	}
}

func auction() fetchbot.Handler {
	return fetchbot.HandlerFunc(func(context *fetchbot.Context, response *http.Response, err error) {

	})
}
