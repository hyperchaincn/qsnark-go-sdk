package model

type CTCode struct {
	CTCode string `json:"CTCode"`
}

type Deploy struct {
	Bin string `json:"Bin"`
	From string `json:"From"`
}

type PayloadReq struct {
	Abi string
	Args []string
	Func string
}

type InvokeReq struct {
	Const bool
	From string
	Payload string
	To string
}

type MaintainReq struct {
	Operation int
	Payload string
	From string
	To string
}
