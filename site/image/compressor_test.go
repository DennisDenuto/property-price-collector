package image

import (
	"bytes"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"strings"
)

type FakeDownloader struct {
	DownloadContents []byte
	DownloadError    error
}

func (fd FakeDownloader) Download(url string, ctx context.Context) (io.Reader, error) {
	if fd.DownloadContents == nil {
		return nil, fd.DownloadError
	}
	return bytes.NewReader(fd.DownloadContents), fd.DownloadError
}

var _ = Describe("Compressor", func() {
	var singleDownloader FakeDownloader

	BeforeEach(func() {
		singleDownloader = FakeDownloader{
			DownloadContents: []byte(strings.Repeat("a", 128)),
			DownloadError:    nil,
		}
	})

	It("should compress wrapped downloader", func() {
		compressor := TryCompress(singleDownloader)
		compressedReader, err := compressor.Download("", context.TODO())
		Expect(err).ToNot(HaveOccurred())

		compressedReadAll, err := ioutil.ReadAll(compressedReader)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(compressedReadAll)).To(BeNumerically("<", len(singleDownloader.DownloadContents)))
	})

	Context("when single downloader returns an error", func() {
		BeforeEach(func() {
			singleDownloader = FakeDownloader{
				DownloadContents: nil,
				DownloadError:    errors.New("some download error"),
			}
		})

		It("should return an error", func() {
			compressor := TryCompress(singleDownloader)
			_, err := compressor.Download("", context.TODO())
			Expect(err).To(Equal(singleDownloader.DownloadError))
		})

	})

})
