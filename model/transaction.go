package model

type Transaction struct {
	Version string
	Hash string
	BlockNumber int
	BlockHash string
	TxIndex int
	From string
	To string
	Amount int
	Timestamp int64
	Nonce int64
	ExecuteTime int
	Payload string
	Invalid bool
	InvalidMsg string
}