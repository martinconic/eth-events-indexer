package data

type Transaction struct {
	ID       int
	ScId     int
	TxAddr   string
	FromAddr string
	ToAddr   string
	Tokens   string
	BlockNr  int
	TxIndex  int
	Remove   bool
	LogIndex int
	LogName  string
}
