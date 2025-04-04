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

type ShopProductPricerHistory struct {
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

func (x *ShopProductPricerHistory) Reset() {
	*x = ShopProductPricerHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopProductPricerHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopProductPricerHistory) ProtoMessage() {}

func (x *ShopProductPricerHistory) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ShopProductPricerHistory.ProtoReflect.Descriptor instead.
func (*ShopProductPricerHistory) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{2}
}

func (x *ShopProductPricerHistory) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *ShopProductPricerHistory) GetShopPrice() int32 {
	if x != nil {
		return x.ShopPrice
	}
	return 0
}

func (x *ShopProductPricerHistory) GetParsePrice() int32 {
	if x != nil {
		return x.ParsePrice
	}
	return 0
}

func (x *ShopProductPricerHistory) GetSetPrice() int32 {
	if x != nil {
		return x.SetPrice
	}
	return 0
}

func (x *ShopProductPricerHistory) GetSendPrice() bool {
	if x != nil {
		return x.SendPrice
	}
	return false
}

func (x *ShopProductPricerHistory) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type ProductPricerHistoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopProduct *ShopProduct                `protobuf:"bytes,1,opt,name=shopProduct,proto3" json:"shopProduct,omitempty"`
	Data        []*ShopProductPricerHistory `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ProductPricerHistoryItem) Reset() {
	*x = ProductPricerHistoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricerHistoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricerHistoryItem) ProtoMessage() {}

func (x *ProductPricerHistoryItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricerHistoryItem.ProtoReflect.Descriptor instead.
func (*ProductPricerHistoryItem) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{3}
}

func (x *ProductPricerHistoryItem) GetShopProduct() *ShopProduct {
	if x != nil {
		return x.ShopProduct
	}
	return nil
}

func (x *ProductPricerHistoryItem) GetData() []*ShopProductPricerHistory {
	if x != nil {
		return x.Data
	}
	return nil
}

type ProductPricerHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shop  string                      `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	Items []*ProductPricerHistoryItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProductPricerHistory) Reset() {
	*x = ProductPricerHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricerHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricerHistory) ProtoMessage() {}

func (x *ProductPricerHistory) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricerHistory.ProtoReflect.Descriptor instead.
func (*ProductPricerHistory) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{4}
}

func (x *ProductPricerHistory) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *ProductPricerHistory) GetItems() []*ProductPricerHistoryItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProductPricer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int64   `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Shop      string  `protobuf:"bytes,2,opt,name=shop,proto3" json:"shop,omitempty"`
	Active    bool    `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Target    float32 `protobuf:"fixed32,4,opt,name=target,proto3" json:"target,omitempty"`
	OffsetMin float32 `protobuf:"fixed32,5,opt,name=offsetMin,proto3" json:"offsetMin,omitempty"`
	OffsetMax float32 `protobuf:"fixed32,6,opt,name=offsetMax,proto3" json:"offsetMax,omitempty"`
	Updated   string  `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	CompanyID int64   `protobuf:"varint,8,opt,name=companyID,proto3" json:"companyID,omitempty"`
}

func (x *ProductPricer) Reset() {
	*x = ProductPricer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricer_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricer) ProtoMessage() {}

func (x *ProductPricer) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricer.ProtoReflect.Descriptor instead.
func (*ProductPricer) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{5}
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
		mi := &file_pricer_v2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPricers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPricers) ProtoMessage() {}

func (x *ProductPricers) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProductPricers.ProtoReflect.Descriptor instead.
func (*ProductPricers) Descriptor() ([]byte, []int) {
	return file_pricer_v2_proto_rawDescGZIP(), []int{6}
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
	0x74, 0x65, 0x72, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x18,
	0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x8d, 0x01, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x38, 0x0a,
	0x0b, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x70,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x65, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x39, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22,
	0x3e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72,
	0x73, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0xe0, 0x05, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69,
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
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a,
	0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*ParamForCounter)(nil),          // 0: cerasusV2.ParamForCounter
	(*ParamsForCounter)(nil),         // 1: cerasusV2.ParamsForCounter
	(*ShopProductPricerHistory)(nil), // 2: cerasusV2.ShopProductPricerHistory
	(*ProductPricerHistoryItem)(nil), // 3: cerasusV2.ProductPricerHistoryItem
	(*ProductPricerHistory)(nil),     // 4: cerasusV2.ProductPricerHistory
	(*ProductPricer)(nil),            // 5: cerasusV2.ProductPricer
	(*ProductPricers)(nil),           // 6: cerasusV2.ProductPricers
	(*ShopProduct)(nil),              // 7: cerasusV2.ShopProduct
	(*PingRequest)(nil),              // 8: cerasusV2.PingRequest
	(*Company)(nil),                  // 9: cerasusV2.Company
	(*ParserGetRequest)(nil),         // 10: cerasusV2.ParserGetRequest
	(*ParserSetRequest)(nil),         // 11: cerasusV2.ParserSetRequest
	(*DetectorGetRequest)(nil),       // 12: cerasusV2.DetectorGetRequest
	(*DetectorSetRequest)(nil),       // 13: cerasusV2.DetectorSetRequest
	(*RequestByID)(nil),              // 14: cerasusV2.RequestByID
	(*PingReply)(nil),                // 15: cerasusV2.PingReply
	(*ParserJob)(nil),                // 16: cerasusV2.ParserJob
	(*StatusReply)(nil),              // 17: cerasusV2.StatusReply
	(*DetectorGetReply)(nil),         // 18: cerasusV2.DetectorGetReply
	(*ReplyID)(nil),                  // 19: cerasusV2.ReplyID
}
var file_pricer_v2_proto_depIdxs = []int32{
	0,  // 0: cerasusV2.ParamsForCounter.params:type_name -> cerasusV2.ParamForCounter
	7,  // 1: cerasusV2.ProductPricerHistoryItem.shopProduct:type_name -> cerasusV2.ShopProduct
	2,  // 2: cerasusV2.ProductPricerHistoryItem.data:type_name -> cerasusV2.ShopProductPricerHistory
	3,  // 3: cerasusV2.ProductPricerHistory.items:type_name -> cerasusV2.ProductPricerHistoryItem
	5,  // 4: cerasusV2.ProductPricers.data:type_name -> cerasusV2.ProductPricer
	8,  // 5: cerasusV2.Pricer.Ping:input_type -> cerasusV2.PingRequest
	9,  // 6: cerasusV2.Pricer.GetParamsForCounter:input_type -> cerasusV2.Company
	10, // 7: cerasusV2.Pricer.GetParserData:input_type -> cerasusV2.ParserGetRequest
	11, // 8: cerasusV2.Pricer.SetParserData:input_type -> cerasusV2.ParserSetRequest
	12, // 9: cerasusV2.Pricer.GetDetectorData:input_type -> cerasusV2.DetectorGetRequest
	13, // 10: cerasusV2.Pricer.SetDetectorData:input_type -> cerasusV2.DetectorSetRequest
	9,  // 11: cerasusV2.Pricer.GetProductsWithPricer:input_type -> cerasusV2.Company
	14, // 12: cerasusV2.Pricer.GetProductPricers:input_type -> cerasusV2.RequestByID
	5,  // 13: cerasusV2.Pricer.SetProductPricer:input_type -> cerasusV2.ProductPricer
	14, // 14: cerasusV2.Pricer.GetProductPricerHistory:input_type -> cerasusV2.RequestByID
	15, // 15: cerasusV2.Pricer.Ping:output_type -> cerasusV2.PingReply
	1,  // 16: cerasusV2.Pricer.GetParamsForCounter:output_type -> cerasusV2.ParamsForCounter
	16, // 17: cerasusV2.Pricer.GetParserData:output_type -> cerasusV2.ParserJob
	17, // 18: cerasusV2.Pricer.SetParserData:output_type -> cerasusV2.StatusReply
	18, // 19: cerasusV2.Pricer.GetDetectorData:output_type -> cerasusV2.DetectorGetReply
	17, // 20: cerasusV2.Pricer.SetDetectorData:output_type -> cerasusV2.StatusReply
	19, // 21: cerasusV2.Pricer.GetProductsWithPricer:output_type -> cerasusV2.ReplyID
	6,  // 22: cerasusV2.Pricer.GetProductPricers:output_type -> cerasusV2.ProductPricers
	17, // 23: cerasusV2.Pricer.SetProductPricer:output_type -> cerasusV2.StatusReply
	4,  // 24: cerasusV2.Pricer.GetProductPricerHistory:output_type -> cerasusV2.ProductPricerHistory
	15, // [15:25] is the sub-list for method output_type
	5,  // [5:15] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
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
			switch v := v.(*ShopProductPricerHistory); i {
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
			switch v := v.(*ProductPricerHistoryItem); i {
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
		file_pricer_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pricer_v2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
