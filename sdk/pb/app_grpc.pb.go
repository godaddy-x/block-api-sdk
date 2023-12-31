// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: app.proto

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

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	// App授权
	AppLogin(ctx context.Context, in *AppLoginReq, opts ...grpc.CallOption) (*AppLoginRes, error)
	// 校验APP
	ValidAppID(ctx context.Context, in *ValidAppIDReq, opts ...grpc.CallOption) (*ValidAppIDRes, error)
	// 校验设备ID
	ValidDeviceRole(ctx context.Context, in *ValidDeviceRoleReq, opts ...grpc.CallOption) (*ValidDeviceRoleRes, error)
	// 绑定设备ID
	AppBindDeviceID(ctx context.Context, in *AppBindDeviceIDReq, opts ...grpc.CallOption) (*AppBindDeviceIDRes, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) AppLogin(ctx context.Context, in *AppLoginReq, opts ...grpc.CallOption) (*AppLoginRes, error) {
	out := new(AppLoginRes)
	err := c.cc.Invoke(ctx, "/service.AppService/AppLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) ValidAppID(ctx context.Context, in *ValidAppIDReq, opts ...grpc.CallOption) (*ValidAppIDRes, error) {
	out := new(ValidAppIDRes)
	err := c.cc.Invoke(ctx, "/service.AppService/ValidAppID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) ValidDeviceRole(ctx context.Context, in *ValidDeviceRoleReq, opts ...grpc.CallOption) (*ValidDeviceRoleRes, error) {
	out := new(ValidDeviceRoleRes)
	err := c.cc.Invoke(ctx, "/service.AppService/ValidDeviceRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) AppBindDeviceID(ctx context.Context, in *AppBindDeviceIDReq, opts ...grpc.CallOption) (*AppBindDeviceIDRes, error) {
	out := new(AppBindDeviceIDRes)
	err := c.cc.Invoke(ctx, "/service.AppService/AppBindDeviceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	// App授权
	AppLogin(context.Context, *AppLoginReq) (*AppLoginRes, error)
	// 校验APP
	ValidAppID(context.Context, *ValidAppIDReq) (*ValidAppIDRes, error)
	// 校验设备ID
	ValidDeviceRole(context.Context, *ValidDeviceRoleReq) (*ValidDeviceRoleRes, error)
	// 绑定设备ID
	AppBindDeviceID(context.Context, *AppBindDeviceIDReq) (*AppBindDeviceIDRes, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) AppLogin(context.Context, *AppLoginReq) (*AppLoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppLogin not implemented")
}
func (UnimplementedAppServiceServer) ValidAppID(context.Context, *ValidAppIDReq) (*ValidAppIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidAppID not implemented")
}
func (UnimplementedAppServiceServer) ValidDeviceRole(context.Context, *ValidDeviceRoleReq) (*ValidDeviceRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidDeviceRole not implemented")
}
func (UnimplementedAppServiceServer) AppBindDeviceID(context.Context, *AppBindDeviceIDReq) (*AppBindDeviceIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppBindDeviceID not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_AppLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).AppLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AppService/AppLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).AppLogin(ctx, req.(*AppLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_ValidAppID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidAppIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).ValidAppID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AppService/ValidAppID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).ValidAppID(ctx, req.(*ValidAppIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_ValidDeviceRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidDeviceRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).ValidDeviceRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AppService/ValidDeviceRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).ValidDeviceRole(ctx, req.(*ValidDeviceRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_AppBindDeviceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppBindDeviceIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).AppBindDeviceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.AppService/AppBindDeviceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).AppBindDeviceID(ctx, req.(*AppBindDeviceIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppLogin",
			Handler:    _AppService_AppLogin_Handler,
		},
		{
			MethodName: "ValidAppID",
			Handler:    _AppService_ValidAppID_Handler,
		},
		{
			MethodName: "ValidDeviceRole",
			Handler:    _AppService_ValidDeviceRole_Handler,
		},
		{
			MethodName: "AppBindDeviceID",
			Handler:    _AppService_AppBindDeviceID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
