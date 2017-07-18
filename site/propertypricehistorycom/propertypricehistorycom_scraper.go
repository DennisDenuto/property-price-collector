package propertypricehistorycom

import (
	"fmt"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

//https://www.propertypricehistory.com/sold/list/NSW/<<POSTCODE>>
//2000 - 2594

func NewPropertyPriceHistoryCom(host string, minPostcode int, maxPostcode int) PropertyPriceHistoryCom {
	var seedUrls []string = make([]string, maxPostcode-minPostcode)
	for postcode := minPostcode; postcode <= maxPostcode; postcode++ {
		seedUrls = append(seedUrls, getListUri(host, postcode))
	}

	return PropertyPriceHistoryCom{
		Host:            host,
		SeedUrls:        seedUrls,
		PropertyChannel: make(chan data.PropertyHistoryData, 100),
	}
}

func getListUri(host string, postcode int) string {
	return fmt.Sprintf("http://%s/sold/list/NSW/%d", host, postcode)
}

type PropertyPriceHistoryCom struct {
	Host            string
	SeedUrls        []string
	PropertyChannel chan data.PropertyHistoryData
}

func (r PropertyPriceHistoryCom) SetupMux(mux *fetchbot.Mux) {
	mux.Response().Path("/sold/list").Handler(historicalPropertyList(r.PropertyChannel, r.Host))
	//mux.Response().Host(r.Host).Path("/sold-price").Handler(auction())
}

func historicalPropertyList(propertyHistoryDataChannel chan data.PropertyHistoryData, host string) fetchbot.Handler {
	return fetchbot.HandlerFunc(func(fc *fetchbot.Context, response *http.Response, _ error) {
		logrus.Debugf("processing host %s", host)
		doc, err := goquery.NewDocumentFromResponse(response)
		if err != nil {
			logrus.WithError(err).Errorf("unable to get document. Skipping")
			return
		}

		doc.Find("#search_result .col-sm-12").Each(func(index int, selection *goquery.Selection) {
			propertyHistoryData := data.PropertyHistoryData{}

			addressSelection := selection.Find(".wx-name a")
			lonLatSelection := selection.Find(".wx-id .init")
			postcode, suburb := parsePostcodeAndSuburb(addressSelection.Children().Text())
			propertyHistoryData.Address = data.Address{
				AddressLine1: strings.Split(addressSelection.Contents().Text(), "\n")[0],
				PostCode:     postcode,
				Suburb:       suburb,
				State:        getStateFromUrl(addressSelection),
				LonLat: data.LonLat{
					Lon: lonLatSelection.AttrOr("data-lng", ""),
					Lat: lonLatSelection.AttrOr("data-lat", ""),
				},
			}

			propertyHistoryData.Type = strings.ToLower(selection.Find(".rank-area .rank-type .label-info").Text())
			propertyHistoryData.DateSold = getDateSold(selection)
			propertyHistoryData.Price = strings.TrimSpace(selection.Find(".rank-area .rank").Text())

			beds, baths, cars := getBedsBathsCars(selection)
			propertyHistoryData.NumBeds = beds
			propertyHistoryData.NumBaths = baths
			propertyHistoryData.NumCars = cars
			propertyHistoryDataChannel <- propertyHistoryData
		})

		if nextPageUrl, found := doc.Find(".goto_nextpage").Attr("href"); found {
			nextPageParseUrl, err := url.Parse(nextPageUrl)
			if err == nil {
				logrus.Debugf("Adding next page: http://%s%s", host, nextPageParseUrl.Path)
				fc.Q.SendStringGet(fmt.Sprintf("http://%s%s", host, nextPageParseUrl.Path))
			} else {
				logrus.WithError(err).Debug("unable to get next page url")
			}
		}
	})
}

func getBedsBathsCars(selection *goquery.Selection) (beds string, baths string, cars string) {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithError(r.(error)).Debug("Recovered getting beds and bath")
		}
	}()

	beds, err := selection.Find(".wx-id .fa-bed").Siblings().Html()
	if err != nil {
		beds = ""
	}

	return strings.TrimSpace(beds),
		strings.TrimSpace(selection.Find(".wx-id .bathrooms").Nodes[0].NextSibling.Data),
		strings.TrimSpace(selection.Find(".wx-id .fa-car").Nodes[0].NextSibling.Data)
}

func getDateSold(selection *goquery.Selection) time.Time {
	dateSold, err := time.Parse("_2 Jan 2006", strings.TrimSpace(selection.Find(".rank-area .rank-date").Text()))
	if err != nil {
		dateSold, err = time.Parse("Jan 2006", strings.TrimSpace(selection.Find(".rank-area .rank-date").Text()))
		if err != nil {
			dateSold = time.Time{}
		}
	}
	return dateSold
}

func parsePostcodeAndSuburb(val string) (postcode string, suburb string) {
	compile, err := regexp.Compile("(.*)(\\d\\d\\d\\d)")
	if err != nil {
		return "", ""
	}
	submatch := compile.FindStringSubmatch(val)
	switch len(submatch) {
	case 3:
		return strings.TrimSpace(submatch[2]), strings.TrimSpace(submatch[1])
	default:
		return "", ""
	}
}

func getStateFromUrl(addressSelection *goquery.Selection) string {
	compile, err := regexp.Compile("/(NSW|QLD|TAS|VIC|SA|WA)/")
	if err != nil {
		return ""
	}
	stateGroupMatch := compile.FindStringSubmatch(addressSelection.AttrOr("href", ""))
	if len(stateGroupMatch) > 1 {
		return stateGroupMatch[1]
	} else {
		return ""
	}
}

func (r PropertyPriceHistoryCom) GetProperties() <-chan data.PropertyHistoryData {
	return r.PropertyChannel
}
