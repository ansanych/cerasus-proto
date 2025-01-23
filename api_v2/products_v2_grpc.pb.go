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

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetProductsCount(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*Count, error)
	GetProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductList, error)
	GetProductSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ProductList, error)
	GetProductsByShopIDs(ctx context.Context, in *RequestByShopIDs, opts ...grpc.CallOption) (*ProductList, error)
	GetProduct(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Product, error)
	GetShopProductsMatch(ctx context.Context, in *RequestByShopIDs, opts ...grpc.CallOption) (*ShopProductsMatches, error)
	GetPurchases(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*ProductsPurchases, error)
	GetProductsUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*UnsortedCount, error)
	ConnectUnsorted(ctx context.Context, in *ConnectUnsortedRequest, opts ...grpc.CallOption) (*StatusReply, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetProductPurchase(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*ProductPurchases, error)
	SetProductPurchase(ctx context.Context, in *SetProductPurchaseRequest, opts ...grpc.CallOption) (*StatusReply, error)
	DeleteProductPurchase(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*StatusReply, error)
	SetProductBrand(ctx context.Context, in *SetProductBrandRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProductsCount(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProductsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProductList(ctx context.Context, in *ProductListRequest, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProductSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProductSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProductsByShopIDs(ctx context.Context, in *RequestByShopIDs, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProductsByShopIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProduct(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetShopProductsMatch(ctx context.Context, in *RequestByShopIDs, opts ...grpc.CallOption) (*ShopProductsMatches, error) {
	out := new(ShopProductsMatches)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetShopProductsMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetPurchases(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*ProductsPurchases, error) {
	out := new(ProductsPurchases)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetPurchases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProductsUnsortedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*UnsortedCount, error) {
	out := new(UnsortedCount)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProductsUnsortedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) ConnectUnsorted(ctx context.Context, in *ConnectUnsortedRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/ConnectUnsorted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetProductPurchase(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*ProductPurchases, error) {
	out := new(ProductPurchases)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/GetProductPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) SetProductPurchase(ctx context.Context, in *SetProductPurchaseRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/SetProductPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) DeleteProductPurchase(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/DeleteProductPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) SetProductBrand(ctx context.Context, in *SetProductBrandRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Products/SetProductBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
// All implementations must embed UnimplementedProductsServer
// for forward compatibility
type ProductsServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetProductsCount(context.Context, *ProductListRequest) (*Count, error)
	GetProductList(context.Context, *ProductListRequest) (*ProductList, error)
	GetProductSearch(context.Context, *SearchRequest) (*ProductList, error)
	GetProductsByShopIDs(context.Context, *RequestByShopIDs) (*ProductList, error)
	GetProduct(context.Context, *RequestByID) (*Product, error)
	GetShopProductsMatch(context.Context, *RequestByShopIDs) (*ShopProductsMatches, error)
	GetPurchases(context.Context, *RequestByIDs) (*ProductsPurchases, error)
	GetProductsUnsortedCount(context.Context, *Auth) (*UnsortedCount, error)
	ConnectUnsorted(context.Context, *ConnectUnsortedRequest) (*StatusReply, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*StatusReply, error)
	GetProductPurchase(context.Context, *RequestByID) (*ProductPurchases, error)
	SetProductPurchase(context.Context, *SetProductPurchaseRequest) (*StatusReply, error)
	DeleteProductPurchase(context.Context, *RequestByID) (*StatusReply, error)
	SetProductBrand(context.Context, *SetProductBrandRequest) (*StatusReply, error)
	mustEmbedUnimplementedProductsServer()
}

// UnimplementedProductsServer must be embedded to have forward compatible implementations.
type UnimplementedProductsServer struct {
}

func (UnimplementedProductsServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProductsServer) GetProductsCount(context.Context, *ProductListRequest) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsCount not implemented")
}
func (UnimplementedProductsServer) GetProductList(context.Context, *ProductListRequest) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedProductsServer) GetProductSearch(context.Context, *SearchRequest) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductSearch not implemented")
}
func (UnimplementedProductsServer) GetProductsByShopIDs(context.Context, *RequestByShopIDs) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsByShopIDs not implemented")
}
func (UnimplementedProductsServer) GetProduct(context.Context, *RequestByID) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductsServer) GetShopProductsMatch(context.Context, *RequestByShopIDs) (*ShopProductsMatches, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopProductsMatch not implemented")
}
func (UnimplementedProductsServer) GetPurchases(context.Context, *RequestByIDs) (*ProductsPurchases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurchases not implemented")
}
func (UnimplementedProductsServer) GetProductsUnsortedCount(context.Context, *Auth) (*UnsortedCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsUnsortedCount not implemented")
}
func (UnimplementedProductsServer) ConnectUnsorted(context.Context, *ConnectUnsortedRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectUnsorted not implemented")
}
func (UnimplementedProductsServer) UpdateProduct(context.Context, *UpdateProductRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductsServer) GetProductPurchase(context.Context, *RequestByID) (*ProductPurchases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductPurchase not implemented")
}
func (UnimplementedProductsServer) SetProductPurchase(context.Context, *SetProductPurchaseRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProductPurchase not implemented")
}
func (UnimplementedProductsServer) DeleteProductPurchase(context.Context, *RequestByID) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductPurchase not implemented")
}
func (UnimplementedProductsServer) SetProductBrand(context.Context, *SetProductBrandRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProductBrand not implemented")
}
func (UnimplementedProductsServer) mustEmbedUnimplementedProductsServer() {}

// UnsafeProductsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServer will
// result in compilation errors.
type UnsafeProductsServer interface {
	mustEmbedUnimplementedProductsServer()
}

func RegisterProductsServer(s grpc.ServiceRegistrar, srv ProductsServer) {
	s.RegisterService(&Products_ServiceDesc, srv)
}

func _Products_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProductsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProductsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductsCount(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductList(ctx, req.(*ProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProductSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProductSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductSearch(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProductsByShopIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByShopIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductsByShopIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProductsByShopIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductsByShopIDs(ctx, req.(*RequestByShopIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProduct(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetShopProductsMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByShopIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetShopProductsMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetShopProductsMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetShopProductsMatch(ctx, req.(*RequestByShopIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetPurchases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetPurchases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetPurchases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetPurchases(ctx, req.(*RequestByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProductsUnsortedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductsUnsortedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProductsUnsortedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductsUnsortedCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_ConnectUnsorted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectUnsortedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).ConnectUnsorted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/ConnectUnsorted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).ConnectUnsorted(ctx, req.(*ConnectUnsortedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetProductPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProductPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/GetProductPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProductPurchase(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_SetProductPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProductPurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).SetProductPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/SetProductPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).SetProductPurchase(ctx, req.(*SetProductPurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_DeleteProductPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).DeleteProductPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/DeleteProductPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).DeleteProductPurchase(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_SetProductBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProductBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).SetProductBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Products/SetProductBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).SetProductBrand(ctx, req.(*SetProductBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Products_ServiceDesc is the grpc.ServiceDesc for Products service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Products_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Products_Ping_Handler,
		},
		{
			MethodName: "GetProductsCount",
			Handler:    _Products_GetProductsCount_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _Products_GetProductList_Handler,
		},
		{
			MethodName: "GetProductSearch",
			Handler:    _Products_GetProductSearch_Handler,
		},
		{
			MethodName: "GetProductsByShopIDs",
			Handler:    _Products_GetProductsByShopIDs_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Products_GetProduct_Handler,
		},
		{
			MethodName: "GetShopProductsMatch",
			Handler:    _Products_GetShopProductsMatch_Handler,
		},
		{
			MethodName: "GetPurchases",
			Handler:    _Products_GetPurchases_Handler,
		},
		{
			MethodName: "GetProductsUnsortedCount",
			Handler:    _Products_GetProductsUnsortedCount_Handler,
		},
		{
			MethodName: "ConnectUnsorted",
			Handler:    _Products_ConnectUnsorted_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Products_UpdateProduct_Handler,
		},
		{
			MethodName: "GetProductPurchase",
			Handler:    _Products_GetProductPurchase_Handler,
		},
		{
			MethodName: "SetProductPurchase",
			Handler:    _Products_SetProductPurchase_Handler,
		},
		{
			MethodName: "DeleteProductPurchase",
			Handler:    _Products_DeleteProductPurchase_Handler,
		},
		{
			MethodName: "SetProductBrand",
			Handler:    _Products_SetProductBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products_v2.proto",
}
