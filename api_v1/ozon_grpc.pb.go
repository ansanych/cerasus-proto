// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ozon.proto

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

// OzonClient is the client API for Ozon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OzonClient interface {
	SetAuth(ctx context.Context, in *SetOzonAuth, opts ...grpc.CallOption) (*BoolReply, error)
	GetAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopOzonAuth, error)
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
	ForCounterDataOZ(ctx context.Context, in *ForCounterRequestOZ, opts ...grpc.CallOption) (*ForCounterReplyOZ, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetProductUrls(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ProductShopUrls, error)
}

type ozonClient struct {
	cc grpc.ClientConnInterface
}

func NewOzonClient(cc grpc.ClientConnInterface) OzonClient {
	return &ozonClient{cc}
}

func (c *ozonClient) SetAuth(ctx context.Context, in *SetOzonAuth, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/SetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopOzonAuth, error) {
	out := new(ShopOzonAuth)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) ErrorAuth(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/ErrorAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetUnsortedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetUnsortedList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error) {
	out := new(ShopProductListReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetUnsortedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetProductCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CountReply, error) {
	out := new(CountReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetProductCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetProductList(ctx context.Context, in *ShopProductListRequest, opts ...grpc.CallOption) (*ShopProductListReply, error) {
	out := new(ShopProductListReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetProduct(ctx context.Context, in *ShopProductRequest, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) UpdateProduct(ctx context.Context, in *ShopProductUpdateRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetDaySales(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DaysSalesReply, error) {
	out := new(DaysSalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetDaySales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetSales(ctx context.Context, in *SalesRequest, opts ...grpc.CallOption) (*SalesReply, error) {
	out := new(SalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetShopServices(ctx context.Context, in *ShopServiceRequest, opts ...grpc.CallOption) (*ShopServiceReply, error) {
	out := new(ShopServiceReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetShopServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetSaleDetail(ctx context.Context, in *SaleDetailsRequest, opts ...grpc.CallOption) (*SaleDetailsReply, error) {
	out := new(SaleDetailsReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetSaleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetProductSales(ctx context.Context, in *ProductSalesRequest, opts ...grpc.CallOption) (*SalesReply, error) {
	out := new(SalesReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetProductSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetMainGraphic(ctx context.Context, in *MainGraphicRequest, opts ...grpc.CallOption) (*MainGraphicReply, error) {
	out := new(MainGraphicReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) CheckShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShopData, error) {
	out := new(CompanyShopData)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/CheckShopData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetDonutGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*DonutGraphic, error) {
	out := new(DonutGraphic)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetDonutGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetWeekGraphics(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*WeekGraphics, error) {
	out := new(WeekGraphics)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetWeekGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) ForCounterDataOZ(ctx context.Context, in *ForCounterRequestOZ, opts ...grpc.CallOption) (*ForCounterReplyOZ, error) {
	out := new(ForCounterReplyOZ)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/ForCounterDataOZ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ozonClient) GetProductUrls(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*ProductShopUrls, error) {
	out := new(ProductShopUrls)
	err := c.cc.Invoke(ctx, "/cerasus.Ozon/GetProductUrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OzonServer is the server API for Ozon service.
// All implementations must embed UnimplementedOzonServer
// for forward compatibility
type OzonServer interface {
	SetAuth(context.Context, *SetOzonAuth) (*BoolReply, error)
	GetAuth(context.Context, *Auth) (*ShopOzonAuth, error)
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
	ForCounterDataOZ(context.Context, *ForCounterRequestOZ) (*ForCounterReplyOZ, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetProductUrls(context.Context, *IDRequest) (*ProductShopUrls, error)
	mustEmbedUnimplementedOzonServer()
}

// UnimplementedOzonServer must be embedded to have forward compatible implementations.
type UnimplementedOzonServer struct {
}

func (UnimplementedOzonServer) SetAuth(context.Context, *SetOzonAuth) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuth not implemented")
}
func (UnimplementedOzonServer) GetAuth(context.Context, *Auth) (*ShopOzonAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuth not implemented")
}
func (UnimplementedOzonServer) ErrorAuth(context.Context, *Auth) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ErrorAuth not implemented")
}
func (UnimplementedOzonServer) GetUnsortedCount(context.Context, *Auth) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedCount not implemented")
}
func (UnimplementedOzonServer) GetUnsortedList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnsortedList not implemented")
}
func (UnimplementedOzonServer) GetProductCount(context.Context, *Auth) (*CountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCount not implemented")
}
func (UnimplementedOzonServer) GetProductList(context.Context, *ShopProductListRequest) (*ShopProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedOzonServer) GetProduct(context.Context, *ShopProductRequest) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedOzonServer) UpdateProduct(context.Context, *ShopProductUpdateRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedOzonServer) GetDaySales(context.Context, *Auth) (*DaysSalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDaySales not implemented")
}
func (UnimplementedOzonServer) GetSales(context.Context, *SalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedOzonServer) GetShopServices(context.Context, *ShopServiceRequest) (*ShopServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopServices not implemented")
}
func (UnimplementedOzonServer) GetSaleDetail(context.Context, *SaleDetailsRequest) (*SaleDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSaleDetail not implemented")
}
func (UnimplementedOzonServer) GetProductSales(context.Context, *ProductSalesRequest) (*SalesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductSales not implemented")
}
func (UnimplementedOzonServer) GetMainGraphic(context.Context, *MainGraphicRequest) (*MainGraphicReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedOzonServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedOzonServer) CheckShopData(context.Context, *Auth) (*CompanyShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckShopData not implemented")
}
func (UnimplementedOzonServer) GetDonutGraphics(context.Context, *Auth) (*DonutGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonutGraphics not implemented")
}
func (UnimplementedOzonServer) GetWeekGraphics(context.Context, *Auth) (*WeekGraphics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekGraphics not implemented")
}
func (UnimplementedOzonServer) ForCounterDataOZ(context.Context, *ForCounterRequestOZ) (*ForCounterReplyOZ, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForCounterDataOZ not implemented")
}
func (UnimplementedOzonServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOzonServer) GetProductUrls(context.Context, *IDRequest) (*ProductShopUrls, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductUrls not implemented")
}
func (UnimplementedOzonServer) mustEmbedUnimplementedOzonServer() {}

// UnsafeOzonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OzonServer will
// result in compilation errors.
type UnsafeOzonServer interface {
	mustEmbedUnimplementedOzonServer()
}

func RegisterOzonServer(s grpc.ServiceRegistrar, srv OzonServer) {
	s.RegisterService(&Ozon_ServiceDesc, srv)
}

func _Ozon_SetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOzonAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).SetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/SetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).SetAuth(ctx, req.(*SetOzonAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetAuth(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_ErrorAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).ErrorAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/ErrorAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).ErrorAuth(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetUnsortedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetUnsortedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetUnsortedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetUnsortedCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetUnsortedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetUnsortedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetUnsortedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetUnsortedList(ctx, req.(*ShopProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetProductCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetProductCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetProductCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetProductCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetProductList(ctx, req.(*ShopProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetProduct(ctx, req.(*ShopProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).UpdateProduct(ctx, req.(*ShopProductUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetDaySales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetDaySales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetDaySales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetDaySales(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetSales(ctx, req.(*SalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetShopServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetShopServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetShopServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetShopServices(ctx, req.(*ShopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetSaleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaleDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetSaleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetSaleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetSaleDetail(ctx, req.(*SaleDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetProductSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetProductSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetProductSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetProductSales(ctx, req.(*ProductSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MainGraphicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetMainGraphic(ctx, req.(*MainGraphicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_CheckShopData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).CheckShopData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/CheckShopData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).CheckShopData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetDonutGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetDonutGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetDonutGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetDonutGraphics(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetWeekGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetWeekGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetWeekGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetWeekGraphics(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_ForCounterDataOZ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForCounterRequestOZ)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).ForCounterDataOZ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/ForCounterDataOZ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).ForCounterDataOZ(ctx, req.(*ForCounterRequestOZ))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ozon_GetProductUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OzonServer).GetProductUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Ozon/GetProductUrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OzonServer).GetProductUrls(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ozon_ServiceDesc is the grpc.ServiceDesc for Ozon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ozon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Ozon",
	HandlerType: (*OzonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAuth",
			Handler:    _Ozon_SetAuth_Handler,
		},
		{
			MethodName: "GetAuth",
			Handler:    _Ozon_GetAuth_Handler,
		},
		{
			MethodName: "ErrorAuth",
			Handler:    _Ozon_ErrorAuth_Handler,
		},
		{
			MethodName: "GetUnsortedCount",
			Handler:    _Ozon_GetUnsortedCount_Handler,
		},
		{
			MethodName: "GetUnsortedList",
			Handler:    _Ozon_GetUnsortedList_Handler,
		},
		{
			MethodName: "GetProductCount",
			Handler:    _Ozon_GetProductCount_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _Ozon_GetProductList_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Ozon_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Ozon_UpdateProduct_Handler,
		},
		{
			MethodName: "GetDaySales",
			Handler:    _Ozon_GetDaySales_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _Ozon_GetSales_Handler,
		},
		{
			MethodName: "GetShopServices",
			Handler:    _Ozon_GetShopServices_Handler,
		},
		{
			MethodName: "GetSaleDetail",
			Handler:    _Ozon_GetSaleDetail_Handler,
		},
		{
			MethodName: "GetProductSales",
			Handler:    _Ozon_GetProductSales_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _Ozon_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Ozon_GetImage_Handler,
		},
		{
			MethodName: "CheckShopData",
			Handler:    _Ozon_CheckShopData_Handler,
		},
		{
			MethodName: "GetDonutGraphics",
			Handler:    _Ozon_GetDonutGraphics_Handler,
		},
		{
			MethodName: "GetWeekGraphics",
			Handler:    _Ozon_GetWeekGraphics_Handler,
		},
		{
			MethodName: "ForCounterDataOZ",
			Handler:    _Ozon_ForCounterDataOZ_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Ozon_Ping_Handler,
		},
		{
			MethodName: "GetProductUrls",
			Handler:    _Ozon_GetProductUrls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozon.proto",
}
