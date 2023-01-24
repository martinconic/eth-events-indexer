package data

import (
	"time"
)

type Contract struct {
	ID           int
	ScAddr       string
	IsIndexing   bool
	LastTxDb     string
	LastIndxDate time.Time
}

func NewContract(address string) *Contract {
	return &Contract{
		ScAddr: address,
	}
}

func UpdateContract(address string, isIndexing bool, lastTx string) *Contract {
	return &Contract{
		ScAddr:       address,
		IsIndexing:   isIndexing,
		LastTxDb:     lastTx,
		LastIndxDate: time.Now(),
	}
}
