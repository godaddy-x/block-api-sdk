package main

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/major"
	"github.com/blocktree/go-openw-sdk/v2/sdk"
	"github.com/godaddy-x/freego/utils"
)

type NewObserver struct {
	sdk.UnimplementedObserver
}

func (s *NewObserver) TransactionNotify(transaction *sdk.Transaction) error {
	str, _ := utils.JsonMarshal(transaction)
	fmt.Println("TransactionNotify: ", string(str))
	return utils.Error("test TransactionNotify error")
}

func (s *NewObserver) BlockNotify(blockHeader *sdk.BlockHeader) error {
	fmt.Println("BlockNotify: ", blockHeader)
	return nil
}

// 余额更新
func (s *NewObserver) BalanceUpdateNotify(balance *sdk.BalanceObject) error {
	str, _ := utils.JsonMarshal(balance)
	fmt.Println("BalanceUpdateNotify: ", string(str))
	return utils.Error("test BalanceUpdateNotify error")
}

func main() {
	config, _ := major.ReadSdkConfig()
	sdk.InitSDK(config, &NewObserver{})
	select {}
}
