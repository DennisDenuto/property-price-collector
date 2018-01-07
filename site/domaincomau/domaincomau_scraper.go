package domaincomau

import (
	"encoding/json"
	"fmt"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training"
	"github.com/DennisDenuto/property-price-collector/data/training/dropbox"
	"github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/robertkrimen/otto"
	"net/http"
	"strings"
	"unicode"
)

type DomainComAuFetcher struct {
	propertyHistoryDataRepo    *dropbox.PropertyHistoryDataRepo
	DomainComAuPropertyChannel chan *data.DomainComAuPropertyListWrapper
	Seeds                      <-chan string
}

func NewDomainComAu(host string,
	propertyHistoryDataRepo training.PropertyHistoryRepo,
	minPostcode int,
	maxPostcode int,
	postcodeSuburbLookup propertypricehistorycom.PostcodeSuburbLookup) DomainComAuFetcher {
	seeds := make(chan string, 100)

	go func() {
		defer func() {
			close(seeds)
		}()
		for postcode := minPostcode; postcode <= maxPostcode; postcode++ {
			suburbs, _ := postcodeSuburbLookup.GetSuburb(postcode)

			for _, suburb := range suburbs {
				historyDataChan, errChan := propertyHistoryDataRepo.List(suburb.State, suburb.Name)
				err := addAddressToCrawler(seeds, host, historyDataChan, errChan)
				if err != nil {
					fmt.Println(fmt.Sprintf("error during downloading property file: %v", err))
				}
			}
		}
	}()

	return DomainComAuFetcher{
		DomainComAuPropertyChannel: make(chan *data.DomainComAuPropertyListWrapper, 100),
		Seeds:                      seeds,
	}
}

func addAddressToCrawler(seedChan chan<- string, host string, historyDataChan <-chan *data.PropertyHistoryData, errChan <-chan error) error {
	for {
		select {
		case propertyHistory, open := <-historyDataChan:
			if !open {
				return nil
			}
			streetAddress := strings.TrimSpace(strings.Replace(propertyHistory.Address.AddressLine1, propertyHistory.Address.Suburb, "", 1))
			streetAddress = strings.TrimSpace(strings.Replace(streetAddress, propertyHistory.Address.PostCode, "", 1))

			seedChan <- getListUri(host,
				sanitizeAddress(streetAddress),
				sanitizeAddress(propertyHistory.Address.State),
				sanitizeAddress(propertyHistory.Address.PostCode),
				sanitizeAddress(propertyHistory.Address.Suburb))

		case err := <-errChan:
			return errors.Wrap(err, "error during getting properties from dropbox %v")
		}
	}
}

func (d DomainComAuFetcher) SetupMux(mux *fetchbot.Mux) {
	mux.Response().Path("/property-profile").Handler(historicalPropertyList(d.DomainComAuPropertyChannel))
}

func historicalPropertyList(historyData chan *data.DomainComAuPropertyListWrapper) fetchbot.Handler {
	return fetchbot.HandlerFunc(func(ctx *fetchbot.Context, response *http.Response, err error) {
		logrus.Debugf("processing host %s", response.Request.URL.String())

		GetDomainComAuPropertyListWrapper(response, err, func(wrapper *data.DomainComAuPropertyListWrapper) {
			historyData <- wrapper
		})
	})
}

func GetDomainComAuPropertyListWrapper(response *http.Response, err error, handler func(wrapper *data.DomainComAuPropertyListWrapper)) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		logrus.WithError(err).Errorf("unable to get document. Skipping")
		return
	}

	vm := otto.New()
	doc.Find("script").Each(func(index int, selection *goquery.Selection) {
		if strings.Contains(selection.Text(), "viewModel.extend") && strings.Contains(selection.Text(), "address") {
			vm.Run(`
				function $(obj){return obj}
				function ViewModel(){}
				ViewModel.prototype.extend = function(json){ return JSON.stringify(json) }
				var viewModel = new ViewModel()
				`)

			domainComAuJsonWrapper, err := vm.Run(selection.Text())
			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to unmarshal json: %s", selection.Text())
				return
			}

			propertyString := domainComAuJsonWrapper.String()
			domainComAuWrapper := &data.DomainComAuPropertyListWrapper{}
			err = json.Unmarshal([]byte(propertyString), domainComAuWrapper)

			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to unmarshal json: %s", propertyString)
				return
			}

			handler(domainComAuWrapper)
		}
	})
}



func GetDomainComAuPropertyListRentWrapper(response *http.Response, err error, handler func(wrapper *data.DomainComAuPropertyListRentWrapper)) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		logrus.WithError(err).Errorf("unable to get document. Skipping")
		return
	}

	vm := otto.New()
	doc.Find("script").Each(func(index int, selection *goquery.Selection) {
		if strings.Contains(selection.Text(), "var digitalData") && strings.Contains(selection.Text(), "medianPrice") {
			_, err := vm.Run(selection.Text())
			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to run javascript json: %s", selection.Text())
				return
			}

			domainComAuJsonWrapper, err := vm.Run(`JSON.stringify(digitalData)`)
			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to unmarshal json: %s", selection.Text())
				return
			}

			propertyString := domainComAuJsonWrapper.String()
			domainComAuWrapper := &data.DomainComAuPropertyListRentWrapper{}
			err = json.Unmarshal([]byte(propertyString), domainComAuWrapper)

			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to unmarshal json: %s", propertyString)
				return
			}

			handler(domainComAuWrapper)
		}
	})
}

func GetDomainComAuPropertyDetailWrapper(response *http.Response, err error, handler func(wrapper *data.DomainComAuPropertyDetailWrapper)) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		logrus.WithError(err).Errorf("unable to get document. Skipping")
		return
	}

	vm := otto.New()
	doc.Find("script").Each(func(index int, selection *goquery.Selection) {
		if strings.Contains(selection.Text(), "var digitalData") && strings.Contains(selection.Text(), "address") {
			_, err := vm.Run(selection.Text())
			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to run javascript json: %s", selection.Text())
				return
			}

			domainComAuJsonWrapper, err := vm.Run(`JSON.stringify(digitalData)`)
			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to unmarshal json: %s", selection.Text())
				return
			}

			propertyString := domainComAuJsonWrapper.String()
			domainComAuWrapper := &data.DomainComAuPropertyDetailWrapper{}
			err = json.Unmarshal([]byte(propertyString), domainComAuWrapper)

			if err != nil {
				logrus.WithField("url", response.Request.URL.String()).WithError(err).Errorf("Unable to unmarshal json: %s", propertyString)
				return
			}

			handler(domainComAuWrapper)
		}
	})
}

func getListUri(host string, streetAddress string, state string, postcode string, suburb string) string {
	return fmt.Sprintf("https://%s/property-profile/%s-%s-%s-%s", host, streetAddress, suburb, state, postcode)
}

func sanitizeAddress(address string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || r == '/' {
			return '-'
		}
		return unicode.ToLower(r)
	}, address)
}

func (d DomainComAuFetcher) GetProperties() <-chan *data.DomainComAuPropertyListWrapper {
	return d.DomainComAuPropertyChannel
}

func (d DomainComAuFetcher) Done() {
	close(d.DomainComAuPropertyChannel)
}
