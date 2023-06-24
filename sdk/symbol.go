package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

func (s *ApiNodeSDK) SymbolBlockList(req *pb.SymbolBlockListReq) ([]*pb.SymbolResult, error) {
	req.AppID = s.appID
	res := &pb.SymbolBlockListRes{}
	if err := s.HttpSDK.PostByAuth("/api/getSymbolBlockList", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}

func (s *ApiNodeSDK) GetBlockStatus(req *pb.GetBlockStatusReq) (*pb.GetBlockStatusRes, error) {
	req.AppID = s.appID
	res := &pb.GetBlockStatusRes{}
	if err := s.HttpSDK.PostByAuth("/api/getBlockStatus", req, res); err != nil {
		return nil, err
	}
	return res, nil
}
