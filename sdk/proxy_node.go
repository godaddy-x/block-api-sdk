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
	// proxy api
	//CreateWallet(ctx *node.Context) error
	//FindWalletByWalletID(ctx *node.Context) error
	//FindWalletByParams(ctx *node.Context) error
	//CreateAccount(ctx *node.Context) error
	//FindAccountByWalletID(ctx *node.Context) error
	//FindAccountByAccountID(ctx *node.Context) error
	//GetBalanceByAccount(ctx *node.Context) error
	//GetAccountBalanceList(ctx *node.Context) error
	//CreateAddress(ctx *node.Context) error
	//ImportAddress(ctx *node.Context) error
	//FindAddressByAddress(ctx *node.Context) error
	//FindAddressByAccountID(ctx *node.Context) error
	//VerifyAddress(ctx *node.Context) error
	//GetBalanceByAddress(ctx *node.Context) error
	//GetAddressBalanceList(ctx *node.Context) error
	//GetSymbolBlockList(ctx *node.Context) error
	//GetBlockStatus(ctx *node.Context) error
	//Subscribe(ctx *node.Context) error
	//FindTradeLog(ctx *node.Context) error
	//CreateTrade(ctx *node.Context) error
	//SubmitTrade(ctx *node.Context) error
	//CreateSummaryTx(ctx *node.Context) error
	//GetContracts(ctx *node.Context) error
	//CreateSmartContractTrade(ctx *node.Context) error
	//SubmitSmartContractTrade(ctx *node.Context) error
	//CallSmartContractABI(ctx *node.Context) error
	//FindSmartContractReceipt(ctx *node.Context) error
	//FollowSmartContractReceipt(ctx *node.Context) error
	//GetNFTCollection(ctx *node.Context) error
	//GetNFTToken(ctx *node.Context) error
	//GetNFTOwnerToken(ctx *node.Context) error
	//GetNFTTransfer(ctx *node.Context) error
	//GetTokenURI(ctx *node.Context) error
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

//func (s *UnimplementedObserver) CreateWallet(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindWalletByWalletID(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindWalletByParams(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) CreateAccount(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindAccountByWalletID(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindAccountByAccountID(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetBalanceByAccount(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetAccountBalanceList(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) CreateAddress(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) ImportAddress(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindAddressByAddress(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindAddressByAccountID(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) VerifyAddress(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetBalanceByAddress(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetAddressBalanceList(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetSymbolBlockList(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetBlockStatus(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) Subscribe(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindTradeLog(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) CreateTrade(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) SubmitTrade(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) CreateSummaryTx(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetContracts(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) CreateSmartContractTrade(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) SubmitSmartContractTrade(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) CallSmartContractABI(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FindSmartContractReceipt(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) FollowSmartContractReceipt(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetNFTCollection(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetNFTToken(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetNFTOwnerToken(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetNFTTransfer(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}
//
//func (s *UnimplementedObserver) GetTokenURI(ctx *node.Context) error {
//	return errors.Errorf("unimplemented method")
//}

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
	return s.Json(ctx, &pb.CreateWalletRes{Result: true, WalletID: res.WalletID, RootPath: res.RootPath, Alias: res.Alias, AccountIndex: res.AccountIndex})
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
	return s.Json(ctx, &pb.FindWalletByWalletIDRes{Result: res})
}
