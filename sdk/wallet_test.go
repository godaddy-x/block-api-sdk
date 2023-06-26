package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/major"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"testing"
)

var symbol = "ETH"
var walletID = "W1ifWSsCxzstxXuNcq2z6PPeQyrJo3y3iM"
var accountID = "Asx7B7tW5fcXvymLZGK8QasmsX8RCWWKPcTKqmYhA6Pv"
var address = "0xd8529d73c8e1911ec45752c4e4a07a1506757a59"

func GetSDK() *ApiNodeSDK {
	api, err := InitSDK(major.SdkConfig{
		AppID:     "87faf36d8a4baa46bc65414aadd9ef07",
		AppKey:    "6a1ddc312116071880ff3ebb7a52e650b5399c710fb7fe7dce37e244dec1d8ee",
		ServeHost: "http://localhost:8422",
		Debug:     true,
	}, nil)
	if err != nil {
		panic(err)
	}
	return api
}

func GetProxySDK() *ApiNodeSDK {
	api, err := InitSDK(major.SdkConfig{
		AppID:     "87faf36d8a4baa46bc65414aadd9ef07",
		AppKey:    "6a1ddc312116071880ff3ebb7a52e650b5399c710fb7fe7dce37e244dec1d8ee",
		ServeHost: "http://localhost:7422",
		Debug:     true,
	}, nil)
	if err != nil {
		panic(err)
	}
	return api
}

func TestFindWalletByWalletID(t *testing.T) {
	res, err := GetSDK().FindWalletByWalletID(&pb.FindWalletByWalletIDReq{WalletID: "W1ifWSsCxzstxXuNcq2z6PPeQyrJo3y3iM"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFindWalletByParams(t *testing.T) {
	res, err := GetSDK().FindWalletByParams(&pb.FindWalletByParamsReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyCreateWallet(t *testing.T) {
	res, err := GetProxySDK().CreateWallet(&pb.CreateWalletReq{
		Alias:    "test",
		WalletID: "test create",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyFindWalletByWalletID(t *testing.T) {
	res, err := GetProxySDK().FindWalletByWalletID(&pb.FindWalletByWalletIDReq{WalletID: "W1ifWSsCxzstxXuNcq2z6PPeQyrJo3y3iM"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestProxyFindWalletByParams(t *testing.T) {
	res, err := GetProxySDK().FindWalletByParams(&pb.FindWalletByParamsReq{WalletID: "W1ifWSsCxzstxXuNcq2z6PPeQyrJo3y3iM"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
