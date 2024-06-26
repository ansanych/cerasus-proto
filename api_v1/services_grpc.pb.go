// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: services.proto

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

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	GetServices(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ServicesReply, error)
	GetCompanyServices(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyServicesReply, error)
	GetCompaniesServicesList(ctx context.Context, in *CompaniesServiesListRequest, opts ...grpc.CallOption) (*CompaniesServiesListReply, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) GetServices(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*ServicesReply, error) {
	out := new(ServicesReply)
	err := c.cc.Invoke(ctx, "/cerasus.Services/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GetCompanyServices(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*CompanyServicesReply, error) {
	out := new(CompanyServicesReply)
	err := c.cc.Invoke(ctx, "/cerasus.Services/GetCompanyServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GetCompaniesServicesList(ctx context.Context, in *CompaniesServiesListRequest, opts ...grpc.CallOption) (*CompaniesServiesListReply, error) {
	out := new(CompaniesServiesListReply)
	err := c.cc.Invoke(ctx, "/cerasus.Services/GetCompaniesServicesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Services/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServer is the server API for Services service.
// All implementations must embed UnimplementedServicesServer
// for forward compatibility
type ServicesServer interface {
	GetServices(context.Context, *Auth) (*ServicesReply, error)
	GetCompanyServices(context.Context, *Auth) (*CompanyServicesReply, error)
	GetCompaniesServicesList(context.Context, *CompaniesServiesListRequest) (*CompaniesServiesListReply, error)
	Ping(context.Context, *PingRequest) (*PingReply, error)
	mustEmbedUnimplementedServicesServer()
}

// UnimplementedServicesServer must be embedded to have forward compatible implementations.
type UnimplementedServicesServer struct {
}

func (UnimplementedServicesServer) GetServices(context.Context, *Auth) (*ServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedServicesServer) GetCompanyServices(context.Context, *Auth) (*CompanyServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyServices not implemented")
}
func (UnimplementedServicesServer) GetCompaniesServicesList(context.Context, *CompaniesServiesListRequest) (*CompaniesServiesListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompaniesServicesList not implemented")
}
func (UnimplementedServicesServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedServicesServer) mustEmbedUnimplementedServicesServer() {}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Services/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GetServices(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GetCompanyServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GetCompanyServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Services/GetCompanyServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GetCompanyServices(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GetCompaniesServicesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompaniesServiesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GetCompaniesServicesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Services/GetCompaniesServicesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GetCompaniesServicesList(ctx, req.(*CompaniesServiesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Services/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _Services_GetServices_Handler,
		},
		{
			MethodName: "GetCompanyServices",
			Handler:    _Services_GetCompanyServices_Handler,
		},
		{
			MethodName: "GetCompaniesServicesList",
			Handler:    _Services_GetCompaniesServicesList_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Services_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
