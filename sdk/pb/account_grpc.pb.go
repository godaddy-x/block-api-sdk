// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: account.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	// 创建账号
	CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error)
	// 通过accountID查询账号
	FindAccountByAccountID(ctx context.Context, in *FindAccountByAccountIDReq, opts ...grpc.CallOption) (*FindAccountByAccountIDRes, error)
	// 通过walletID查询账号
	FindAccountByWalletID(ctx context.Context, in *FindAccountByWalletIDReq, opts ...grpc.CallOption) (*FindAccountByWalletIDRes, error)
	// 获取账号的token余额
	GetBalanceByAccount(ctx context.Context, in *GetBalanceByAccountReq, opts ...grpc.CallOption) (*GetBalanceByAccountRes, error)
	// 获取账号的余额列表
	GetAccountBalanceList(ctx context.Context, in *GetAccountBalanceListReq, opts ...grpc.CallOption) (*GetAccountBalanceListRes, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error) {
	out := new(CreateAccountRes)
	err := c.cc.Invoke(ctx, "/service.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) FindAccountByAccountID(ctx context.Context, in *FindAccountByAccountIDReq, opts ...grpc.CallOption) (*FindAccountByAccountIDRes, error) {
	out := new(FindAccountByAccountIDRes)
	err := c.cc.Invoke(ctx, "/service.AccountService/FindAccountByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) FindAccountByWalletID(ctx context.Context, in *FindAccountByWalletIDReq, opts ...grpc.CallOption) (*FindAccountByWalletIDRes, error) {
	out := new(FindAccountByWalletIDRes)
	err := c.cc.Invoke(ctx, "/service.AccountService/FindAccountByWalletID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetBalanceByAccount(ctx context.Context, in *GetBalanceByAccountReq, opts ...grpc.CallOption) (*GetBalanceByAccountRes, error) {
	out := new(GetBalanceByAccountRes)
	err := c.cc.Invoke(ctx, "/service.AccountService/GetBalanceByAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountBalanceList(ctx context.Context, in *GetAccountBalanceListReq, opts ...grpc.CallOption) (*GetAccountBalanceListRes, error) {
	out := new(GetAccountBalanceListRes)
	err := c.cc.Invoke(ctx, "/service.AccountService/GetAccountBalanceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	// 创建账号
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRes, error)
	// 通过accountID查询账号
	FindAccountByAccountID(context.Context, *FindAccountByAccountIDReq) (*FindAccountByAccountIDRes, error)
	// 通过walletID查询账号
	FindAccountByWalletID(context.Context, *FindAccountByWalletIDReq) (*FindAccountByWalletIDRes, error)
	// 获取账号的token余额
	GetBalanceByAccount(context.Context, *GetBalanceByAccountReq) (*GetBalanceByAccountRes, error)
	// 获取账号的余额列表
	GetAccountBalanceList(context.Context, *GetAccountBalanceListReq) (*GetAccountBalanceListRes, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) FindAccountByAccountID(context.Context, *FindAccountByAccountIDReq) (*FindAccountByAccountIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccountByAccountID not implemented")
}
func (UnimplementedAccountServiceServer) FindAccountByWalletID(context.Context, *FindAccountByWalletIDReq) (*FindAccountByWalletIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccountByWalletID not implemented")
}
func (UnimplementedAccountServiceServer) GetBalanceByAccount(context.Context, *GetBalanceByAccountReq) (*GetBalanceByAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalanceByAccount not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountBalanceList(context.Context, *GetAccountBalanceListReq) (*GetAccountBalanceListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalanceList not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_FindAccountByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAccountByAccountIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).FindAccountByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountService/FindAccountByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).FindAccountByAccountID(ctx, req.(*FindAccountByAccountIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_FindAccountByWalletID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAccountByWalletIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).FindAccountByWalletID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountService/FindAccountByWalletID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).FindAccountByWalletID(ctx, req.(*FindAccountByWalletIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetBalanceByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceByAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetBalanceByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountService/GetBalanceByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetBalanceByAccount(ctx, req.(*GetBalanceByAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountBalanceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBalanceListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountBalanceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AccountService/GetAccountBalanceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountBalanceList(ctx, req.(*GetAccountBalanceListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "FindAccountByAccountID",
			Handler:    _AccountService_FindAccountByAccountID_Handler,
		},
		{
			MethodName: "FindAccountByWalletID",
			Handler:    _AccountService_FindAccountByWalletID_Handler,
		},
		{
			MethodName: "GetBalanceByAccount",
			Handler:    _AccountService_GetBalanceByAccount_Handler,
		},
		{
			MethodName: "GetAccountBalanceList",
			Handler:    _AccountService_GetAccountBalanceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
