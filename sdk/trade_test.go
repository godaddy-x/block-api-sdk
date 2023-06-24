package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"github.com/godaddy-x/freego/utils"
	"testing"
)

func TestCreateTrade(t *testing.T) {
	//scanner.InitRedis()
	req := &pb.CreateTradeReq{
		AccountID: accountID,
		Sid:       utils.NextSID(),
		Coin:      &pb.CoinInfo{Symbol: symbol, IsContract: false, ContractID: ""},
		To:        map[string]string{address: "2"},
	}
	res, err := GetSDK().CreateTrade(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestSummaryTx(t *testing.T) {
	req := &pb.CreateSummaryTxReq{
		AccountID: accountID,
		Address:   address,
		//FeeRate:           "0.001",
		Sid:               utils.NextSID(),
		AddressStartIndex: 0,
		AddressLimit:      10,
		MinTransfer:       "0",
		RetainedBalance:   "0",
		Coin:              &pb.CoinInfo{Symbol: symbol, IsContract: false, ContractID: ""},
		FeesSupportAccount: &pb.FeesSupportAccount{
			AccountID:        "HHieCKBaqkscY2BvXp3hby69KZo1ppkKWbd8BmngZkDH",
			FixSupportAmount: "0.01",
			FeesSupportScale: "0.01",
		},
	}
	res, err := GetSDK().CreateSummaryTx(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func SubmitTrade(tx *RawTx) {
	req := &SubmitTradeReq{
		RawTx: []*RawTx{tx},
	}
	res, err := GetSDK().SubmitTrade(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

//func SignTradeTx(v *RawTx) *RawTx {
//	symbol, err := tradeutil.GetSymbol(v.Coin.Symbol)
//	if err != nil {
//		panic(err)
//	}
//	if symbol.Id == 0 {
//		panic(err)
//	}
//	account, err := tradeutil.GetAccount(v.AccountID)
//	if err != nil {
//		panic(err)
//	}
//	if account.Id == 0 {
//		panic(err)
//	}
//	tx, err := impl.GetRawTransaction(symbol, account, v)
//	if err != nil {
//		panic(err)
//	}
//	// 数据签名
//	tm := testInitWalletManager()
//	defer tm.CloseDB(testApp)
//	_, err = tm.SignTransaction(testApp, walletID, accountID, "12345678", tx)
//	if err != nil {
//		panic(err)
//	}
//	//_, err = tm.VerifyTransaction(testApp, walletID, accountID, tx)
//	//if err != nil {
//	//	panic(err)
//	//}
//	sign := tx.Signatures[accountID][0].Signature
//	fmt.Println("签名结果:")
//	fmt.Println(sign)
//	v.SigParts[accountID][0].Signed = sign
//	return v
//}
