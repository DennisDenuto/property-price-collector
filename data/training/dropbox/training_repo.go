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
