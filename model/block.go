package model

type Block struct {
	Number       int
	Hash         string
	ParentHash   string
	WriteTime    int64
	AvgTime      int
	Txcounts     int
	MerkleRoot   string
	Transactions []Transaction
}
