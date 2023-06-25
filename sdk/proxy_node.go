package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/major"
	"github.com/godaddy-x/freego/node"
	"github.com/godaddy-x/freego/utils/crypto"
	"github.com/pkg/errors"
)

// ******************************************************* proxy node *******************************************************

type ProxyApi interface {
	Key(ctx *node.Context) error
	Login(ctx *node.Context) error
	CreateWallet(ctx *node.Context) error
	FindWalletByWalletID(ctx *node.Context) error
	FindWalletByParams(ctx *node.Context) error
	CreateAccount(ctx *node.Context) error
	FindAccountByWalletID(ctx *node.Context) error
	FindAccountByAccountID(ctx *node.Context) error
	GetBalanceByAccount(ctx *node.Context) error
	GetAccountBalanceList(ctx *node.Context) error
	CreateAddress(ctx *node.Context) error
	ImportAddress(ctx *node.Context) error
	FindAddressByAddress(ctx *node.Context) error
	FindAddressByAccountID(ctx *node.Context) error
	VerifyAddress(ctx *node.Context) error
	GetBalanceByAddress(ctx *node.Context) error
	GetAddressBalanceList(ctx *node.Context) error
	GetSymbolBlockList(ctx *node.Context) error
	GetBlockStatus(ctx *node.Context) error
	Subscribe(ctx *node.Context) error
	FindTradeLog(ctx *node.Context) error
	CreateTrade(ctx *node.Context) error
	SubmitTrade(ctx *node.Context) error
	CreateSummaryTx(ctx *node.Context) error
	GetContracts(ctx *node.Context) error
	CreateSmartContractTrade(ctx *node.Context) error
	SubmitSmartContractTrade(ctx *node.Context) error
	CallSmartContractABI(ctx *node.Context) error
	FindSmartContractReceipt(ctx *node.Context) error
	FollowSmartContractReceipt(ctx *node.Context) error
	GetNFTCollection(ctx *node.Context) error
	GetNFTToken(ctx *node.Context) error
	GetNFTOwnerToken(ctx *node.Context) error
	GetNFTTransfer(ctx *node.Context) error
	GetTokenURI(ctx *node.Context) error
}

type UnimplementedProxyApi struct{}

func (s *UnimplementedProxyApi) Key(ctx *node.Context) error {
	return errors.Errorf("unimplemented method: key")
}

// NFT合约交易数据通知
func (s *UnimplementedProxyApi) Login(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CreateWallet(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindWalletByWalletID(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindWalletByParams(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CreateAccount(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindAccountByWalletID(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindAccountByAccountID(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetBalanceByAccount(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetAccountBalanceList(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CreateAddress(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) ImportAddress(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindAddressByAddress(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindAddressByAccountID(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) VerifyAddress(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetBalanceByAddress(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetAddressBalanceList(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetSymbolBlockList(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetBlockStatus(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) Subscribe(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindTradeLog(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CreateTrade(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) SubmitTrade(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CreateSummaryTx(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetContracts(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CreateSmartContractTrade(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) SubmitSmartContractTrade(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) CallSmartContractABI(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FindSmartContractReceipt(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) FollowSmartContractReceipt(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetNFTCollection(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetNFTToken(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetNFTOwnerToken(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetNFTTransfer(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

func (s *UnimplementedProxyApi) GetTokenURI(ctx *node.Context) error {
	return errors.Errorf("unimplemented method")
}

// ******************************************************* webapp node *******************************************************

type ProxyNode struct {
	node.HttpNode
	proxy ProxyApi
	auth  AuthInfo
}

func NewProxyNode() *ProxyNode {
	var my = &ProxyNode{}
	my.EnableECC(true)
	keyConfig, _ := major.ReadKeyConfig()
	if len(keyConfig.EcdsaKey) > 0 {
		cipher := &crypto.EccObj{}
		if err := cipher.LoadS256ECC(keyConfig.EcdsaKey); err != nil {
			panic("ECC certificate generation failed")
		}
		my.AddCipher(cipher)
	}
	jwtConfig := major.ReadJwtConfig()
	my.AddJwtConfig(jwtConfig)
	return my
}

func StartProxyNode(proxy ProxyApi, appID, appKey, proxyHost string) {
	if proxy == nil {
		panic("proxy api instance can't be empty")
	}
	my := NewProxyNode()
	my.proxy = proxy
	my.auth = AuthInfo{appID: appID, appKey: appKey}
	my.GET("/api/key", my.key, &node.RouterConfig{Guest: true})
	my.POST("/api/login", my.login, &node.RouterConfig{UseRSA: true})
	my.StartServer(proxyHost)
}

// ******************************************************* webapp api *******************************************************

func (s *ProxyNode) key(ctx *node.Context) error {
	return s.proxy.Key(ctx)
}

func (s *ProxyNode) login(ctx *node.Context) error {
	return s.proxy.Login(ctx)
}
