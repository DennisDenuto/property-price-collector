package dropbox

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
)

//go:generate counterfeiter . client
type client interface {
	Upload(arg *files.CommitInfo, content io.Reader) (res *files.FileMetadata, err error)
	Download(arg *files.DownloadArg) (res *files.FileMetadata, content io.ReadCloser, err error)
	ListFolder(arg *files.ListFolderArg) (res *files.ListFolderResult, err error)
	ListFolderContinue(arg *files.ListFolderContinueArg) (res *files.ListFolderResult, err error)
}
