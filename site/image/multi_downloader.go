package image

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"io"
	"sync"
	"time"
)

type MultiDownloader struct {
	urls []string
}

func MultiDownload(urls []string, ctx context.Context) ([]io.Reader, error) {
	downloadChannel := make(chan io.Reader, len(urls))

	downloader := TryCompress(NewSingleDownloader())
	wg := sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		go func() {
			defer wg.Done()
			ctxTimeout, timeoutFunc := context.WithTimeout(ctx, 5*time.Second)
			compressedReader, err := downloader.Download(url, ctxTimeout)
			if err != nil {
				log.WithError(err).Errorf("Unable to download %s", url)
				return
			}
			timeoutFunc()
			downloadChannel <- compressedReader
		}()
	}

	wg.Wait()
	close(downloadChannel)

	returnedReaders := []io.Reader{}
	for downloadReader := range downloadChannel {
		returnedReaders = append(returnedReaders, downloadReader)
	}

	return returnedReaders, nil
}
