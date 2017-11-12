package domaincomau_test

import (
	. "github.com/DennisDenuto/property-price-collector/site/domaincomau"

	"crypto/tls"
	"fmt"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training/trainingfakes"
	"github.com/DennisDenuto/property-price-collector/site"
	"github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom/propertypricehistorycomfakes"
	"github.com/PuerkitoBio/fetchbot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"net/url"
	"time"
)

var _ = Describe("DomaincomauScraper", func() {
	var scraper DomainComAuFetcher
	var testMux *fetchbot.Mux
	var fetcher *fetchbot.Fetcher
	var server *ghttp.Server

	var fakePropertyHistoryRepo *trainingfakes.FakePropertyHistoryRepo
	var results chan *data.PropertyHistoryData
	BeforeEach(func() {
		server = ghttp.NewTLSServer()

		testMux = fetchbot.NewMux()
		urlParsed, err := url.Parse(server.URL())
		Expect(err).ToNot(HaveOccurred())

		lookup := &propertypricehistorycomfakes.FakePostcodeSuburbLookup{}
		By("Only looking up addresses for greystanes", func() {
			lookup.GetSuburbReturns([]site.Suburb{
				{Name: "Greystanes", State: "NSW"},
			}, true)
		})

		By("only two address being returned for greystanes", func() {
			fakePropertyHistoryRepo = &trainingfakes.FakePropertyHistoryRepo{}
			results = make(chan *data.PropertyHistoryData, 2)
			results <- &data.PropertyHistoryData{
				Address: data.Address{
					AddressLine1: "7 Bilpin Street Greystanes 2145",
					State:        "NSW",
					Suburb:       "Greystanes",
					PostCode:     "2145",
				},
			}
			results <- &data.PropertyHistoryData{
				Address: data.Address{
					AddressLine1: "3 Bilpin Street Greystanes 2145",
					State:        "NSW",
					Suburb:       "Greystanes",
					PostCode:     "2145",
				},
			}
			close(results)
			fakePropertyHistoryRepo.ListReturns(results, make(chan error))
		})

		scraper = NewDomainComAu(fmt.Sprintf("localhost:%s", urlParsed.Port()), fakePropertyHistoryRepo, 2155, 2155, lookup)
		scraper.SetupMux(testMux)

		fetcher = fetchbot.New(testMux)
		fetcher.CrawlDelay = 0
		fetcher.HttpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		fetcher.AutoClose = true
		fetcher.WorkerIdleTTL = 1 * time.Second
	})

	AfterEach(func() {
		server.Close()
	})

	Context("property list", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/robots.txt"),
					ghttp.RespondWith(200, `
User-agent: *
Disallow: /deny`,
					),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/property-profile/7-bilpin-street-greystanes-nsw-2145"),
					ghttp.RespondWith(200, DomainComAuPropertyProfile),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/property-profile/3-bilpin-street-greystanes-nsw-2145"),
					ghttp.RespondWith(200, DomainComAuPropertyProfileWithMultipleSoldAndRented),
				),
			)
		})

		It("should map to properties", func() {
			queue := fetcher.Start()

			for seed := range scraper.Seeds {
				queueCount, err := queue.SendStringGet(seed)
				Expect(err).ToNot(HaveOccurred())
				Expect(queueCount).To(Equal(1))
			}

			var properties *data.DomainComAuPropertyWrapper
			Eventually(scraper.GetProperties(), 2).Should(Receive(&properties))
			Expect(properties.Property.Address).To(Equal("7 Bilpin Street, Greystanes NSW 2145"))
			Expect(properties.PropertyObject.StreetAddress).To(Equal("7 Bilpin Street"))
			Expect(properties.Valuation.LowerPrice).To(Equal(800000))
			Expect(properties.Valuation.UpperPrice).To(Equal(1055000))
			Expect(properties.PropertyObject.History.Sales).To(HaveLen(2))
			Expect(properties.PropertyObject.History.Rentals).To(HaveLen(0))

			Eventually(scraper.GetProperties(), 2).Should(Receive(&properties))
			Expect(properties.Property.Address).To(Equal("3 Bilpin Street, Greystanes NSW 2145"))
			Expect(properties.Valuation.LowerPrice).To(Equal(725000))
			Expect(properties.Valuation.UpperPrice).To(Equal(955000))
			Expect(properties.PropertyObject.History.Sales).To(HaveLen(3))
			Expect(properties.PropertyObject.History.Rentals).To(HaveLen(4))
		})
	})

})
