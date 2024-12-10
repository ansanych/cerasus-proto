// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: products_v2.proto

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

// ProductsV2Client is the client API for ProductsV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsV2Client interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetProductByID(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Product, error)
	GetProductsByIDs(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*Products, error)
	GetProductsByShopIDs(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*Products, error)
}

type productsV2Client struct {
	cc grpc.ClientConnInterface
}

func NewProductsV2Client(cc grpc.ClientConnInterface) ProductsV2Client {
	return &productsV2Client{cc}
}

func (c *productsV2Client) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.ProductsV2/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsV2Client) GetProductByID(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/cerasusV2.ProductsV2/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsV2Client) GetProductsByIDs(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/cerasusV2.ProductsV2/GetProductsByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsV2Client) GetProductsByShopIDs(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, "/cerasusV2.ProductsV2/GetProductsByShopIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsV2Server is the server API for ProductsV2 service.
// All implementations must embed UnimplementedProductsV2Server
// for forward compatibility
type ProductsV2Server interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetProductByID(context.Context, *RequestByID) (*Product, error)
	GetProductsByIDs(context.Context, *RequestByIDs) (*Products, error)
	GetProductsByShopIDs(context.Context, *RequestByIDs) (*Products, error)
	mustEmbedUnimplementedProductsV2Server()
}

// UnimplementedProductsV2Server must be embedded to have forward compatible implementations.
type UnimplementedProductsV2Server struct {
}

func (UnimplementedProductsV2Server) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProductsV2Server) GetProductByID(context.Context, *RequestByID) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedProductsV2Server) GetProductsByIDs(context.Context, *RequestByIDs) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsByIDs not implemented")
}
func (UnimplementedProductsV2Server) GetProductsByShopIDs(context.Context, *RequestByIDs) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsByShopIDs not implemented")
}
func (UnimplementedProductsV2Server) mustEmbedUnimplementedProductsV2Server() {}

// UnsafeProductsV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsV2Server will
// result in compilation errors.
type UnsafeProductsV2Server interface {
	mustEmbedUnimplementedProductsV2Server()
}

func RegisterProductsV2Server(s grpc.ServiceRegistrar, srv ProductsV2Server) {
	s.RegisterService(&ProductsV2_ServiceDesc, srv)
}

func _ProductsV2_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsV2Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.ProductsV2/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsV2Server).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsV2_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsV2Server).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.ProductsV2/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsV2Server).GetProductByID(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsV2_GetProductsByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsV2Server).GetProductsByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.ProductsV2/GetProductsByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsV2Server).GetProductsByIDs(ctx, req.(*RequestByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsV2_GetProductsByShopIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsV2Server).GetProductsByShopIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.ProductsV2/GetProductsByShopIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsV2Server).GetProductsByShopIDs(ctx, req.(*RequestByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductsV2_ServiceDesc is the grpc.ServiceDesc for ProductsV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductsV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.ProductsV2",
	HandlerType: (*ProductsV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ProductsV2_Ping_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _ProductsV2_GetProductByID_Handler,
		},
		{
			MethodName: "GetProductsByIDs",
			Handler:    _ProductsV2_GetProductsByIDs_Handler,
		},
		{
			MethodName: "GetProductsByShopIDs",
			Handler:    _ProductsV2_GetProductsByShopIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products_v2.proto",
}
