// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pricer.proto

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

// PricerClient is the client API for Pricer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PricerClient interface {
	SetSystemSettings(ctx context.Context, in *PricerSystemSettings, opts ...grpc.CallOption) (*BoolReply, error)
	GetSystemSettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*PricerSystemSettings, error)
	SetCompanySettings(ctx context.Context, in *PricerCompanySettings, opts ...grpc.CallOption) (*BoolReply, error)
	GetCompanySettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*PricerCompanySettingsReply, error)
	SetCompanyParams(ctx context.Context, in *SetPricerParams, opts ...grpc.CallOption) (*BoolReply, error)
	GetCompanyParams(ctx context.Context, in *CompanyParamsRequest, opts ...grpc.CallOption) (*PricerParams, error)
	GetCompanyParamProduct(ctx context.Context, in *CompanyParamProductRequest, opts ...grpc.CallOption) (*PricerParam, error)
	GetPricerItem(ctx context.Context, in *GetPricerItemRequest, opts ...grpc.CallOption) (*GetPricerItemReply, error)
	SetPricerItem(ctx context.Context, in *SetPricerItemRequest, opts ...grpc.CallOption) (*BoolReply, error)
	GetProductsInPricer(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ProductsInPricer, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetForDumping(ctx context.Context, in *ForDumpingDataRequest, opts ...grpc.CallOption) (*ForDumpingDataReply, error)
}

type pricerClient struct {
	cc grpc.ClientConnInterface
}

func NewPricerClient(cc grpc.ClientConnInterface) PricerClient {
	return &pricerClient{cc}
}

func (c *pricerClient) SetSystemSettings(ctx context.Context, in *PricerSystemSettings, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/SetSystemSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetSystemSettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*PricerSystemSettings, error) {
	out := new(PricerSystemSettings)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetSystemSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) SetCompanySettings(ctx context.Context, in *PricerCompanySettings, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/SetCompanySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetCompanySettings(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*PricerCompanySettingsReply, error) {
	out := new(PricerCompanySettingsReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetCompanySettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) SetCompanyParams(ctx context.Context, in *SetPricerParams, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/SetCompanyParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetCompanyParams(ctx context.Context, in *CompanyParamsRequest, opts ...grpc.CallOption) (*PricerParams, error) {
	out := new(PricerParams)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetCompanyParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetCompanyParamProduct(ctx context.Context, in *CompanyParamProductRequest, opts ...grpc.CallOption) (*PricerParam, error) {
	out := new(PricerParam)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetCompanyParamProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetPricerItem(ctx context.Context, in *GetPricerItemRequest, opts ...grpc.CallOption) (*GetPricerItemReply, error) {
	out := new(GetPricerItemReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetPricerItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) SetPricerItem(ctx context.Context, in *SetPricerItemRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/SetPricerItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetProductsInPricer(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ProductsInPricer, error) {
	out := new(ProductsInPricer)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetProductsInPricer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetForDumping(ctx context.Context, in *ForDumpingDataRequest, opts ...grpc.CallOption) (*ForDumpingDataReply, error) {
	out := new(ForDumpingDataReply)
	err := c.cc.Invoke(ctx, "/cerasus.Pricer/GetForDumping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricerServer is the server API for Pricer service.
// All implementations must embed UnimplementedPricerServer
// for forward compatibility
type PricerServer interface {
	SetSystemSettings(context.Context, *PricerSystemSettings) (*BoolReply, error)
	GetSystemSettings(context.Context, *Auth) (*PricerSystemSettings, error)
	SetCompanySettings(context.Context, *PricerCompanySettings) (*BoolReply, error)
	GetCompanySettings(context.Context, *Auth) (*PricerCompanySettingsReply, error)
	SetCompanyParams(context.Context, *SetPricerParams) (*BoolReply, error)
	GetCompanyParams(context.Context, *CompanyParamsRequest) (*PricerParams, error)
	GetCompanyParamProduct(context.Context, *CompanyParamProductRequest) (*PricerParam, error)
	GetPricerItem(context.Context, *GetPricerItemRequest) (*GetPricerItemReply, error)
	SetPricerItem(context.Context, *SetPricerItemRequest) (*BoolReply, error)
	GetProductsInPricer(context.Context, *Auth) (*ProductsInPricer, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetForDumping(context.Context, *ForDumpingDataRequest) (*ForDumpingDataReply, error)
	mustEmbedUnimplementedPricerServer()
}

// UnimplementedPricerServer must be embedded to have forward compatible implementations.
type UnimplementedPricerServer struct {
}

func (UnimplementedPricerServer) SetSystemSettings(context.Context, *PricerSystemSettings) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSystemSettings not implemented")
}
func (UnimplementedPricerServer) GetSystemSettings(context.Context, *Auth) (*PricerSystemSettings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemSettings not implemented")
}
func (UnimplementedPricerServer) SetCompanySettings(context.Context, *PricerCompanySettings) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCompanySettings not implemented")
}
func (UnimplementedPricerServer) GetCompanySettings(context.Context, *Auth) (*PricerCompanySettingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanySettings not implemented")
}
func (UnimplementedPricerServer) SetCompanyParams(context.Context, *SetPricerParams) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCompanyParams not implemented")
}
func (UnimplementedPricerServer) GetCompanyParams(context.Context, *CompanyParamsRequest) (*PricerParams, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyParams not implemented")
}
func (UnimplementedPricerServer) GetCompanyParamProduct(context.Context, *CompanyParamProductRequest) (*PricerParam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyParamProduct not implemented")
}
func (UnimplementedPricerServer) GetPricerItem(context.Context, *GetPricerItemRequest) (*GetPricerItemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPricerItem not implemented")
}
func (UnimplementedPricerServer) SetPricerItem(context.Context, *SetPricerItemRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPricerItem not implemented")
}
func (UnimplementedPricerServer) GetProductsInPricer(context.Context, *Auth) (*ProductsInPricer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsInPricer not implemented")
}
func (UnimplementedPricerServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPricerServer) GetForDumping(context.Context, *ForDumpingDataRequest) (*ForDumpingDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForDumping not implemented")
}
func (UnimplementedPricerServer) mustEmbedUnimplementedPricerServer() {}

// UnsafePricerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PricerServer will
// result in compilation errors.
type UnsafePricerServer interface {
	mustEmbedUnimplementedPricerServer()
}

func RegisterPricerServer(s grpc.ServiceRegistrar, srv PricerServer) {
	s.RegisterService(&Pricer_ServiceDesc, srv)
}

func _Pricer_SetSystemSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricerSystemSettings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).SetSystemSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/SetSystemSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).SetSystemSettings(ctx, req.(*PricerSystemSettings))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetSystemSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetSystemSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetSystemSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetSystemSettings(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_SetCompanySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricerCompanySettings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).SetCompanySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/SetCompanySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).SetCompanySettings(ctx, req.(*PricerCompanySettings))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetCompanySettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetCompanySettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetCompanySettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetCompanySettings(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_SetCompanyParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPricerParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).SetCompanyParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/SetCompanyParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).SetCompanyParams(ctx, req.(*SetPricerParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetCompanyParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetCompanyParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetCompanyParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetCompanyParams(ctx, req.(*CompanyParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetCompanyParamProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyParamProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetCompanyParamProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetCompanyParamProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetCompanyParamProduct(ctx, req.(*CompanyParamProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetPricerItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPricerItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetPricerItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetPricerItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetPricerItem(ctx, req.(*GetPricerItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_SetPricerItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPricerItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).SetPricerItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/SetPricerItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).SetPricerItem(ctx, req.(*SetPricerItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetProductsInPricer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetProductsInPricer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetProductsInPricer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetProductsInPricer(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetForDumping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForDumpingDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetForDumping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Pricer/GetForDumping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetForDumping(ctx, req.(*ForDumpingDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pricer_ServiceDesc is the grpc.ServiceDesc for Pricer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pricer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Pricer",
	HandlerType: (*PricerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetSystemSettings",
			Handler:    _Pricer_SetSystemSettings_Handler,
		},
		{
			MethodName: "GetSystemSettings",
			Handler:    _Pricer_GetSystemSettings_Handler,
		},
		{
			MethodName: "SetCompanySettings",
			Handler:    _Pricer_SetCompanySettings_Handler,
		},
		{
			MethodName: "GetCompanySettings",
			Handler:    _Pricer_GetCompanySettings_Handler,
		},
		{
			MethodName: "SetCompanyParams",
			Handler:    _Pricer_SetCompanyParams_Handler,
		},
		{
			MethodName: "GetCompanyParams",
			Handler:    _Pricer_GetCompanyParams_Handler,
		},
		{
			MethodName: "GetCompanyParamProduct",
			Handler:    _Pricer_GetCompanyParamProduct_Handler,
		},
		{
			MethodName: "GetPricerItem",
			Handler:    _Pricer_GetPricerItem_Handler,
		},
		{
			MethodName: "SetPricerItem",
			Handler:    _Pricer_SetPricerItem_Handler,
		},
		{
			MethodName: "GetProductsInPricer",
			Handler:    _Pricer_GetProductsInPricer_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Pricer_Ping_Handler,
		},
		{
			MethodName: "GetForDumping",
			Handler:    _Pricer_GetForDumping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricer.proto",
}
