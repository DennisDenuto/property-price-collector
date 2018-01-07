package training

import (
	"github.com/DennisDenuto/property-price-collector/data"
)

//go:generate counterfeiter . PropertyHistoryRepo
type PropertyHistoryRepo interface {
	Add(data.PropertyHistoryData) error
	List(state, suburb string) (<-chan *data.PropertyHistoryData, <-chan error)
}

//go:generate counterfeiter . DomainComAuHistoryRepo
type DomainComAuHistoryRepo interface {
	Add(history *data.DomainComAuPropertyListWrapper) error
}

//go:generate counterfeiter . TxnRepo
type TxnRepo interface {
	Create() error
	Add(interface{}) error
	StartTxn() (string, error)
	Commit(commitId string) error
}
