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
	UpdateProduct(ctx context.Context, in *ShopProductUpdateRequest, opts ...grpc.CallOption) (*BoolReply, error)
	GetDaySales(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DaysSalesReply, error)
	GetSales(ctx context.Context, in *SalesRequest, opts ...grpc.CallOption) (*SalesReply, error)
	GetShopServices(ctx context.Context, in *ShopServiceRequest, opts ...grpc.CallOption) (*ShopServiceReply, error)
	GetSaleDetail(ctx context.Context, in *SaleDetailsRequest, opts ...grpc.CallOption) (*SaleDetailsReply, error)
	GetProductSales(ctx context.Context, in *ProductSalesRequest, opts ...grpc.CallOption) (*SalesReply, error)
	GetMainGraphic(ctx context.Context, in *MainGraphicRequest, opts ...grpc.CallOption) (*MainGraphicReply, error)
	GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error)
	CheckShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShopData, error)
	GetDonutGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DonutGraphic, error)
	GetWeekGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*WeekGraphics, error)
	ForCounterData(ctx context.Context, in *ForCounterRequest, opts ...grpc.CallOption) (*ForCounterReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
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

func (c *wBClient) UpdateProduct(ctx context.Context, in *ShopProductUpdateRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/UpdateProduct", in, out, opts...)
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

func (c *wBClient) GetShopServices(ctx context.Context, in *ShopServiceRequest, opts ...grpc.CallOption) (*ShopServiceReply, error) {
	out := new(ShopServiceReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetShopServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetSaleDetail(ctx context.Context, in *SaleDetailsRequest, opts ...grpc.CallOption) (*SaleDetailsReply, error) {
	out := new(SaleDetailsReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetSaleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetProductSales(ctx context.Context, in *ProductSalesRequest, opts ...grpc.CallOption) (*SalesReply, error) {
	out := new(SalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetProductSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetMainGraphic(ctx context.Context, in *MainGraphicRequest, opts ...grpc.CallOption) (*MainGraphicReply, error) {
	out := new(MainGraphicReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) CheckShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShopData, error) {
	out := new(CompanyShopData)
	err := c.cc.Invoke(ctx, "/cerasus.WB/CheckShopData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetDonutGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DonutGraphic, error) {
	out := new(DonutGraphic)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetDonutGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) GetWeekGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*WeekGraphics, error) {
	out := new(WeekGraphics)
	err := c.cc.Invoke(ctx, "/cerasus.WB/GetWeekGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) ForCounterData(ctx context.Context, in *ForCounterRequest, opts ...grpc.CallOption) (*ForCounterReply, error) {
	out := new(ForCounterReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/ForCounterData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wBClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.WB/Ping", in, out, opts...)
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
	UpdateProduct(context.Context, *ShopProductUpdateRequest) (*BoolReply, error)
	GetDaySales(context.Context, *Auth) (*DaysSalesReply, error)
	GetSales(context.Context, *SalesRequest) (*SalesReply, error)
	GetShopServices(context.Context, *ShopServiceRequest) (*ShopServiceReply, error)
	GetSaleDetail(context.Context, *SaleDetailsRequest) (*SaleDetailsReply, error)
	GetProductSales(context.Context, *ProductSalesRequest) (*SalesReply, error)
	GetMainGraphic(context.Context, *MainGraphicRequest) (*MainGraphicReply, error)
	GetImage(context.Context, *ImageRequest) (*ImageReply, error)
	CheckShopData(context.Context, *Auth) (*CompanyShopData, error)
	GetDonutGraphics(context.Context, *Auth) (*DonutGraphic, error)
	GetWeekGraphics(context.Context, *Auth) (*WeekGraphics, error)
	ForCounterData(context.Context, *ForCounterRequest) (*ForCounterReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
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
func (UnimplementedWBServer) UpdateProduct(context.Context, *ShopProductUpdateRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedWBServer) GetDaySales(context.Context, *Auth) (*DaysSalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDaySales not implemented")
}
func (UnimplementedWBServer) GetSales(context.Context, *SalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedWBServer) GetShopServices(context.Context, *ShopServiceRequest) (*ShopServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopServices not implemented")
}
func (UnimplementedWBServer) GetSaleDetail(context.Context, *SaleDetailsRequest) (*SaleDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaleDetail not implemented")
}
func (UnimplementedWBServer) GetProductSales(context.Context, *ProductSalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductSales not implemented")
}
func (UnimplementedWBServer) GetMainGraphic(context.Context, *MainGraphicRequest) (*MainGraphicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedWBServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedWBServer) CheckShopData(context.Context, *Auth) (*CompanyShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckShopData not implemented")
}
func (UnimplementedWBServer) GetDonutGraphics(context.Context, *Auth) (*DonutGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonutGraphics not implemented")
}
func (UnimplementedWBServer) GetWeekGraphics(context.Context, *Auth) (*WeekGraphics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekGraphics not implemented")
}
func (UnimplementedWBServer) ForCounterData(context.Context, *ForCounterRequest) (*ForCounterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForCounterData not implemented")
}
func (UnimplementedWBServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
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

func _WB_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).UpdateProduct(ctx, req.(*ShopProductUpdateRequest))
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

func _WB_GetShopServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetShopServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetShopServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetShopServices(ctx, req.(*ShopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetSaleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaleDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetSaleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetSaleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetSaleDetail(ctx, req.(*SaleDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetProductSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetProductSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetProductSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetProductSales(ctx, req.(*ProductSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MainGraphicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetMainGraphic(ctx, req.(*MainGraphicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_CheckShopData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).CheckShopData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/CheckShopData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).CheckShopData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetDonutGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetDonutGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetDonutGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetDonutGraphics(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_GetWeekGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).GetWeekGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/GetWeekGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).GetWeekGraphics(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _WB_ForCounterData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WBServer).ForCounterData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.WB/ForCounterData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).ForCounterData(ctx, req.(*ForCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/cerasus.WB/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WBServer).Ping(ctx, req.(*PingRequest))
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
			MethodName: "UpdateProduct",
			Handler:    _WB_UpdateProduct_Handler,
		},
		{
			MethodName: "GetDaySales",
			Handler:    _WB_GetDaySales_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _WB_GetSales_Handler,
		},
		{
			MethodName: "GetShopServices",
			Handler:    _WB_GetShopServices_Handler,
		},
		{
			MethodName: "GetSaleDetail",
			Handler:    _WB_GetSaleDetail_Handler,
		},
		{
			MethodName: "GetProductSales",
			Handler:    _WB_GetProductSales_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _WB_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _WB_GetImage_Handler,
		},
		{
			MethodName: "CheckShopData",
			Handler:    _WB_CheckShopData_Handler,
		},
		{
			MethodName: "GetDonutGraphics",
			Handler:    _WB_GetDonutGraphics_Handler,
		},
		{
			MethodName: "GetWeekGraphics",
			Handler:    _WB_GetWeekGraphics_Handler,
		},
		{
			MethodName: "ForCounterData",
			Handler:    _WB_ForCounterData_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _WB_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wb.proto",
}
