// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: settings_v2.proto

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

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetUserAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*UserAppData, error)
	GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error)
	GetFlowGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error)
	GetMarginGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error)
	GetWeekGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*WeekGraphic, error)
	GetOrderLeaders(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*OrderLeaders, error)
	GetCompanyBrands(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Brands, error)
	GetBrand(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Brand, error)
	GetProductGraphics(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*LineGraphics, error)
	GetTaxes(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Taxes, error)
	GetMarginLevels(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*MarginLevels, error)
	GetProductWidget(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error)
	GetProductWidgetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error)
	SetGeoPlace(ctx context.Context, in *SetGeoPlaceRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetCompanyShops(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShops, error)
	GetMargin(ctx context.Context, in *GetMarginRequest, opts ...grpc.CallOption) (*MarginSettings, error)
	SetMargin(ctx context.Context, in *SetMarginRequest, opts ...grpc.CallOption) (*StatusReply, error)
	DeleteMargin(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*StatusReply, error)
	GetAppTaxes(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppTaxes, error)
	SetTax(ctx context.Context, in *SetTaxRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error)
	SearchBrand(ctx context.Context, in *SearchBrandRequest, opts ...grpc.CallOption) (*Brands, error)
	SetBrand(ctx context.Context, in *SetBrandRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetUserAppData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*UserAppData, error) {
	out := new(UserAppData)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetUserAppData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetMainGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*LineGraph, error) {
	out := new(LineGraph)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetMainGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetFlowGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error) {
	out := new(RoundGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetFlowGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetMarginGraphic(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*RoundGraphic, error) {
	out := new(RoundGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetMarginGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetWeekGraphic(ctx context.Context, in *LineGraphRequest, opts ...grpc.CallOption) (*WeekGraphic, error) {
	out := new(WeekGraphic)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetWeekGraphic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetOrderLeaders(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*OrderLeaders, error) {
	out := new(OrderLeaders)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetOrderLeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetCompanyBrands(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Brands, error) {
	out := new(Brands)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetCompanyBrands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetBrand(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetProductGraphics(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*LineGraphics, error) {
	out := new(LineGraphics)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetProductGraphics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetTaxes(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Taxes, error) {
	out := new(Taxes)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetTaxes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetMarginLevels(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*MarginLevels, error) {
	out := new(MarginLevels)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetMarginLevels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetProductWidget(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error) {
	out := new(ProductWidgets)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetProductWidget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetProductWidgetOrders(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*ProductWidgets, error) {
	out := new(ProductWidgets)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetProductWidgetOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetGeoPlace(ctx context.Context, in *SetGeoPlaceRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/SetGeoPlace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetCompanyShops(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyShops, error) {
	out := new(CompanyShops)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetCompanyShops", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetMargin(ctx context.Context, in *GetMarginRequest, opts ...grpc.CallOption) (*MarginSettings, error) {
	out := new(MarginSettings)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetMargin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetMargin(ctx context.Context, in *SetMarginRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/SetMargin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) DeleteMargin(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/DeleteMargin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetAppTaxes(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*AppTaxes, error) {
	out := new(AppTaxes)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetAppTaxes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetTax(ctx context.Context, in *SetTaxRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/SetTax", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetImage(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SearchBrand(ctx context.Context, in *SearchBrandRequest, opts ...grpc.CallOption) (*Brands, error) {
	out := new(Brands)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/SearchBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetBrand(ctx context.Context, in *SetBrandRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Settings/SetBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
// All implementations must embed UnimplementedSettingsServer
// for forward compatibility
type SettingsServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetUserAppData(context.Context, *Auth) (*UserAppData, error)
	GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error)
	GetFlowGraphic(context.Context, *Auth) (*RoundGraphic, error)
	GetMarginGraphic(context.Context, *Auth) (*RoundGraphic, error)
	GetWeekGraphic(context.Context, *LineGraphRequest) (*WeekGraphic, error)
	GetOrderLeaders(context.Context, *Auth) (*OrderLeaders, error)
	GetCompanyBrands(context.Context, *Auth) (*Brands, error)
	GetBrand(context.Context, *RequestByID) (*Brand, error)
	GetProductGraphics(context.Context, *RequestByDates) (*LineGraphics, error)
	GetTaxes(context.Context, *Auth) (*Taxes, error)
	GetMarginLevels(context.Context, *Auth) (*MarginLevels, error)
	GetProductWidget(context.Context, *RequestByDates) (*ProductWidgets, error)
	GetProductWidgetOrders(context.Context, *RequestByDates) (*ProductWidgets, error)
	SetGeoPlace(context.Context, *SetGeoPlaceRequest) (*StatusReply, error)
	GetCompanyShops(context.Context, *Auth) (*CompanyShops, error)
	GetMargin(context.Context, *GetMarginRequest) (*MarginSettings, error)
	SetMargin(context.Context, *SetMarginRequest) (*StatusReply, error)
	DeleteMargin(context.Context, *RequestByID) (*StatusReply, error)
	GetAppTaxes(context.Context, *Auth) (*AppTaxes, error)
	SetTax(context.Context, *SetTaxRequest) (*StatusReply, error)
	GetImage(context.Context, *ImageRequest) (*ImageReply, error)
	SearchBrand(context.Context, *SearchBrandRequest) (*Brands, error)
	SetBrand(context.Context, *SetBrandRequest) (*StatusReply, error)
	mustEmbedUnimplementedSettingsServer()
}

// UnimplementedSettingsServer must be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (UnimplementedSettingsServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSettingsServer) GetUserAppData(context.Context, *Auth) (*UserAppData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAppData not implemented")
}
func (UnimplementedSettingsServer) GetMainGraphic(context.Context, *LineGraphRequest) (*LineGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainGraphic not implemented")
}
func (UnimplementedSettingsServer) GetFlowGraphic(context.Context, *Auth) (*RoundGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlowGraphic not implemented")
}
func (UnimplementedSettingsServer) GetMarginGraphic(context.Context, *Auth) (*RoundGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarginGraphic not implemented")
}
func (UnimplementedSettingsServer) GetWeekGraphic(context.Context, *LineGraphRequest) (*WeekGraphic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeekGraphic not implemented")
}
func (UnimplementedSettingsServer) GetOrderLeaders(context.Context, *Auth) (*OrderLeaders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderLeaders not implemented")
}
func (UnimplementedSettingsServer) GetCompanyBrands(context.Context, *Auth) (*Brands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyBrands not implemented")
}
func (UnimplementedSettingsServer) GetBrand(context.Context, *RequestByID) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedSettingsServer) GetProductGraphics(context.Context, *RequestByDates) (*LineGraphics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGraphics not implemented")
}
func (UnimplementedSettingsServer) GetTaxes(context.Context, *Auth) (*Taxes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaxes not implemented")
}
func (UnimplementedSettingsServer) GetMarginLevels(context.Context, *Auth) (*MarginLevels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarginLevels not implemented")
}
func (UnimplementedSettingsServer) GetProductWidget(context.Context, *RequestByDates) (*ProductWidgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWidget not implemented")
}
func (UnimplementedSettingsServer) GetProductWidgetOrders(context.Context, *RequestByDates) (*ProductWidgets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductWidgetOrders not implemented")
}
func (UnimplementedSettingsServer) SetGeoPlace(context.Context, *SetGeoPlaceRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGeoPlace not implemented")
}
func (UnimplementedSettingsServer) GetCompanyShops(context.Context, *Auth) (*CompanyShops, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyShops not implemented")
}
func (UnimplementedSettingsServer) GetMargin(context.Context, *GetMarginRequest) (*MarginSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMargin not implemented")
}
func (UnimplementedSettingsServer) SetMargin(context.Context, *SetMarginRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMargin not implemented")
}
func (UnimplementedSettingsServer) DeleteMargin(context.Context, *RequestByID) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMargin not implemented")
}
func (UnimplementedSettingsServer) GetAppTaxes(context.Context, *Auth) (*AppTaxes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppTaxes not implemented")
}
func (UnimplementedSettingsServer) SetTax(context.Context, *SetTaxRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTax not implemented")
}
func (UnimplementedSettingsServer) GetImage(context.Context, *ImageRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedSettingsServer) SearchBrand(context.Context, *SearchBrandRequest) (*Brands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBrand not implemented")
}
func (UnimplementedSettingsServer) SetBrand(context.Context, *SetBrandRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBrand not implemented")
}
func (UnimplementedSettingsServer) mustEmbedUnimplementedSettingsServer() {}

// UnsafeSettingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServer will
// result in compilation errors.
type UnsafeSettingsServer interface {
	mustEmbedUnimplementedSettingsServer()
}

func RegisterSettingsServer(s grpc.ServiceRegistrar, srv SettingsServer) {
	s.RegisterService(&Settings_ServiceDesc, srv)
}

func _Settings_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetUserAppData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetUserAppData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetUserAppData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetUserAppData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetMainGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetMainGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetMainGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetMainGraphic(ctx, req.(*LineGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetFlowGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetFlowGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetFlowGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetFlowGraphic(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetMarginGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetMarginGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetMarginGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetMarginGraphic(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetWeekGraphic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LineGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetWeekGraphic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetWeekGraphic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetWeekGraphic(ctx, req.(*LineGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetOrderLeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetOrderLeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetOrderLeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetOrderLeaders(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetCompanyBrands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetCompanyBrands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetCompanyBrands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetCompanyBrands(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetBrand(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetProductGraphics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetProductGraphics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetProductGraphics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetProductGraphics(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetTaxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetTaxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetTaxes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetTaxes(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetMarginLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetMarginLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetMarginLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetMarginLevels(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetProductWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetProductWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetProductWidget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetProductWidget(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetProductWidgetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetProductWidgetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetProductWidgetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetProductWidgetOrders(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetGeoPlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGeoPlaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetGeoPlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/SetGeoPlace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetGeoPlace(ctx, req.(*SetGeoPlaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetCompanyShops_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetCompanyShops(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetCompanyShops",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetCompanyShops(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetMargin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetMargin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetMargin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetMargin(ctx, req.(*GetMarginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetMargin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMarginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetMargin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/SetMargin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetMargin(ctx, req.(*SetMarginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_DeleteMargin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).DeleteMargin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/DeleteMargin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).DeleteMargin(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetAppTaxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetAppTaxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetAppTaxes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetAppTaxes(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetTax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTaxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetTax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/SetTax",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetTax(ctx, req.(*SetTaxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetImage(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SearchBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SearchBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/SearchBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SearchBrand(ctx, req.(*SearchBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Settings/SetBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetBrand(ctx, req.(*SetBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Settings_ServiceDesc is the grpc.ServiceDesc for Settings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Settings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Settings_Ping_Handler,
		},
		{
			MethodName: "GetUserAppData",
			Handler:    _Settings_GetUserAppData_Handler,
		},
		{
			MethodName: "GetMainGraphic",
			Handler:    _Settings_GetMainGraphic_Handler,
		},
		{
			MethodName: "GetFlowGraphic",
			Handler:    _Settings_GetFlowGraphic_Handler,
		},
		{
			MethodName: "GetMarginGraphic",
			Handler:    _Settings_GetMarginGraphic_Handler,
		},
		{
			MethodName: "GetWeekGraphic",
			Handler:    _Settings_GetWeekGraphic_Handler,
		},
		{
			MethodName: "GetOrderLeaders",
			Handler:    _Settings_GetOrderLeaders_Handler,
		},
		{
			MethodName: "GetCompanyBrands",
			Handler:    _Settings_GetCompanyBrands_Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _Settings_GetBrand_Handler,
		},
		{
			MethodName: "GetProductGraphics",
			Handler:    _Settings_GetProductGraphics_Handler,
		},
		{
			MethodName: "GetTaxes",
			Handler:    _Settings_GetTaxes_Handler,
		},
		{
			MethodName: "GetMarginLevels",
			Handler:    _Settings_GetMarginLevels_Handler,
		},
		{
			MethodName: "GetProductWidget",
			Handler:    _Settings_GetProductWidget_Handler,
		},
		{
			MethodName: "GetProductWidgetOrders",
			Handler:    _Settings_GetProductWidgetOrders_Handler,
		},
		{
			MethodName: "SetGeoPlace",
			Handler:    _Settings_SetGeoPlace_Handler,
		},
		{
			MethodName: "GetCompanyShops",
			Handler:    _Settings_GetCompanyShops_Handler,
		},
		{
			MethodName: "GetMargin",
			Handler:    _Settings_GetMargin_Handler,
		},
		{
			MethodName: "SetMargin",
			Handler:    _Settings_SetMargin_Handler,
		},
		{
			MethodName: "DeleteMargin",
			Handler:    _Settings_DeleteMargin_Handler,
		},
		{
			MethodName: "GetAppTaxes",
			Handler:    _Settings_GetAppTaxes_Handler,
		},
		{
			MethodName: "SetTax",
			Handler:    _Settings_SetTax_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Settings_GetImage_Handler,
		},
		{
			MethodName: "SearchBrand",
			Handler:    _Settings_SearchBrand_Handler,
		},
		{
			MethodName: "SetBrand",
			Handler:    _Settings_SetBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings_v2.proto",
}
