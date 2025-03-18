// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: board_v2.proto

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

// BoardClient is the client API for Board service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	GetQueues(ctx context.Context, in *QueuesRequest, opts ...grpc.CallOption) (*Queues, error)
	GetQueue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*Queue, error)
	ReQueue(ctx context.Context, in *ReQueueRequest, opts ...grpc.CallOption) (*StatusReply, error)
	SetQueueParams(ctx context.Context, in *QueueParamsSet, opts ...grpc.CallOption) (*StatusReply, error)
	GetLogsCount(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Count, error)
	GetCompanies(ctx context.Context, in *BoardCompaniesRequest, opts ...grpc.CallOption) (*CompanyList, error)
	SearchCompanies(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*CompanyList, error)
	GetCompany(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*BoardCompanyData, error)
	UpdateCompany(ctx context.Context, in *UpdateBoardCompanyRequest, opts ...grpc.CallOption) (*StatusReply, error)
	UpdateCompanyCounter(ctx context.Context, in *UpdateBoardCompanyServiceRequest, opts ...grpc.CallOption) (*StatusReply, error)
	UpdateCompanyPricer(ctx context.Context, in *UpdateBoardCompanyServiceRequest, opts ...grpc.CallOption) (*StatusReply, error)
	SearchProduct(ctx context.Context, in *SearchProductRequest, opts ...grpc.CallOption) (*BoardProductData, error)
	PingServices(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingServicesReply, error)
}

type boardClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardClient(cc grpc.ClientConnInterface) BoardClient {
	return &boardClient{cc}
}

func (c *boardClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) GetQueues(ctx context.Context, in *QueuesRequest, opts ...grpc.CallOption) (*Queues, error) {
	out := new(Queues)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/GetQueues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) GetQueue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*Queue, error) {
	out := new(Queue)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/GetQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) ReQueue(ctx context.Context, in *ReQueueRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/ReQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) SetQueueParams(ctx context.Context, in *QueueParamsSet, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/SetQueueParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) GetLogsCount(ctx context.Context, in *RequestByDates, opts ...grpc.CallOption) (*Count, error) {
	out := new(Count)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/GetLogsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) GetCompanies(ctx context.Context, in *BoardCompaniesRequest, opts ...grpc.CallOption) (*CompanyList, error) {
	out := new(CompanyList)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/GetCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) SearchCompanies(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*CompanyList, error) {
	out := new(CompanyList)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/SearchCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) GetCompany(ctx context.Context, in *RequestByID, opts ...grpc.CallOption) (*BoardCompanyData, error) {
	out := new(BoardCompanyData)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) UpdateCompany(ctx context.Context, in *UpdateBoardCompanyRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) UpdateCompanyCounter(ctx context.Context, in *UpdateBoardCompanyServiceRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/UpdateCompanyCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) UpdateCompanyPricer(ctx context.Context, in *UpdateBoardCompanyServiceRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/UpdateCompanyPricer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) SearchProduct(ctx context.Context, in *SearchProductRequest, opts ...grpc.CallOption) (*BoardProductData, error) {
	out := new(BoardProductData)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/SearchProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardClient) PingServices(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingServicesReply, error) {
	out := new(PingServicesReply)
	err := c.cc.Invoke(ctx, "/cerasusV2.Board/PingServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServer is the server API for Board service.
// All implementations must embed UnimplementedBoardServer
// for forward compatibility
type BoardServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	GetQueues(context.Context, *QueuesRequest) (*Queues, error)
	GetQueue(context.Context, *QueueRequest) (*Queue, error)
	ReQueue(context.Context, *ReQueueRequest) (*StatusReply, error)
	SetQueueParams(context.Context, *QueueParamsSet) (*StatusReply, error)
	GetLogsCount(context.Context, *RequestByDates) (*Count, error)
	GetCompanies(context.Context, *BoardCompaniesRequest) (*CompanyList, error)
	SearchCompanies(context.Context, *SearchRequest) (*CompanyList, error)
	GetCompany(context.Context, *RequestByID) (*BoardCompanyData, error)
	UpdateCompany(context.Context, *UpdateBoardCompanyRequest) (*StatusReply, error)
	UpdateCompanyCounter(context.Context, *UpdateBoardCompanyServiceRequest) (*StatusReply, error)
	UpdateCompanyPricer(context.Context, *UpdateBoardCompanyServiceRequest) (*StatusReply, error)
	SearchProduct(context.Context, *SearchProductRequest) (*BoardProductData, error)
	PingServices(context.Context, *PingRequest) (*PingServicesReply, error)
	mustEmbedUnimplementedBoardServer()
}

// UnimplementedBoardServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServer struct {
}

func (UnimplementedBoardServer) Ping(context.Context, *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBoardServer) GetQueues(context.Context, *QueuesRequest) (*Queues, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueues not implemented")
}
func (UnimplementedBoardServer) GetQueue(context.Context, *QueueRequest) (*Queue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueue not implemented")
}
func (UnimplementedBoardServer) ReQueue(context.Context, *ReQueueRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReQueue not implemented")
}
func (UnimplementedBoardServer) SetQueueParams(context.Context, *QueueParamsSet) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetQueueParams not implemented")
}
func (UnimplementedBoardServer) GetLogsCount(context.Context, *RequestByDates) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogsCount not implemented")
}
func (UnimplementedBoardServer) GetCompanies(context.Context, *BoardCompaniesRequest) (*CompanyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanies not implemented")
}
func (UnimplementedBoardServer) SearchCompanies(context.Context, *SearchRequest) (*CompanyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCompanies not implemented")
}
func (UnimplementedBoardServer) GetCompany(context.Context, *RequestByID) (*BoardCompanyData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedBoardServer) UpdateCompany(context.Context, *UpdateBoardCompanyRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedBoardServer) UpdateCompanyCounter(context.Context, *UpdateBoardCompanyServiceRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanyCounter not implemented")
}
func (UnimplementedBoardServer) UpdateCompanyPricer(context.Context, *UpdateBoardCompanyServiceRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanyPricer not implemented")
}
func (UnimplementedBoardServer) SearchProduct(context.Context, *SearchProductRequest) (*BoardProductData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProduct not implemented")
}
func (UnimplementedBoardServer) PingServices(context.Context, *PingRequest) (*PingServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingServices not implemented")
}
func (UnimplementedBoardServer) mustEmbedUnimplementedBoardServer() {}

// UnsafeBoardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServer will
// result in compilation errors.
type UnsafeBoardServer interface {
	mustEmbedUnimplementedBoardServer()
}

func RegisterBoardServer(s grpc.ServiceRegistrar, srv BoardServer) {
	s.RegisterService(&Board_ServiceDesc, srv)
}

func _Board_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_GetQueues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetQueues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/GetQueues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetQueues(ctx, req.(*QueuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_GetQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/GetQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetQueue(ctx, req.(*QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_ReQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).ReQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/ReQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).ReQueue(ctx, req.(*ReQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_SetQueueParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueParamsSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).SetQueueParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/SetQueueParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).SetQueueParams(ctx, req.(*QueueParamsSet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_GetLogsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByDates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetLogsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/GetLogsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetLogsCount(ctx, req.(*RequestByDates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_GetCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/GetCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetCompanies(ctx, req.(*BoardCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_SearchCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).SearchCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/SearchCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).SearchCompanies(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).GetCompany(ctx, req.(*RequestByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoardCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).UpdateCompany(ctx, req.(*UpdateBoardCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_UpdateCompanyCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoardCompanyServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).UpdateCompanyCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/UpdateCompanyCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).UpdateCompanyCounter(ctx, req.(*UpdateBoardCompanyServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_UpdateCompanyPricer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoardCompanyServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).UpdateCompanyPricer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/UpdateCompanyPricer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).UpdateCompanyPricer(ctx, req.(*UpdateBoardCompanyServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_SearchProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).SearchProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/SearchProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).SearchProduct(ctx, req.(*SearchProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Board_PingServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServer).PingServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cerasusV2.Board/PingServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServer).PingServices(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Board_ServiceDesc is the grpc.ServiceDesc for Board service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Board_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cerasusV2.Board",
	HandlerType: (*BoardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Board_Ping_Handler,
		},
		{
			MethodName: "GetQueues",
			Handler:    _Board_GetQueues_Handler,
		},
		{
			MethodName: "GetQueue",
			Handler:    _Board_GetQueue_Handler,
		},
		{
			MethodName: "ReQueue",
			Handler:    _Board_ReQueue_Handler,
		},
		{
			MethodName: "SetQueueParams",
			Handler:    _Board_SetQueueParams_Handler,
		},
		{
			MethodName: "GetLogsCount",
			Handler:    _Board_GetLogsCount_Handler,
		},
		{
			MethodName: "GetCompanies",
			Handler:    _Board_GetCompanies_Handler,
		},
		{
			MethodName: "SearchCompanies",
			Handler:    _Board_SearchCompanies_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _Board_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Board_UpdateCompany_Handler,
		},
		{
			MethodName: "UpdateCompanyCounter",
			Handler:    _Board_UpdateCompanyCounter_Handler,
		},
		{
			MethodName: "UpdateCompanyPricer",
			Handler:    _Board_UpdateCompanyPricer_Handler,
		},
		{
			MethodName: "SearchProduct",
			Handler:    _Board_SearchProduct_Handler,
		},
		{
			MethodName: "PingServices",
			Handler:    _Board_PingServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "board_v2.proto",
}
