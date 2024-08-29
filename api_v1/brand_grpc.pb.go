// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: brand.proto

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

// BrandsClient is the client API for Brands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandsClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetBrandLink(ctx context.Context, in *BrandLinkRequest, opts ...grpc.CallOption) (*BBrandLink, error)
}

type brandsClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandsClient(cc grpc.ClientConnInterface) BrandsClient {
	return &brandsClient{cc}
}

func (c *brandsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBrandLink(ctx context.Context, in *BrandLinkRequest, opts ...grpc.CallOption) (*BBrandLink, error) {
	out := new(BBrandLink)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBrandLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandsServer is the server API for Brands service.
// All implementations must embed UnimplementedBrandsServer
// for forward compatibility
type BrandsServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetBrandLink(context.Context, *BrandLinkRequest) (*BBrandLink, error)
	mustEmbedUnimplementedBrandsServer()
}

// UnimplementedBrandsServer must be embedded to have forward compatible implementations.
type UnimplementedBrandsServer struct {
}

func (UnimplementedBrandsServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBrandsServer) GetBrandLink(context.Context, *BrandLinkRequest) (*BBrandLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandLink not implemented")
}
func (UnimplementedBrandsServer) mustEmbedUnimplementedBrandsServer() {}

// UnsafeBrandsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandsServer will
// result in compilation errors.
type UnsafeBrandsServer interface {
	mustEmbedUnimplementedBrandsServer()
}

func RegisterBrandsServer(s grpc.ServiceRegistrar, srv BrandsServer) {
	s.RegisterService(&Brands_ServiceDesc, srv)
}

func _Brands_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBrandLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBrandLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBrandLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBrandLink(ctx, req.(*BrandLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Brands_ServiceDesc is the grpc.ServiceDesc for Brands service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Brands_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Brands",
	HandlerType: (*BrandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Brands_Ping_Handler,
		},
		{
			MethodName: "GetBrandLink",
			Handler:    _Brands_GetBrandLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brand.proto",
}
