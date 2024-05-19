// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: counter.proto

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

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterClient interface {
	GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ProductsCounter, error)
	GetProductCount(ctx context.Context, in *ProductCountRequest, opts ...grpc.CallOption) (*ProductCounter, error)
	SetProductCount(ctx context.Context, in *ProductCountSetter, opts ...grpc.CallOption) (*BoolReply, error)
	SysCounter(ctx context.Context, in *SysCounterRequest, opts ...grpc.CallOption) (*SysCounterReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type counterClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterClient(cc grpc.ClientConnInterface) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) GetProductsCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ProductsCounter, error) {
	out := new(ProductsCounter)
	err := c.cc.Invoke(ctx, "/cerasus.Counter/GetProductsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) GetProductCount(ctx context.Context, in *ProductCountRequest, opts ...grpc.CallOption) (*ProductCounter, error) {
	out := new(ProductCounter)
	err := c.cc.Invoke(ctx, "/cerasus.Counter/GetProductCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) SetProductCount(ctx context.Context, in *ProductCountSetter, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Counter/SetProductCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) SysCounter(ctx context.Context, in *SysCounterRequest, opts ...grpc.CallOption) (*SysCounterReply, error) {
	out := new(SysCounterReply)
	err := c.cc.Invoke(ctx, "/cerasus.Counter/SysCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Counter/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServer is the server API for Counter service.
// All implementations must embed UnimplementedCounterServer
// for forward compatibility
type CounterServer interface {
	GetProductsCount(context.Context, *Auth) (*ProductsCounter, error)
	GetProductCount(context.Context, *ProductCountRequest) (*ProductCounter, error)
	SetProductCount(context.Context, *ProductCountSetter) (*BoolReply, error)
	SysCounter(context.Context, *SysCounterRequest) (*SysCounterReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	mustEmbedUnimplementedCounterServer()
}

// UnimplementedCounterServer must be embedded to have forward compatible implementations.
type UnimplementedCounterServer struct {
}

func (UnimplementedCounterServer) GetProductsCount(context.Context, *Auth) (*ProductsCounter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCount not implemented")
}
func (UnimplementedCounterServer) GetProductCount(context.Context, *ProductCountRequest) (*ProductCounter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductCount not implemented")
}
func (UnimplementedCounterServer) SetProductCount(context.Context, *ProductCountSetter) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProductCount not implemented")
}
func (UnimplementedCounterServer) SysCounter(context.Context, *SysCounterRequest) (*SysCounterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysCounter not implemented")
}
func (UnimplementedCounterServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCounterServer) mustEmbedUnimplementedCounterServer() {}

// UnsafeCounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServer will
// result in compilation errors.
type UnsafeCounterServer interface {
	mustEmbedUnimplementedCounterServer()
}

func RegisterCounterServer(s grpc.ServiceRegistrar, srv CounterServer) {
	s.RegisterService(&Counter_ServiceDesc, srv)
}

func _Counter_GetProductsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).GetProductsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Counter/GetProductsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).GetProductsCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_GetProductCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).GetProductCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Counter/GetProductCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).GetProductCount(ctx, req.(*ProductCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_SetProductCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCountSetter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).SetProductCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Counter/SetProductCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).SetProductCount(ctx, req.(*ProductCountSetter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_SysCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).SysCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Counter/SysCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).SysCounter(ctx, req.(*SysCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Counter/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Counter_ServiceDesc is the grpc.ServiceDesc for Counter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Counter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductsCount",
			Handler:    _Counter_GetProductsCount_Handler,
		},
		{
			MethodName: "GetProductCount",
			Handler:    _Counter_GetProductCount_Handler,
		},
		{
			MethodName: "SetProductCount",
			Handler:    _Counter_SetProductCount_Handler,
		},
		{
			MethodName: "SysCounter",
			Handler:    _Counter_SysCounter_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Counter_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
