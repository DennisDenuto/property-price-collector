package dropbox

import (
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"strings"
	"unicode"
	"path/filepath"
	"github.com/golang/go/src/pkg/encoding/json"
	"io"
)

type PropertyHistoryDataRepo struct {
	token         string
	dropboxClient client
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
	_, err := p.dropboxClient.Upload(files.NewCommitInfo(fileName), pr)
	if err != nil {
		panic(err)
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
