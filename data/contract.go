package data

import "time"

type Contract struct {
	ID           int
	ScAddr       string
	IsIndexing   bool
	LastTxDb     string
	LastIndxDate time.Time
}
