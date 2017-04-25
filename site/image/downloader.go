package image

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/pkg/errors"
	"bytes"
)

type Downloader interface {
	Download(url string, ctx context.Context) (io.Reader, error)
}

type DownloadFunc func(url string, ctx context.Context) (io.Reader, error)

func (df DownloadFunc) Download(url string, ctx context.Context) (io.Reader, error) {
	return df(url, ctx)
}

type SingleDownload struct{}

func NewSingleDownloader() Downloader {
	return SingleDownload{}
}
func (SingleDownload) Download(url string, ctx context.Context) (io.Reader, error) {
	request, err := http.NewRequest("", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build http request when downloading")
	}

	requestWithCtx := request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(requestWithCtx)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read response body when downloading")
	}

	return bytes.NewReader(all), nil
}
