package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"testing"
)

func TestSymbolBlockList(t *testing.T) {
	res, err := GetSDK().SymbolBlockList(&pb.SymbolBlockListReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetBlockStatus(t *testing.T) {
	res, err := GetSDK().GetBlockStatus(&pb.GetBlockStatusReq{
		Symbol: symbol,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
