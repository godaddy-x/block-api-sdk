package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
)

func (s *ApiNodeSDK) CreateSubscribe(req *pb.CreateSubscribeReq) (bool, error) {
	req.AppID = s.appID
	res := &pb.CreateSubscribeRes{}
	if err := s.HttpSDK.PostByAuth("/api/subscribe", req, &res.Result); err != nil {
		return false, err
	}
	return res.Result, nil
}
