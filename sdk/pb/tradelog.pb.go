// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: tradelog.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TradeLogResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppID        string   `protobuf:"bytes,2,opt,name=appID,proto3" json:"appID,omitempty"`
	WalletID     string   `protobuf:"bytes,3,opt,name=walletID,proto3" json:"walletID,omitempty"`
	AccountID    string   `protobuf:"bytes,4,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Sid          string   `protobuf:"bytes,5,opt,name=sid,proto3" json:"sid,omitempty"`
	TxID         string   `protobuf:"bytes,6,opt,name=txID,proto3" json:"txID,omitempty"`
	TxType       int64    `protobuf:"varint,7,opt,name=txType,proto3" json:"txType,omitempty"`
	TxAction     string   `protobuf:"bytes,8,opt,name=txAction,proto3" json:"txAction,omitempty"`
	WxID         string   `protobuf:"bytes,9,opt,name=wxID,proto3" json:"wxID,omitempty"`
	FromAddress  []string `protobuf:"bytes,10,rep,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	FromAddressV []string `protobuf:"bytes,11,rep,name=fromAddressV,proto3" json:"fromAddressV,omitempty"`
	ToAddress    []string `protobuf:"bytes,12,rep,name=toAddress,proto3" json:"toAddress,omitempty"`
	ToAddressV   []string `protobuf:"bytes,13,rep,name=toAddressV,proto3" json:"toAddressV,omitempty"`
	Amount       string   `protobuf:"bytes,14,opt,name=amount,proto3" json:"amount,omitempty"`
	Fees         string   `protobuf:"bytes,15,opt,name=fees,proto3" json:"fees,omitempty"`
	Type         int64    `protobuf:"varint,16,opt,name=type,proto3" json:"type,omitempty"`
	Symbol       string   `protobuf:"bytes,17,opt,name=symbol,proto3" json:"symbol,omitempty"`
	ContractID   string   `protobuf:"bytes,18,opt,name=contractID,proto3" json:"contractID,omitempty"`
	IsContract   int64    `protobuf:"varint,19,opt,name=isContract,proto3" json:"isContract,omitempty"`
	Confirm      int64    `protobuf:"varint,20,opt,name=confirm,proto3" json:"confirm,omitempty"`
	BlockHash    string   `protobuf:"bytes,21,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockHeight  int64    `protobuf:"varint,22,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	IsMemo       int64    `protobuf:"varint,23,opt,name=isMemo,proto3" json:"isMemo,omitempty"`
	IsMain       int64    `protobuf:"varint,24,opt,name=isMain,proto3" json:"isMain,omitempty"`
	Memo         string   `protobuf:"bytes,25,opt,name=memo,proto3" json:"memo,omitempty"`
	Applytime    int64    `protobuf:"varint,26,opt,name=applytime,proto3" json:"applytime,omitempty"`
	Decimals     int64    `protobuf:"varint,27,opt,name=decimals,proto3" json:"decimals,omitempty"`
	ContractName string   `protobuf:"bytes,28,opt,name=contractName,proto3" json:"contractName,omitempty"`
	ContractAddr string   `protobuf:"bytes,29,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
	Succtime     int64    `protobuf:"varint,30,opt,name=succtime,proto3" json:"succtime,omitempty"`
	Dealstate    int64    `protobuf:"varint,31,opt,name=dealstate,proto3" json:"dealstate,omitempty"`
	Notifystate  int64    `protobuf:"varint,32,opt,name=notifystate,proto3" json:"notifystate,omitempty"`
	Success      string   `protobuf:"bytes,33,opt,name=success,proto3" json:"success,omitempty"`
	BalanceMode  int64    `protobuf:"varint,34,opt,name=balanceMode,proto3" json:"balanceMode,omitempty"`
}

func (x *TradeLogResult) Reset() {
	*x = TradeLogResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tradelog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeLogResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeLogResult) ProtoMessage() {}

func (x *TradeLogResult) ProtoReflect() protoreflect.Message {
	mi := &file_tradelog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeLogResult.ProtoReflect.Descriptor instead.
func (*TradeLogResult) Descriptor() ([]byte, []int) {
	return file_tradelog_proto_rawDescGZIP(), []int{0}
}

func (x *TradeLogResult) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TradeLogResult) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *TradeLogResult) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *TradeLogResult) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *TradeLogResult) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *TradeLogResult) GetTxID() string {
	if x != nil {
		return x.TxID
	}
	return ""
}

func (x *TradeLogResult) GetTxType() int64 {
	if x != nil {
		return x.TxType
	}
	return 0
}

func (x *TradeLogResult) GetTxAction() string {
	if x != nil {
		return x.TxAction
	}
	return ""
}

func (x *TradeLogResult) GetWxID() string {
	if x != nil {
		return x.WxID
	}
	return ""
}

func (x *TradeLogResult) GetFromAddress() []string {
	if x != nil {
		return x.FromAddress
	}
	return nil
}

func (x *TradeLogResult) GetFromAddressV() []string {
	if x != nil {
		return x.FromAddressV
	}
	return nil
}

func (x *TradeLogResult) GetToAddress() []string {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *TradeLogResult) GetToAddressV() []string {
	if x != nil {
		return x.ToAddressV
	}
	return nil
}

func (x *TradeLogResult) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TradeLogResult) GetFees() string {
	if x != nil {
		return x.Fees
	}
	return ""
}

func (x *TradeLogResult) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *TradeLogResult) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TradeLogResult) GetContractID() string {
	if x != nil {
		return x.ContractID
	}
	return ""
}

func (x *TradeLogResult) GetIsContract() int64 {
	if x != nil {
		return x.IsContract
	}
	return 0
}

func (x *TradeLogResult) GetConfirm() int64 {
	if x != nil {
		return x.Confirm
	}
	return 0
}

func (x *TradeLogResult) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *TradeLogResult) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *TradeLogResult) GetIsMemo() int64 {
	if x != nil {
		return x.IsMemo
	}
	return 0
}

func (x *TradeLogResult) GetIsMain() int64 {
	if x != nil {
		return x.IsMain
	}
	return 0
}

func (x *TradeLogResult) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TradeLogResult) GetApplytime() int64 {
	if x != nil {
		return x.Applytime
	}
	return 0
}

func (x *TradeLogResult) GetDecimals() int64 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TradeLogResult) GetContractName() string {
	if x != nil {
		return x.ContractName
	}
	return ""
}

func (x *TradeLogResult) GetContractAddr() string {
	if x != nil {
		return x.ContractAddr
	}
	return ""
}

func (x *TradeLogResult) GetSucctime() int64 {
	if x != nil {
		return x.Succtime
	}
	return 0
}

func (x *TradeLogResult) GetDealstate() int64 {
	if x != nil {
		return x.Dealstate
	}
	return 0
}

func (x *TradeLogResult) GetNotifystate() int64 {
	if x != nil {
		return x.Notifystate
	}
	return 0
}

func (x *TradeLogResult) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

func (x *TradeLogResult) GetBalanceMode() int64 {
	if x != nil {
		return x.BalanceMode
	}
	return 0
}

type FindTradeLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID       string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	WalletID    string `protobuf:"bytes,2,opt,name=walletID,proto3" json:"walletID,omitempty"`
	AccountID   string `protobuf:"bytes,3,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TxID        string `protobuf:"bytes,4,opt,name=txID,proto3" json:"txID,omitempty"`
	WxID        string `protobuf:"bytes,5,opt,name=wxID,proto3" json:"wxID,omitempty"`
	Sid         string `protobuf:"bytes,6,opt,name=sid,proto3" json:"sid,omitempty"`
	Symbol      string `protobuf:"bytes,7,opt,name=symbol,proto3" json:"symbol,omitempty"`
	BlockHeight int64  `protobuf:"varint,8,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Type        int64  `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	IsMain      int64  `protobuf:"varint,10,opt,name=isMain,proto3" json:"isMain,omitempty"`
	IsContract  int64  `protobuf:"varint,11,opt,name=isContract,proto3" json:"isContract,omitempty"`
	ContractID  string `protobuf:"bytes,12,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Address     string `protobuf:"bytes,13,opt,name=address,proto3" json:"address,omitempty"`
	Limit       int64  `protobuf:"varint,14,opt,name=limit,proto3" json:"limit,omitempty"`
	LastID      int64  `protobuf:"varint,15,opt,name=lastID,proto3" json:"lastID,omitempty"`
	PrevID      int64  `protobuf:"varint,16,opt,name=prevID,proto3" json:"prevID,omitempty"`
}

func (x *FindTradeLogReq) Reset() {
	*x = FindTradeLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tradelog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTradeLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTradeLogReq) ProtoMessage() {}

func (x *FindTradeLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_tradelog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTradeLogReq.ProtoReflect.Descriptor instead.
func (*FindTradeLogReq) Descriptor() ([]byte, []int) {
	return file_tradelog_proto_rawDescGZIP(), []int{1}
}

func (x *FindTradeLogReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *FindTradeLogReq) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *FindTradeLogReq) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *FindTradeLogReq) GetTxID() string {
	if x != nil {
		return x.TxID
	}
	return ""
}

func (x *FindTradeLogReq) GetWxID() string {
	if x != nil {
		return x.WxID
	}
	return ""
}

func (x *FindTradeLogReq) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *FindTradeLogReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *FindTradeLogReq) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *FindTradeLogReq) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FindTradeLogReq) GetIsMain() int64 {
	if x != nil {
		return x.IsMain
	}
	return 0
}

func (x *FindTradeLogReq) GetIsContract() int64 {
	if x != nil {
		return x.IsContract
	}
	return 0
}

func (x *FindTradeLogReq) GetContractID() string {
	if x != nil {
		return x.ContractID
	}
	return ""
}

func (x *FindTradeLogReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *FindTradeLogReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindTradeLogReq) GetLastID() int64 {
	if x != nil {
		return x.LastID
	}
	return 0
}

func (x *FindTradeLogReq) GetPrevID() int64 {
	if x != nil {
		return x.PrevID
	}
	return 0
}

type FindTradeLogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastID int64             `protobuf:"varint,1,opt,name=lastID,proto3" json:"lastID,omitempty"`
	Result []*TradeLogResult `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	PrevID int64             `protobuf:"varint,3,opt,name=prevID,proto3" json:"prevID,omitempty"`
}

func (x *FindTradeLogRes) Reset() {
	*x = FindTradeLogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tradelog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTradeLogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTradeLogRes) ProtoMessage() {}

func (x *FindTradeLogRes) ProtoReflect() protoreflect.Message {
	mi := &file_tradelog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTradeLogRes.ProtoReflect.Descriptor instead.
func (*FindTradeLogRes) Descriptor() ([]byte, []int) {
	return file_tradelog_proto_rawDescGZIP(), []int{2}
}

func (x *FindTradeLogRes) GetLastID() int64 {
	if x != nil {
		return x.LastID
	}
	return 0
}

func (x *FindTradeLogRes) GetResult() []*TradeLogResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *FindTradeLogRes) GetPrevID() int64 {
	if x != nil {
		return x.PrevID
	}
	return 0
}

var File_tradelog_proto protoreflect.FileDescriptor

var file_tradelog_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb2, 0x07, 0x0a, 0x0e, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x78, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x78,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x78,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x78, 0x49, 0x44, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x78, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x65, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20,
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x69, 0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4d, 0x61,
	0x69, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x61, 0x69, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x63, 0x63, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x75, 0x63, 0x63, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x22, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xa1,
	0x03, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x78, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x78, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x78, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73,
	0x4d, 0x61, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x73, 0x4d, 0x61,
	0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x44, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x76, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x72, 0x65, 0x76,
	0x49, 0x44, 0x22, 0x72, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x44, 0x12, 0x2f, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x72, 0x65, 0x76, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x70, 0x72, 0x65, 0x76, 0x49, 0x44, 0x32, 0x57, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c,
	0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x46, 0x69, 0x6e,
	0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tradelog_proto_rawDescOnce sync.Once
	file_tradelog_proto_rawDescData = file_tradelog_proto_rawDesc
)

func file_tradelog_proto_rawDescGZIP() []byte {
	file_tradelog_proto_rawDescOnce.Do(func() {
		file_tradelog_proto_rawDescData = protoimpl.X.CompressGZIP(file_tradelog_proto_rawDescData)
	})
	return file_tradelog_proto_rawDescData
}

var file_tradelog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tradelog_proto_goTypes = []interface{}{
	(*TradeLogResult)(nil),  // 0: service.TradeLogResult
	(*FindTradeLogReq)(nil), // 1: service.FindTradeLogReq
	(*FindTradeLogRes)(nil), // 2: service.FindTradeLogRes
}
var file_tradelog_proto_depIdxs = []int32{
	0, // 0: service.FindTradeLogRes.result:type_name -> service.TradeLogResult
	1, // 1: service.TradeLogService.FindTradeLog:input_type -> service.FindTradeLogReq
	2, // 2: service.TradeLogService.FindTradeLog:output_type -> service.FindTradeLogRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tradelog_proto_init() }
func file_tradelog_proto_init() {
	if File_tradelog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tradelog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeLogResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tradelog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTradeLogReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tradelog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTradeLogRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tradelog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tradelog_proto_goTypes,
		DependencyIndexes: file_tradelog_proto_depIdxs,
		MessageInfos:      file_tradelog_proto_msgTypes,
	}.Build()
	File_tradelog_proto = out.File
	file_tradelog_proto_rawDesc = nil
	file_tradelog_proto_goTypes = nil
	file_tradelog_proto_depIdxs = nil
}
