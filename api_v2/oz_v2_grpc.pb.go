// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: oz_v2.proto

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
	GetCounterParams(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*OZCounterParams, error)
	GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error)
	GetShopProductByCode(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ShopProduct, error)
	GetOrdersForBrand(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*Orders, error)
	SetAuthData(ctx context.Context, in *OZAuthData, opts ...grpc.CallOption) (*StatusReply, error)
}

type oZClient struct {
	cc grpc.ClientConnInterface
}

func NewOZClient(cc grpc.ClientConnInterface) OZClient {
	return &oZClient{cc}
}

func (c *oZClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppShopData, error) {
	out := new(AppShopData)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetAppData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetShopData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopData, error) {
	out := new(ShopData)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetShopData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetShopWidget(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopWidget, error) {
	out := new(ShopWidget)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetShopWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error) {
	out := new(LineGraph)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetProductsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductsCountUnsorted(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetProductsCountUnsorted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetFlowGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetFlowGraphicData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetMarginGraphicData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetMarginGraphicData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetWeekGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*WeekGraphic, error) {
	out := new(WeekGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetWeekGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetPayRoundGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error) {
	out := new(RoundGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetPayRoundGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetCountRoundGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error) {
	out := new(RoundGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetCountRoundGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetOrderLeaders(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*OrderLeaders, error) {
	out := new(OrderLeaders)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetOrderLeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetShopProducts(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetShopProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetShopProduct(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetShopProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductGraphics(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*LineGraphics, error) {
	out := new(LineGraphics)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetProductGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetSales(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Sales, error) {
	out := new(Sales)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetSales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductWidget(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error) {
	out := new(ProductWidgets)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetProductWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductWidgetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error) {
	out := new(ProductWidgets)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetProductWidgetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetSale(ctx context.Context, in *SaleRequest, opts ...grpc.CallOption) (*Sale, error) {
	out := new(Sale)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetSale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetProductsUnsortedList(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ShopProductList, error) {
	out := new(ShopProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetProductsUnsortedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) SetShopProductUrl(ctx context.Context, in *ShopProductUrlSetter, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/SetShopProductUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetCounterParams(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*OZCounterParams, error) {
	out := new(OZCounterParams)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetCounterParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetShopProductByCode(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ShopProduct, error) {
	out := new(ShopProduct)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetShopProductByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) GetOrdersForBrand(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/GetOrdersForBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oZClient) SetAuthData(ctx context.Context, in *OZAuthData, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.OZ/SetAuthData", in, out, opts...)
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
	GetCounterParams(context.Context, *RequestByIDs) (*OZCounterParams, error)
	GetImage(context.Context, *ImageRequest) (*ImageReply, error)
	GetShopProductByCode(context.Context, *SearchRequest) (*ShopProduct, error)
	GetOrdersForBrand(context.Context, *OrdersRequest) (*Orders, error)
	SetAuthData(context.Context, *OZAuthData) (*StatusReply, error)
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
func (UnimplementedOZServer) GetShopData(context.Context, *Auth) (*ShopData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopData not implemented")
}
func (UnimplementedOZServer) GetShopWidget(context.Context, *Auth) (*ShopWidget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopWidget not implemented")
}
func (UnimplementedOZServer) GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedOZServer) GetProductsCount(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCount not implemented")
}
func (UnimplementedOZServer) GetProductsCountUnsorted(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCountUnsorted not implemented")
}
func (UnimplementedOZServer) GetFlowGraphicData(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlowGraphicData not implemented")
}
func (UnimplementedOZServer) GetMarginGraphicData(context.Context, *Auth) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarginGraphicData not implemented")
}
func (UnimplementedOZServer) GetWeekGraphic(context.Context, *LineGraphRequest) (*WeekGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekGraphic not implemented")
}
func (UnimplementedOZServer) GetPayRoundGraphic(context.Context, *Auth) (*RoundGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayRoundGraphic not implemented")
}
func (UnimplementedOZServer) GetCountRoundGraphic(context.Context, *Auth) (*RoundGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountRoundGraphic not implemented")
}
func (UnimplementedOZServer) GetOrderLeaders(context.Context, *Auth) (*OrderLeaders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderLeaders not implemented")
}
func (UnimplementedOZServer) GetShopProducts(context.Context, *RequestByIDs) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProducts not implemented")
}
func (UnimplementedOZServer) GetShopProduct(context.Context, *RequestByID) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProduct not implemented")
}
func (UnimplementedOZServer) GetProductGraphics(context.Context, *RequestByDates) (*LineGraphics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGraphics not implemented")
}
func (UnimplementedOZServer) GetSales(context.Context, *RequestByDates) (*Sales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSales not implemented")
}
func (UnimplementedOZServer) GetProductWidget(context.Context, *RequestByDates) (*ProductWidgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWidget not implemented")
}
func (UnimplementedOZServer) GetProductWidgetOrders(context.Context, *RequestByDates) (*ProductWidgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWidgetOrders not implemented")
}
func (UnimplementedOZServer) GetSale(context.Context, *SaleRequest) (*Sale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSale not implemented")
}
func (UnimplementedOZServer) GetProductsUnsortedList(context.Context, *Auth) (*ShopProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsUnsortedList not implemented")
}
func (UnimplementedOZServer) SetShopProductUrl(context.Context, *ShopProductUrlSetter) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetShopProductUrl not implemented")
}
func (UnimplementedOZServer) GetCounterParams(context.Context, *RequestByIDs) (*OZCounterParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCounterParams not implemented")
}
func (UnimplementedOZServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedOZServer) GetShopProductByCode(context.Context, *SearchRequest) (*ShopProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProductByCode not implemented")
}
func (UnimplementedOZServer) GetOrdersForBrand(context.Context, *OrdersRequest) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersForBrand not implemented")
}
func (UnimplementedOZServer) SetAuthData(context.Context, *OZAuthData) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuthData not implemented")
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
		FullMethod: "/cerasusV2.OZ/Ping",
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
		FullMethod: "/cerasusV2.OZ/GetAppData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetAppData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetShopData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetShopData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetShopData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetShopData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetShopWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetShopWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetShopWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetShopWidget(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetMainGraphic(ctx, req.(*LineGraphRequest))
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
		FullMethod: "/cerasusV2.OZ/GetProductsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductsCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetProductsCountUnsorted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetProductsCountUnsorted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetProductsCountUnsorted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductsCountUnsorted(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetFlowGraphicData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetFlowGraphicData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetFlowGraphicData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetFlowGraphicData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetMarginGraphicData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetMarginGraphicData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetMarginGraphicData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetMarginGraphicData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetWeekGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetWeekGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetWeekGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetWeekGraphic(ctx, req.(*LineGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetPayRoundGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetPayRoundGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetPayRoundGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetPayRoundGraphic(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetCountRoundGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetCountRoundGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetCountRoundGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetCountRoundGraphic(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetOrderLeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetOrderLeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetOrderLeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetOrderLeaders(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetShopProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetShopProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetShopProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetShopProducts(ctx, req.(*RequestByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetShopProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetShopProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetShopProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetShopProduct(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetProductGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetProductGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetProductGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductGraphics(ctx, req.(*RequestByDates))
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
		FullMethod: "/cerasusV2.OZ/GetSales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetSales(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetProductWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetProductWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetProductWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductWidget(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetProductWidgetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetProductWidgetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetProductWidgetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductWidgetOrders(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetSale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetSale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetSale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetSale(ctx, req.(*SaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetProductsUnsortedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetProductsUnsortedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetProductsUnsortedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetProductsUnsortedList(ctx, req.(*Auth))
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
		FullMethod: "/cerasusV2.OZ/SetShopProductUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).SetShopProductUrl(ctx, req.(*ShopProductUrlSetter))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetCounterParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetCounterParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetCounterParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetCounterParams(ctx, req.(*RequestByIDs))
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
		FullMethod: "/cerasusV2.OZ/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetShopProductByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetShopProductByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetShopProductByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetShopProductByCode(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_GetOrdersForBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).GetOrdersForBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/GetOrdersForBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).GetOrdersForBrand(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OZ_SetAuthData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OZAuthData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OZServer).SetAuthData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.OZ/SetAuthData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OZServer).SetAuthData(ctx, req.(*OZAuthData))
	}
	return interceptor(ctx, in, info, handler)
}

// OZ_ServiceDesc is the grpc.ServiceDesc for OZ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OZ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.OZ",
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
			MethodName: "GetShopData",
			Handler:    _OZ_GetShopData_Handler,
		},
		{
			MethodName: "GetShopWidget",
			Handler:    _OZ_GetShopWidget_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _OZ_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetProductsCount",
			Handler:    _OZ_GetProductsCount_Handler,
		},
		{
			MethodName: "GetProductsCountUnsorted",
			Handler:    _OZ_GetProductsCountUnsorted_Handler,
		},
		{
			MethodName: "GetFlowGraphicData",
			Handler:    _OZ_GetFlowGraphicData_Handler,
		},
		{
			MethodName: "GetMarginGraphicData",
			Handler:    _OZ_GetMarginGraphicData_Handler,
		},
		{
			MethodName: "GetWeekGraphic",
			Handler:    _OZ_GetWeekGraphic_Handler,
		},
		{
			MethodName: "GetPayRoundGraphic",
			Handler:    _OZ_GetPayRoundGraphic_Handler,
		},
		{
			MethodName: "GetCountRoundGraphic",
			Handler:    _OZ_GetCountRoundGraphic_Handler,
		},
		{
			MethodName: "GetOrderLeaders",
			Handler:    _OZ_GetOrderLeaders_Handler,
		},
		{
			MethodName: "GetShopProducts",
			Handler:    _OZ_GetShopProducts_Handler,
		},
		{
			MethodName: "GetShopProduct",
			Handler:    _OZ_GetShopProduct_Handler,
		},
		{
			MethodName: "GetProductGraphics",
			Handler:    _OZ_GetProductGraphics_Handler,
		},
		{
			MethodName: "GetSales",
			Handler:    _OZ_GetSales_Handler,
		},
		{
			MethodName: "GetProductWidget",
			Handler:    _OZ_GetProductWidget_Handler,
		},
		{
			MethodName: "GetProductWidgetOrders",
			Handler:    _OZ_GetProductWidgetOrders_Handler,
		},
		{
			MethodName: "GetSale",
			Handler:    _OZ_GetSale_Handler,
		},
		{
			MethodName: "GetProductsUnsortedList",
			Handler:    _OZ_GetProductsUnsortedList_Handler,
		},
		{
			MethodName: "SetShopProductUrl",
			Handler:    _OZ_SetShopProductUrl_Handler,
		},
		{
			MethodName: "GetCounterParams",
			Handler:    _OZ_GetCounterParams_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _OZ_GetImage_Handler,
		},
		{
			MethodName: "GetShopProductByCode",
			Handler:    _OZ_GetShopProductByCode_Handler,
		},
		{
			MethodName: "GetOrdersForBrand",
			Handler:    _OZ_GetOrdersForBrand_Handler,
		},
		{
			MethodName: "SetAuthData",
			Handler:    _OZ_SetAuthData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oz_v2.proto",
}
