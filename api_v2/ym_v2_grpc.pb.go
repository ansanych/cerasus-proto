// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ym_v2.proto

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

// YMClient is the client API for YM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YMClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error)
	GetShopWidget(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWidget, error)
	GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error)
	GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetProductsCountUnsorted(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetFlowGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetMarginGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetWeekGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*WeekGraphic, error)
	GetPayRoundGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error)
	GetCountRoundGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error)
	GetOrderLeaders(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*OrderLeaders, error)
	GetShopProducts(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*ShopProductList, error)
	GetSales(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Sales, error)
}

type yMClient struct {
	cc grpc.ClientConnInterface
}

func NewYMClient(cc grpc.ClientConnInterface) YMClient {
	return &yMClient{cc}
}

func (c *yMClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error) {
	out := new(AppShopData)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetAppData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetShopWidget(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWidget, error) {
	out := new(ShopWidget)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetShopWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error) {
	out := new(LineGraph)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetProductsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductsCountUnsorted(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetProductsCountUnsorted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetFlowGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetFlowGraphicData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetMarginGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetMarginGraphicData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetWeekGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*WeekGraphic, error) {
	out := new(WeekGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetWeekGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetPayRoundGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error) {
	out := new(RoundGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetPayRoundGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetCountRoundGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error) {
	out := new(RoundGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetCountRoundGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetOrderLeaders(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*OrderLeaders, error) {
	out := new(OrderLeaders)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetOrderLeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetShopProducts(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetShopProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetSales(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Sales, error) {
	out := new(Sales)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YMServer is the server API for YM service.
// All implementations must embed UnimplementedYMServer
// for forward compatibility
type YMServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetAppData(context.Context, *Auth) (*AppShopData, error)
	GetShopWidget(context.Context, *Auth) (*ShopWidget, error)
	GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error)
	GetProductsCount(context.Context, *Auth) (*Count, error)
	GetProductsCountUnsorted(context.Context, *Auth) (*Count, error)
	GetFlowGraphicData(context.Context, *Auth) (*Count, error)
	GetMarginGraphicData(context.Context, *Auth) (*Count, error)
	GetWeekGraphic(context.Context, *LineGraphRequest) (*WeekGraphic, error)
	GetPayRoundGraphic(context.Context, *Auth) (*RoundGraphic, error)
	GetCountRoundGraphic(context.Context, *Auth) (*RoundGraphic, error)
	GetOrderLeaders(context.Context, *Auth) (*OrderLeaders, error)
	GetShopProducts(context.Context, *RequestByIDs) (*ShopProductList, error)
	GetSales(context.Context, *RequestByDates) (*Sales, error)
	mustEmbedUnimplementedYMServer()
}

// UnimplementedYMServer must be embedded to have forward compatible implementations.
type UnimplementedYMServer struct {
}

func (UnimplementedYMServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedYMServer) GetAppData(context.Context, *Auth) (*AppShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppData not implemented")
}
func (UnimplementedYMServer) GetShopWidget(context.Context, *Auth) (*ShopWidget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopWidget not implemented")
}
func (UnimplementedYMServer) GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedYMServer) GetProductsCount(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCount not implemented")
}
func (UnimplementedYMServer) GetProductsCountUnsorted(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCountUnsorted not implemented")
}
func (UnimplementedYMServer) GetFlowGraphicData(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlowGraphicData not implemented")
}
func (UnimplementedYMServer) GetMarginGraphicData(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarginGraphicData not implemented")
}
func (UnimplementedYMServer) GetWeekGraphic(context.Context, *LineGraphRequest) (*WeekGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekGraphic not implemented")
}
func (UnimplementedYMServer) GetPayRoundGraphic(context.Context, *Auth) (*RoundGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayRoundGraphic not implemented")
}
func (UnimplementedYMServer) GetCountRoundGraphic(context.Context, *Auth) (*RoundGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountRoundGraphic not implemented")
}
func (UnimplementedYMServer) GetOrderLeaders(context.Context, *Auth) (*OrderLeaders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderLeaders not implemented")
}
func (UnimplementedYMServer) GetShopProducts(context.Context, *RequestByIDs) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProducts not implemented")
}
func (UnimplementedYMServer) GetSales(context.Context, *RequestByDates) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedYMServer) mustEmbedUnimplementedYMServer() {}

// UnsafeYMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YMServer will
// result in compilation errors.
type UnsafeYMServer interface {
	mustEmbedUnimplementedYMServer()
}

func RegisterYMServer(s grpc.ServiceRegistrar, srv YMServer) {
	s.RegisterService(&YM_ServiceDesc, srv)
}

func _YM_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetAppData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetAppData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetAppData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetAppData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetShopWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetShopWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetShopWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetShopWidget(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetMainGraphic(ctx, req.(*LineGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetProductsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductsCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductsCountUnsorted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductsCountUnsorted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetProductsCountUnsorted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductsCountUnsorted(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetFlowGraphicData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetFlowGraphicData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetFlowGraphicData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetFlowGraphicData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetMarginGraphicData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetMarginGraphicData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetMarginGraphicData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetMarginGraphicData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetWeekGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetWeekGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetWeekGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetWeekGraphic(ctx, req.(*LineGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetPayRoundGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetPayRoundGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetPayRoundGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetPayRoundGraphic(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetCountRoundGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetCountRoundGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetCountRoundGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetCountRoundGraphic(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetOrderLeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetOrderLeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetOrderLeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetOrderLeaders(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetShopProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetShopProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetShopProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetShopProducts(ctx, req.(*RequestByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetSales(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

// YM_ServiceDesc is the grpc.ServiceDesc for YM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.YM",
	HandlerType: (*YMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _YM_Ping_Handler,
		},
		{
			MethodName: "GetAppData",
			Handler:    _YM_GetAppData_Handler,
		},
		{
			MethodName: "GetShopWidget",
			Handler:    _YM_GetShopWidget_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _YM_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetProductsCount",
			Handler:    _YM_GetProductsCount_Handler,
		},
		{
			MethodName: "GetProductsCountUnsorted",
			Handler:    _YM_GetProductsCountUnsorted_Handler,
		},
		{
			MethodName: "GetFlowGraphicData",
			Handler:    _YM_GetFlowGraphicData_Handler,
		},
		{
			MethodName: "GetMarginGraphicData",
			Handler:    _YM_GetMarginGraphicData_Handler,
		},
		{
			MethodName: "GetWeekGraphic",
			Handler:    _YM_GetWeekGraphic_Handler,
		},
		{
			MethodName: "GetPayRoundGraphic",
			Handler:    _YM_GetPayRoundGraphic_Handler,
		},
		{
			MethodName: "GetCountRoundGraphic",
			Handler:    _YM_GetCountRoundGraphic_Handler,
		},
		{
			MethodName: "GetOrderLeaders",
			Handler:    _YM_GetOrderLeaders_Handler,
		},
		{
			MethodName: "GetShopProducts",
			Handler:    _YM_GetShopProducts_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _YM_GetSales_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ym_v2.proto",
}
