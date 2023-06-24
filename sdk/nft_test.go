package sdk

import (
	"fmt"
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"testing"
)

func TestGetNFTCollection(t *testing.T) {
	res, err := GetSDK().GetNFTCollection(&pb.GetNFTCollectionReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetNFTToken(t *testing.T) {
	res, err := GetSDK().GetNFTToken(&pb.GetNFTTokenReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetNFTOwnerToken(t *testing.T) {
	res, err := GetSDK().GetNFTOwnerToken(&pb.GetNFTOwnerTokenReq{
		Owner: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetNFTTransfer(t *testing.T) {
	res, err := GetSDK().GetNFTTransfer(&pb.GetNFTTransferReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestGetTokenURI(t *testing.T) {
	res, err := GetSDK().GetTokenURI(&pb.GetTokenURIReq{
		Symbol:  symbol,
		TokenID: address,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
