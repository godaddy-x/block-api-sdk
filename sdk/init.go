package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/major"
	"github.com/godaddy-x/freego/utils"
	"github.com/godaddy-x/freego/utils/sdk"
)

type ApiNodeSDK struct {
	appID   string
	appKey  string
	HttpSDK *sdk.HttpSDK
}

func InitSDK(config major.SdkConfig, observer Observer) (*ApiNodeSDK, error) {
	if len(config.AppID) == 0 {
		return nil, utils.Error("config appID is nil")
	}
	if len(config.AppKey) == 0 {
		return nil, utils.Error("config appKey is nil")
	}
	if len(config.ServeHost) == 0 {
		return nil, utils.Error("config serveHost is nil")
	}
	object := map[string]string{
		"appID":     config.AppID,
		"appSecret": utils.PasswordHash(config.AppKey, utils.GetLocalSecretKey()),
	}
	httpSDK := &sdk.HttpSDK{
		Debug:     config.Debug,
		Domain:    config.ServeHost,
		KeyPath:   "/api/key",
		LoginPath: "/api/login",
	}
	httpSDK.AuthObject(&object)
	apiNodeSDK := &ApiNodeSDK{}
	apiNodeSDK.appID = config.AppID
	apiNodeSDK.appKey = config.AppKey
	apiNodeSDK.HttpSDK = httpSDK
	if err := apiNodeSDK.checkReady(); err != nil {
		return nil, err
	}
	if observer != nil { // TODO init subscribe call, asynchronous service, needs to keep the main thread from ending
		if len(config.AddrHost) == 0 {
			panic("config addrHost is nil")
		}
		go StartSubscribeNode(observer, apiNodeSDK, config.AddrHost, config.EnableProxy)
	}
	return apiNodeSDK, nil
}

func (s *ApiNodeSDK) checkReady() error {
	if len(s.appID) == 0 || len(s.appKey) == 0 {
		return utils.Error("appID/appKey invalid")
	}
	if s.HttpSDK == nil {
		return utils.Error("http sdk not initialized")
	}
	if len(s.HttpSDK.Domain) == 0 {
		return utils.Error("http sdk domain invalid")
	}
	if len(s.HttpSDK.KeyPath) == 0 {
		return utils.Error("http sdk keyPath invalid")
	}
	if len(s.HttpSDK.LoginPath) == 0 {
		return utils.Error("http sdk loginPath invalid")
	}
	return nil
}
