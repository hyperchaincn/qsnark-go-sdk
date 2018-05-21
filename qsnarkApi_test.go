package qsnark

import (
	"testing"
	"qsnark-go-sdk/options"
	"qsnark-go-sdk/model"
)

func initApi(){
	opt := options.Options{
		Phone:"",
		Password:"",
		ClientId:"",
		ClientSecret:"",
	}
	if err := InitQsnackApi(opt);err != nil{
		panic(err)
	}
}

func TestInitQsnackApi(t *testing.T) {
	initApi()
	if len(QApi.Token.AccessToken) == 0 {
		t.Fatal("token 为空")
	}
}

func TestQsnackApi_RefreshAccessToken(t *testing.T) {
	initApi()
	token,err := QApi.RefreshAccessToken()
	if err != nil{
		t.Fatal()
	}
	if len(token.AccessToken) == 0 {
		t.Fatal("token 为空")
	}
}

func TestQsnackApi_QueryBlock(t *testing.T) {
	initApi()
	resp,err := QApi.QueryBlock(QueryTypeNumber,"10")
	if err != nil{
		t.Fatal(err)
	}
	if resp.Block.Hash != "0xd39abace80b6e09d1f7fcc88b218ea3303d88358124e11401fa6fc5a9eb9403f"{
		t.Fatal("区块Hash不正确", resp.Block.Hash)
	}
}

func TestQsnackApi_QueryBlockByPage(t *testing.T) {
	initApi()
	resp,err := QApi.QueryBlockByPage(1,10)
	if err != nil{
		t.Fatal(err)
	}
	if len(resp.List) == 0{
		t.Fatal("未取到区块")
	}
}

func TestQsnackApi_QueryBlockByRange(t *testing.T) {
	initApi()
	resp,err := QApi.QueryBlockByRange("1","10")
	if err != nil{
		t.Fatal(err)
	}
	if len(resp.Blocks) == 0{
		t.Fatal("未取到区块")
	}
}

func TestQsnackApi_TransactionCount(t *testing.T) {
	initApi()
	resp,err := QApi.TransactionCount()
	if err != nil{
		t.Fatal(err)
	}
	if resp.Count == 0{
		t.Fatal("不正确的交易数目",resp.Count)
	}
}

func TestQsnackApi_QueryTransaction(t *testing.T) {
	hash := `0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827`
	initApi()
	resp,err := QApi.QueryTransaction(hash)
	if err != nil{
		t.Fatal(err)
	}
	if resp.Transaction.BlockNumber != 3646{
		t.Fatal("交易Number不正确",resp.Transaction.BlockNumber)
	}
}

func TestQsnackApi_TransactionTxreceipt(t *testing.T) {
	txhash := `0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827`
	initApi()
	resp,err := QApi.TransactionTxreceipt(txhash)
	if err != nil{
		t.Fatal(err)
	}
	if resp.Ret != "0x606060405263ffffffff60e060020a600035041663867904b48114610037578063a9059cbb14610058578063f8b2cb4f14610079575bfe5b341561003f57fe5b610056600160a060020a03600435166024356100a7565b005b341561006057fe5b610056600160a060020a03600435166024356100e6565b005b341561008157fe5b610095600160a060020a036004351661017f565b60408051918252519081900360200190f35b60005433600160a060020a039081169116146100c35760006000fd5b600160a060020a03821660009081526001602052604090208054820190555b5050565b600160a060020a0333166000908152600160205260409020548190101561010d5760006000fd5b600160a060020a0333811660008181526001602090815260408083208054879003905593861680835291849020805486019055835192835282015280820183905290517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360600190a15b5050565b600160a060020a0381166000908152600160205260409020545b9190505600a165627a7a72305820299e9bb6a492d60cb690d97c76ac26d821ff6bba1b863ce1b8720e449789692c0029"{
		t.Fatal("交易Ret不正确",resp.Ret)
	}
}

func TestQsnackApi_DiscardTransaction(t *testing.T) {
	initApi()
	resp,err := QApi.DiscardTransaction(1515140350903604865,1515140350903604867)
	if err != nil{
		t.Fatal(err)
	}
	if resp.Transactions[0].Hash != "0xff24e05ac6119b87cdbf6fb224b5797b7f1dd66a9f42e8c1e98152d9ec38709b"{
		t.Fatal("交易Hash不正确",resp.Transactions[0].Hash)
	}
}

func TestQsnackApi_CompileContract(t *testing.T) {
	initApi()
	resp,err := QApi.CompileContract("contract test{}")
	if err != nil{
		t.Fatal(err)
	}
	if resp.Cts[0].Abi != "[]"{
		t.Fatal("合约Abi不正确",resp.Cts[0].Abi)
	}
}

func TestQsnackApi_DeployContract(t *testing.T) {
	initApi()
	bin := "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029"
	from := "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d"
	_,err := QApi.DeployContract(bin,from)
	if err == nil{
		t.Fatal(err)
	}
	if err.Error() != "account doesn't exist or disabled."{
		t.Fatal("结果不正确")
	}
}

func TestQsnackApi_GetPayload(t *testing.T) {
	abi := "[{\"constant\":false,\"inputs\":[{\"name\":\"num1\",\"type\":\"uint32\"},{\"name\":\"num2\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"
	initApi()
	resp,err := QApi.GetPayload(abi,[]string{"1","2"},"add")
	if err != nil{
		t.Fatal(err)
	}
	if resp != "0x3ad14af300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"{
		t.Fatal("Payload不正确",resp)
	}
}

func TestQsnackApi_InvokeContract(t *testing.T) {
	initApi()
	req := model.InvokeReq{
		Const:false,
		From:"fcf359e0069562e3120b3c9065706b52371d63cd",
		Payload:"0x3ad14af300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002",
		To:"0xbe246d87df310d73b665129ebe5d9542b1fd8044",
	}
	resp,err := QApi.InvokeContract(req)
	if err != nil{
		t.Fatal(err)
	}
	if len(resp.TxHash) == 0{
		t.Fatal("txHash不正确",resp)
	}
}

func TestQsnackApi_InvokeContractSync(t *testing.T) {
	initApi()
	req := model.InvokeReq{
		Const:false,
		From:"fcf359e0069562e3120b3c9065706b52371d63cd",
		Payload:"0x3ad14af300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002",
		To:"0xbe246d87df310d73b665129ebe5d9542b1fd8044",
	}
	resp,err := QApi.InvokeContractSync(req)
	if err != nil{
		t.Fatal(err)
	}
	if resp.Ret == ""{
		t.Fatal("txHash不正确",resp)
	}
}

func TestQsnackApi_MaintainContract(t *testing.T) {
	initApi()
	req := model.MaintainReq{
		Operation:1,
		From:"0x19a170a0413096a9b18f2ca4066faa127f4d6f4a",
		Payload:"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029",
		To:"0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8",
	}
	resp,err := QApi.MaintainContract(req)
	if err != nil{
		t.Fatal(err)
	}
	if resp.TxHash == ""{
		t.Fatal("txHash不正确",resp)
	}
}

func TestQsnackApi_ContractStatus(t *testing.T) {
	initApi()

	resp,err := QApi.ContractStatus("0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8")
	if err != nil{
		t.Fatal(err)
	}
	if resp.CtStatus != "normal"{
		t.Fatal("CtStatus不正确",resp)
	}
}

func TestQsnackApi_CreateAccount(t *testing.T) {
	initApi()

	resp,err := QApi.CreateAccount()
	if err != nil{
		t.Fatal(err)
	}

	if len(resp.Address) == 0{
		t.Fatal("Address不正确",resp)
	}

}













