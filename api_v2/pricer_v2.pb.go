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
		mi := &file_pricer_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamForCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamForCounter) ProtoMessage() {}

func (x *ParamForCounter) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ParamForCounter.ProtoReflect.Descriptor instead.
func (*ParamForCounter) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pricer_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamsForCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamsForCounter) ProtoMessage() {}

func (x *ParamsForCounter) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ParamsForCounter.ProtoReflect.Descriptor instead.
func (*ParamsForCounter) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{1}
}

func (x *ParamsForCounter) GetParams() []*ParamForCounter {
	if x != nil {
		return x.Params
	}
	return nil
}

type ProductPricerHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated    string `protobuf:"bytes,1,opt,name=updated,proto3" json:"updated,omitempty"`
	ShopPrice  int32  `protobuf:"varint,2,opt,name=shopPrice,proto3" json:"shopPrice,omitempty"`
	ParsePrice int32  `protobuf:"varint,3,opt,name=parsePrice,proto3" json:"parsePrice,omitempty"`
	SetPrice   int32  `protobuf:"varint,4,opt,name=setPrice,proto3" json:"setPrice,omitempty"`
	SendPrice  bool   `protobuf:"varint,5,opt,name=sendPrice,proto3" json:"sendPrice,omitempty"`
	Error      bool   `protobuf:"varint,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ProductPricerHistory) Reset() {
	*x = ProductPricerHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricerHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricerHistory) ProtoMessage() {}

func (x *ProductPricerHistory) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricerHistory.ProtoReflect.Descriptor instead.
func (*ProductPricerHistory) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{2}
}

func (x *ProductPricerHistory) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *ProductPricerHistory) GetShopPrice() int32 {
	if x != nil {
		return x.ShopPrice
	}
	return 0
}

func (x *ProductPricerHistory) GetParsePrice() int32 {
	if x != nil {
		return x.ParsePrice
	}
	return 0
}

func (x *ProductPricerHistory) GetSetPrice() int32 {
	if x != nil {
		return x.SetPrice
	}
	return 0
}

func (x *ProductPricerHistory) GetSendPrice() bool {
	if x != nil {
		return x.SendPrice
	}
	return false
}

func (x *ProductPricerHistory) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type ProductPricer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int64                   `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Shop      string                  `protobuf:"bytes,2,opt,name=shop,proto3" json:"shop,omitempty"`
	Active    bool                    `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Target    float32                 `protobuf:"fixed32,4,opt,name=target,proto3" json:"target,omitempty"`
	OffsetMin float32                 `protobuf:"fixed32,5,opt,name=offsetMin,proto3" json:"offsetMin,omitempty"`
	OffsetMax float32                 `protobuf:"fixed32,6,opt,name=offsetMax,proto3" json:"offsetMax,omitempty"`
	Updated   string                  `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	History   []*ProductPricerHistory `protobuf:"bytes,8,rep,name=history,proto3" json:"history,omitempty"`
	CompanyID int64                   `protobuf:"varint,9,opt,name=companyID,proto3" json:"companyID,omitempty"`
}

func (x *ProductPricer) Reset() {
	*x = ProductPricer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricer) ProtoMessage() {}

func (x *ProductPricer) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricer.ProtoReflect.Descriptor instead.
func (*ProductPricer) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{3}
}

func (x *ProductPricer) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *ProductPricer) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *ProductPricer) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *ProductPricer) GetTarget() float32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *ProductPricer) GetOffsetMin() float32 {
	if x != nil {
		return x.OffsetMin
	}
	return 0
}

func (x *ProductPricer) GetOffsetMax() float32 {
	if x != nil {
		return x.OffsetMax
	}
	return 0
}

func (x *ProductPricer) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *ProductPricer) GetHistory() []*ProductPricerHistory {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *ProductPricer) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

type ProductPricers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ProductPricer `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProductPricers) Reset() {
	*x = ProductPricers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricers) ProtoMessage() {}

func (x *ProductPricers) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricers.ProtoReflect.Descriptor instead.
func (*ProductPricers) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{4}
}

func (x *ProductPricers) GetData() []*ProductPricer {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pricer_v2_proto protoreflect.FileDescriptor

var file_pricer_v2_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43,
	0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x68, 0x6f, 0x70, 0x22, 0x46, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x46, 0x6f, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x14,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa0, 0x02, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x68, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22,
	0x3e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72,
	0x73, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x8a, 0x05, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x1b, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x62,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x53,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72,
	0x12, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x44, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x73, 0x12,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x72, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x72, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e,
	0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pricer_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pricer_v2_proto_goTypes = []interface{}{
	(*ParamForCounter)(nil),      // 0: cerasusV2.ParamForCounter
	(*ParamsForCounter)(nil),     // 1: cerasusV2.ParamsForCounter
	(*ProductPricerHistory)(nil), // 2: cerasusV2.ProductPricerHistory
	(*ProductPricer)(nil),        // 3: cerasusV2.ProductPricer
	(*ProductPricers)(nil),       // 4: cerasusV2.ProductPricers
	(*PingRequest)(nil),          // 5: cerasusV2.PingRequest
	(*Company)(nil),              // 6: cerasusV2.Company
	(*ParserGetRequest)(nil),     // 7: cerasusV2.ParserGetRequest
	(*ParserSetRequest)(nil),     // 8: cerasusV2.ParserSetRequest
	(*DetectorGetRequest)(nil),   // 9: cerasusV2.DetectorGetRequest
	(*DetectorSetRequest)(nil),   // 10: cerasusV2.DetectorSetRequest
	(*RequestByID)(nil),          // 11: cerasusV2.RequestByID
	(*PingReply)(nil),            // 12: cerasusV2.PingReply
	(*ParserJob)(nil),            // 13: cerasusV2.ParserJob
	(*StatusReply)(nil),          // 14: cerasusV2.StatusReply
	(*DetectorGetReply)(nil),     // 15: cerasusV2.DetectorGetReply
	(*ReplyID)(nil),              // 16: cerasusV2.ReplyID
}
var file_pricer_v2_proto_depIdxs = []int32{
	0,  // 0: cerasusV2.ParamsForCounter.params:type_name -> cerasusV2.ParamForCounter
	2,  // 1: cerasusV2.ProductPricer.history:type_name -> cerasusV2.ProductPricerHistory
	3,  // 2: cerasusV2.ProductPricers.data:type_name -> cerasusV2.ProductPricer
	5,  // 3: cerasusV2.Pricer.Ping:input_type -> cerasusV2.PingRequest
	6,  // 4: cerasusV2.Pricer.GetParamsForCounter:input_type -> cerasusV2.Company
	7,  // 5: cerasusV2.Pricer.GetParserData:input_type -> cerasusV2.ParserGetRequest
	8,  // 6: cerasusV2.Pricer.SetParserData:input_type -> cerasusV2.ParserSetRequest
	9,  // 7: cerasusV2.Pricer.GetDetectorData:input_type -> cerasusV2.DetectorGetRequest
	10, // 8: cerasusV2.Pricer.SetDetectorData:input_type -> cerasusV2.DetectorSetRequest
	6,  // 9: cerasusV2.Pricer.GetProductsWithPricer:input_type -> cerasusV2.Company
	11, // 10: cerasusV2.Pricer.GetProductPricers:input_type -> cerasusV2.RequestByID
	3,  // 11: cerasusV2.Pricer.SetProductPricer:input_type -> cerasusV2.ProductPricer
	12, // 12: cerasusV2.Pricer.Ping:output_type -> cerasusV2.PingReply
	1,  // 13: cerasusV2.Pricer.GetParamsForCounter:output_type -> cerasusV2.ParamsForCounter
	13, // 14: cerasusV2.Pricer.GetParserData:output_type -> cerasusV2.ParserJob
	14, // 15: cerasusV2.Pricer.SetParserData:output_type -> cerasusV2.StatusReply
	15, // 16: cerasusV2.Pricer.GetDetectorData:output_type -> cerasusV2.DetectorGetReply
	14, // 17: cerasusV2.Pricer.SetDetectorData:output_type -> cerasusV2.StatusReply
	16, // 18: cerasusV2.Pricer.GetProductsWithPricer:output_type -> cerasusV2.ReplyID
	4,  // 19: cerasusV2.Pricer.GetProductPricers:output_type -> cerasusV2.ProductPricers
	14, // 20: cerasusV2.Pricer.SetProductPricer:output_type -> cerasusV2.StatusReply
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pricer_v2_proto_init() }
func file_pricer_v2_proto_init() {
	if File_pricer_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pricer_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pricer_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pricer_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPricerHistory); i {
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
			switch v := v.(*ProductPricer); i {
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
			switch v := v.(*ProductPricers); i {
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
			NumMessages:   5,
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
