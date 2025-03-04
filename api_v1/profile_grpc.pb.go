// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: profile.proto

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

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	UpdateCompany(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetRoles(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Roles, error)
	CreateUser(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*StatusReply, error)
	UpdateUser(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetUser(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*User, error)
	GetCompanyUsers(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Users, error)
	GetCompaniesList(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*CompanyesList, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateCompany(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetRoles(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Roles, error) {
	out := new(Roles)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CreateUser(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateUser(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetUser(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetCompanyUsers(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/GetCompanyUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetCompaniesList(ctx context.Context, in *RequestByIDs, opts ...grpc.CallOption) (*CompanyesList, error) {
	out := new(CompanyesList)
	err := c.cc.Invoke(ctx, "/cerasus.Profile/GetCompaniesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	UpdateCompany(context.Context, *SetProfileRequest) (*StatusReply, error)
	GetRoles(context.Context, *Auth) (*Roles, error)
	CreateUser(context.Context, *SetProfileRequest) (*StatusReply, error)
	UpdateUser(context.Context, *SetProfileRequest) (*StatusReply, error)
	GetUser(context.Context, *RequestByID) (*User, error)
	GetCompanyUsers(context.Context, *Auth) (*Users, error)
	GetCompaniesList(context.Context, *RequestByIDs) (*CompanyesList, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProfileServer) UpdateCompany(context.Context, *SetProfileRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedProfileServer) GetRoles(context.Context, *Auth) (*Roles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedProfileServer) CreateUser(context.Context, *SetProfileRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedProfileServer) UpdateUser(context.Context, *SetProfileRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedProfileServer) GetUser(context.Context, *RequestByID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedProfileServer) GetCompanyUsers(context.Context, *Auth) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyUsers not implemented")
}
func (UnimplementedProfileServer) GetCompaniesList(context.Context, *RequestByIDs) (*CompanyesList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompaniesList not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateCompany(ctx, req.(*SetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetRoles(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CreateUser(ctx, req.(*SetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateUser(ctx, req.(*SetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetUser(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetCompanyUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetCompanyUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/GetCompanyUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetCompanyUsers(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetCompaniesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetCompaniesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasus.Profile/GetCompaniesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetCompaniesList(ctx, req.(*RequestByIDs))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasus.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Profile_Ping_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Profile_UpdateCompany_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _Profile_GetRoles_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Profile_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Profile_UpdateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Profile_GetUser_Handler,
		},
		{
			MethodName: "GetCompanyUsers",
			Handler:    _Profile_GetCompanyUsers_Handler,
		},
		{
			MethodName: "GetCompaniesList",
			Handler:    _Profile_GetCompaniesList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
