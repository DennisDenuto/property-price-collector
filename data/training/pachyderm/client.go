package pachyderm

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"io"
)

//go:generate counterfeiter . APIClient
type APIClient interface {
	CreateRepo(repoName string) error
	ListRepo(provenance []string) ([]*pfs.RepoInfo, error)

	StartCommit(repoName string, branch string) (*pfs.Commit, error)
	FinishCommit(repoName string, commitID string) error
	ListCommitByRepo(repoName string) ([]*pfs.CommitInfo, error)
	FlushCommit(commits []*pfs.Commit, toRepos []*pfs.Repo) (client.CommitInfoIterator, error)

	PutFile(repoName string, commitID string, path string, reader io.Reader) (_ int, retErr error)
	GetFileReader(repoName string, commitID string, path string, offset int64, size int64) (io.Reader, error)
}
