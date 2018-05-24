# 趣链开发者平台 Go SDK
帮助go开发者便利的使用趣链开发者平台API
# 安装
```
go get github.com/hyperchaincn/qsnark-go-sdk
```
# 使用
```
	opt := options.Options{
		Phone:        "18702604793",
		Password:     "me123456",
		ClientId:     "c4bc5803-8f2c-43e6-9961-9ff3cbd1f733",
		ClientSecret: "9f5qzgZval3O7YKqp44m4jM74m03IN9S",
	}
	if err := qsnark.InitQsnackApi(opt); err != nil {
		panic(err)
	}

    resp, err := qsnark.QApi.QueryBlock(qsnark.QueryTypeNumber, "10")
```

# 接口
## 授权
### 1.获取Token
无特殊需求不需要主动获取，`InitQsnackApi()`会自动获取token并保持。
```
qsnark.QApi.GetAccessToken()
```

### 2.刷新Token
Token失效可主动刷新
```
qsnark.QApi.RefreshAccessToken()
```

## 区块
### 1.查询区块信息
```
resp, err := qsnark.QApi.QueryBlock(qsnark.QueryTypeNumber, "10")
```

### 2.分页获取区块信息
```
resp, err := qsnark.QApi.QueryBlockByPage(1, 10)
```

### 3.查询某范围内区块信息
```
resp, err := qsnark.QApi.QueryBlockByRange("1", "10")
```

## 交易
### 1.获取交易总数
```
	resp, err := qsnark.QApi.TransactionCount()
```

### 2.通过hash查询某交易
```
	hash := `0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827`
	resp, err := qsnark.QApi.QueryTransaction(hash)
```

### 3.查询交易回执
```
	txhash := `0xed70377c261bfdc7dd7f4fc15c8961c145f9457186d6ff95f60907e9fb63d827`
	resp, err := qsnark.QApi.TransactionTxreceipt(txhash)
```

### 4.查询无效交易
```
	resp, err := qsnark.QApi.DiscardTransaction(1515140350903604865, 1515140350903604867)
```

## 合约
### 1.编译合约
```
	resp, err := qsnark.QApi.CompileContract("contract test{}")
```

### 2.部署合约
```
	bin := "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029"
	from := "0x3713c3d01ae09cf32787c9c66c9c0781cf4b613d"
	_, err := qsnark.QApi.DeployContract(bin, from)
```

### 3.获取调用方法Payload
```
	abi := "[{\"constant\":false,\"inputs\":[{\"name\":\"num1\",\"type\":\"uint32\"},{\"name\":\"num2\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getSum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"
	resp, err := qsnark.QApi.GetPayload(abi, []string{"1", "2"}, "add")
```

### 4.调用合约方法
```
	req := model.InvokeReq{
		Const:   false,
		From:    "fcf359e0069562e3120b3c9065706b52371d63cd",
		Payload: "0x3ad14af300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002",
		To:      "0xbe246d87df310d73b665129ebe5d9542b1fd8044",
	}
	resp, err := qsnark.QApi.InvokeContract(req)
```

### 5.同步调用合约方法
```
	req := model.InvokeReq{
		Const:   false,
		From:    "fcf359e0069562e3120b3c9065706b52371d63cd",
		Payload: "0x3ad14af300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002",
		To:      "0xbe246d87df310d73b665129ebe5d9542b1fd8044",
	}
	resp, err := qsnark.QApi.InvokeContractSync(req)
```

### 6.合约维护
```
	req := model.MaintainReq{
		Operation: 1,
		From:      "0x19a170a0413096a9b18f2ca4066faa127f4d6f4a",
		Payload:   "0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820b4c36b8b61723f302432d246407a061599017f8607ed26f1c053b5ecc63a54200029",
		To:        "0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8",
	}
	resp, err := qsnark.QApi.MaintainContract(req)
```

### 7.查询合约状态
```
	resp, err := qsnark.QApi.ContractStatus("0xd3a7bdd391f6aa13b28a72690e19d2ab3d845ac8")
```

## 账户
### 1.新建区块链账户
```
	resp, err := qsnark.QApi.CreateAccount()
```




