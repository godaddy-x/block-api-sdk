package sdk

import (
	"github.com/blocktree/go-openw-sdk/v2/sdk/pb"
	"github.com/godaddy-x/freego/node"
	"github.com/godaddy-x/freego/utils"
	"github.com/godaddy-x/freego/utils/jwt"
	"github.com/godaddy-x/freego/utils/sdk"
)

// ******************************************************* proxy node *******************************************************

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

// ******************************************************* notify node *******************************************************

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

// ******************************************************* webapp api *******************************************************

func (s *SubscribeNode) key(ctx *node.Context) error {
	_, publicKey := ctx.RSA.GetPublicKey()
	return s.Text(ctx, publicKey)
}

func (s *SubscribeNode) login(ctx *node.Context) error {
	req := &AppLogin{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	if !utils.PasswordVerify(s.api.appKey, utils.GetLocalSecretKey(), req.AppSecret) {
		return utils.Error("appSecret invalid")
	}
	subject := &jwt.Subject{}
	config := ctx.GetJwtConfig()
	token := subject.Create(s.api.appID).Dev("API").Generate(config)
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

func (s *SubscribeNode) createWallet(ctx *node.Context) error {
	req := &pb.CreateWalletReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateWallet(req)
	if err != nil {
		return err
	}
	res1, err := s.api.FindWalletByWalletID(&pb.FindWalletByWalletIDReq{WalletID: res.WalletID})
	if err != nil {
		return err
	}
	return s.Json(ctx, res1)
}

func (s *SubscribeNode) findWalletByWalletID(ctx *node.Context) error {
	req := &pb.FindWalletByWalletIDReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindWalletByWalletID(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findWalletByParams(ctx *node.Context) error {
	req := &pb.FindWalletByParamsReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindWalletByParams(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) createAccount(ctx *node.Context) error {
	req := &pb.CreateAccountReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateAccount(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findAccountByAccountID(ctx *node.Context) error {
	req := &pb.FindAccountByAccountIDReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindAccountByAccountID(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findAccountByWalletID(ctx *node.Context) error {
	req := &pb.FindAccountByWalletIDReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindAccountByWalletID(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getBalanceByAccount(ctx *node.Context) error {
	req := &pb.GetBalanceByAccountReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetBalanceByAccount(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getAccountBalanceList(ctx *node.Context) error {
	req := &pb.GetAccountBalanceListReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetAccountBalanceList(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) createAddress(ctx *node.Context) error {
	req := &pb.CreateAddressReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateAddress(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) importAddress(ctx *node.Context) error {
	req := &pb.ImportAddressReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.ImportAddress(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findAddressByAddress(ctx *node.Context) error {
	req := &pb.FindAddressByAddressReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindAddressByAddress(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findAddressByAccountID(ctx *node.Context) error {
	req := &pb.FindAddressByAccountIDReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindAddressByAccountID(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) verifyAddress(ctx *node.Context) error {
	req := &pb.VerifyAddressReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.VerifyAddress(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getBalanceByAddress(ctx *node.Context) error {
	req := &pb.GetBalanceByAddressReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetBalanceByAddress(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getAddressBalanceList(ctx *node.Context) error {
	req := &pb.GetAddressBalanceListReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetAddressBalanceList(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getSymbolBlockList(ctx *node.Context) error {
	req := &pb.SymbolBlockListReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.SymbolBlockList(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getBlockStatus(ctx *node.Context) error {
	req := &pb.GetBlockStatusReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetBlockStatus(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) createSubscribe(ctx *node.Context) error {
	req := &pb.CreateSubscribeReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateSubscribe(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findTradeLog(ctx *node.Context) error {
	req := &pb.FindTradeLogReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindTradeLog(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) createTrade(ctx *node.Context) error {
	req := &pb.CreateTradeReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateTrade(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) submitTrade(ctx *node.Context) error {
	req := &SubmitTradeReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.SubmitTrade(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) createSummaryTx(ctx *node.Context) error {
	req := &pb.CreateSummaryTxReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateSummaryTx(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getContracts(ctx *node.Context) error {
	req := &pb.GetContractsReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetContracts(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) createSmartContractTrade(ctx *node.Context) error {
	req := &pb.CreateSmartContractTradeReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CreateSmartContractTrade(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) submitSmartContractTrade(ctx *node.Context) error {
	req := &SubmitSmartContractTradeReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.SubmitSmartContractTrade(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) callSmartContractABI(ctx *node.Context) error {
	req := &pb.CallSmartContractABIReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.CallSmartContractABI(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) findSmartContractReceipt(ctx *node.Context) error {
	req := &pb.FindSmartContractReceiptReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FindSmartContractReceipt(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) followSmartContractReceipt(ctx *node.Context) error {
	req := &pb.FollowSmartContractReceiptReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.FollowSmartContractReceipt(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getNFTCollection(ctx *node.Context) error {
	req := &pb.GetNFTCollectionReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetNFTCollection(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getNFTToken(ctx *node.Context) error {
	req := &pb.GetNFTTokenReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetNFTToken(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getNFTOwnerToken(ctx *node.Context) error {
	req := &pb.GetNFTOwnerTokenReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetNFTOwnerToken(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getNFTTransfer(ctx *node.Context) error {
	req := &pb.GetNFTTransferReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetNFTTransfer(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}

func (s *SubscribeNode) getTokenURI(ctx *node.Context) error {
	req := &pb.GetTokenURIReq{}
	if err := ctx.JsonBody.ParseData(req); err != nil {
		return err
	}
	res, err := s.api.GetTokenURI(req)
	if err != nil {
		return err
	}
	return s.Json(ctx, res)
}
