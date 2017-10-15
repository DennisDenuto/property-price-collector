package pachyderm

import (
	"bytes"
	"encoding/json"
	"github.com/DennisDenuto/property-price-collector/data"
	"github.com/DennisDenuto/property-price-collector/data/training"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"
	"unicode"
)

const training_data_repo_name = "training-data-properties"

var _ training.Repo = TrainingDataRepo{}

type TrainingDataRepo struct {
	client APIClient
}

func NewTrainingDataRepo(client APIClient) *TrainingDataRepo {
	return &TrainingDataRepo{
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

func (t TrainingDataRepo) StartTxn() (string, error) {
	commits, err := t.client.ListCommitByRepo(training_data_repo_name)
	if err != nil {
		return "", errors.Wrap(err, "listing txn")
	}

	for _, value := range commits {
		if value.Finished == nil {
			return value.Commit.ID, nil
		}
	}

	startedCommit, err := t.client.StartCommit(training_data_repo_name, "master")
	if err != nil {
		return "", errors.Wrap(err, "starting txn")
	}
	return startedCommit.ID, nil
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
			sanitizeAddress(address.State),
			sanitizeAddress(address.Suburb),
			sanitizeAddress(address.AddressLine1),
		),
		bytes.NewReader(json))
	return nil
}

func sanitizeAddress(address string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}
		return '_'
	}, address)
}

func (t TrainingDataRepo) Commit(commitId string) error {
	err := t.client.FinishCommit(training_data_repo_name, commitId)
	if err != nil {
		return errors.Wrap(err, "committing")
	}
	return nil
}
