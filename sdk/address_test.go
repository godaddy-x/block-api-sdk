package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	res, err := GetSDK().CreateAddress(&pb.CreateAddressReq{
		Symbol:    symbol,
		AccountID: accountID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestImportAddress(t *testing.T) {
	res, err := GetSDK().ImportAddress(&pb.ImportAddressReq{
		Symbol:   symbol,
		WalletID: walletID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFindAddressByAddress(t *testing.T) {
	res, err := GetSDK().FindAddressByAddress(&pb.FindAddressByAddressReq{
		Symbol:  symbol,
		Address: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFindAddressByAccountID(t *testing.T) {
	res, err := GetSDK().FindAddressByAccountID(&pb.FindAddressByAccountIDReq{
		Symbol:    symbol,
		AccountID: accountID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestVerifyAddress(t *testing.T) {
	res, err := GetSDK().VerifyAddress(&pb.VerifyAddressReq{
		Symbol: symbol,
		//Address: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetBalanceByAddress(t *testing.T) {
	res, err := GetSDK().GetBalanceByAddress(&pb.GetBalanceByAddressReq{
		Symbol:  symbol,
		Address: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetAddressBalanceList(t *testing.T) {
	res, err := GetSDK().GetAddressBalanceList(&pb.GetAddressBalanceListReq{
		Symbol: symbol,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyFindAddressByAddress(t *testing.T) {
	res, err := GetProxySDK().FindAddressByAddress(&pb.FindAddressByAddressReq{
		Symbol:  symbol,
		Address: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyFindAddressByAccountID(t *testing.T) {
	res, err := GetProxySDK().FindAddressByAccountID(&pb.FindAddressByAccountIDReq{
		Symbol:    symbol,
		AccountID: accountID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyVerifyAddress(t *testing.T) {
	res, err := GetProxySDK().VerifyAddress(&pb.VerifyAddressReq{
		Symbol: symbol,
		//Address: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyGetBalanceByAddress(t *testing.T) {
	res, err := GetProxySDK().GetBalanceByAddress(&pb.GetBalanceByAddressReq{
		Symbol:  symbol,
		Address: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyGetAddressBalanceList(t *testing.T) {
	res, err := GetProxySDK().GetAddressBalanceList(&pb.GetAddressBalanceListReq{
		Symbol: symbol,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
