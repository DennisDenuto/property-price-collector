package realestatecomau_test

import (
	. "github.com/DennisDenuto/property-price-collector/site/realestatecomau"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/PuerkitoBio/fetchbot"
	"github.com/onsi/gomega/ghttp"
	"net/url"
	"fmt"
	"time"
	"github.com/DennisDenuto/property-price-collector/data"
)

var _ = Describe("PropertyScraper", func() {
	var scraper RealEstateComAu
	var testMux *fetchbot.Mux
	var fetcher *fetchbot.Fetcher
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()

		testMux = fetchbot.NewMux()
		urlParsed, err := url.Parse(server.URL())
		Expect(err).ToNot(HaveOccurred())

		scraper = NewRealEstateComAu(fmt.Sprintf("localhost:%s", urlParsed.Port()))
		scraper.SetupMux(testMux)

		fetcher = fetchbot.New(testMux)
		fetcher.CrawlDelay = 0
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
					ghttp.VerifyRequest("GET", "/buy/in-nsw/list-1"),
					ghttp.RespondWith(200, ReadRealEstateComAu_Buy_list_1),
				),

				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/4960a40e2d6596f45311eb640fe32ebd242b866c146f84c612b0c3810eb2494d/main.jpg"),
					ghttp.RespondWith(200, []byte(`pic1`)),
				),

				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/abcd8d01fa8543ab85d188d0888d6f579f7d50dfdeea57abbdecc5a2c6beb872/image2.jpg"),
					ghttp.RespondWith(200, []byte(`pic2`)),
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

			receivedProperty := &data.Property{}
			Eventually(scraper.GetProperties(), 2*time.Second).Should(Receive(receivedProperty))

			Expect(receivedProperty.Address).To(Equal(data.Address{
				AddressLine1: "82 Orana Avenue",
				State:        "NSW",
				Suburb:       "Seven Hills",
				PostCode:     "2147",
				LonLat: data.LonLat{
					Lon: "150.9258728027",
					Lat: "-33.7793083191",
				},
			}))

			Expect(receivedProperty.Type).To(Equal("house"))
			Expect(receivedProperty.Price).To(Equal("Price guide $700,000 - $760,000"))
			//Expect(receivedProperty.Images).To(HaveLen(7))
		})
	})

})
