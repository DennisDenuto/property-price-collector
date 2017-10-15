package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
)

//go:generate counterfeiter . client
type client interface {
	Upload(arg *files.CommitInfo, content io.Reader) (res *files.FileMetadata, err error)
}
