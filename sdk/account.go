package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

func (s *ApiNodeSDK) CreateAccount(req *pb.CreateAccountReq) (*pb.CreateAccountRes, error) {
	req.AppID = s.appID
	res := &pb.CreateAccountRes{}
	if err := s.HttpSDK.PostByAuth("/api/createAccount", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ApiNodeSDK) FindAccountByAccountID(req *pb.FindAccountByAccountIDReq) (*pb.AccountResult, error) {
	req.AppID = s.appID
	res := &pb.FindAccountByAccountIDRes{}
	if err := s.HttpSDK.PostByAuth("/api/findAccountByAccountID", req, &res.Account); err != nil {
		return nil, err
	}
	return res.Account, nil
}

func (s *ApiNodeSDK) FindAccountByWalletID(req *pb.FindAccountByWalletIDReq) ([]*pb.AccountResult, error) {
	req.AppID = s.appID
	res := &pb.FindAccountByWalletIDRes{}
	if err := s.HttpSDK.PostByAuth("/api/findAccountByWalletID", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetBalanceByAccount(req *pb.GetBalanceByAccountReq) (*pb.BalanceResult, error) {
	req.AppID = s.appID
	res := &pb.GetBalanceByAccountRes{}
	if err := s.HttpSDK.PostByAuth("/api/getBalanceByAccount", req, &res.Balance); err != nil {
		return nil, err
	}
	return res.Balance, nil
}

func (s *ApiNodeSDK) GetAccountBalanceList(req *pb.GetAccountBalanceListReq) ([]*pb.BalanceResult, error) {
	req.AppID = s.appID
	res := &pb.GetAccountBalanceListRes{}
	if err := s.HttpSDK.PostByAuth("/api/getAccountBalanceList", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}
