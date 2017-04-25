package image_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"github.com/DennisDenuto/property-price-collector/site/image"
	"context"
	"io"
	"io/ioutil"
	"compress/gzip"
)

var _ = Describe("MultiDownloader", func() {

	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()

		server.AppendHandlers(
			ghttp.RespondWith(200, "content"),
			ghttp.RespondWith(200, "content"),
			ghttp.RespondWith(200, "content"),
		)
	})

	It("should download every url provoded", func() {
		readers, err := image.MultiDownload([]string{server.URL(), server.URL(), server.URL()}, context.Background())
		Expect(err).ToNot(HaveOccurred())

		Expect(readers).To(HaveLen(3))
		Expect(readers[0]).To(WithTransform(decompressToString, Equal("content")))
		Expect(readers[1]).To(WithTransform(decompressToString, Equal("content")))
		Expect(readers[2]).To(WithTransform(decompressToString, Equal("content")))
	})

})

func decompressToString(reader io.Reader) string {
	gzipReader, err := gzip.NewReader(reader)
	Expect(err).ToNot(HaveOccurred())
	all, err := ioutil.ReadAll(gzipReader)

	Expect(err).ToNot(HaveOccurred())
	return string(all)
}
