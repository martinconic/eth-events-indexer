package storage

import "github.com/martinconic/eth-events-indexer/data"

type Database interface {
	Insert(contract string) (int64, error)
	Get(contract string) (int, error)
	Update(contract *data.Contract) (string, error)
	UpdateIndexing(contract string, isIndexing bool) (string, error)
	InsertEvent(tx *data.Transaction) (string, error)
	GetEvents(cid int) ([]data.Transaction, error)
}
