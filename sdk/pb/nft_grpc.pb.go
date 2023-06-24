// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: nft.proto

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

// NFTServiceClient is the client API for NFTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NFTServiceClient interface {
	// 获取NFT集合
	GetNFTCollection(ctx context.Context, in *GetNFTCollectionReq, opts ...grpc.CallOption) (*GetNFTCollectionRes, error)
	// 获取NFT集合的TokenID
	GetNFTToken(ctx context.Context, in *GetNFTTokenReq, opts ...grpc.CallOption) (*GetNFTTokenRes, error)
	// 通过owner获取NFT集合的TokenID
	GetNFTOwnerToken(ctx context.Context, in *GetNFTOwnerTokenReq, opts ...grpc.CallOption) (*GetNFTOwnerTokenRes, error)
	// 获取NFT交易日志
	GetNFTTransfer(ctx context.Context, in *GetNFTTransferReq, opts ...grpc.CallOption) (*GetNFTTransferRes, error)
	// 获取tokenURI
	GetTokenURI(ctx context.Context, in *GetTokenURIReq, opts ...grpc.CallOption) (*GetTokenURIRes, error)
	// 获取tokenAmount
	GetTokenAmount(ctx context.Context, in *GetTokenAmountReq, opts ...grpc.CallOption) (*GetTokenAmountRes, error)
}

type nFTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNFTServiceClient(cc grpc.ClientConnInterface) NFTServiceClient {
	return &nFTServiceClient{cc}
}

func (c *nFTServiceClient) GetNFTCollection(ctx context.Context, in *GetNFTCollectionReq, opts ...grpc.CallOption) (*GetNFTCollectionRes, error) {
	out := new(GetNFTCollectionRes)
	err := c.cc.Invoke(ctx, "/service.NFTService/GetNFTCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nFTServiceClient) GetNFTToken(ctx context.Context, in *GetNFTTokenReq, opts ...grpc.CallOption) (*GetNFTTokenRes, error) {
	out := new(GetNFTTokenRes)
	err := c.cc.Invoke(ctx, "/service.NFTService/GetNFTToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nFTServiceClient) GetNFTOwnerToken(ctx context.Context, in *GetNFTOwnerTokenReq, opts ...grpc.CallOption) (*GetNFTOwnerTokenRes, error) {
	out := new(GetNFTOwnerTokenRes)
	err := c.cc.Invoke(ctx, "/service.NFTService/GetNFTOwnerToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nFTServiceClient) GetNFTTransfer(ctx context.Context, in *GetNFTTransferReq, opts ...grpc.CallOption) (*GetNFTTransferRes, error) {
	out := new(GetNFTTransferRes)
	err := c.cc.Invoke(ctx, "/service.NFTService/GetNFTTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nFTServiceClient) GetTokenURI(ctx context.Context, in *GetTokenURIReq, opts ...grpc.CallOption) (*GetTokenURIRes, error) {
	out := new(GetTokenURIRes)
	err := c.cc.Invoke(ctx, "/service.NFTService/GetTokenURI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nFTServiceClient) GetTokenAmount(ctx context.Context, in *GetTokenAmountReq, opts ...grpc.CallOption) (*GetTokenAmountRes, error) {
	out := new(GetTokenAmountRes)
	err := c.cc.Invoke(ctx, "/service.NFTService/GetTokenAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NFTServiceServer is the server API for NFTService service.
// All implementations must embed UnimplementedNFTServiceServer
// for forward compatibility
type NFTServiceServer interface {
	// 获取NFT集合
	GetNFTCollection(context.Context, *GetNFTCollectionReq) (*GetNFTCollectionRes, error)
	// 获取NFT集合的TokenID
	GetNFTToken(context.Context, *GetNFTTokenReq) (*GetNFTTokenRes, error)
	// 通过owner获取NFT集合的TokenID
	GetNFTOwnerToken(context.Context, *GetNFTOwnerTokenReq) (*GetNFTOwnerTokenRes, error)
	// 获取NFT交易日志
	GetNFTTransfer(context.Context, *GetNFTTransferReq) (*GetNFTTransferRes, error)
	// 获取tokenURI
	GetTokenURI(context.Context, *GetTokenURIReq) (*GetTokenURIRes, error)
	// 获取tokenAmount
	GetTokenAmount(context.Context, *GetTokenAmountReq) (*GetTokenAmountRes, error)
	mustEmbedUnimplementedNFTServiceServer()
}

// UnimplementedNFTServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNFTServiceServer struct {
}

func (UnimplementedNFTServiceServer) GetNFTCollection(context.Context, *GetNFTCollectionReq) (*GetNFTCollectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNFTCollection not implemented")
}
func (UnimplementedNFTServiceServer) GetNFTToken(context.Context, *GetNFTTokenReq) (*GetNFTTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNFTToken not implemented")
}
func (UnimplementedNFTServiceServer) GetNFTOwnerToken(context.Context, *GetNFTOwnerTokenReq) (*GetNFTOwnerTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNFTOwnerToken not implemented")
}
func (UnimplementedNFTServiceServer) GetNFTTransfer(context.Context, *GetNFTTransferReq) (*GetNFTTransferRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNFTTransfer not implemented")
}
func (UnimplementedNFTServiceServer) GetTokenURI(context.Context, *GetTokenURIReq) (*GetTokenURIRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenURI not implemented")
}
func (UnimplementedNFTServiceServer) GetTokenAmount(context.Context, *GetTokenAmountReq) (*GetTokenAmountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenAmount not implemented")
}
func (UnimplementedNFTServiceServer) mustEmbedUnimplementedNFTServiceServer() {}

// UnsafeNFTServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NFTServiceServer will
// result in compilation errors.
type UnsafeNFTServiceServer interface {
	mustEmbedUnimplementedNFTServiceServer()
}

func RegisterNFTServiceServer(s grpc.ServiceRegistrar, srv NFTServiceServer) {
	s.RegisterService(&NFTService_ServiceDesc, srv)
}

func _NFTService_GetNFTCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNFTCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetNFTCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NFTService/GetNFTCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetNFTCollection(ctx, req.(*GetNFTCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NFTService_GetNFTToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNFTTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetNFTToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NFTService/GetNFTToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetNFTToken(ctx, req.(*GetNFTTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NFTService_GetNFTOwnerToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNFTOwnerTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetNFTOwnerToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NFTService/GetNFTOwnerToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetNFTOwnerToken(ctx, req.(*GetNFTOwnerTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NFTService_GetNFTTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNFTTransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetNFTTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NFTService/GetNFTTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetNFTTransfer(ctx, req.(*GetNFTTransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NFTService_GetTokenURI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenURIReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetTokenURI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NFTService/GetTokenURI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetTokenURI(ctx, req.(*GetTokenURIReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NFTService_GetTokenAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenAmountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NFTServiceServer).GetTokenAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NFTService/GetTokenAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NFTServiceServer).GetTokenAmount(ctx, req.(*GetTokenAmountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NFTService_ServiceDesc is the grpc.ServiceDesc for NFTService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NFTService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.NFTService",
	HandlerType: (*NFTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNFTCollection",
			Handler:    _NFTService_GetNFTCollection_Handler,
		},
		{
			MethodName: "GetNFTToken",
			Handler:    _NFTService_GetNFTToken_Handler,
		},
		{
			MethodName: "GetNFTOwnerToken",
			Handler:    _NFTService_GetNFTOwnerToken_Handler,
		},
		{
			MethodName: "GetNFTTransfer",
			Handler:    _NFTService_GetNFTTransfer_Handler,
		},
		{
			MethodName: "GetTokenURI",
			Handler:    _NFTService_GetTokenURI_Handler,
		},
		{
			MethodName: "GetTokenAmount",
			Handler:    _NFTService_GetTokenAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nft.proto",
}
