package propertypricehistorycom_test

import (
	. "github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"

	"fmt"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/PuerkitoBio/fetchbot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/url"
	"time"
	"github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom/propertypricehistorycomfakes"
)

var _ = Describe("HistoricalPropertyScraper", func() {
	var scraper PropertyPriceHistoryCom
	var testMux *fetchbot.Mux
	var fetcher *fetchbot.Fetcher
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()

		testMux = fetchbot.NewMux()
		urlParsed, err := url.Parse(server.URL())
		Expect(err).ToNot(HaveOccurred())

		lookup := &propertypricehistorycomfakes.FakePostcodeSuburbLookup{}
		lookup.GetSuburbReturns([]string{"Kellyville Ridge"}, true)

		scraper = NewPropertyPriceHistoryCom(fmt.Sprintf("localhost:%s", urlParsed.Port()), 2155, 2155, lookup)
		scraper.SetupMux(testMux)

		fetcher = fetchbot.New(testMux)
		fetcher.CrawlDelay = 0
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
					ghttp.VerifyRequest("GET", "/sold/list/NSW/2155/Kellyville+Ridge"),
					ghttp.RespondWith(200, PropertyPriceHistory_list_nsw_2155),
				),
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/sold/list/NSW/2155/Kellyville+Ridge/2/"),
					ghttp.RespondWith(200, PropertyPriceHistory_list_nsw_2155_last_page),
				),
			)
		})

		It("should map to properties", func() {
			queue := fetcher.Start()

			for _, seed := range scraper.SeedUrls {
				queueCount, err := queue.SendStringGet(seed)
				Expect(err).ToNot(HaveOccurred())
				Expect(queueCount).To(Equal(1))
			}

			receivedProperty := &data.PropertyHistoryData{}

			Eventually(scraper.GetProperties(), 2*time.Second).Should(Receive(receivedProperty))

			Expect(receivedProperty.Address).To(Equal(data.Address{
				AddressLine1: "16/11 Kilbenny Street",
				State:        "NSW",
				Suburb:       "Kellyville Ridge",
				PostCode:     "2155",
				LonLat: data.LonLat{
					Lon: "150.9276190",
					Lat: "-33.6984460",
				},
			}))
			Expect(receivedProperty.Type).To(Equal("apartment"))
			Expect(receivedProperty.DateSold.String()).To(ContainSubstring("2017-06-08"))
			Expect(receivedProperty.Price).To(Equal("N/A"))
			Expect(receivedProperty.NumBeds).To(Equal("2"))
			Expect(receivedProperty.NumBaths).To(Equal("3"))
			Expect(receivedProperty.NumCars).To(Equal("1"))

			Eventually(func() bool {
				for p := range scraper.GetProperties() {
					if p.Address.AddressLine1 == "1075 West Jindalee Road" {
						return true
					}
				}
				return false

			}, 2*time.Second).Should(BeTrue())

		})
	})

})
