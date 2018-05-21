package qsnark

import (
	"qsnark-go-sdk/options"
	"github.com/astaxie/beego/httplib"
	"qsnark-go-sdk/model"
	"errors"
	"strconv"
	"fmt"
	"strings"
	"encoding/json"
)

var (
	QApi *QsnackApi
	remoteHost = "https://api.hyperchain.cn"
)

type QsnackApi struct {
	opts options.Options
	Token model.TokenResp
}

func InitQsnackApi(opts options.Options) error{
	if QApi == nil{
		QApi = &QsnackApi{
			opts:opts,
		}
	}
	_,err := QApi.GetAccessToken()
	if err != nil{
		return err
	}
	return nil
}

func SetRemoteHost(host string){
	remoteHost = host
}

func (q *QsnackApi)GetAccessToken() (tokenResp model.TokenResp,err error){
	api := remoteHost+GeTokenAddr
	req := httplib.Post(api)
	req.Param("phone",q.opts.Phone)
	req.Param("password",q.opts.Password)
	req.Param("client_id",q.opts.ClientId)
	req.Param("client_secret",q.opts.ClientSecret)
	if err = req.ToJSON(&tokenResp);err != nil{
		return tokenResp,err
	}
	QApi.Token = tokenResp
	return tokenResp,nil
}

func (q *QsnackApi)RefreshAccessToken() (tokenResp model.TokenResp,err error){
	api := remoteHost+RefreshTokenAddr
	req := httplib.Post(api)
	req.Param("refresh_token",q.Token.RefreshToken)
	req.Param("client_id",q.opts.ClientId)
	req.Param("client_secret",q.opts.ClientSecret)
	if err = req.ToJSON(&tokenResp);err != nil{
		return tokenResp,err
	}
	QApi.Token = tokenResp
	return tokenResp,nil
}

//  =================== 区块API ===================
const (
	QueryTypeNumber = "number"
	QueryTypeHash = "hash"
)

// 查询区块信息
func (q *QsnackApi)QueryBlock(queryType string,value string) (queryBlockResp model.QueryBlockResp,err error){
	api := remoteHost+QueryBlockAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("type",queryType)
	req.Param("value",value)
	if err = req.ToJSON(&queryBlockResp);err != nil{
		return queryBlockResp,err
	}
	if queryBlockResp.Code != 0{
		return queryBlockResp,errors.New(queryBlockResp.Status)
	}
	return queryBlockResp,nil
}

// 查询区块信息(分页)
func (q *QsnackApi)QueryBlockByPage(page int,pageSize int) (queryBlockResp model.QueryBlockByPageResp,err error){
	api := remoteHost+QueryBlockByPageAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("index",strconv.Itoa(page))
	req.Param("pageSize",strconv.Itoa(pageSize))
	if err = req.ToJSON(&queryBlockResp);err != nil{
		return queryBlockResp,err
	}
	if queryBlockResp.Code != 0{
		return queryBlockResp,errors.New(queryBlockResp.Status)
	}
	return queryBlockResp,nil
}

// 按区间查询区块
func (q *QsnackApi)QueryBlockByRange(from string,to string) (queryBlockResp model.QueryBlockByRangeResp,err error){
	api := remoteHost+QueryBlockByRangeAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("from",from)
	req.Param("to",to)
	if err = req.ToJSON(&queryBlockResp);err != nil{
		return queryBlockResp,err
	}
	if queryBlockResp.Code != 0{
		return queryBlockResp,errors.New(queryBlockResp.Status)
	}
	return queryBlockResp,nil
}

//  ================  合约API ===================

// 编译合约
func (q *QsnackApi)CompileContract(ctCode string) (compileResp model.CompileResp,err error){
	api := remoteHost+CompileContractAddr
	ctCodeReq := model.CTCode{
		CTCode:ctCode,
	}
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(ctCodeReq)
	if err != nil{
		return compileResp,err
	}
	if err = req.ToJSON(&compileResp);err != nil{
		return compileResp,err
	}
	if compileResp.Code != 0{
		return compileResp,errors.New(compileResp.Status)
	}
	return compileResp,nil
}

// 部署合约
func (q *QsnackApi)DeployContract(bin string,from string) (deployResp model.DeployResp,err error){
	api := remoteHost+DeployContractAddr
	deployReq := model.Deploy{
		Bin:bin,
		From:from,
	}
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(deployReq)
	if err != nil{
		return deployResp,err
	}
	if err = req.ToJSON(&deployResp);err != nil{
		return deployResp,err
	}
	if deployResp.Code != 0{
		return deployResp,errors.New(deployResp.Status)
	}
	return deployResp,nil
}

// 同步部署合约
func (q *QsnackApi)DeployContractSync(bin string,from string) (deploySyncResp model.DeploySyncResp,err error){
	api := remoteHost+DeployContractSyncAddr
	deployReq := model.Deploy{
		Bin:bin,
		From:from,
	}
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(deployReq)
	if err != nil{
		return deploySyncResp,err
	}
	if err = req.ToJSON(&deploySyncResp);err != nil{
		return deploySyncResp,err
	}
	if deploySyncResp.Code != 0{
		return deploySyncResp,errors.New(deploySyncResp.Status)
	}
	return deploySyncResp,nil
}

//获取 Payload
//调用合约需要转入合约方法与合约参数编码后的input字节码,该接口可为用户返回Payload
func (q *QsnackApi)GetPayload(abi string,args []string,method string) (payload string,err error){
	api := remoteHost+GetPayloadAddr
	payloadReq := model.PayloadReq{
		Abi:abi,
		Args:args,
		Func:method,
	}
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(payloadReq)
	if err != nil{
		return "",err
	}
	bytes,err := req.Bytes()
	if err != nil{
		return "",err
	}

	if err := json.Unmarshal(bytes,&payload);err != nil{
		return "",err
	}

	if len(payload) == 0 || strings.Contains(payload,"Status"){
		return "",errors.New("获取payload失败:"+payload)
	}
	return payload,nil
}

// 调用合约
func (q *QsnackApi)InvokeContract(invokeReq model.InvokeReq) (invokeResp model.InvokeResp,err error){
	api := remoteHost+InvokeContractAddr
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(invokeReq)
	if err != nil{
		return invokeResp,err
	}
	if err = req.ToJSON(&invokeResp);err != nil{
		return invokeResp,err
	}
	if invokeResp.Code != 0{
		return invokeResp,errors.New(invokeResp.Status)
	}
	return invokeResp,nil
}

// 同步调用合约
func (q *QsnackApi)InvokeContractSync(invokeReq model.InvokeReq) (invokeSyncResp model.InvokeSyncResp,err error){
	api := remoteHost+InvokeContractSyncAddr
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(invokeReq)
	if err != nil{
		return invokeSyncResp,err
	}
	fmt.Println(req.String())
	if err = req.ToJSON(&invokeSyncResp);err != nil{
		return invokeSyncResp,err
	}
	if invokeSyncResp.Code != 0{
		return invokeSyncResp,errors.New(invokeSyncResp.Status)
	}
	return invokeSyncResp,nil
}

// 合约维护
func (q *QsnackApi)MaintainContract(maintainReq model.MaintainReq) (maintainResp model.MaintainResp,err error){
	api := remoteHost+MaintainContractAddr
	req := httplib.Post(api)
	req.Header("Authorization",q.Token.AccessToken)
	req,err = req.JSONBody(maintainReq)
	if err != nil{
		return maintainResp,err
	}
	if err = req.ToJSON(&maintainResp);err != nil{
		return maintainResp,err
	}
	if maintainResp.Code != 0{
		return maintainResp,errors.New(maintainResp.Status)
	}
	return maintainResp,nil
}

// 合约状态
func (q *QsnackApi)ContractStatus(address string) (statusResp model.StatusResp,err error){
	api :=remoteHost+ContractStatusAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("address",address)
	if err = req.ToJSON(&statusResp);err != nil{
		return statusResp,err
	}
	if statusResp.Code != 0{
		return statusResp,errors.New(statusResp.Status)
	}
	return statusResp,nil
}

//  ================  交易API ===================

// 查询交易总数
func (q *QsnackApi)TransactionCount() (txCountResp model.TxCountResp,err error){
	api := remoteHost+TransactionCountAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	if err = req.ToJSON(&txCountResp);err != nil{
		return txCountResp,err
	}
	if txCountResp.Code != 0{
		return txCountResp,errors.New(txCountResp.Status)
	}
	return txCountResp,nil
}

// 查询交易信息
func (q *QsnackApi)QueryTransaction(hash string) (queryTxResp model.QueryTxResp,err error){
	api := remoteHost+QueryTransactionAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("hash",hash)
	if err = req.ToJSON(&queryTxResp);err != nil{
		return queryTxResp,err
	}
	if queryTxResp.Code != 0{
		return queryTxResp,errors.New(queryTxResp.Status)
	}
	return queryTxResp,nil
}

// 交易信息回执
func (q *QsnackApi)TransactionTxreceipt(txhash string) (transactionTxreceiptResp model.TransactionTxreceiptResp,err error){
	api := remoteHost+TransactionTxreceiptAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("txhash",txhash)
	if err = req.ToJSON(&transactionTxreceiptResp);err != nil{
		return transactionTxreceiptResp,err
	}
	if transactionTxreceiptResp.Code != 0{
		return transactionTxreceiptResp,errors.New(transactionTxreceiptResp.Status)
	}
	return transactionTxreceiptResp,nil
}

// 交易信息回执
func (q *QsnackApi)DiscardTransaction(start int64,end int64) (discardTransactionResp model.DiscardTransactionResp,err error){
	api := remoteHost+DiscardTransactionAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	req.Param("start",strconv.FormatInt(start,10))
	req.Param("end",strconv.FormatInt(end,10))
	if err = req.ToJSON(&discardTransactionResp);err != nil{
		return discardTransactionResp,err
	}
	if discardTransactionResp.Code != 0{
		return discardTransactionResp,errors.New(discardTransactionResp.Status)
	}
	return discardTransactionResp,nil
}

//  ================  账户API ===================

// 新建区块链账户
func (q *QsnackApi)CreateAccount() (accountResp model.AccountResp,err error){
	api := remoteHost+CreateAccountAddr
	req := httplib.Get(api)
	req.Header("Authorization",q.Token.AccessToken)
	if err = req.ToJSON(&accountResp);err != nil{
		return accountResp,err
	}
	if accountResp.Code != 0{
		return accountResp,errors.New(accountResp.Status)
	}
	return accountResp,nil
}





















