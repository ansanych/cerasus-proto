// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: wb-api_v2.proto

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

// WB_APIClient is the client API for WB_API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WB_APIClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetWarehouses(ctx context.Context, in *WBAuthParams, opts ...grpc.CallOption) (*Warehouses, error)
}

type wB_APIClient struct {
	cc grpc.ClientConnInterface
}

func NewWB_APIClient(cc grpc.ClientConnInterface) WB_APIClient {
	return &wB_APIClient{cc}
}

func (c *wB_APIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB_API/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wB_APIClient) GetWarehouses(ctx context.Context, in *WBAuthParams, opts ...grpc.CallOption) (*Warehouses, error) {
	out := new(Warehouses)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB_API/GetWarehouses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WB_APIServer is the server API for WB_API service.
// All implementations must embed UnimplementedWB_APIServer
// for forward compatibility
type WB_APIServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetWarehouses(context.Context, *WBAuthParams) (*Warehouses, error)
	mustEmbedUnimplementedWB_APIServer()
}

// UnimplementedWB_APIServer must be embedded to have forward compatible implementations.
type UnimplementedWB_APIServer struct {
}

func (UnimplementedWB_APIServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedWB_APIServer) GetWarehouses(context.Context, *WBAuthParams) (*Warehouses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouses not implemented")
}
func (UnimplementedWB_APIServer) mustEmbedUnimplementedWB_APIServer() {}

// UnsafeWB_APIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WB_APIServer will
// result in compilation errors.
type UnsafeWB_APIServer interface {
	mustEmbedUnimplementedWB_APIServer()
}

func RegisterWB_APIServer(s grpc.ServiceRegistrar, srv WB_APIServer) {
	s.RegisterService(&WB_API_ServiceDesc, srv)
}

func _WB_API_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WB_APIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB_API/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WB_APIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_API_GetWarehouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WBAuthParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WB_APIServer).GetWarehouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB_API/GetWarehouses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WB_APIServer).GetWarehouses(ctx, req.(*WBAuthParams))
	}
	return interceptor(ctx, in, info, handler)
}

// WB_API_ServiceDesc is the grpc.ServiceDesc for WB_API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WB_API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.WB_API",
	HandlerType: (*WB_APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _WB_API_Ping_Handler,
		},
		{
			MethodName: "GetWarehouses",
			Handler:    _WB_API_GetWarehouses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb-api_v2.proto",
}
