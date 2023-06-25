package main

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/major"
	"github.com/blocktree/go-openw-sdk/v2/sdk"
	"github.com/godaddy-x/freego/node"
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

func (s *NewObserver) BalanceUpdateNotify(balance *sdk.BalanceObject) error {
	str, _ := utils.JsonMarshal(balance)
	fmt.Println("BalanceUpdateNotify: ", string(str))
	return utils.Error("test BalanceUpdateNotify error")
}

type NewProxyApi struct {
	sdk.UnimplementedProxyApi
}

func (s *NewProxyApi) Key(ctx *node.Context) error {
	return nil
}

func (s *NewProxyApi) Login(ctx *node.Context) error {
	return nil
}

func main() {
	config, _ := major.ReadSdkConfig()
	sdk.InitSDK(config, &NewObserver{}, &NewProxyApi{})
	select {}
}
