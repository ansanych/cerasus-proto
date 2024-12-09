// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: wb_v2.proto

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

// WBClient is the client API for WB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WBClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error)
	GetShopWidget(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWidget, error)
	GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error)
	GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetFlowGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetMarginGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
}

type wBClient struct {
	cc grpc.ClientConnInterface
}

func NewWBClient(cc grpc.ClientConnInterface) WBClient {
	return &wBClient{cc}
}

func (c *wBClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error) {
	out := new(AppShopData)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/GetAppData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetShopWidget(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWidget, error) {
	out := new(ShopWidget)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/GetShopWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error) {
	out := new(LineGraph)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/GetProductsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetFlowGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/GetFlowGraphicData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetMarginGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.WB/GetMarginGraphicData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WBServer is the server API for WB service.
// All implementations must embed UnimplementedWBServer
// for forward compatibility
type WBServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetAppData(context.Context, *Auth) (*AppShopData, error)
	GetShopWidget(context.Context, *Auth) (*ShopWidget, error)
	GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error)
	GetProductsCount(context.Context, *Auth) (*Count, error)
	GetFlowGraphicData(context.Context, *Auth) (*Count, error)
	GetMarginGraphicData(context.Context, *Auth) (*Count, error)
	mustEmbedUnimplementedWBServer()
}

// UnimplementedWBServer must be embedded to have forward compatible implementations.
type UnimplementedWBServer struct {
}

func (UnimplementedWBServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedWBServer) GetAppData(context.Context, *Auth) (*AppShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppData not implemented")
}
func (UnimplementedWBServer) GetShopWidget(context.Context, *Auth) (*ShopWidget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopWidget not implemented")
}
func (UnimplementedWBServer) GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedWBServer) GetProductsCount(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCount not implemented")
}
func (UnimplementedWBServer) GetFlowGraphicData(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlowGraphicData not implemented")
}
func (UnimplementedWBServer) GetMarginGraphicData(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarginGraphicData not implemented")
}
func (UnimplementedWBServer) mustEmbedUnimplementedWBServer() {}

// UnsafeWBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WBServer will
// result in compilation errors.
type UnsafeWBServer interface {
	mustEmbedUnimplementedWBServer()
}

func RegisterWBServer(s grpc.ServiceRegistrar, srv WBServer) {
	s.RegisterService(&WB_ServiceDesc, srv)
}

func _WB_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetAppData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetAppData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/GetAppData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetAppData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetShopWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetShopWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/GetShopWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetShopWidget(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetMainGraphic(ctx, req.(*LineGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProductsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProductsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/GetProductsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProductsCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetFlowGraphicData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetFlowGraphicData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/GetFlowGraphicData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetFlowGraphicData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetMarginGraphicData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetMarginGraphicData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.WB/GetMarginGraphicData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetMarginGraphicData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

// WB_ServiceDesc is the grpc.ServiceDesc for WB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.WB",
	HandlerType: (*WBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _WB_Ping_Handler,
		},
		{
			MethodName: "GetAppData",
			Handler:    _WB_GetAppData_Handler,
		},
		{
			MethodName: "GetShopWidget",
			Handler:    _WB_GetShopWidget_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _WB_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetProductsCount",
			Handler:    _WB_GetProductsCount_Handler,
		},
		{
			MethodName: "GetFlowGraphicData",
			Handler:    _WB_GetFlowGraphicData_Handler,
		},
		{
			MethodName: "GetMarginGraphicData",
			Handler:    _WB_GetMarginGraphicData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb_v2.proto",
}
