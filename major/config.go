package major

import (
	"github.com/godaddy-x/freego/utils"
	"github.com/godaddy-x/freego/utils/jwt"
	"runtime"
)

type KeyConfig struct {
	DataKey  string `json:"dataKey"`
	EcdsaKey string `json:"ecdsaKey"`
}

type SdkConfig struct {
	AppID     string `json:"appID"`
	AppKey    string `json:"appKey"`
	Debug     bool   `json:"debug"`
	ServeHost string `json:"serveHost"` // 服务端API
	AddrHost  string `json:"addrHost"`  // 服务启动ip:port
	ProxyHost string `json:"proxyHost"` // 服务启动ip:port
}

func GetPath(path string) string {
	result := "resource/" + path
	if runtime.GOOS == "linux" {
		result = "/conf/ops/" + path
	}
	return result
}

func ReadJwtConfig() jwt.JwtConfig {
	var config jwt.JwtConfig
	if err := utils.ReadLocalJsonConfig(GetPath("jwt"), &config); err != nil {
		panic(utils.AddStr("读取jwt配置失败: ", err.Error()))
	}
	return config
}

func ReadKeyConfig() (KeyConfig, error) {
	var config KeyConfig
	if err := utils.ReadLocalJsonConfig(GetPath("secret"), &config); err != nil {
		return KeyConfig{}, nil
	}
	config.DataKey = utils.HMAC_SHA256(config.DataKey, utils.GetLocalSecretKey())
	return config, nil
}

func ReadSdkConfig() (SdkConfig, error) {
	var config SdkConfig
	if err := utils.ReadLocalJsonConfig(GetPath("sdk"), &config); err != nil {
		return SdkConfig{}, nil
	}
	return config, nil
}
