package image_test

import (
	. "github.com/DennisDenuto/property-price-collector/site/image"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

var _ = Describe("Downloader", func() {

	var server *ghttp.Server
	var responseBody *string
	var httpStatus *int

	BeforeEach(func() {
		tmpBody := ""
		tmpStatus := 200
		responseBody = &tmpBody
		httpStatus = &tmpStatus
		server = ghttp.NewServer()
		server.AppendHandlers(
			ghttp.RespondWithPtr(httpStatus, responseBody, nil),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("given a valid url", func() {
		BeforeEach(func() {
			*responseBody = "some content"
		})

		It("should download", func() {
			content, err := Download(server.URL(), context.TODO())
			Expect(err).ToNot(HaveOccurred())

			all, err := ioutil.ReadAll(content)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(all)).To(Equal(*responseBody))
		})
	})
	Context("given a valid url that takes a long time", func() {
		BeforeEach(func() {
			server.WrapHandler(0, http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
				time.Sleep(3 * time.Second)
			}))
		})

		It("should download", func() {
			context, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			_, err := Download(server.URL(), context)
			Expect(err).To(HaveOccurred())
		})
	})
	Context("given an invalid url that takes a long time", func() {
		It("should return an error", func() {
			_, err := Download(":", context.TODO())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unable to build http request"))
		})
	})
})
