package training

//go:generate counterfeiter . Repo
type Repo interface {
	Create() error
	Add(interface{}) error
}

//go:generate counterfeiter . TxnRepo
type TxnRepo interface {
	Repo
	StartTxn() (string, error)
	Commit(commitId string) error
}
