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
	GetCerasusBrandData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Brand, error)
	GetBProducts(ctx context.Context, in *BProductsRequest, opts ...grpc.CallOption) (*BProductsReply, error)
	GetBProduct(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BProduct, error)
	CreateBProducts(ctx context.Context, in *BProductCreateRequest, opts ...grpc.CallOption) (*BProductCreateReply, error)
	UpdateBProducts(ctx context.Context, in *BProductCreateRequest, opts ...grpc.CallOption) (*BoolReply, error)
	GetBPrices(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BProductPrices, error)
	GetBPrice(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BProductPrice, error)
	SetBPrice(ctx context.Context, in *BProductCreatePrice, opts ...grpc.CallOption) (*BProductPrice, error)
	GetBSellers(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BSellers, error)
	GetBSeller(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BSeller, error)
	CreateBSeller(ctx context.Context, in *BSellerCreateRequest, opts ...grpc.CallOption) (*BSeller, error)
	UpdateBSeller(ctx context.Context, in *BSellerCreateRequest, opts ...grpc.CallOption) (*BoolReply, error)
	UploadBFile(ctx context.Context, in *BUploadRequest, opts ...grpc.CallOption) (*ImageReply, error)
	GetBFile(ctx context.Context, in *BFileRequest, opts ...grpc.CallOption) (*ImageReply, error)
	GetBrandLink(ctx context.Context, in *BrandLinkRequest, opts ...grpc.CallOption) (*BBrandLink, error)
	GetSellerLinked(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BSellerCompanies, error)
	GetSellerCompanyProducts(ctx context.Context, in *BSCProductsRequest, opts ...grpc.CallOption) (*BSCProductsReply, error)
	SetSellerCompanyProductLink(ctx context.Context, in *BSCProductLinkRequest, opts ...grpc.CallOption) (*BoolReply, error)
	GetSellerUnlinkedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*SellerUnlinkedCountReply, error)
	ConnectBSellerCompanyShop(ctx context.Context, in *ConnectShopRequest, opts ...grpc.CallOption) (*BoolReply, error)
	SetSellerNullCompanyProductLink(ctx context.Context, in *BSCProductLinkRequest, opts ...grpc.CallOption) (*InsertReply, error)
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

func (c *brandsClient) GetCerasusBrandData(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetCerasusBrandData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBProducts(ctx context.Context, in *BProductsRequest, opts ...grpc.CallOption) (*BProductsReply, error) {
	out := new(BProductsReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBProduct(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BProduct, error) {
	out := new(BProduct)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) CreateBProducts(ctx context.Context, in *BProductCreateRequest, opts ...grpc.CallOption) (*BProductCreateReply, error) {
	out := new(BProductCreateReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/CreateBProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) UpdateBProducts(ctx context.Context, in *BProductCreateRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/UpdateBProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBPrices(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BProductPrices, error) {
	out := new(BProductPrices)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBPrice(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BProductPrice, error) {
	out := new(BProductPrice)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) SetBPrice(ctx context.Context, in *BProductCreatePrice, opts ...grpc.CallOption) (*BProductPrice, error) {
	out := new(BProductPrice)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/SetBPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBSellers(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*BSellers, error) {
	out := new(BSellers)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBSellers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBSeller(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BSeller, error) {
	out := new(BSeller)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBSeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) CreateBSeller(ctx context.Context, in *BSellerCreateRequest, opts ...grpc.CallOption) (*BSeller, error) {
	out := new(BSeller)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/CreateBSeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) UpdateBSeller(ctx context.Context, in *BSellerCreateRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/UpdateBSeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) UploadBFile(ctx context.Context, in *BUploadRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/UploadBFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetBFile(ctx context.Context, in *BFileRequest, opts ...grpc.CallOption) (*ImageReply, error) {
	out := new(ImageReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetBFile", in, out, opts...)
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

func (c *brandsClient) GetSellerLinked(ctx context.Context, in *BIDRequest, opts ...grpc.CallOption) (*BSellerCompanies, error) {
	out := new(BSellerCompanies)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetSellerLinked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetSellerCompanyProducts(ctx context.Context, in *BSCProductsRequest, opts ...grpc.CallOption) (*BSCProductsReply, error) {
	out := new(BSCProductsReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetSellerCompanyProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) SetSellerCompanyProductLink(ctx context.Context, in *BSCProductLinkRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/SetSellerCompanyProductLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) GetSellerUnlinkedCount(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*SellerUnlinkedCountReply, error) {
	out := new(SellerUnlinkedCountReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/GetSellerUnlinkedCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) ConnectBSellerCompanyShop(ctx context.Context, in *ConnectShopRequest, opts ...grpc.CallOption) (*BoolReply, error) {
	out := new(BoolReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/ConnectBSellerCompanyShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandsClient) SetSellerNullCompanyProductLink(ctx context.Context, in *BSCProductLinkRequest, opts ...grpc.CallOption) (*InsertReply, error) {
	out := new(InsertReply)
	err := c.cc.Invoke(ctx, "/cerasus.Brands/SetSellerNullCompanyProductLink", in, out, opts...)
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
	GetCerasusBrandData(context.Context, *Auth) (*Brand, error)
	GetBProducts(context.Context, *BProductsRequest) (*BProductsReply, error)
	GetBProduct(context.Context, *BIDRequest) (*BProduct, error)
	CreateBProducts(context.Context, *BProductCreateRequest) (*BProductCreateReply, error)
	UpdateBProducts(context.Context, *BProductCreateRequest) (*BoolReply, error)
	GetBPrices(context.Context, *Auth) (*BProductPrices, error)
	GetBPrice(context.Context, *BIDRequest) (*BProductPrice, error)
	SetBPrice(context.Context, *BProductCreatePrice) (*BProductPrice, error)
	GetBSellers(context.Context, *Auth) (*BSellers, error)
	GetBSeller(context.Context, *BIDRequest) (*BSeller, error)
	CreateBSeller(context.Context, *BSellerCreateRequest) (*BSeller, error)
	UpdateBSeller(context.Context, *BSellerCreateRequest) (*BoolReply, error)
	UploadBFile(context.Context, *BUploadRequest) (*ImageReply, error)
	GetBFile(context.Context, *BFileRequest) (*ImageReply, error)
	GetBrandLink(context.Context, *BrandLinkRequest) (*BBrandLink, error)
	GetSellerLinked(context.Context, *BIDRequest) (*BSellerCompanies, error)
	GetSellerCompanyProducts(context.Context, *BSCProductsRequest) (*BSCProductsReply, error)
	SetSellerCompanyProductLink(context.Context, *BSCProductLinkRequest) (*BoolReply, error)
	GetSellerUnlinkedCount(context.Context, *Auth) (*SellerUnlinkedCountReply, error)
	ConnectBSellerCompanyShop(context.Context, *ConnectShopRequest) (*BoolReply, error)
	SetSellerNullCompanyProductLink(context.Context, *BSCProductLinkRequest) (*InsertReply, error)
	mustEmbedUnimplementedBrandsServer()
}

// UnimplementedBrandsServer must be embedded to have forward compatible implementations.
type UnimplementedBrandsServer struct {
}

func (UnimplementedBrandsServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBrandsServer) GetCerasusBrandData(context.Context, *Auth) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCerasusBrandData not implemented")
}
func (UnimplementedBrandsServer) GetBProducts(context.Context, *BProductsRequest) (*BProductsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBProducts not implemented")
}
func (UnimplementedBrandsServer) GetBProduct(context.Context, *BIDRequest) (*BProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBProduct not implemented")
}
func (UnimplementedBrandsServer) CreateBProducts(context.Context, *BProductCreateRequest) (*BProductCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBProducts not implemented")
}
func (UnimplementedBrandsServer) UpdateBProducts(context.Context, *BProductCreateRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBProducts not implemented")
}
func (UnimplementedBrandsServer) GetBPrices(context.Context, *Auth) (*BProductPrices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBPrices not implemented")
}
func (UnimplementedBrandsServer) GetBPrice(context.Context, *BIDRequest) (*BProductPrice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBPrice not implemented")
}
func (UnimplementedBrandsServer) SetBPrice(context.Context, *BProductCreatePrice) (*BProductPrice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBPrice not implemented")
}
func (UnimplementedBrandsServer) GetBSellers(context.Context, *Auth) (*BSellers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBSellers not implemented")
}
func (UnimplementedBrandsServer) GetBSeller(context.Context, *BIDRequest) (*BSeller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBSeller not implemented")
}
func (UnimplementedBrandsServer) CreateBSeller(context.Context, *BSellerCreateRequest) (*BSeller, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBSeller not implemented")
}
func (UnimplementedBrandsServer) UpdateBSeller(context.Context, *BSellerCreateRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBSeller not implemented")
}
func (UnimplementedBrandsServer) UploadBFile(context.Context, *BUploadRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadBFile not implemented")
}
func (UnimplementedBrandsServer) GetBFile(context.Context, *BFileRequest) (*ImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBFile not implemented")
}
func (UnimplementedBrandsServer) GetBrandLink(context.Context, *BrandLinkRequest) (*BBrandLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandLink not implemented")
}
func (UnimplementedBrandsServer) GetSellerLinked(context.Context, *BIDRequest) (*BSellerCompanies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSellerLinked not implemented")
}
func (UnimplementedBrandsServer) GetSellerCompanyProducts(context.Context, *BSCProductsRequest) (*BSCProductsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSellerCompanyProducts not implemented")
}
func (UnimplementedBrandsServer) SetSellerCompanyProductLink(context.Context, *BSCProductLinkRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSellerCompanyProductLink not implemented")
}
func (UnimplementedBrandsServer) GetSellerUnlinkedCount(context.Context, *Auth) (*SellerUnlinkedCountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSellerUnlinkedCount not implemented")
}
func (UnimplementedBrandsServer) ConnectBSellerCompanyShop(context.Context, *ConnectShopRequest) (*BoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectBSellerCompanyShop not implemented")
}
func (UnimplementedBrandsServer) SetSellerNullCompanyProductLink(context.Context, *BSCProductLinkRequest) (*InsertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSellerNullCompanyProductLink not implemented")
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

func _Brands_GetCerasusBrandData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetCerasusBrandData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetCerasusBrandData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetCerasusBrandData(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBProducts(ctx, req.(*BProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBProduct(ctx, req.(*BIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_CreateBProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).CreateBProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/CreateBProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).CreateBProducts(ctx, req.(*BProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_UpdateBProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BProductCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).UpdateBProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/UpdateBProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).UpdateBProducts(ctx, req.(*BProductCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBPrices(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBPrice(ctx, req.(*BIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_SetBPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BProductCreatePrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).SetBPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/SetBPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).SetBPrice(ctx, req.(*BProductCreatePrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBSellers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBSellers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBSellers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBSellers(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBSeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBSeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBSeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBSeller(ctx, req.(*BIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_CreateBSeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BSellerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).CreateBSeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/CreateBSeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).CreateBSeller(ctx, req.(*BSellerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_UpdateBSeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BSellerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).UpdateBSeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/UpdateBSeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).UpdateBSeller(ctx, req.(*BSellerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_UploadBFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).UploadBFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/UploadBFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).UploadBFile(ctx, req.(*BUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetBFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetBFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetBFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetBFile(ctx, req.(*BFileRequest))
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

func _Brands_GetSellerLinked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetSellerLinked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetSellerLinked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetSellerLinked(ctx, req.(*BIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetSellerCompanyProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BSCProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetSellerCompanyProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetSellerCompanyProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetSellerCompanyProducts(ctx, req.(*BSCProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_SetSellerCompanyProductLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BSCProductLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).SetSellerCompanyProductLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/SetSellerCompanyProductLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).SetSellerCompanyProductLink(ctx, req.(*BSCProductLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_GetSellerUnlinkedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).GetSellerUnlinkedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/GetSellerUnlinkedCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).GetSellerUnlinkedCount(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_ConnectBSellerCompanyShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).ConnectBSellerCompanyShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/ConnectBSellerCompanyShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).ConnectBSellerCompanyShop(ctx, req.(*ConnectShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brands_SetSellerNullCompanyProductLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BSCProductLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandsServer).SetSellerNullCompanyProductLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Brands/SetSellerNullCompanyProductLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandsServer).SetSellerNullCompanyProductLink(ctx, req.(*BSCProductLinkRequest))
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
			MethodName: "GetCerasusBrandData",
			Handler:    _Brands_GetCerasusBrandData_Handler,
		},
		{
			MethodName: "GetBProducts",
			Handler:    _Brands_GetBProducts_Handler,
		},
		{
			MethodName: "GetBProduct",
			Handler:    _Brands_GetBProduct_Handler,
		},
		{
			MethodName: "CreateBProducts",
			Handler:    _Brands_CreateBProducts_Handler,
		},
		{
			MethodName: "UpdateBProducts",
			Handler:    _Brands_UpdateBProducts_Handler,
		},
		{
			MethodName: "GetBPrices",
			Handler:    _Brands_GetBPrices_Handler,
		},
		{
			MethodName: "GetBPrice",
			Handler:    _Brands_GetBPrice_Handler,
		},
		{
			MethodName: "SetBPrice",
			Handler:    _Brands_SetBPrice_Handler,
		},
		{
			MethodName: "GetBSellers",
			Handler:    _Brands_GetBSellers_Handler,
		},
		{
			MethodName: "GetBSeller",
			Handler:    _Brands_GetBSeller_Handler,
		},
		{
			MethodName: "CreateBSeller",
			Handler:    _Brands_CreateBSeller_Handler,
		},
		{
			MethodName: "UpdateBSeller",
			Handler:    _Brands_UpdateBSeller_Handler,
		},
		{
			MethodName: "UploadBFile",
			Handler:    _Brands_UploadBFile_Handler,
		},
		{
			MethodName: "GetBFile",
			Handler:    _Brands_GetBFile_Handler,
		},
		{
			MethodName: "GetBrandLink",
			Handler:    _Brands_GetBrandLink_Handler,
		},
		{
			MethodName: "GetSellerLinked",
			Handler:    _Brands_GetSellerLinked_Handler,
		},
		{
			MethodName: "GetSellerCompanyProducts",
			Handler:    _Brands_GetSellerCompanyProducts_Handler,
		},
		{
			MethodName: "SetSellerCompanyProductLink",
			Handler:    _Brands_SetSellerCompanyProductLink_Handler,
		},
		{
			MethodName: "GetSellerUnlinkedCount",
			Handler:    _Brands_GetSellerUnlinkedCount_Handler,
		},
		{
			MethodName: "ConnectBSellerCompanyShop",
			Handler:    _Brands_ConnectBSellerCompanyShop_Handler,
		},
		{
			MethodName: "SetSellerNullCompanyProductLink",
			Handler:    _Brands_SetSellerNullCompanyProductLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brand.proto",
}
