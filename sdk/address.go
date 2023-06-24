package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

func (s *ApiNodeSDK) CreateAddress(req *pb.CreateAddressReq) ([]*pb.AddressResult, error) {
	req.AppID = s.appID
	res := &pb.CreateAccountRes{}
	if err := s.HttpSDK.PostByAuth("/api/createAddress", req, &res.Addresses); err != nil {
		return nil, err
	}
	return res.Addresses, nil
}

func (s *ApiNodeSDK) ImportAddress(req *pb.ImportAddressReq) ([]*pb.AddressResult, error) {
	req.AppID = s.appID
	res := &pb.ImportAddressRes{}
	if err := s.HttpSDK.PostByAuth("/api/importAddress", req, &res.Addresses); err != nil {
		return nil, err
	}
	return res.Addresses, nil
}

func (s *ApiNodeSDK) FindAddressByAddress(req *pb.FindAddressByAddressReq) (*pb.AddressResult, error) {
	req.AppID = s.appID
	res := &pb.FindAddressByAddressRes{}
	if err := s.HttpSDK.PostByAuth("/api/findAddressByAddress", req, &res.Address); err != nil {
		return nil, err
	}
	return res.Address, nil
}

func (s *ApiNodeSDK) FindAddressByAccountID(req *pb.FindAddressByAccountIDReq) ([]*pb.AddressResult, error) {
	req.AppID = s.appID
	res := &pb.FindAddressByAccountIDRes{}
	if err := s.HttpSDK.PostByAuth("/api/findAddressByAccountID", req, &res.Addresses); err != nil {
		return nil, err
	}
	return res.Addresses, nil
}

func (s *ApiNodeSDK) VerifyAddress(req *pb.VerifyAddressReq) (bool, error) {
	req.AppID = s.appID
	res := &pb.VerifyAddressRes{}
	if err := s.HttpSDK.PostByAuth("/api/verifyAddress", req, &res.Result); err != nil {
		return false, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetBalanceByAddress(req *pb.GetBalanceByAddressReq) (*pb.BalanceResult, error) {
	req.AppID = s.appID
	res := &pb.GetBalanceByAddressRes{}
	if err := s.HttpSDK.PostByAuth("/api/getBalanceByAddress", req, &res.Balance); err != nil {
		return nil, err
	}
	return res.Balance, nil
}

func (s *ApiNodeSDK) GetAddressBalanceList(req *pb.GetAddressBalanceListReq) ([]*pb.BalanceResult, error) {
	req.AppID = s.appID
	res := &pb.GetAddressBalanceListRes{}
	if err := s.HttpSDK.PostByAuth("/api/getAddressBalanceList", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}
