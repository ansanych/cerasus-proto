// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: availabler.proto

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

// AvailablerClient is the client API for Availabler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvailablerClient interface {
	SetSystemSettings(ctx context.Context, in *AvailablerSystemSettings, opts ...grpc.CallOption) (*BoolReply, error)
	GetSystemSettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AvailablerSystemSettings, error)
	SetCompanySettings(ctx context.Context, in *AvailablerCompanySettings, opts ...grpc.CallOption) (*BoolReply, error)
	GetCompanySettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AvailablerCompanySettingsReply, error)
	SetCompanyParams(ctx context.Context, in *SetAvailablerParams, opts ...grpc.CallOption) (*BoolReply, error)
	GetCompanyParams(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AvailablerParams, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type availablerClient struct {
	cc grpc.ClientConnInterface
}

func NewAvailablerClient(cc grpc.ClientConnInterface) AvailablerClient {
	return &availablerClient{cc}
}

func (c *availablerClient) SetSystemSettings(ctx context.Context, in *AvailablerSystemSettings, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/SetSystemSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availablerClient) GetSystemSettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AvailablerSystemSettings, error) {
	out := new(AvailablerSystemSettings)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/GetSystemSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availablerClient) SetCompanySettings(ctx context.Context, in *AvailablerCompanySettings, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/SetCompanySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availablerClient) GetCompanySettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AvailablerCompanySettingsReply, error) {
	out := new(AvailablerCompanySettingsReply)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/GetCompanySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availablerClient) SetCompanyParams(ctx context.Context, in *SetAvailablerParams, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/SetCompanyParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availablerClient) GetCompanyParams(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AvailablerParams, error) {
	out := new(AvailablerParams)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/GetCompanyParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *availablerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Availabler/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AvailablerServer is the server API for Availabler service.
// All implementations must embed UnimplementedAvailablerServer
// for forward compatibility
type AvailablerServer interface {
	SetSystemSettings(context.Context, *AvailablerSystemSettings) (*BoolReply, error)
	GetSystemSettings(context.Context, *Auth) (*AvailablerSystemSettings, error)
	SetCompanySettings(context.Context, *AvailablerCompanySettings) (*BoolReply, error)
	GetCompanySettings(context.Context, *Auth) (*AvailablerCompanySettingsReply, error)
	SetCompanyParams(context.Context, *SetAvailablerParams) (*BoolReply, error)
	GetCompanyParams(context.Context, *Auth) (*AvailablerParams, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	mustEmbedUnimplementedAvailablerServer()
}

// UnimplementedAvailablerServer must be embedded to have forward compatible implementations.
type UnimplementedAvailablerServer struct {
}

func (UnimplementedAvailablerServer) SetSystemSettings(context.Context, *AvailablerSystemSettings) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSystemSettings not implemented")
}
func (UnimplementedAvailablerServer) GetSystemSettings(context.Context, *Auth) (*AvailablerSystemSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemSettings not implemented")
}
func (UnimplementedAvailablerServer) SetCompanySettings(context.Context, *AvailablerCompanySettings) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCompanySettings not implemented")
}
func (UnimplementedAvailablerServer) GetCompanySettings(context.Context, *Auth) (*AvailablerCompanySettingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanySettings not implemented")
}
func (UnimplementedAvailablerServer) SetCompanyParams(context.Context, *SetAvailablerParams) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCompanyParams not implemented")
}
func (UnimplementedAvailablerServer) GetCompanyParams(context.Context, *Auth) (*AvailablerParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyParams not implemented")
}
func (UnimplementedAvailablerServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAvailablerServer) mustEmbedUnimplementedAvailablerServer() {}

// UnsafeAvailablerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvailablerServer will
// result in compilation errors.
type UnsafeAvailablerServer interface {
	mustEmbedUnimplementedAvailablerServer()
}

func RegisterAvailablerServer(s grpc.ServiceRegistrar, srv AvailablerServer) {
	s.RegisterService(&Availabler_ServiceDesc, srv)
}

func _Availabler_SetSystemSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailablerSystemSettings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).SetSystemSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/SetSystemSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).SetSystemSettings(ctx, req.(*AvailablerSystemSettings))
	}
	return interceptor(ctx, in, info, handler)
}

func _Availabler_GetSystemSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).GetSystemSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/GetSystemSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).GetSystemSettings(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Availabler_SetCompanySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AvailablerCompanySettings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).SetCompanySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/SetCompanySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).SetCompanySettings(ctx, req.(*AvailablerCompanySettings))
	}
	return interceptor(ctx, in, info, handler)
}

func _Availabler_GetCompanySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).GetCompanySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/GetCompanySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).GetCompanySettings(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Availabler_SetCompanyParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAvailablerParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).SetCompanyParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/SetCompanyParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).SetCompanyParams(ctx, req.(*SetAvailablerParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Availabler_GetCompanyParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).GetCompanyParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/GetCompanyParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).GetCompanyParams(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Availabler_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AvailablerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Availabler/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AvailablerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Availabler_ServiceDesc is the grpc.ServiceDesc for Availabler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Availabler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Availabler",
	HandlerType: (*AvailablerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSystemSettings",
			Handler:    _Availabler_SetSystemSettings_Handler,
		},
		{
			MethodName: "GetSystemSettings",
			Handler:    _Availabler_GetSystemSettings_Handler,
		},
		{
			MethodName: "SetCompanySettings",
			Handler:    _Availabler_SetCompanySettings_Handler,
		},
		{
			MethodName: "GetCompanySettings",
			Handler:    _Availabler_GetCompanySettings_Handler,
		},
		{
			MethodName: "SetCompanyParams",
			Handler:    _Availabler_SetCompanyParams_Handler,
		},
		{
			MethodName: "GetCompanyParams",
			Handler:    _Availabler_GetCompanyParams_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Availabler_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "availabler.proto",
}
