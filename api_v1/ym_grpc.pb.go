// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ym.proto

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
	SetAuth(ctx context.Context, in *SetYMAuth, opts ...grpc.CallOption) (*BoolReply, error)
	GetAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopYMAuth, error)
	ErrorAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BoolReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	CheckShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShopData, error)
	GetUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error)
	GetUnsortedList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error)
	GetProductCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error)
	GetProductList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error)
	GetProduct(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error)
	UpdateProduct(ctx context.Context, in *ShopProductUpdateRequest, opts ...grpc.CallOption) (*BoolReply, error)
	GetDaySales(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DaysSalesReply, error)
	GetSales(ctx context.Context, in *SalesRequest, opts ...grpc.CallOption) (*SalesReply, error)
	GetShopServices(ctx context.Context, in *ShopServiceRequest, opts ...grpc.CallOption) (*ShopServiceReply, error)
	GetSaleDetail(ctx context.Context, in *SaleDetailsRequest, opts ...grpc.CallOption) (*SaleDetailsReply, error)
	GetProductSales(ctx context.Context, in *ProductSalesRequest, opts ...grpc.CallOption) (*SalesReply, error)
	GetMainGraphic(ctx context.Context, in *MainGraphicRequest, opts ...grpc.CallOption) (*MainGraphicReply, error)
	GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error)
	GetDonutGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DonutGraphic, error)
	GetWeekGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*WeekGraphics, error)
	ForCounterDataYM(ctx context.Context, in *ForCounterRequestYM, opts ...grpc.CallOption) (*ForCounterReplyYM, error)
}

type yMClient struct {
	cc grpc.ClientConnInterface
}

func NewYMClient(cc grpc.ClientConnInterface) YMClient {
	return &yMClient{cc}
}

func (c *yMClient) SetAuth(ctx context.Context, in *SetYMAuth, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/SetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopYMAuth, error) {
	out := new(ShopYMAuth)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) ErrorAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/ErrorAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) CheckShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShopData, error) {
	out := new(CompanyShopData)
	err := c.cc.Invoke(ctx, "/cerasus.YM/CheckShopData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetUnsortedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetUnsortedList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error) {
	out := new(ShopProductListReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetUnsortedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetProductCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error) {
	out := new(ShopProductListReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProduct(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) UpdateProduct(ctx context.Context, in *ShopProductUpdateRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetDaySales(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DaysSalesReply, error) {
	out := new(DaysSalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetDaySales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetSales(ctx context.Context, in *SalesRequest, opts ...grpc.CallOption) (*SalesReply, error) {
	out := new(SalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetShopServices(ctx context.Context, in *ShopServiceRequest, opts ...grpc.CallOption) (*ShopServiceReply, error) {
	out := new(ShopServiceReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetShopServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetSaleDetail(ctx context.Context, in *SaleDetailsRequest, opts ...grpc.CallOption) (*SaleDetailsReply, error) {
	out := new(SaleDetailsReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetSaleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductSales(ctx context.Context, in *ProductSalesRequest, opts ...grpc.CallOption) (*SalesReply, error) {
	out := new(SalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetProductSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetMainGraphic(ctx context.Context, in *MainGraphicRequest, opts ...grpc.CallOption) (*MainGraphicReply, error) {
	out := new(MainGraphicReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetDonutGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DonutGraphic, error) {
	out := new(DonutGraphic)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetDonutGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetWeekGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*WeekGraphics, error) {
	out := new(WeekGraphics)
	err := c.cc.Invoke(ctx, "/cerasus.YM/GetWeekGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) ForCounterDataYM(ctx context.Context, in *ForCounterRequestYM, opts ...grpc.CallOption) (*ForCounterReplyYM, error) {
	out := new(ForCounterReplyYM)
	err := c.cc.Invoke(ctx, "/cerasus.YM/ForCounterDataYM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YMServer is the server API for YM service.
// All implementations must embed UnimplementedYMServer
// for forward compatibility
type YMServer interface {
	SetAuth(context.Context, *SetYMAuth) (*BoolReply, error)
	GetAuth(context.Context, *Auth) (*ShopYMAuth, error)
	ErrorAuth(context.Context, *Auth) (*BoolReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	CheckShopData(context.Context, *Auth) (*CompanyShopData, error)
	GetUnsortedCount(context.Context, *Auth) (*CountReply, error)
	GetUnsortedList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error)
	GetProductCount(context.Context, *Auth) (*CountReply, error)
	GetProductList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error)
	GetProduct(context.Context, *ShopProductRequest) (*ShopProduct, error)
	UpdateProduct(context.Context, *ShopProductUpdateRequest) (*BoolReply, error)
	GetDaySales(context.Context, *Auth) (*DaysSalesReply, error)
	GetSales(context.Context, *SalesRequest) (*SalesReply, error)
	GetShopServices(context.Context, *ShopServiceRequest) (*ShopServiceReply, error)
	GetSaleDetail(context.Context, *SaleDetailsRequest) (*SaleDetailsReply, error)
	GetProductSales(context.Context, *ProductSalesRequest) (*SalesReply, error)
	GetMainGraphic(context.Context, *MainGraphicRequest) (*MainGraphicReply, error)
	GetImage(context.Context, *ImageRequest) (*ImageReply, error)
	GetDonutGraphics(context.Context, *Auth) (*DonutGraphic, error)
	GetWeekGraphics(context.Context, *Auth) (*WeekGraphics, error)
	ForCounterDataYM(context.Context, *ForCounterRequestYM) (*ForCounterReplyYM, error)
	mustEmbedUnimplementedYMServer()
}

// UnimplementedYMServer must be embedded to have forward compatible implementations.
type UnimplementedYMServer struct {
}

func (UnimplementedYMServer) SetAuth(context.Context, *SetYMAuth) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuth not implemented")
}
func (UnimplementedYMServer) GetAuth(context.Context, *Auth) (*ShopYMAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
}
func (UnimplementedYMServer) ErrorAuth(context.Context, *Auth) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ErrorAuth not implemented")
}
func (UnimplementedYMServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedYMServer) CheckShopData(context.Context, *Auth) (*CompanyShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckShopData not implemented")
}
func (UnimplementedYMServer) GetUnsortedCount(context.Context, *Auth) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedCount not implemented")
}
func (UnimplementedYMServer) GetUnsortedList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedList not implemented")
}
func (UnimplementedYMServer) GetProductCount(context.Context, *Auth) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCount not implemented")
}
func (UnimplementedYMServer) GetProductList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedYMServer) GetProduct(context.Context, *ShopProductRequest) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedYMServer) UpdateProduct(context.Context, *ShopProductUpdateRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedYMServer) GetDaySales(context.Context, *Auth) (*DaysSalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDaySales not implemented")
}
func (UnimplementedYMServer) GetSales(context.Context, *SalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedYMServer) GetShopServices(context.Context, *ShopServiceRequest) (*ShopServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopServices not implemented")
}
func (UnimplementedYMServer) GetSaleDetail(context.Context, *SaleDetailsRequest) (*SaleDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaleDetail not implemented")
}
func (UnimplementedYMServer) GetProductSales(context.Context, *ProductSalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductSales not implemented")
}
func (UnimplementedYMServer) GetMainGraphic(context.Context, *MainGraphicRequest) (*MainGraphicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedYMServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedYMServer) GetDonutGraphics(context.Context, *Auth) (*DonutGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonutGraphics not implemented")
}
func (UnimplementedYMServer) GetWeekGraphics(context.Context, *Auth) (*WeekGraphics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekGraphics not implemented")
}
func (UnimplementedYMServer) ForCounterDataYM(context.Context, *ForCounterRequestYM) (*ForCounterReplyYM, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForCounterDataYM not implemented")
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

func _YM_SetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetYMAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).SetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/SetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).SetAuth(ctx, req.(*SetYMAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetAuth(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_ErrorAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).ErrorAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/ErrorAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).ErrorAuth(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/cerasus.YM/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_CheckShopData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).CheckShopData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/CheckShopData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).CheckShopData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetUnsortedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetUnsortedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetUnsortedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetUnsortedCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetUnsortedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetUnsortedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetUnsortedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetUnsortedList(ctx, req.(*ShopProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetProductCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductList(ctx, req.(*ShopProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProduct(ctx, req.(*ShopProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).UpdateProduct(ctx, req.(*ShopProductUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetDaySales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetDaySales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetDaySales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetDaySales(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetSales(ctx, req.(*SalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetShopServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetShopServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetShopServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetShopServices(ctx, req.(*ShopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetSaleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaleDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetSaleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetSaleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetSaleDetail(ctx, req.(*SaleDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetProductSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductSales(ctx, req.(*ProductSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MainGraphicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetMainGraphic(ctx, req.(*MainGraphicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetDonutGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetDonutGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetDonutGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetDonutGraphics(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetWeekGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetWeekGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/GetWeekGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetWeekGraphics(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_ForCounterDataYM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForCounterRequestYM)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).ForCounterDataYM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.YM/ForCounterDataYM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).ForCounterDataYM(ctx, req.(*ForCounterRequestYM))
	}
	return interceptor(ctx, in, info, handler)
}

// YM_ServiceDesc is the grpc.ServiceDesc for YM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.YM",
	HandlerType: (*YMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAuth",
			Handler:    _YM_SetAuth_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _YM_GetAuth_Handler,
		},
		{
			MethodName: "ErrorAuth",
			Handler:    _YM_ErrorAuth_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _YM_Ping_Handler,
		},
		{
			MethodName: "CheckShopData",
			Handler:    _YM_CheckShopData_Handler,
		},
		{
			MethodName: "GetUnsortedCount",
			Handler:    _YM_GetUnsortedCount_Handler,
		},
		{
			MethodName: "GetUnsortedList",
			Handler:    _YM_GetUnsortedList_Handler,
		},
		{
			MethodName: "GetProductCount",
			Handler:    _YM_GetProductCount_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _YM_GetProductList_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _YM_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _YM_UpdateProduct_Handler,
		},
		{
			MethodName: "GetDaySales",
			Handler:    _YM_GetDaySales_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _YM_GetSales_Handler,
		},
		{
			MethodName: "GetShopServices",
			Handler:    _YM_GetShopServices_Handler,
		},
		{
			MethodName: "GetSaleDetail",
			Handler:    _YM_GetSaleDetail_Handler,
		},
		{
			MethodName: "GetProductSales",
			Handler:    _YM_GetProductSales_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _YM_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _YM_GetImage_Handler,
		},
		{
			MethodName: "GetDonutGraphics",
			Handler:    _YM_GetDonutGraphics_Handler,
		},
		{
			MethodName: "GetWeekGraphics",
			Handler:    _YM_GetWeekGraphics_Handler,
		},
		{
			MethodName: "ForCounterDataYM",
			Handler:    _YM_ForCounterDataYM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ym.proto",
}
