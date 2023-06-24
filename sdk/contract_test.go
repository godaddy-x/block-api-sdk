package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"testing"
)

func TestGetContracts(t *testing.T) {
	res, err := GetSDK().GetContracts(&pb.GetContractsReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFollowSmartContractReceipt(t *testing.T) {
	req := &pb.FollowSmartContractReceiptReq{
		Type:            1,
		FollowContracts: []string{"test", "test1", "test1"},
	}
	res, err := GetSDK().FollowSmartContractReceipt(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestFindSmartContractReceipt(t *testing.T) {
	req := &pb.FindSmartContractReceiptReq{
		Symbol: symbol,
	}
	res, err := GetSDK().FindSmartContractReceipt(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestCreateSmartContractTrade(t *testing.T) {
	req := &pb.CreateSmartContractTradeReq{}
	res, err := GetSDK().CreateSmartContractTrade(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestSubmitSmartContractTrade(t *testing.T) {
	req := &SubmitSmartContractTradeReq{}
	res, err := GetSDK().SubmitSmartContractTrade(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestCallSmartContractABI(t *testing.T) {
	req := &pb.CallSmartContractABIReq{}
	res, err := GetSDK().CallSmartContractABI(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
