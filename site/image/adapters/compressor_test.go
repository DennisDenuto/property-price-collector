package adapters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"context"
	"bytes"
	"github.com/DennisDenuto/property-price-collector/site/image/adapters"
	"io/ioutil"
	"github.com/pkg/errors"
	"strings"
)

type FakeDownloader struct {
	DownloadContents []byte
	DownloadError    error
}

func (fd FakeDownloader) Download(url string, ctx context.Context) (io.Reader, error) {
	return bytes.NewReader(fd.DownloadContents), fd.DownloadError
}

var _ = Describe("Compressor", func() {
	var singleDownloader FakeDownloader

	BeforeEach(func() {
		singleDownloader = FakeDownloader{
			DownloadContents: []byte(strings.Repeat("a", 128)),
			DownloadError:    errors.New("some download error"),
		}
	})

	It("should compress wrapped downloader", func() {
		compressor := adapters.Compress(singleDownloader)
		compressedReader, err := compressor.Download("", context.TODO())
		Expect(err).To(Equal(singleDownloader.DownloadError))

		compressedReadAll, err := ioutil.ReadAll(compressedReader)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(compressedReadAll)).To(BeNumerically("<", len(singleDownloader.DownloadContents)))
	})

})
