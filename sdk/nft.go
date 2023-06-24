package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

func (s *ApiNodeSDK) GetNFTCollection(req *pb.GetNFTCollectionReq) ([]*pb.NFTResult, error) {
	req.AppID = s.appID
	res := &pb.GetNFTCollectionRes{}
	if err := s.HttpSDK.PostByAuth("/api/getNFTCollection", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetNFTToken(req *pb.GetNFTTokenReq) ([]*pb.NFTResult, error) {
	req.AppID = s.appID
	res := &pb.GetNFTTokenRes{}
	if err := s.HttpSDK.PostByAuth("/api/getNFTToken", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetNFTOwnerToken(req *pb.GetNFTOwnerTokenReq) ([]*pb.NFTResult, error) {
	req.AppID = s.appID
	res := &pb.GetNFTOwnerTokenRes{}
	if err := s.HttpSDK.PostByAuth("/api/getNFTOwnerToken", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetNFTTransfer(req *pb.GetNFTTransferReq) ([]*pb.NFTTransferResult, error) {
	req.AppID = s.appID
	res := &pb.GetNFTTransferRes{}
	if err := s.HttpSDK.PostByAuth("/api/getNFTTransfer", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetTokenURI(req *pb.GetTokenURIReq) (string, error) {
	req.AppID = s.appID
	res := &pb.GetTokenURIRes{}
	if err := s.HttpSDK.PostByAuth("/api/getTokenURI", req, res); err != nil {
		return "", err
	}
	return res.URI, nil
}
