// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth.proto

package cerasus_proto

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

// AuthentyClient is the client API for Authenty service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthentyClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*StatusReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*LoginReply, error)
	CheckAccess(ctx context.Context, in *AccessRequest, opts ...grpc.CallOption) (*Auth, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type authentyClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthentyClient(cc grpc.ClientConnInterface) AuthentyClient {
	return &authentyClient{cc}
}

func (c *authentyClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasus.Authenty/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentyClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/cerasus.Authenty/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentyClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/cerasus.Authenty/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentyClient) CheckAccess(ctx context.Context, in *AccessRequest, opts ...grpc.CallOption) (*Auth, error) {
	out := new(Auth)
	err := c.cc.Invoke(ctx, "/cerasus.Authenty/CheckAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentyClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Authenty/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthentyServer is the server API for Authenty service.
// All implementations must embed UnimplementedAuthentyServer
// for forward compatibility
type AuthentyServer interface {
	Register(context.Context, *RegisterRequest) (*StatusReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Refresh(context.Context, *RefreshRequest) (*LoginReply, error)
	CheckAccess(context.Context, *AccessRequest) (*Auth, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	mustEmbedUnimplementedAuthentyServer()
}

// UnimplementedAuthentyServer must be embedded to have forward compatible implementations.
type UnimplementedAuthentyServer struct {
}

func (UnimplementedAuthentyServer) Register(context.Context, *RegisterRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthentyServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthentyServer) Refresh(context.Context, *RefreshRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthentyServer) CheckAccess(context.Context, *AccessRequest) (*Auth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccess not implemented")
}
func (UnimplementedAuthentyServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAuthentyServer) mustEmbedUnimplementedAuthentyServer() {}

// UnsafeAuthentyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthentyServer will
// result in compilation errors.
type UnsafeAuthentyServer interface {
	mustEmbedUnimplementedAuthentyServer()
}

func RegisterAuthentyServer(s grpc.ServiceRegistrar, srv AuthentyServer) {
	s.RegisterService(&Authenty_ServiceDesc, srv)
}

func _Authenty_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentyServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Authenty/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentyServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenty_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentyServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Authenty/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentyServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenty_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentyServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Authenty/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentyServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenty_CheckAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentyServer).CheckAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Authenty/CheckAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentyServer).CheckAccess(ctx, req.(*AccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenty_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentyServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Authenty/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentyServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authenty_ServiceDesc is the grpc.ServiceDesc for Authenty service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authenty_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Authenty",
	HandlerType: (*AuthentyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Authenty_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authenty_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Authenty_Refresh_Handler,
		},
		{
			MethodName: "CheckAccess",
			Handler:    _Authenty_CheckAccess_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Authenty_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
