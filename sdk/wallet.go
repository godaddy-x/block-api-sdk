package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"github.com/blocktree/openwallet/v2/hdkeystore"
)

func (s *ApiNodeSDK) CreateWallet(req *pb.CreateWalletReq) (*pb.WalletResult, error) {
	if len(req.RootPath) == 0 {
		req.RootPath = hdkeystore.OpenwCoinTypePath
	}
	req.AppID = s.appID
	res := &pb.CreateWalletRes{}
	if err := s.HttpSDK.PostByAuth("/api/createWallet", req, res); err != nil {
		return nil, err
	}
	return s.FindWalletByWalletID(req.WalletID)
}

func (s *ApiNodeSDK) FindWalletByWalletID(walletID string) (*pb.WalletResult, error) {
	req := &pb.FindWalletByWalletIDReq{
		AppID:    s.appID,
		WalletID: walletID,
	}
	res := &pb.FindWalletByWalletIDRes{}
	if err := s.HttpSDK.PostByAuth("/api/findWalletByWalletID", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) FindWalletByParams(req *pb.FindWalletByParamsReq) ([]*pb.WalletResult, error) {
	req.AppID = s.appID
	res := &pb.FindWalletByParamsRes{}
	if err := s.HttpSDK.PostByAuth("/api/findWalletByParams", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}
