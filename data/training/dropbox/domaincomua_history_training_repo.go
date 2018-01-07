package dropbox

import (
	"encoding/json"
	"github.com/DennisDenuto/property-price-collector/data"
	dropboxclient "github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
	"path/filepath"
)

type DomainComAuHistoryDataRepo struct {
	dropboxClient client
}

func NewDomainComAuHistoryDataRepo(token string) *DomainComAuHistoryDataRepo {
	config := dropboxclient.Config{
		Token:    token,
		LogLevel: dropboxclient.LogInfo,
	}

	return &DomainComAuHistoryDataRepo{
		dropboxClient: files.New(config),
	}
}

func (repo DomainComAuHistoryDataRepo) Add(history *data.DomainComAuPropertyListWrapper) error {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		err := json.NewEncoder(pw).Encode(history)
		if err != nil {
			pw.CloseWithError(err)
		}
	}()

	fileName := filepath.Join(
		"/domaincomau/",
		sanitizeAddress(history.PropertyObject.State),
		sanitizeAddress(history.PropertyObject.Suburb),
		sanitizeAddress(history.PropertyObject.StreetAddress),
	)
	commitInfo := files.NewCommitInfo(fileName)
	commitInfo.Mode = &files.WriteMode{Tagged: dropboxclient.Tagged{files.WriteModeOverwrite}}

	_, err := repo.dropboxClient.Upload(commitInfo, pr)
	if err != nil {
		return err
	}

	return nil

}
