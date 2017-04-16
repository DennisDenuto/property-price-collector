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
			)
		})

		It("should map to properties", func() {
			queue := fetcher.Start()

			for _, seed := range scraper.SeedUrls {
				queueCount, err := queue.SendStringGet(seed)
				Expect(err).ToNot(HaveOccurred())
				Expect(queueCount).To(Equal(1))
			}

			Eventually(scraper.GetProperties(), 2*time.Second).Should(Receive())
		})
	})

})
