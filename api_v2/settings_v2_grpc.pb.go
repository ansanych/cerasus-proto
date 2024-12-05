// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: settings_v2.proto

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

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetUserAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*UserAppData, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetUserAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*UserAppData, error) {
	out := new(UserAppData)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetUserAppData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
// All implementations must embed UnimplementedSettingsServer
// for forward compatibility
type SettingsServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetUserAppData(context.Context, *Auth) (*UserAppData, error)
	mustEmbedUnimplementedSettingsServer()
}

// UnimplementedSettingsServer must be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (UnimplementedSettingsServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSettingsServer) GetUserAppData(context.Context, *Auth) (*UserAppData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAppData not implemented")
}
func (UnimplementedSettingsServer) mustEmbedUnimplementedSettingsServer() {}

// UnsafeSettingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServer will
// result in compilation errors.
type UnsafeSettingsServer interface {
	mustEmbedUnimplementedSettingsServer()
}

func RegisterSettingsServer(s grpc.ServiceRegistrar, srv SettingsServer) {
	s.RegisterService(&Settings_ServiceDesc, srv)
}

func _Settings_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetUserAppData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetUserAppData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetUserAppData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetUserAppData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

// Settings_ServiceDesc is the grpc.ServiceDesc for Settings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Settings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Settings_Ping_Handler,
		},
		{
			MethodName: "GetUserAppData",
			Handler:    _Settings_GetUserAppData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings_v2.proto",
}