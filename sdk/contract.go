package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

type SmartRawTx struct {
	Coin         pb.CoinInfo         `json:"coin"`         //@required 区块链类型标识
	WalletID     string              `json:"walletID"`     //@required 创建交易单的账户
	AccountID    string              `json:"accountID"`    //@required 创建交易单的账户
	TxID         string              `json:"txID"`         //交易单ID，广播后会生成
	Sid          string              `json:"sid"`          //@required 业务订单号，保证业务不重复交易而用
	Raw          string              `json:"raw"`          //交易单调用参数，根据RawType填充数据
	RawType      int                 `json:"rawType"`      // 0：hex字符串，1：json字符串，2：base64字符串
	FeeRate      string              `json:"feeRate"`      //自定义费率
	Fees         string              `json:"fees"`         //手续费
	Value        string              `json:"value"`        //主币数量
	SigParts     map[string][]KeySig `json:"sigParts"`     //拥有者accountID: []未花签名
	AwaitResult  bool                `json:"awaitResult"`  //是否广播后同时等待结果
	AwaitTimeout int64               `json:"awaitTimeout"` //广播后等待超时秒，0 = 默认超时90秒
}

type SubmitSmartContractTradeReq struct {
	AppID string        `json:"appID"`
	RawTx []*SmartRawTx `json:"rawTx"`
}

type SubmitSmartContractTradeRes struct {
	Success []*SmartContractReceipt    `json:"success"`
	Failure []*FailureSmartContractLog `json:"failure"`
}

type FailureSmartContractLog struct {
	Error string      `json:"error"`
	RawTx *SmartRawTx `json:"rawTx"`
}

func (s *ApiNodeSDK) GetContracts(req *pb.GetContractsReq) ([]*pb.ContractResult, error) {
	req.AppID = s.appID
	res := &pb.GetContractsRes{}
	if err := s.HttpSDK.PostByAuth("/api/getContracts", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) CreateSmartContractTrade(req *pb.CreateSmartContractTradeReq) ([]*SmartRawTx, error) {
	req.AppID = s.appID
	var res []*SmartRawTx
	if err := s.HttpSDK.PostByAuth("/api/createSmartContractTrade", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ApiNodeSDK) SubmitSmartContractTrade(req *SubmitSmartContractTradeReq) (*SubmitSmartContractTradeRes, error) {
	req.AppID = s.appID
	res := &SubmitSmartContractTradeRes{}
	if err := s.HttpSDK.PostByAuth("/api/submitSmartContractTrade", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ApiNodeSDK) CallSmartContractABI(req *pb.CallSmartContractABIReq) (*pb.CallSmartContractABIRes, error) {
	req.AppID = s.appID
	res := &pb.CallSmartContractABIRes{}
	if err := s.HttpSDK.PostByAuth("/api/callSmartContractABI", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ApiNodeSDK) FindSmartContractReceipt(req *pb.FindSmartContractReceiptReq) ([]*pb.ReceiptResult, error) {
	req.AppID = s.appID
	res := &pb.FindSmartContractReceiptRes{}
	if err := s.HttpSDK.PostByAuth("/api/findSmartContractReceipt", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) FollowSmartContractReceipt(req *pb.FollowSmartContractReceiptReq) (bool, error) {
	req.AppID = s.appID
	res := &pb.FollowSmartContractReceiptRes{}
	if err := s.HttpSDK.PostByAuth("/api/followSmartContractReceipt", req, &res.Result); err != nil {
		return false, err
	}
	return res.Result, nil
}
