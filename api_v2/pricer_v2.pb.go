// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pricer_v2.proto

package cerasus_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParserSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parser      string `protobuf:"bytes,1,opt,name=parser,proto3" json:"parser,omitempty"`
	QueueDataID int64  `protobuf:"varint,2,opt,name=queueDataID,proto3" json:"queueDataID,omitempty"`
	File        []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *ParserSetRequest) Reset() {
	*x = ParserSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserSetRequest) ProtoMessage() {}

func (x *ParserSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserSetRequest.ProtoReflect.Descriptor instead.
func (*ParserSetRequest) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{0}
}

func (x *ParserSetRequest) GetParser() string {
	if x != nil {
		return x.Parser
	}
	return ""
}

func (x *ParserSetRequest) GetQueueDataID() int64 {
	if x != nil {
		return x.QueueDataID
	}
	return 0
}

func (x *ParserSetRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type ParserGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parser string `protobuf:"bytes,1,opt,name=parser,proto3" json:"parser,omitempty"`
}

func (x *ParserGetRequest) Reset() {
	*x = ParserGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserGetRequest) ProtoMessage() {}

func (x *ParserGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserGetRequest.ProtoReflect.Descriptor instead.
func (*ParserGetRequest) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{1}
}

func (x *ParserGetRequest) GetParser() string {
	if x != nil {
		return x.Parser
	}
	return ""
}

type ParserJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parser      string `protobuf:"bytes,1,opt,name=parser,proto3" json:"parser,omitempty"`
	QueueDataID int64  `protobuf:"varint,2,opt,name=queueDataID,proto3" json:"queueDataID,omitempty"`
	Url         string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ParserJob) Reset() {
	*x = ParserJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParserJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParserJob) ProtoMessage() {}

func (x *ParserJob) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParserJob.ProtoReflect.Descriptor instead.
func (*ParserJob) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{2}
}

func (x *ParserJob) GetParser() string {
	if x != nil {
		return x.Parser
	}
	return ""
}

func (x *ParserJob) GetQueueDataID() int64 {
	if x != nil {
		return x.QueueDataID
	}
	return 0
}

func (x *ParserJob) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ParamForCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int64    `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Shop      []string `protobuf:"bytes,2,rep,name=shop,proto3" json:"shop,omitempty"`
}

func (x *ParamForCounter) Reset() {
	*x = ParamForCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamForCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamForCounter) ProtoMessage() {}

func (x *ParamForCounter) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamForCounter.ProtoReflect.Descriptor instead.
func (*ParamForCounter) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{3}
}

func (x *ParamForCounter) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *ParamForCounter) GetShop() []string {
	if x != nil {
		return x.Shop
	}
	return nil
}

type ParamsForCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params []*ParamForCounter `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *ParamsForCounter) Reset() {
	*x = ParamsForCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsForCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsForCounter) ProtoMessage() {}

func (x *ParamsForCounter) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamsForCounter.ProtoReflect.Descriptor instead.
func (*ParamsForCounter) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{4}
}

func (x *ParamsForCounter) GetParams() []*ParamForCounter {
	if x != nil {
		return x.Params
	}
	return nil
}

type PricerJobParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shop string `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
}

func (x *PricerJobParamsRequest) Reset() {
	*x = PricerJobParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricerJobParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricerJobParamsRequest) ProtoMessage() {}

func (x *PricerJobParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricerJobParamsRequest.ProtoReflect.Descriptor instead.
func (*PricerJobParamsRequest) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{5}
}

func (x *PricerJobParamsRequest) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

type PricerJobParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID int64  `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Shop  string `protobuf:"bytes,2,opt,name=shop,proto3" json:"shop,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PricerJobParams) Reset() {
	*x = PricerJobParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricerJobParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricerJobParams) ProtoMessage() {}

func (x *PricerJobParams) ProtoReflect() protoreflect.Message {
	mi := &file_pricer_v2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricerJobParams.ProtoReflect.Descriptor instead.
func (*PricerJobParams) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{6}
}

func (x *PricerJobParams) GetJobID() int64 {
	if x != nil {
		return x.JobID
	}
	return 0
}

func (x *PricerJobParams) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *PricerJobParams) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_pricer_v2_proto protoreflect.FileDescriptor

var file_pricer_v2_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60,
	0x0a, 0x10, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x2a, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x09,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x43, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x46, 0x6f,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x46, 0x0a, 0x10, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x2c, 0x0a, 0x16, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x22, 0x4d, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32,
	0xb7, 0x03, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x1b, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x21, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68,
	0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pricer_v2_proto_rawDescOnce sync.Once
	file_pricer_v2_proto_rawDescData = file_pricer_v2_proto_rawDesc
)

func file_pricer_v2_proto_rawDescGZIP() []byte {
	file_pricer_v2_proto_rawDescOnce.Do(func() {
		file_pricer_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_pricer_v2_proto_rawDescData)
	})
	return file_pricer_v2_proto_rawDescData
}

var file_pricer_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pricer_v2_proto_goTypes = []interface{}{
	(*ParserSetRequest)(nil),       // 0: cerasusV2.ParserSetRequest
	(*ParserGetRequest)(nil),       // 1: cerasusV2.ParserGetRequest
	(*ParserJob)(nil),              // 2: cerasusV2.ParserJob
	(*ParamForCounter)(nil),        // 3: cerasusV2.ParamForCounter
	(*ParamsForCounter)(nil),       // 4: cerasusV2.ParamsForCounter
	(*PricerJobParamsRequest)(nil), // 5: cerasusV2.PricerJobParamsRequest
	(*PricerJobParams)(nil),        // 6: cerasusV2.PricerJobParams
	(*PingRequest)(nil),            // 7: cerasusV2.PingRequest
	(*Company)(nil),                // 8: cerasusV2.Company
	(*DetectedPrice)(nil),          // 9: cerasusV2.DetectedPrice
	(*PingReply)(nil),              // 10: cerasusV2.PingReply
	(*StatusReply)(nil),            // 11: cerasusV2.StatusReply
}
var file_pricer_v2_proto_depIdxs = []int32{
	3,  // 0: cerasusV2.ParamsForCounter.params:type_name -> cerasusV2.ParamForCounter
	7,  // 1: cerasusV2.Pricer.Ping:input_type -> cerasusV2.PingRequest
	8,  // 2: cerasusV2.Pricer.GetParamsForCounter:input_type -> cerasusV2.Company
	5,  // 3: cerasusV2.Pricer.GetPricerJobParams:input_type -> cerasusV2.PricerJobParamsRequest
	9,  // 4: cerasusV2.Pricer.SetDetectedPrice:input_type -> cerasusV2.DetectedPrice
	1,  // 5: cerasusV2.Pricer.GetParserData:input_type -> cerasusV2.ParserGetRequest
	0,  // 6: cerasusV2.Pricer.SetParserData:input_type -> cerasusV2.ParserSetRequest
	10, // 7: cerasusV2.Pricer.Ping:output_type -> cerasusV2.PingReply
	4,  // 8: cerasusV2.Pricer.GetParamsForCounter:output_type -> cerasusV2.ParamsForCounter
	6,  // 9: cerasusV2.Pricer.GetPricerJobParams:output_type -> cerasusV2.PricerJobParams
	11, // 10: cerasusV2.Pricer.SetDetectedPrice:output_type -> cerasusV2.StatusReply
	2,  // 11: cerasusV2.Pricer.GetParserData:output_type -> cerasusV2.ParserJob
	11, // 12: cerasusV2.Pricer.SetParserData:output_type -> cerasusV2.StatusReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pricer_v2_proto_init() }
func file_pricer_v2_proto_init() {
	if File_pricer_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pricer_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserSetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pricer_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pricer_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParserJob); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pricer_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamForCounter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pricer_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamsForCounter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pricer_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricerJobParamsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pricer_v2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricerJobParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pricer_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pricer_v2_proto_goTypes,
		DependencyIndexes: file_pricer_v2_proto_depIdxs,
		MessageInfos:      file_pricer_v2_proto_msgTypes,
	}.Build()
	File_pricer_v2_proto = out.File
	file_pricer_v2_proto_rawDesc = nil
	file_pricer_v2_proto_goTypes = nil
	file_pricer_v2_proto_depIdxs = nil
}
