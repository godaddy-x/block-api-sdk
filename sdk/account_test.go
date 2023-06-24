package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"testing"
)

func TestFindAccountByAccountID(t *testing.T) {
	res, err := GetSDK().FindAccountByAccountID(&pb.FindAccountByAccountIDReq{
		Symbol:    symbol,
		AccountID: accountID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFindAccountByWalletID(t *testing.T) {
	res, err := GetSDK().FindAccountByWalletID(&pb.FindAccountByWalletIDReq{
		Symbol:   symbol,
		WalletID: walletID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetBalanceByAccount(t *testing.T) {
	res, err := GetSDK().GetBalanceByAccount(&pb.GetBalanceByAccountReq{
		Symbol:    symbol,
		AccountID: accountID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetAccountBalanceList(t *testing.T) {
	res, err := GetSDK().GetAccountBalanceList(&pb.GetAccountBalanceListReq{
		Symbol:    symbol,
		AccountID: accountID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
