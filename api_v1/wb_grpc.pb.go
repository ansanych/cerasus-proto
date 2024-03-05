// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: wb.proto

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
	SetAuth(ctx context.Context, in *SetWBAuth, opts ...grpc.CallOption) (*BoolReply, error)
	GetAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWBAuth, error)
	ErrorAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BoolReply, error)
	GetUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error)
	GetUnsortedList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error)
	GetProductCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error)
	GetProductList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error)
	GetProduct(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error)
	GetDaySales(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DaysSalesReply, error)
	GetSales(ctx context.Context, in *SalesRequest, opts ...grpc.CallOption) (*SalesReply, error)
}

type wBClient struct {
	cc grpc.ClientConnInterface
}

func NewWBClient(cc grpc.ClientConnInterface) WBClient {
	return &wBClient{cc}
}

func (c *wBClient) SetAuth(ctx context.Context, in *SetWBAuth, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/SetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWBAuth, error) {
	out := new(ShopWBAuth)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) ErrorAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/ErrorAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetUnsortedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetUnsortedList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error) {
	out := new(ShopProductListReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetUnsortedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProductCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetProductCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProductList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error) {
	out := new(ShopProductListReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProduct(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetDaySales(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DaysSalesReply, error) {
	out := new(DaysSalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetDaySales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetSales(ctx context.Context, in *SalesRequest, opts ...grpc.CallOption) (*SalesReply, error) {
	out := new(SalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WBServer is the server API for WB service.
// All implementations must embed UnimplementedWBServer
// for forward compatibility
type WBServer interface {
	SetAuth(context.Context, *SetWBAuth) (*BoolReply, error)
	GetAuth(context.Context, *Auth) (*ShopWBAuth, error)
	ErrorAuth(context.Context, *Auth) (*BoolReply, error)
	GetUnsortedCount(context.Context, *Auth) (*CountReply, error)
	GetUnsortedList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error)
	GetProductCount(context.Context, *Auth) (*CountReply, error)
	GetProductList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error)
	GetProduct(context.Context, *ShopProductRequest) (*ShopProduct, error)
	GetDaySales(context.Context, *Auth) (*DaysSalesReply, error)
	GetSales(context.Context, *SalesRequest) (*SalesReply, error)
	mustEmbedUnimplementedWBServer()
}

// UnimplementedWBServer must be embedded to have forward compatible implementations.
type UnimplementedWBServer struct {
}

func (UnimplementedWBServer) SetAuth(context.Context, *SetWBAuth) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuth not implemented")
}
func (UnimplementedWBServer) GetAuth(context.Context, *Auth) (*ShopWBAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
}
func (UnimplementedWBServer) ErrorAuth(context.Context, *Auth) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ErrorAuth not implemented")
}
func (UnimplementedWBServer) GetUnsortedCount(context.Context, *Auth) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedCount not implemented")
}
func (UnimplementedWBServer) GetUnsortedList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedList not implemented")
}
func (UnimplementedWBServer) GetProductCount(context.Context, *Auth) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCount not implemented")
}
func (UnimplementedWBServer) GetProductList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedWBServer) GetProduct(context.Context, *ShopProductRequest) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedWBServer) GetDaySales(context.Context, *Auth) (*DaysSalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDaySales not implemented")
}
func (UnimplementedWBServer) GetSales(context.Context, *SalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
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

func _WB_SetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWBAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).SetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/SetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).SetAuth(ctx, req.(*SetWBAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetAuth(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_ErrorAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).ErrorAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/ErrorAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).ErrorAuth(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetUnsortedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetUnsortedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetUnsortedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetUnsortedCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetUnsortedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetUnsortedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetUnsortedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetUnsortedList(ctx, req.(*ShopProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProductCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProductCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetProductCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProductCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProductList(ctx, req.(*ShopProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProduct(ctx, req.(*ShopProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetDaySales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetDaySales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetDaySales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetDaySales(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetSales(ctx, req.(*SalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WB_ServiceDesc is the grpc.ServiceDesc for WB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.WB",
	HandlerType: (*WBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAuth",
			Handler:    _WB_SetAuth_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _WB_GetAuth_Handler,
		},
		{
			MethodName: "ErrorAuth",
			Handler:    _WB_ErrorAuth_Handler,
		},
		{
			MethodName: "GetUnsortedCount",
			Handler:    _WB_GetUnsortedCount_Handler,
		},
		{
			MethodName: "GetUnsortedList",
			Handler:    _WB_GetUnsortedList_Handler,
		},
		{
			MethodName: "GetProductCount",
			Handler:    _WB_GetProductCount_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _WB_GetProductList_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _WB_GetProduct_Handler,
		},
		{
			MethodName: "GetDaySales",
			Handler:    _WB_GetDaySales_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _WB_GetSales_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb.proto",
}
