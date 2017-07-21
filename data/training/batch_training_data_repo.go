package training

type BatchTrainingDataRepo struct {
	repo             *TrainingDataRepo
	batchSize        int
	currentBatchSize int
}

func NewBatchTrainingDataRepo(repo *TrainingDataRepo, batchSize int) *BatchTrainingDataRepo {
	return &BatchTrainingDataRepo{
		repo:      repo,
		batchSize: batchSize,
	}
}

func (t *BatchTrainingDataRepo) Create() error {
	return t.repo.Create()
}

func (t *BatchTrainingDataRepo) StartTxn() (string, error) {
	return t.repo.StartTxn()
}

func (t *BatchTrainingDataRepo) Add(file interface{}) error {
	t.currentBatchSize++
	return t.repo.Add(file)
}

func (t *BatchTrainingDataRepo) Commit(commitId string) error {
	if t.currentBatchSize >= t.batchSize {
		return t.repo.Commit(commitId)
	}

	return nil
}

func (t *BatchTrainingDataRepo) ForceCommit(commitId string) error {
	return t.repo.Commit(commitId)
}
