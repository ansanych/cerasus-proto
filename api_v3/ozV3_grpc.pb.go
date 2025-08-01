// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ozV3.proto

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

// OZClient is the client API for OZ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OZClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error)
	SetAuthData(ctx context.Context, in *OZAuthParams, opts ...grpc.CallOption) (*StatusReply, error)
	SetShopProductUrl(ctx context.Context, in *ShopProductUrlSetter, opts ...grpc.CallOption) (*StatusReply, error)
	GetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Orders, error)
	GetSales(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Sales, error)
	SetQueueJob(ctx context.Context, in *QueuerJob, opts ...grpc.CallOption) (*StatusReply, error)
	GetConnectedCompanies(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*CompanyList, error)
	GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error)
	GetShopProducts(ctx context.Context, in *RequestByIDS, opts ...grpc.CallOption) (*ShopProductList, error)
	GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error)
	GetUnsortedList(ctx context.Context, in *RequestByIDS, opts ...grpc.CallOption) (*ShopProductList, error)
}

type oZClient struct {
	cc grpc.ClientConnInterface
}

func NewOZClient(cc grpc.ClientConnInterface) OZClient {
	return &oZClient{cc}
}

func (c *oZClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error) {
	out := new(AppShopData)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetAppData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) SetAuthData(ctx context.Context, in *OZAuthParams, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/SetAuthData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) SetShopProductUrl(ctx context.Context, in *ShopProductUrlSetter, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/SetShopProductUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetSales(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Sales, error) {
	out := new(Sales)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) SetQueueJob(ctx context.Context, in *QueuerJob, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/SetQueueJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetConnectedCompanies(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*CompanyList, error) {
	out := new(CompanyList)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetConnectedCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetShopProducts(ctx context.Context, in *RequestByIDS, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetShopProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetProductsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetUnsortedList(ctx context.Context, in *RequestByIDS, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/cerasusV3.OZ/GetUnsortedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OZServer is the server API for OZ service.
// All implementations must embed UnimplementedOZServer
// for forward compatibility
type OZServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetAppData(context.Context, *Auth) (*AppShopData, error)
	SetAuthData(context.Context, *OZAuthParams) (*StatusReply, error)
	SetShopProductUrl(context.Context, *ShopProductUrlSetter) (*StatusReply, error)
	GetOrders(context.Context, *RequestByDates) (*Orders, error)
	GetSales(context.Context, *RequestByDates) (*Sales, error)
	SetQueueJob(context.Context, *QueuerJob) (*StatusReply, error)
	GetConnectedCompanies(context.Context, *PingRequest) (*CompanyList, error)
	GetImage(context.Context, *ImageRequest) (*ImageReply, error)
	GetShopProducts(context.Context, *RequestByIDS) (*ShopProductList, error)
	GetProductsCount(context.Context, *Auth) (*Count, error)
	GetUnsortedList(context.Context, *RequestByIDS) (*ShopProductList, error)
	mustEmbedUnimplementedOZServer()
}

// UnimplementedOZServer must be embedded to have forward compatible implementations.
type UnimplementedOZServer struct {
}

func (UnimplementedOZServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOZServer) GetAppData(context.Context, *Auth) (*AppShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppData not implemented")
}
func (UnimplementedOZServer) SetAuthData(context.Context, *OZAuthParams) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuthData not implemented")
}
func (UnimplementedOZServer) SetShopProductUrl(context.Context, *ShopProductUrlSetter) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetShopProductUrl not implemented")
}
func (UnimplementedOZServer) GetOrders(context.Context, *RequestByDates) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOZServer) GetSales(context.Context, *RequestByDates) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedOZServer) SetQueueJob(context.Context, *QueuerJob) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetQueueJob not implemented")
}
func (UnimplementedOZServer) GetConnectedCompanies(context.Context, *PingRequest) (*CompanyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectedCompanies not implemented")
}
func (UnimplementedOZServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedOZServer) GetShopProducts(context.Context, *RequestByIDS) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProducts not implemented")
}
func (UnimplementedOZServer) GetProductsCount(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCount not implemented")
}
func (UnimplementedOZServer) GetUnsortedList(context.Context, *RequestByIDS) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedList not implemented")
}
func (UnimplementedOZServer) mustEmbedUnimplementedOZServer() {}

// UnsafeOZServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OZServer will
// result in compilation errors.
type UnsafeOZServer interface {
	mustEmbedUnimplementedOZServer()
}

func RegisterOZServer(s grpc.ServiceRegistrar, srv OZServer) {
	s.RegisterService(&OZ_ServiceDesc, srv)
}

func _OZ_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetAppData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetAppData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetAppData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetAppData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_SetAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OZAuthParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).SetAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/SetAuthData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).SetAuthData(ctx, req.(*OZAuthParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_SetShopProductUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductUrlSetter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).SetShopProductUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/SetShopProductUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).SetShopProductUrl(ctx, req.(*ShopProductUrlSetter))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetOrders(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetSales(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_SetQueueJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueuerJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).SetQueueJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/SetQueueJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).SetQueueJob(ctx, req.(*QueuerJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetConnectedCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetConnectedCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetConnectedCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetConnectedCompanies(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetShopProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetShopProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetShopProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetShopProducts(ctx, req.(*RequestByIDS))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetProductsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetProductsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetProductsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductsCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetUnsortedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetUnsortedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV3.OZ/GetUnsortedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetUnsortedList(ctx, req.(*RequestByIDS))
	}
	return interceptor(ctx, in, info, handler)
}

// OZ_ServiceDesc is the grpc.ServiceDesc for OZ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OZ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV3.OZ",
	HandlerType: (*OZServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _OZ_Ping_Handler,
		},
		{
			MethodName: "GetAppData",
			Handler:    _OZ_GetAppData_Handler,
		},
		{
			MethodName: "SetAuthData",
			Handler:    _OZ_SetAuthData_Handler,
		},
		{
			MethodName: "SetShopProductUrl",
			Handler:    _OZ_SetShopProductUrl_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OZ_GetOrders_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _OZ_GetSales_Handler,
		},
		{
			MethodName: "SetQueueJob",
			Handler:    _OZ_SetQueueJob_Handler,
		},
		{
			MethodName: "GetConnectedCompanies",
			Handler:    _OZ_GetConnectedCompanies_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _OZ_GetImage_Handler,
		},
		{
			MethodName: "GetShopProducts",
			Handler:    _OZ_GetShopProducts_Handler,
		},
		{
			MethodName: "GetProductsCount",
			Handler:    _OZ_GetProductsCount_Handler,
		},
		{
			MethodName: "GetUnsortedList",
			Handler:    _OZ_GetUnsortedList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozV3.proto",
}
