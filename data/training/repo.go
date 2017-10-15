package training

import "github.com/DennisDenuto/property-price-collector/data"

//go:generate counterfeiter . PropertyHistoryRepo
type PropertyHistoryRepo interface {
	Add(data data.PropertyHistoryData) error
}

//go:generate counterfeiter . TxnRepo
type TxnRepo interface {
	Create() error
	Add(interface{}) error
	StartTxn() (string, error)
	Commit(commitId string) error
}
