package storage

import "github.com/martinconic/eth-events-indexer/data"

type Contracts interface {
	GetSmartContracts() ([]data.Contract, error)
	UpdateIndexing(contract string, isIndexing bool) (string, error)
	Insert(contract string) (int64, error)
	Get(contract string) (int, error)
	Update(contract *data.Contract) (string, error)
}

type Events interface {
	GetEvents(cid int) ([]data.Transaction, error)
	InsertEvent(tx *data.Transaction) (string, error)
}

type Database interface {
	Contracts
	Events
}
