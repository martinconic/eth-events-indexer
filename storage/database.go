package storage

type Database interface {
	Insert(contract string) error
	Get(contract string) (string, error)
	Update(contract string) error
}
