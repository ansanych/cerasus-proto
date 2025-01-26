// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: counter_v2.proto

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

type ProductCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth      *Auth              `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ProductID int64              `protobuf:"varint,2,opt,name=productID,proto3" json:"productID,omitempty"`
	Count     int32              `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	ID        int64              `protobuf:"varint,4,opt,name=ID,proto3" json:"ID,omitempty"`
	ShopData  []*ShopCounterData `protobuf:"bytes,5,rep,name=shopData,proto3" json:"shopData,omitempty"`
}

func (x *ProductCount) Reset() {
	*x = ProductCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCount) ProtoMessage() {}

func (x *ProductCount) ProtoReflect() protoreflect.Message {
	mi := &file_counter_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCount.ProtoReflect.Descriptor instead.
func (*ProductCount) Descriptor() ([]byte, []int) {
	return file_counter_v2_proto_rawDescGZIP(), []int{0}
}

func (x *ProductCount) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ProductCount) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *ProductCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ProductCount) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductCount) GetShopData() []*ShopCounterData {
	if x != nil {
		return x.ShopData
	}
	return nil
}

type ShopCounterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shop   string `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	Active bool   `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	Count  int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Date   string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ShopCounterData) Reset() {
	*x = ShopCounterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopCounterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopCounterData) ProtoMessage() {}

func (x *ShopCounterData) ProtoReflect() protoreflect.Message {
	mi := &file_counter_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopCounterData.ProtoReflect.Descriptor instead.
func (*ShopCounterData) Descriptor() ([]byte, []int) {
	return file_counter_v2_proto_rawDescGZIP(), []int{1}
}

func (x *ShopCounterData) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *ShopCounterData) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *ShopCounterData) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ShopCounterData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_counter_v2_proto protoreflect.FileDescriptor

var file_counter_v2_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaf, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x08, 0x73, 0x68, 0x6f,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x67, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63,
	0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counter_v2_proto_rawDescOnce sync.Once
	file_counter_v2_proto_rawDescData = file_counter_v2_proto_rawDesc
)

func file_counter_v2_proto_rawDescGZIP() []byte {
	file_counter_v2_proto_rawDescOnce.Do(func() {
		file_counter_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_counter_v2_proto_rawDescData)
	})
	return file_counter_v2_proto_rawDescData
}

var file_counter_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_counter_v2_proto_goTypes = []interface{}{
	(*ProductCount)(nil),    // 0: cerasusV2.ProductCount
	(*ShopCounterData)(nil), // 1: cerasusV2.ShopCounterData
	(*Auth)(nil),            // 2: cerasusV2.Auth
	(*PingRequest)(nil),     // 3: cerasusV2.PingRequest
	(*RequestByID)(nil),     // 4: cerasusV2.RequestByID
	(*PingReply)(nil),       // 5: cerasusV2.PingReply
	(*StatusReply)(nil),     // 6: cerasusV2.StatusReply
}
var file_counter_v2_proto_depIdxs = []int32{
	2, // 0: cerasusV2.ProductCount.auth:type_name -> cerasusV2.Auth
	1, // 1: cerasusV2.ProductCount.shopData:type_name -> cerasusV2.ShopCounterData
	3, // 2: cerasusV2.Counter.Ping:input_type -> cerasusV2.PingRequest
	4, // 3: cerasusV2.Counter.GetProductCount:input_type -> cerasusV2.RequestByID
	0, // 4: cerasusV2.Counter.SetProductCount:input_type -> cerasusV2.ProductCount
	5, // 5: cerasusV2.Counter.Ping:output_type -> cerasusV2.PingReply
	0, // 6: cerasusV2.Counter.GetProductCount:output_type -> cerasusV2.ProductCount
	6, // 7: cerasusV2.Counter.SetProductCount:output_type -> cerasusV2.StatusReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_counter_v2_proto_init() }
func file_counter_v2_proto_init() {
	if File_counter_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_counter_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCount); i {
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
		file_counter_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopCounterData); i {
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
			RawDescriptor: file_counter_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counter_v2_proto_goTypes,
		DependencyIndexes: file_counter_v2_proto_depIdxs,
		MessageInfos:      file_counter_v2_proto_msgTypes,
	}.Build()
	File_counter_v2_proto = out.File
	file_counter_v2_proto_rawDesc = nil
	file_counter_v2_proto_goTypes = nil
	file_counter_v2_proto_depIdxs = nil
}
