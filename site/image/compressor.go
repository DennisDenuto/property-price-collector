package image

import (
	"bytes"
	"compress/gzip"
	"context"
	"github.com/pkg/errors"
	"io"
)

func TryCompress(downloader Downloader) Downloader {
	return DownloadFunc(func(url string, ctx context.Context) (io.Reader, error) {
		downloadReader, downloadErr := downloader.Download(url, ctx)
		if downloadErr != nil {
			return nil, downloadErr
		}

		var buffer bytes.Buffer
		compressWriter := gzip.NewWriter(&buffer)
		defer compressWriter.Close()

		_, err := io.Copy(compressWriter, downloadReader)
		if err != nil {
			return downloadReader, errors.Wrap(downloadErr, err.Error())
		}

		return &buffer, downloadErr
	})
}
