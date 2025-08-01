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
	GetShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopData, error)
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
	GetShopProduct(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*ShopProduct, error)
	GetProductGraphics(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*LineGraphics, error)
	GetSales(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Sales, error)
	GetProductWidget(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error)
	GetProductWidgetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error)
	GetSale(ctx context.Context, in *SaleRequest, opts ...grpc.CallOption) (*Sale, error)
	GetProductsUnsortedList(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopProductList, error)
	SetShopProductUrl(ctx context.Context, in *ShopProductUrlSetter, opts ...grpc.CallOption) (*StatusReply, error)
	GetCounterParams(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*YMCounterParams, error)
	GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error)
	GetShopProductByCode(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ShopProduct, error)
	GetOrdersForBrand(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*Orders, error)
	SetAuthData(ctx context.Context, in *YMAuthData, opts ...grpc.CallOption) (*StatusReply, error)
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

func (c *yMClient) GetShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopData, error) {
	out := new(ShopData)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetShopData", in, out, opts...)
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

func (c *yMClient) GetShopProduct(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetShopProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductGraphics(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*LineGraphics, error) {
	out := new(LineGraphics)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetProductGraphics", in, out, opts...)
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

func (c *yMClient) GetProductWidget(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error) {
	out := new(ProductWidgets)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetProductWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductWidgetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error) {
	out := new(ProductWidgets)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetProductWidgetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetSale(ctx context.Context, in *SaleRequest, opts ...grpc.CallOption) (*Sale, error) {
	out := new(Sale)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetSale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetProductsUnsortedList(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetProductsUnsortedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) SetShopProductUrl(ctx context.Context, in *ShopProductUrlSetter, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/SetShopProductUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetCounterParams(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*YMCounterParams, error) {
	out := new(YMCounterParams)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetCounterParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetShopProductByCode(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetShopProductByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) GetOrdersForBrand(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/GetOrdersForBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yMClient) SetAuthData(ctx context.Context, in *YMAuthData, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.YM/SetAuthData", in, out, opts...)
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
	GetShopData(context.Context, *Auth) (*ShopData, error)
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
	GetShopProduct(context.Context, *RequestByID) (*ShopProduct, error)
	GetProductGraphics(context.Context, *RequestByDates) (*LineGraphics, error)
	GetSales(context.Context, *RequestByDates) (*Sales, error)
	GetProductWidget(context.Context, *RequestByDates) (*ProductWidgets, error)
	GetProductWidgetOrders(context.Context, *RequestByDates) (*ProductWidgets, error)
	GetSale(context.Context, *SaleRequest) (*Sale, error)
	GetProductsUnsortedList(context.Context, *Auth) (*ShopProductList, error)
	SetShopProductUrl(context.Context, *ShopProductUrlSetter) (*StatusReply, error)
	GetCounterParams(context.Context, *RequestByIDs) (*YMCounterParams, error)
	GetImage(context.Context, *ImageRequest) (*ImageReply, error)
	GetShopProductByCode(context.Context, *SearchRequest) (*ShopProduct, error)
	GetOrdersForBrand(context.Context, *OrdersRequest) (*Orders, error)
	SetAuthData(context.Context, *YMAuthData) (*StatusReply, error)
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
func (UnimplementedYMServer) GetShopData(context.Context, *Auth) (*ShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopData not implemented")
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
func (UnimplementedYMServer) GetShopProduct(context.Context, *RequestByID) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProduct not implemented")
}
func (UnimplementedYMServer) GetProductGraphics(context.Context, *RequestByDates) (*LineGraphics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGraphics not implemented")
}
func (UnimplementedYMServer) GetSales(context.Context, *RequestByDates) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedYMServer) GetProductWidget(context.Context, *RequestByDates) (*ProductWidgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWidget not implemented")
}
func (UnimplementedYMServer) GetProductWidgetOrders(context.Context, *RequestByDates) (*ProductWidgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWidgetOrders not implemented")
}
func (UnimplementedYMServer) GetSale(context.Context, *SaleRequest) (*Sale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSale not implemented")
}
func (UnimplementedYMServer) GetProductsUnsortedList(context.Context, *Auth) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsUnsortedList not implemented")
}
func (UnimplementedYMServer) SetShopProductUrl(context.Context, *ShopProductUrlSetter) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetShopProductUrl not implemented")
}
func (UnimplementedYMServer) GetCounterParams(context.Context, *RequestByIDs) (*YMCounterParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounterParams not implemented")
}
func (UnimplementedYMServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedYMServer) GetShopProductByCode(context.Context, *SearchRequest) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProductByCode not implemented")
}
func (UnimplementedYMServer) GetOrdersForBrand(context.Context, *OrdersRequest) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersForBrand not implemented")
}
func (UnimplementedYMServer) SetAuthData(context.Context, *YMAuthData) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuthData not implemented")
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

func _YM_GetShopData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetShopData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetShopData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetShopData(ctx, req.(*Auth))
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

func _YM_GetShopProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetShopProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetShopProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetShopProduct(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetProductGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductGraphics(ctx, req.(*RequestByDates))
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

func _YM_GetProductWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetProductWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductWidget(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductWidgetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductWidgetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetProductWidgetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductWidgetOrders(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetSale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetSale(ctx, req.(*SaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetProductsUnsortedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetProductsUnsortedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetProductsUnsortedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetProductsUnsortedList(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_SetShopProductUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopProductUrlSetter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).SetShopProductUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/SetShopProductUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).SetShopProductUrl(ctx, req.(*ShopProductUrlSetter))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetCounterParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetCounterParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetCounterParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetCounterParams(ctx, req.(*RequestByIDs))
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
		FullMethod: "/cerasusV2.YM/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetShopProductByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetShopProductByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetShopProductByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetShopProductByCode(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_GetOrdersForBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).GetOrdersForBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/GetOrdersForBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).GetOrdersForBrand(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YM_SetAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YMAuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YMServer).SetAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.YM/SetAuthData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YMServer).SetAuthData(ctx, req.(*YMAuthData))
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
			MethodName: "GetShopData",
			Handler:    _YM_GetShopData_Handler,
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
			MethodName: "GetShopProduct",
			Handler:    _YM_GetShopProduct_Handler,
		},
		{
			MethodName: "GetProductGraphics",
			Handler:    _YM_GetProductGraphics_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _YM_GetSales_Handler,
		},
		{
			MethodName: "GetProductWidget",
			Handler:    _YM_GetProductWidget_Handler,
		},
		{
			MethodName: "GetProductWidgetOrders",
			Handler:    _YM_GetProductWidgetOrders_Handler,
		},
		{
			MethodName: "GetSale",
			Handler:    _YM_GetSale_Handler,
		},
		{
			MethodName: "GetProductsUnsortedList",
			Handler:    _YM_GetProductsUnsortedList_Handler,
		},
		{
			MethodName: "SetShopProductUrl",
			Handler:    _YM_SetShopProductUrl_Handler,
		},
		{
			MethodName: "GetCounterParams",
			Handler:    _YM_GetCounterParams_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _YM_GetImage_Handler,
		},
		{
			MethodName: "GetShopProductByCode",
			Handler:    _YM_GetShopProductByCode_Handler,
		},
		{
			MethodName: "GetOrdersForBrand",
			Handler:    _YM_GetOrdersForBrand_Handler,
		},
		{
			MethodName: "SetAuthData",
			Handler:    _YM_SetAuthData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ym_v2.proto",
}
