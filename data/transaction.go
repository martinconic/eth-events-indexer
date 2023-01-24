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
	Removed  bool
	LogIndex int
	LogName  string
}

func NewTransaction(txaddr string, from string, to string, tokens string,
	block int, txindex int, remove bool, logindex int, logname string) *Transaction {
	return &Transaction{
		TxAddr:   txaddr,
		FromAddr: from,
		ToAddr:   to,
		Tokens:   tokens,
		BlockNr:  block,
		TxIndex:  txindex,
		Removed:  remove,
		LogIndex: logindex,
		LogName:  logname,
	}
}
