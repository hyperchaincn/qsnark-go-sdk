package model

/*
{
  "access_token": "P5EALRU-O_EIT1PEJAHYQG",
  "expires_in": 7200,
  "refresh_token": "HAGQZRWXWNYL3FNMC_UAXG",
  "scope": "all",
  "token_type": "
}
 */
type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
	TokenType string `json:"token_type"`
}

type Resp struct {
	Code int `json:"code"`
	Status string `json:"status"`
}

type QueryBlockResp struct {
	Resp
	Block Block `json:"block"`
}

type QueryBlockByPageResp struct {
	Resp
	List []Block `json:"List"`
	Count int64 `json:"Count"`
}

type QueryBlockByRangeResp struct {
	Resp
	Blocks []Block `json:"Blocks"`
}

type CompileResp struct {
	Resp
	Cts []CompileResult
}

type DeployResp struct {
	Resp
	TxHash string
}

type DeploySyncResp struct {
	Resp
	TxHash string
	PostState string
	ContractAddress string
	Ret string
}

type CompileResult struct {
	Id int
	Status string
	Bin string
	Abi string
	Name string
	OK bool `json:"OK"`
}

type InvokeResp struct {
	Resp
	TxHash string
}

type InvokeSyncResp struct {
	Resp
	TxHash string
	PostState string
	ContractAddress string
	Ret string
	DecodeRet string
}

type MaintainResp struct {
	Resp
	TxHash string
}

type StatusResp struct {
	Resp
	CtStatus string `json:"ctStatus"`
}

type TxCountResp struct {
	Resp
	Count int64
	Timestamp int64
}

type QueryTxResp struct {
	Resp
	Transaction Transaction
}

type TransactionTxreceiptResp struct {
	Resp
	TxHash string
	PostState string
	ContractAddress string
	Ret string
}

type DiscardTransactionResp struct {
	Resp
	Transactions []Transaction
}

type AccountResp struct {
	Resp
	Id int `json:"id"`
	Address string `json:"address"`
	Time string `json:"time"`
	IsDisabled bool `json:"isDisabled"`
	AppName string `json:"appName"`
}


