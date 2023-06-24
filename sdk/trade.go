package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"github.com/blocktree/openwallet/v2/openwallet"
)

type RawTx struct {
	AccountID string              `json:"accountID"`
	Coin      openwallet.Coin     `json:"coin"`
	RawHex    string              `json:"rawHex"`
	RawHexSig string              `json:"rawHexSig"`
	ReqSigs   uint64              `json:"reqSigs"`
	Sid       string              `json:"sid"`
	SigCount  int64               `json:"sigCount"`
	ExtParam  string              `json:"extParam"`
	Fees      string              `json:"fees"`
	To        map[string]string   `json:"to"`
	FeeRate   string              `json:"feeRate"`
	SigParts  map[string][]KeySig `json:"sigParts"`
	ErrorMsg  map[string]string   `json:"errorMsg"`
}

type KeySig struct {
	Address     string `json:"address"`
	DerivedPath string `json:"derivedPath"`
	EccType     uint32 `json:"eccType"`
	InputIndex  int64  `json:"inputIndex"`
	Msg         string `json:"msg"`
	Nonce       string `json:"nonce"`
	Signed      string `json:"signed"`
	WalletID    string `json:"walletID"`
	IsImport    int64  `json:"isImport"`
	Publickey   string `json:"publickey"`
	RSV         bool   `json:"rsv"`
}

type SubmitTradeReq struct {
	AppID string   `json:"appID"`
	RawTx []*RawTx `json:"rawTx"`
}

type SubmitTradeRes struct {
	Success []*pb.TradeLogResult `json:"success"`
	Failure []*FailedRawTx       `json:"failure"`
}

type FailedRawTx struct {
	RawTx  *RawTx `json:"rawTx"`
	Reason string `json:"error"`
}

func (s *ApiNodeSDK) CreateTrade(req *pb.CreateTradeReq) (*RawTx, error) {
	req.AppID = s.appID
	res := &RawTx{}
	if err := s.HttpSDK.PostByAuth("/api/createTrade", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ApiNodeSDK) SubmitTrade(req *SubmitTradeReq) (*SubmitTradeRes, error) {
	req.AppID = s.appID
	res := &SubmitTradeRes{}
	if err := s.HttpSDK.PostByAuth("/api/submitTrade", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ApiNodeSDK) CreateSummaryTx(req *pb.CreateSummaryTxReq) ([]*RawTx, error) {
	req.AppID = s.appID
	var res []*RawTx
	if err := s.HttpSDK.PostByAuth("/api/createSummaryTx", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}
