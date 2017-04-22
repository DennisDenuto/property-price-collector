package adapters

import (
	"github.com/DennisDenuto/property-price-collector/site/image"
	"io"
	"context"
	"compress/gzip"
	"bytes"
)

func Compress(downloader image.Downloader) image.Downloader {
	return image.DownloadFunc(func(url string, ctx context.Context) (io.Reader, error) {
		downloadReader, downloadErr := downloader.Download(url, ctx)

		var buffer bytes.Buffer
		compressWriter := gzip.NewWriter(&buffer)
		defer compressWriter.Close()

		_, err := io.Copy(compressWriter, downloadReader)
		if err != nil {
			panic(err)
		}

		return &buffer, downloadErr
	})
}
