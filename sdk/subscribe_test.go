package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"github.com/godaddy-x/freego/utils"
	"testing"
)

func TestCreateSubscribe(t *testing.T) {
	res, err := GetSDK().CreateSubscribe(&pb.CreateSubscribeReq{
		CallbackMode: 1,
		CallbackNode: &pb.CallbackNode{
			NodeID:      utils.NextSID(),
			Address:     "http://127.0.0.1:7422",
			ConnectType: "http",
		},
		SubscribeMethod: []string{TransactionNotify, BalanceUpdateNotify, BlockNotify, SmartContractReceiptNotify, NFTTransferNotify},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
