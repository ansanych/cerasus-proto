// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pricer_v2.proto

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
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetParamsForCounter(ctx context.Context, in *Company, opts ...grpc.CallOption) (*ParamsForCounter, error)
	GetParserData(ctx context.Context, in *ParserGetRequest, opts ...grpc.CallOption) (*ParserJob, error)
	SetParserData(ctx context.Context, in *ParserSetRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetDetectorData(ctx context.Context, in *DetectorGetRequest, opts ...grpc.CallOption) (*DetectorGetReply, error)
	SetDetectorData(ctx context.Context, in *DetectorSetRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type pricerClient struct {
	cc grpc.ClientConnInterface
}

func NewPricerClient(cc grpc.ClientConnInterface) PricerClient {
	return &pricerClient{cc}
}

func (c *pricerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Pricer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetParamsForCounter(ctx context.Context, in *Company, opts ...grpc.CallOption) (*ParamsForCounter, error) {
	out := new(ParamsForCounter)
	err := c.cc.Invoke(ctx, "/cerasusV2.Pricer/GetParamsForCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetParserData(ctx context.Context, in *ParserGetRequest, opts ...grpc.CallOption) (*ParserJob, error) {
	out := new(ParserJob)
	err := c.cc.Invoke(ctx, "/cerasusV2.Pricer/GetParserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) SetParserData(ctx context.Context, in *ParserSetRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Pricer/SetParserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) GetDetectorData(ctx context.Context, in *DetectorGetRequest, opts ...grpc.CallOption) (*DetectorGetReply, error) {
	out := new(DetectorGetReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Pricer/GetDetectorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricerClient) SetDetectorData(ctx context.Context, in *DetectorSetRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Pricer/SetDetectorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricerServer is the server API for Pricer service.
// All implementations must embed UnimplementedPricerServer
// for forward compatibility
type PricerServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetParamsForCounter(context.Context, *Company) (*ParamsForCounter, error)
	GetParserData(context.Context, *ParserGetRequest) (*ParserJob, error)
	SetParserData(context.Context, *ParserSetRequest) (*StatusReply, error)
	GetDetectorData(context.Context, *DetectorGetRequest) (*DetectorGetReply, error)
	SetDetectorData(context.Context, *DetectorSetRequest) (*StatusReply, error)
	mustEmbedUnimplementedPricerServer()
}

// UnimplementedPricerServer must be embedded to have forward compatible implementations.
type UnimplementedPricerServer struct {
}

func (UnimplementedPricerServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPricerServer) GetParamsForCounter(context.Context, *Company) (*ParamsForCounter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParamsForCounter not implemented")
}
func (UnimplementedPricerServer) GetParserData(context.Context, *ParserGetRequest) (*ParserJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParserData not implemented")
}
func (UnimplementedPricerServer) SetParserData(context.Context, *ParserSetRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetParserData not implemented")
}
func (UnimplementedPricerServer) GetDetectorData(context.Context, *DetectorGetRequest) (*DetectorGetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetectorData not implemented")
}
func (UnimplementedPricerServer) SetDetectorData(context.Context, *DetectorSetRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDetectorData not implemented")
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
		FullMethod: "/cerasusV2.Pricer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetParamsForCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetParamsForCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Pricer/GetParamsForCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetParamsForCounter(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetParserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetParserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Pricer/GetParserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetParserData(ctx, req.(*ParserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_SetParserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).SetParserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Pricer/SetParserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).SetParserData(ctx, req.(*ParserSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_GetDetectorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectorGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).GetDetectorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Pricer/GetDetectorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).GetDetectorData(ctx, req.(*DetectorGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pricer_SetDetectorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricerServer).SetDetectorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Pricer/SetDetectorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricerServer).SetDetectorData(ctx, req.(*DetectorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pricer_ServiceDesc is the grpc.ServiceDesc for Pricer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pricer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.Pricer",
	HandlerType: (*PricerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Pricer_Ping_Handler,
		},
		{
			MethodName: "GetParamsForCounter",
			Handler:    _Pricer_GetParamsForCounter_Handler,
		},
		{
			MethodName: "GetParserData",
			Handler:    _Pricer_GetParserData_Handler,
		},
		{
			MethodName: "SetParserData",
			Handler:    _Pricer_SetParserData_Handler,
		},
		{
			MethodName: "GetDetectorData",
			Handler:    _Pricer_GetDetectorData_Handler,
		},
		{
			MethodName: "SetDetectorData",
			Handler:    _Pricer_SetDetectorData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricer_v2.proto",
}
