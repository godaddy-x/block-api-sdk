package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/major"
	"github.com/godaddy-x/freego/node"
	"github.com/godaddy-x/freego/utils"
	"github.com/godaddy-x/freego/utils/crypto"
	"github.com/godaddy-x/freego/utils/jwt"
	"github.com/godaddy-x/freego/utils/sdk"
)

const (
	TransactionNotify          = "transactionNotify"          //订阅区块交易单通知
	BlockNotify                = "blockNotify"                //订阅区块头数据通知
	BalanceUpdateNotify        = "balanceUpdateNotify"        //订阅余额更新通信
	SmartContractReceiptNotify = "smartContractReceiptNotify" //订阅智能合约交易回执通知
	NFTTransferNotify          = "nftTransferNotify"          //订阅NFT交易数据通知
)

// *************************************************** 新版本订阅回调 ***************************************************

type AppLogin struct {
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
}

type CallResult struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errorMsg"`
}

type AuthInfo struct {
	appID  string
	appKey string
}

// ******************************************************* observer object *******************************************************

type Transaction struct {
	Id           int64                  `json:"id" bson:"_id"`
	AppID        string                 `json:"appID" bson:"appID"`
	WalletID     string                 `json:"walletID" bson:"walletID"`
	AccountID    string                 `json:"accountID" bson:"accountID"`
	Sid          string                 `json:"sid" bson:"sid"`
	TxID         string                 `json:"txID" bson:"txID"`
	WxID         string                 `json:"wxID" bson:"wxID"`
	FromAddress  []string               `json:"fromAddress" bson:"fromAddress"`
	FromAddressV []string               `json:"fromAddressV" bson:"fromAddressV"`
	ToAddress    []string               `json:"toAddress" bson:"toAddress"`
	ToAddressV   []string               `json:"toAddressV" bson:"toAddressV"`
	Amount       string                 `json:"amount" bson:"amount"`
	Fees         string                 `json:"fees" bson:"fees"`
	Type         int64                  `json:"type" bson:"type"`
	Symbol       string                 `json:"symbol" bson:"symbol"`
	ContractID   string                 `json:"contractID" bson:"contractID"`
	IsContract   int64                  `json:"isContract" bson:"isContract"`
	Confirm      int64                  `json:"confirm" bson:"confirm"`
	BlockHash    string                 `json:"blockHash" bson:"blockHash"`
	BlockHeight  int64                  `json:"blockHeight" bson:"blockHeight"`
	IsMemo       int64                  `json:"isMemo" bson:"isMemo"`
	IsMain       int64                  `json:"isMain" bson:"isMain"`
	Memo         string                 `json:"memo" bson:"memo"`
	Applytime    int64                  `json:"applytime" bson:"applytime"`
	SubmitTime   int64                  `json:"submitTime" bson:"submitTime"`
	ConfirmTime  int64                  `json:"confirmTime" bson:"confirmTime"`
	Decimals     int64                  `json:"decimals" bson:"decimals"`
	Succtime     int64                  `json:"succtime" bson:"succtime"`
	Dealstate    int64                  `json:"dealstate" bson:"dealstate"`
	Notifystate  int64                  `json:"notifystate" bson:"notifystate"`
	ContractID2  string                 `json:"contractID2" bson:"contractID2"`
	ContractName string                 `json:"contractName" bson:"contractName"`
	ContractAddr string                 `json:"contractAddr" bson:"contractAddr"`
	Contract     map[string]interface{} `json:"contract" bson:"contract"`
	Success      string                 `json:"success"`                        //用于判断交易单链上的真实状态，0：失败，1：成功
	TxType       int64                  `json:"txType"`                         //0:转账, 1:合约调用(发生于主链), >100: 自定义
	TxAction     string                 `json:"txAction"`                       //执行事件, 例如：合约的Transfer事
	BalanceMode  uint64                 `json:"balanceMode" bson:"balanceMode"` //余额模型 0.地址 1.账户
}

type BlockHeader struct {
	Hash              string `json:"hash"`
	Confirmations     uint64 `json:"confirmations"`
	Merkleroot        string `json:"merkleroot"`
	Previousblockhash string `json:"previousblockhash"`
	Height            uint64 `json:"height"`
	Version           uint64 `json:"version"`
	Time              uint64 `json:"time"`
	Fork              bool   `json:"fork"`
	Symbol            string `json:"symbol"`
}

type SmartContractReceipt struct {
	Id             int64                 `json:"id"`
	AppID          string                `json:"appID"`
	WalletID       string                `json:"walletID"`
	AccountID      string                `json:"accountID"`
	Sid            string                `json:"sid"`
	TxID           string                `json:"txID"`
	WxID           string                `json:"wxID"`
	FromAddress    string                `json:"fromAddress"`
	ToAddress      string                `json:"toAddress"`
	Value          string                `json:"value"`
	Fees           string                `json:"fees"`
	Symbol         string                `json:"symbol"`
	ContractID     string                `json:"contractID"`
	ContractName   string                `json:"contractName"`
	ContractAddr   string                `json:"contractAddr"`
	BlockHash      string                `json:"blockHash"`
	BlockHeight    int64                 `json:"blockHeight"`
	IsMain         int64                 `json:"isMain"`
	Applytime      int64                 `json:"applytime"`
	Succtime       int64                 `json:"succtime"`
	Dealstate      int64                 `json:"dealstate"`
	Notifystate    int64                 `json:"notifystate"`
	Success        string                `json:"success"`
	RawReceipt     string                `json:"rawReceipt"`
	Events         []*SmartContractEvent `json:"events"`
	SubscribeToken string                `json:"subscribeToken"` // 客户端令牌
}

type SmartContractEvent struct {
	Symbol       string `json:"symbol"`
	ContractID   string `json:"contractID"`
	ContractName string `json:"contractName"`
	ContractAddr string `json:"contractAddr"`
	Event        string `json:"event"`
	Value        string `json:"value"`
}

type NFTTransfer struct {
	Symbol      string `json:"symbol"`   //@required 主币的symbol
	Address     string `json:"address"`  // NFT集合地址
	Token       string `json:"token"`    //@required NFT的symbol
	Protocol    string `json:"protocol"` // NFT协议, ERC1155 ERC721
	Name        string `json:"name"`
	TokenID     string `json:"tokenID"`     // NFT集合mint具体ID
	From        string `json:"from"`        // 原始拥有者
	To          string `json:"to"`          // 最新拥有者
	URI         string `json:"uri"`         // metadata uri
	Operator    string `json:"operator"`    //required 被授权转账的操作者
	Amount      string `json:"amount"`      // 转移tokenID数量
	Fees        string `json:"fees"`        //手续费
	BlockHash   string `json:"blockHash"`   //@required
	BlockHeight uint64 `json:"blockHeight"` //@required
	TxID        string `json:"txID"`        //@required
	Status      string `json:"status"`      //@required 链上状态，0：失败，1：成功
}

// ******************************************************* notify node *******************************************************

type Observer interface {
	// 交易单通知
	TransactionNotify(transaction *Transaction) error
	// 区块头通知
	BlockNotify(blockHeader *BlockHeader) error
	// 余额更新
	BalanceUpdateNotify(balance *BalanceObject) error
	// 智能合约交易回执通知
	SmartContractReceiptNotify(receipt *SmartContractReceipt) error
	// NFT合约交易数据通知
	NFTTransferNotify(transfer *NFTTransfer) error
}

type UnimplementedObserver struct{}

func (s *UnimplementedObserver) TransactionNotify(transaction *Transaction) error {
	return nil
}

// 区块头通知
func (s *UnimplementedObserver) BlockNotify(blockHeader *BlockHeader) error {
	return nil
}

// 余额更新
func (s *UnimplementedObserver) BalanceUpdateNotify(balance *BalanceObject) error {
	return nil
}

// 智能合约交易回执通知
func (s *UnimplementedObserver) SmartContractReceiptNotify(receipt *SmartContractReceipt) error {
	return nil
}

// NFT合约交易数据通知
func (s *UnimplementedObserver) NFTTransferNotify(transfer *NFTTransfer) error {
	return nil
}

// ******************************************************* webapp node *******************************************************

type SubscribeNode struct {
	node.HttpNode
	obs  Observer
	auth AuthInfo
}

func NewSubscribeNode() *SubscribeNode {
	var my = &SubscribeNode{}
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

func StartSubscribeNode(ob Observer, appID, appKey, addrHost string) {
	if ob == nil {
		panic("observer instance can't be empty")
	}
	my := NewSubscribeNode()
	my.obs = ob
	my.auth = AuthInfo{appID: appID, appKey: appKey}
	my.GET("/api/key", my.key, &node.RouterConfig{Guest: true})
	my.POST("/api/login", my.login, &node.RouterConfig{UseRSA: true})
	my.POST("/api/transactionNotify", my.transactionNotify, &node.RouterConfig{})
	my.POST("/api/blockNotify", my.blockNotify, &node.RouterConfig{})
	my.POST("/api/balanceUpdateNotify", my.balanceUpdateNotify, &node.RouterConfig{})
	my.POST("/api/smartContractReceiptNotify", my.smartContractReceiptNotify, &node.RouterConfig{})
	my.POST("/api/nftTransferNotify", my.nftTransferNotify, &node.RouterConfig{})
	my.StartServer(addrHost)
}

// ******************************************************* webapp api *******************************************************

type BalanceObject struct {
	Type             int64               `json:"type"` // 1.地址 2.账户
	Symbol           string              `json:"symbol"`
	AccountID        string              `json:"accountID"`
	Address          string              `json:"address"`
	Balance          string              `json:"balance"`
	ConfirmBalance   string              `json:"confirmBalance"`
	UnconfirmBalance string              `json:"unconfirmBalance"`
	TokenBalance     *TokenBalanceObject `json:"tokenBalance"`
	SubscribeToken   string              `json:"subscribeToken"`
}

type TokenBalanceObject struct {
	IsContract       int64  `json:"isContract"`
	Balance          string `json:"balance"`
	ConfirmBalance   string `json:"confirmBalance"`
	UnconfirmBalance string `json:"unconfirmBalance"`
	ContractID       string `json:"contractID"`
}

func (s *SubscribeNode) key(ctx *node.Context) error {
	_, publicKey := ctx.RSA.GetPublicKey()
	return s.Text(ctx, publicKey)
}

func (s *SubscribeNode) login(ctx *node.Context) error {
	req := &AppLogin{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	if !utils.PasswordVerify(s.auth.appKey, utils.GetLocalSecretKey(), req.AppSecret) {
		return utils.Error("appSecret invalid")
	}
	subject := &jwt.Subject{}
	config := ctx.GetJwtConfig()
	token := subject.Create(s.auth.appID).Dev("API").Generate(config)
	secret := jwt.GetTokenSecret(token, config.TokenKey)
	return s.Json(ctx, &sdk.AuthToken{Token: token, Secret: secret, Expired: subject.Payload.Exp})
}

func (s *SubscribeNode) transactionNotify(ctx *node.Context) error {
	tx := &Transaction{}
	if err := ctx.JsonBody.ParseData(tx); err != nil {
		return err
	}
	if err := s.obs.TransactionNotify(tx); err != nil {
		return s.Json(ctx, &CallResult{Success: false, ErrorMsg: err.Error()})
	}
	return s.Json(ctx, &CallResult{Success: true})
}

func (s *SubscribeNode) blockNotify(ctx *node.Context) error {
	head := &BlockHeader{}
	if err := ctx.JsonBody.ParseData(head); err != nil {
		return err
	}
	if err := s.obs.BlockNotify(head); err != nil {
		return s.Json(ctx, &CallResult{Success: false, ErrorMsg: err.Error()})
	}
	return s.Json(ctx, &CallResult{Success: true})
}

func (s *SubscribeNode) balanceUpdateNotify(ctx *node.Context) error {
	balance := &BalanceObject{}
	if err := ctx.JsonBody.ParseData(balance); err != nil {
		return err
	}
	if err := s.obs.BalanceUpdateNotify(balance); err != nil {
		return s.Json(ctx, &CallResult{Success: false, ErrorMsg: err.Error()})
	}
	return s.Json(ctx, &CallResult{Success: true})
}

func (s *SubscribeNode) smartContractReceiptNotify(ctx *node.Context) error {
	receipt := &SmartContractReceipt{}
	if err := ctx.JsonBody.ParseData(receipt); err != nil {
		return err
	}
	if err := s.obs.SmartContractReceiptNotify(receipt); err != nil {
		return s.Json(ctx, &CallResult{Success: false, ErrorMsg: err.Error()})
	}
	return s.Json(ctx, &CallResult{Success: true})
}

func (s *SubscribeNode) nftTransferNotify(ctx *node.Context) error {
	transfer := &NFTTransfer{}
	if err := ctx.JsonBody.ParseData(transfer); err != nil {
		return err
	}
	if err := s.obs.NFTTransferNotify(transfer); err != nil {
		return s.Json(ctx, &CallResult{Success: false, ErrorMsg: err.Error()})
	}
	return s.Json(ctx, &CallResult{Success: true})
}
