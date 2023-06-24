package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

func (s *ApiNodeSDK) FindTradeLog(req *pb.FindTradeLogReq) ([]*pb.TradeLogResult, error) {
	req.AppID = s.appID
	res := &pb.FindTradeLogRes{}
	if err := s.HttpSDK.PostByAuth("/api/findTradeLog", req, &res.Result); err != nil {
		return nil, err
	}
	return res.Result, nil
}
