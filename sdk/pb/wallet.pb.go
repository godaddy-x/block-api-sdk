// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: wallet.proto

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

type CreateWalletReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID    string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	Alias    string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	WalletID string `protobuf:"bytes,3,opt,name=walletID,proto3" json:"walletID,omitempty"`
	RootPath string `protobuf:"bytes,4,opt,name=rootPath,proto3" json:"rootPath,omitempty"`
}

func (x *CreateWalletReq) Reset() {
	*x = CreateWalletReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletReq) ProtoMessage() {}

func (x *CreateWalletReq) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletReq.ProtoReflect.Descriptor instead.
func (*CreateWalletReq) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWalletReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateWalletReq) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CreateWalletReq) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *CreateWalletReq) GetRootPath() string {
	if x != nil {
		return x.RootPath
	}
	return ""
}

type CreateWalletRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result       bool   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	WalletID     string `protobuf:"bytes,2,opt,name=walletID,proto3" json:"walletID,omitempty"`
	RootPath     string `protobuf:"bytes,3,opt,name=rootPath,proto3" json:"rootPath,omitempty"`
	Alias        string `protobuf:"bytes,4,opt,name=alias,proto3" json:"alias,omitempty"`
	AccountIndex int64  `protobuf:"varint,5,opt,name=accountIndex,proto3" json:"accountIndex,omitempty"`
}

func (x *CreateWalletRes) Reset() {
	*x = CreateWalletRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRes) ProtoMessage() {}

func (x *CreateWalletRes) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRes.ProtoReflect.Descriptor instead.
func (*CreateWalletRes) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWalletRes) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *CreateWalletRes) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *CreateWalletRes) GetRootPath() string {
	if x != nil {
		return x.RootPath
	}
	return ""
}

func (x *CreateWalletRes) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *CreateWalletRes) GetAccountIndex() int64 {
	if x != nil {
		return x.AccountIndex
	}
	return 0
}

type FindWalletByWalletIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID    string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	WalletID string `protobuf:"bytes,2,opt,name=walletID,proto3" json:"walletID,omitempty"`
}

func (x *FindWalletByWalletIDReq) Reset() {
	*x = FindWalletByWalletIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWalletByWalletIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWalletByWalletIDReq) ProtoMessage() {}

func (x *FindWalletByWalletIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWalletByWalletIDReq.ProtoReflect.Descriptor instead.
func (*FindWalletByWalletIDReq) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *FindWalletByWalletIDReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *FindWalletByWalletIDReq) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

type FindWalletByWalletIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *WalletResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FindWalletByWalletIDRes) Reset() {
	*x = FindWalletByWalletIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWalletByWalletIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWalletByWalletIDRes) ProtoMessage() {}

func (x *FindWalletByWalletIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWalletByWalletIDRes.ProtoReflect.Descriptor instead.
func (*FindWalletByWalletIDRes) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *FindWalletByWalletIDRes) GetResult() *WalletResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type FindWalletByParamsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID    string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	WalletID string `protobuf:"bytes,2,opt,name=walletID,proto3" json:"walletID,omitempty"`
	LastID   int64  `protobuf:"varint,3,opt,name=lastID,proto3" json:"lastID,omitempty"`
	Limit    int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	PrevID   int64  `protobuf:"varint,5,opt,name=prevID,proto3" json:"prevID,omitempty"`
}

func (x *FindWalletByParamsReq) Reset() {
	*x = FindWalletByParamsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWalletByParamsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWalletByParamsReq) ProtoMessage() {}

func (x *FindWalletByParamsReq) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWalletByParamsReq.ProtoReflect.Descriptor instead.
func (*FindWalletByParamsReq) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *FindWalletByParamsReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *FindWalletByParamsReq) GetWalletID() string {
	if x != nil {
		return x.WalletID
	}
	return ""
}

func (x *FindWalletByParamsReq) GetLastID() int64 {
	if x != nil {
		return x.LastID
	}
	return 0
}

func (x *FindWalletByParamsReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindWalletByParamsReq) GetPrevID() int64 {
	if x != nil {
		return x.PrevID
	}
	return 0
}

type FindWalletByParamsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*WalletResult `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *FindWalletByParamsRes) Reset() {
	*x = FindWalletByParamsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindWalletByParamsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindWalletByParamsRes) ProtoMessage() {}

func (x *FindWalletByParamsRes) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindWalletByParamsRes.ProtoReflect.Descriptor instead.
func (*FindWalletByParamsRes) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *FindWalletByParamsRes) GetResult() []*WalletResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x9b, 0x01,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x4b, 0x0a, 0x17, 0x46,
	0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6c, 0x61, 0x73, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x76, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x76, 0x49, 0x44, 0x22, 0x46, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x8b, 0x02, 0x0a,
	0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x42, 0x79, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x42, 0x79, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x42, 0x79, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x42, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wallet_proto_goTypes = []interface{}{
	(*CreateWalletReq)(nil),         // 0: service.CreateWalletReq
	(*CreateWalletRes)(nil),         // 1: service.CreateWalletRes
	(*FindWalletByWalletIDReq)(nil), // 2: service.FindWalletByWalletIDReq
	(*FindWalletByWalletIDRes)(nil), // 3: service.FindWalletByWalletIDRes
	(*FindWalletByParamsReq)(nil),   // 4: service.FindWalletByParamsReq
	(*FindWalletByParamsRes)(nil),   // 5: service.FindWalletByParamsRes
	(*WalletResult)(nil),            // 6: service.WalletResult
}
var file_wallet_proto_depIdxs = []int32{
	6, // 0: service.FindWalletByWalletIDRes.result:type_name -> service.WalletResult
	6, // 1: service.FindWalletByParamsRes.result:type_name -> service.WalletResult
	0, // 2: service.WalletService.CreateWallet:input_type -> service.CreateWalletReq
	2, // 3: service.WalletService.FindWalletByWalletID:input_type -> service.FindWalletByWalletIDReq
	4, // 4: service.WalletService.FindWalletByParams:input_type -> service.FindWalletByParamsReq
	1, // 5: service.WalletService.CreateWallet:output_type -> service.CreateWalletRes
	3, // 6: service.WalletService.FindWalletByWalletID:output_type -> service.FindWalletByWalletIDRes
	5, // 7: service.WalletService.FindWalletByParams:output_type -> service.FindWalletByParamsRes
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	file_account_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletReq); i {
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
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRes); i {
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
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWalletByWalletIDReq); i {
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
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWalletByWalletIDRes); i {
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
		file_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWalletByParamsReq); i {
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
		file_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindWalletByParamsRes); i {
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
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
