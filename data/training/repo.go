package training

import (
	"github.com/pachyderm/pachyderm/src/client"
	_ "github.com/gogo/protobuf/gogoproto"

	"github.com/pachyderm/pachyderm/src/client/pfs"
	"io"
	"github.com/pkg/errors"
	"path/filepath"
	"github.com/DennisDenuto/property-price-collector/data"
	"strings"
	"encoding/json"
	"bytes"
)

const training_data_repo_name = "training-data-properties"

//go:generate counterfeiter . Repo
type Repo interface {
	Create() error
	StartTxn() error
	Add(interface{}) error
	Commit() error
}

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

type TrainingDataRepo struct {
	client APIClient
}

func NewTrainingDataRepo(client APIClient) TrainingDataRepo {
	return TrainingDataRepo{
		client: client,
	}
}

func (t TrainingDataRepo) Create() error {
	repos, err := t.client.ListRepo(nil)
	if err != nil {
		return errors.Wrap(err, "listing repos")
	}
	if len(repos) <= 0 {
		return errors.Wrap(t.client.CreateRepo(training_data_repo_name), "creating repo")
	}

	return nil
}

func (t TrainingDataRepo) StartTxn() error {
	commits, err := t.client.ListCommitByRepo(training_data_repo_name)
	if err != nil {
		return errors.Wrap(err, "listing txn")
	}

	for _, value := range commits {
		if value.Finished == nil {
			return nil
		}
	}

	_, err = t.client.StartCommit(training_data_repo_name, "master")
	if err != nil {
		return errors.Wrap(err, "starting txn")
	}
	return nil
}

func (t TrainingDataRepo) Add(file interface{}) error {
	historyData := file.(data.PropertyHistoryData)

	json, err := json.Marshal(file)
	if err != nil {
		return err
	}

	address := historyData.Address
	t.client.PutFile(
		training_data_repo_name,
		"master",
		filepath.Join(
			"/",
			strings.Replace(strings.ToLower(address.State), " ", "_", -1),
			strings.Replace(strings.ToLower(address.Suburb), " ", "_", -1),
			strings.Replace(strings.ToLower(address.AddressLine1), " ", "_", -1),
		),
		bytes.NewReader(json))
	return nil
}

func (t TrainingDataRepo) Commit() error {
	err := t.client.FinishCommit(training_data_repo_name, "master")
	if err != nil {
		return errors.Wrap(err, "committing")
	}
	return nil
}
