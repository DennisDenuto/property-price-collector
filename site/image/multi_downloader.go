package image

import "io"

type MultiDownloader struct {
	urls []string
}

func MultiDownload(urls []string) ([]io.Reader, error) {
	return nil, nil
}
