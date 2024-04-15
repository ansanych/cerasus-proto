// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: counter.proto

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

type ShopRegulator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CounterID int64 `protobuf:"varint,2,opt,name=counterID,proto3" json:"counterID,omitempty"`
	ProductID int64 `protobuf:"varint,3,opt,name=productID,proto3" json:"productID,omitempty"`
	ShopCode  int64 `protobuf:"varint,4,opt,name=shopCode,proto3" json:"shopCode,omitempty"`
	Active    bool  `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *ShopRegulator) Reset() {
	*x = ShopRegulator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopRegulator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopRegulator) ProtoMessage() {}

func (x *ShopRegulator) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopRegulator.ProtoReflect.Descriptor instead.
func (*ShopRegulator) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{0}
}

func (x *ShopRegulator) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShopRegulator) GetCounterID() int64 {
	if x != nil {
		return x.CounterID
	}
	return 0
}

func (x *ShopRegulator) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *ShopRegulator) GetShopCode() int64 {
	if x != nil {
		return x.ShopCode
	}
	return 0
}

func (x *ShopRegulator) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type ShopCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CounterID int64  `protobuf:"varint,2,opt,name=counterID,proto3" json:"counterID,omitempty"`
	ProductID int64  `protobuf:"varint,3,opt,name=productID,proto3" json:"productID,omitempty"`
	ShopCode  int64  `protobuf:"varint,4,opt,name=shopCode,proto3" json:"shopCode,omitempty"`
	Count     int64  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Created   string `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *ShopCounter) Reset() {
	*x = ShopCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopCounter) ProtoMessage() {}

func (x *ShopCounter) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopCounter.ProtoReflect.Descriptor instead.
func (*ShopCounter) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{1}
}

func (x *ShopCounter) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShopCounter) GetCounterID() int64 {
	if x != nil {
		return x.CounterID
	}
	return 0
}

func (x *ShopCounter) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *ShopCounter) GetShopCode() int64 {
	if x != nil {
		return x.ShopCode
	}
	return 0
}

func (x *ShopCounter) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ShopCounter) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

type ProductCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ProductID     int64            `protobuf:"varint,2,opt,name=productID,proto3" json:"productID,omitempty"`
	CompanyID     int64            `protobuf:"varint,3,opt,name=companyID,proto3" json:"companyID,omitempty"`
	ProductCount  int64            `protobuf:"varint,4,opt,name=productCount,proto3" json:"productCount,omitempty"`
	ShopCounter   []*ShopCounter   `protobuf:"bytes,5,rep,name=shopCounter,proto3" json:"shopCounter,omitempty"`
	ShopRegulator []*ShopRegulator `protobuf:"bytes,6,rep,name=ShopRegulator,proto3" json:"ShopRegulator,omitempty"`
}

func (x *ProductCounter) Reset() {
	*x = ProductCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCounter) ProtoMessage() {}

func (x *ProductCounter) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCounter.ProtoReflect.Descriptor instead.
func (*ProductCounter) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{2}
}

func (x *ProductCounter) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductCounter) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *ProductCounter) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *ProductCounter) GetProductCount() int64 {
	if x != nil {
		return x.ProductCount
	}
	return 0
}

func (x *ProductCounter) GetShopCounter() []*ShopCounter {
	if x != nil {
		return x.ShopCounter
	}
	return nil
}

func (x *ProductCounter) GetShopRegulator() []*ShopRegulator {
	if x != nil {
		return x.ShopRegulator
	}
	return nil
}

type ProductsCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCounter []*ProductCounter `protobuf:"bytes,1,rep,name=productCounter,proto3" json:"productCounter,omitempty"`
}

func (x *ProductsCounter) Reset() {
	*x = ProductsCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsCounter) ProtoMessage() {}

func (x *ProductsCounter) ProtoReflect() protoreflect.Message {
	mi := &file_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsCounter.ProtoReflect.Descriptor instead.
func (*ProductsCounter) Descriptor() ([]byte, []int) {
	return file_counter_proto_rawDescGZIP(), []int{3}
}

func (x *ProductsCounter) GetProductCounter() []*ProductCounter {
	if x != nil {
		return x.ProductCounter
	}
	return nil
}

var File_counter_proto protoreflect.FileDescriptor

var file_counter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x70,
	0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x53, 0x68,
	0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x22, 0xf6, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x0b, 0x73, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0d,
	0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x53, 0x68, 0x6f,
	0x70, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x52, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x32, 0x4f,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x00, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e,
	0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counter_proto_rawDescOnce sync.Once
	file_counter_proto_rawDescData = file_counter_proto_rawDesc
)

func file_counter_proto_rawDescGZIP() []byte {
	file_counter_proto_rawDescOnce.Do(func() {
		file_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_counter_proto_rawDescData)
	})
	return file_counter_proto_rawDescData
}

var file_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_counter_proto_goTypes = []interface{}{
	(*ShopRegulator)(nil),   // 0: cerasus.ShopRegulator
	(*ShopCounter)(nil),     // 1: cerasus.ShopCounter
	(*ProductCounter)(nil),  // 2: cerasus.ProductCounter
	(*ProductsCounter)(nil), // 3: cerasus.ProductsCounter
	(*Auth)(nil),            // 4: cerasus.Auth
}
var file_counter_proto_depIdxs = []int32{
	1, // 0: cerasus.ProductCounter.shopCounter:type_name -> cerasus.ShopCounter
	0, // 1: cerasus.ProductCounter.ShopRegulator:type_name -> cerasus.ShopRegulator
	2, // 2: cerasus.ProductsCounter.productCounter:type_name -> cerasus.ProductCounter
	4, // 3: cerasus.Counter.GetCompanyProductsCount:input_type -> cerasus.Auth
	3, // 4: cerasus.Counter.GetCompanyProductsCount:output_type -> cerasus.ProductsCounter
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_counter_proto_init() }
func file_counter_proto_init() {
	if File_counter_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopRegulator); i {
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
		file_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopCounter); i {
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
		file_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCounter); i {
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
		file_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsCounter); i {
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
			RawDescriptor: file_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counter_proto_goTypes,
		DependencyIndexes: file_counter_proto_depIdxs,
		MessageInfos:      file_counter_proto_msgTypes,
	}.Build()
	File_counter_proto = out.File
	file_counter_proto_rawDesc = nil
	file_counter_proto_goTypes = nil
	file_counter_proto_depIdxs = nil
}
