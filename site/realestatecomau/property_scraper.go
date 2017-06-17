package realestatecomau

import (
	"context"
	"fmt"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/site/image"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/robertkrimen/otto"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	return fetchbot.HandlerFunc(func(fetchbotCtx *fetchbot.Context, response *http.Response, err error) {
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
			id := toString(val["id"])
			propertyPrice := doc.Find(fmt.Sprintf("#t%s .priceText", id))
			propertyFeatures := doc.Find(fmt.Sprintf("#t%s .rui-property-features dd", id))

			numBeds := getPropertyFeature(propertyFeatures, 0)
			numBaths := getPropertyFeature(propertyFeatures, 1)
			numCars := getPropertyFeature(propertyFeatures, 2)

			//TODO find out the size of the property
			property := data.Property{
				Images:   getPropertyPhotos(val),
				Type:     getPropertyType(toString(val["prettyDetailsUrl"])),
				Price:    propertyPrice.Text(),
				NumBeds:  numBeds,
				NumBaths: numBaths,
				NumCars:  numCars,
				Address: data.Address{
					AddressLine1: toString(val["streetAddress"]),
					State:        toString(val["state"]),
					PostCode:     toString(val["postalCode"]),
					Suburb:       toString(val["city"]),
					LonLat: data.LonLat{
						Lat: toString(val["latitude"]),
						Lon: toString(val["longitude"]),
					},
				},
			}
			propertyChannel <- property
		}
	})
}
func getPropertyPhotos(val map[string]interface{}) []data.Image {
	var imageReaders []io.Reader
	if photos, ok := val["photos"].([]map[string]interface{}); ok {
		var urls []string
		for _, photo := range photos {
			urls = append(urls, "https://i1.au.reastatic.net/640x480"+photo["src"].(string))
		}

		imageReaders, _ = image.MultiDownload(urls, context.Background())
	}
	var propertyImages []data.Image
	for _, imageReader := range imageReaders {
		all, err := ioutil.ReadAll(imageReader)
		if err != nil {
			logrus.WithError(err).Error("Unable to read downloaded property image.")
			continue
		}
		propertyImages = append(propertyImages, all)
	}
	return propertyImages
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

func toString(obj interface{}) string {
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
