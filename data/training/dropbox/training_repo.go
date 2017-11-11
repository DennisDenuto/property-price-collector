package dropbox

import (
	"encoding/json"
	"github.com/DennisDenuto/property-price-collector/data"
	dropboxclient "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
	"path/filepath"
	"strings"
	"unicode"
	"github.com/pkg/errors"
	"sync"
)

type PropertyHistoryDataRepo struct {
	dropboxClient client
}

func NewPropertyHistoryDataRepo(token string) *PropertyHistoryDataRepo {
	config := dropboxclient.Config{
		Token:    token,
		LogLevel: dropboxclient.LogInfo,
	}

	return &PropertyHistoryDataRepo{
		dropboxClient: files.New(config),
	}
}

func (p PropertyHistoryDataRepo) List(state, suburb string) (<-chan *data.PropertyHistoryData, <-chan error) {
	propertyHistoryChan := make(chan *data.PropertyHistoryData, 100)
	entries := make(chan files.IsMetadata, 10)
	errChan := make(chan error, 1)

	go func() {
		listFolderRes, err := p.dropboxClient.ListFolder(&files.ListFolderArg{
			Path: filepath.Join("/", sanitizeAddress(state), sanitizeAddress(suburb)),
		})
		if err != nil {
			errChan <- errors.Wrap(err, "Unable to list files in directory")
			close(entries)
			return
		}

		for _, entry := range listFolderRes.Entries {
			entries <- entry
		}

		for listFolderRes.HasMore {
			listFolderRes, err = p.dropboxClient.ListFolderContinue(&files.ListFolderContinueArg{
				Cursor: listFolderRes.Cursor,
			})
			if err != nil {
				errChan <- errors.Wrap(err, "Unable to list files in directory")
				close(entries)
				return
			}
			for _, entry := range listFolderRes.Entries {
				entries <- entry
			}
		}

		close(entries)
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Wait()
		close(propertyHistoryChan)
	}()

	for {
		select {
		case entry, ok := <-entries:
			if !ok {
				wg.Done()
				return propertyHistoryChan, errChan
			}

			if _, isAFile := entry.(*files.FileMetadata); !isAFile {
				continue
			}

			wg.Add(1)
			go func(fileMetadata *files.FileMetadata) {
				defer wg.Done()

				_, dropboxDownloadedFile, err := p.dropboxClient.Download(files.NewDownloadArg(filepath.Join(fileMetadata.PathLower, fileMetadata.Name)))
				if err != nil {
					errChan <- errors.Wrap(err, "Unable to download file")
					return
				}
				defer dropboxDownloadedFile.Close()

				historyData := &data.PropertyHistoryData{}
				err = json.NewDecoder(dropboxDownloadedFile).Decode(historyData)
				if err != nil && err != io.EOF {
					errChan <- errors.Wrap(err, "Unable to unmarshal dropbox file")
					return
				}

				propertyHistoryChan <- historyData
			}(entry.(*files.FileMetadata))
		}
	}
}

func (p PropertyHistoryDataRepo) Add(data data.PropertyHistoryData) error {

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		err := json.NewEncoder(pw).Encode(data)
		if err != nil {
			pw.CloseWithError(err)
		}
	}()

	fileName := filepath.Join(
		"/",
		sanitizeAddress(data.Address.State),
		sanitizeAddress(data.Address.Suburb),
		sanitizeAddress(data.Address.AddressLine1),
	)
	commitInfo := files.NewCommitInfo(fileName)
	commitInfo.Mode = &files.WriteMode{Tagged: dropboxclient.Tagged{files.WriteModeOverwrite}}

	_, err := p.dropboxClient.Upload(commitInfo, pr)
	if err != nil {
		return err
	}

	return nil
}

func sanitizeAddress(address string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' {
			return unicode.ToLower(r)
		}
		return '_'
	}, address)
}
