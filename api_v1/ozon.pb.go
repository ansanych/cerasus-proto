// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ozon.proto

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

type ShopOzonAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID int32  `protobuf:"varint,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ApiKey   string `protobuf:"bytes,2,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Error    bool   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ShopOzonAuth) Reset() {
	*x = ShopOzonAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopOzonAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopOzonAuth) ProtoMessage() {}

func (x *ShopOzonAuth) ProtoReflect() protoreflect.Message {
	mi := &file_ozon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopOzonAuth.ProtoReflect.Descriptor instead.
func (*ShopOzonAuth) Descriptor() ([]byte, []int) {
	return file_ozon_proto_rawDescGZIP(), []int{0}
}

func (x *ShopOzonAuth) GetClientID() int32 {
	if x != nil {
		return x.ClientID
	}
	return 0
}

func (x *ShopOzonAuth) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *ShopOzonAuth) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type SetOzonAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *Auth         `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ShopAuth *ShopOzonAuth `protobuf:"bytes,2,opt,name=shopAuth,proto3" json:"shopAuth,omitempty"`
}

func (x *SetOzonAuth) Reset() {
	*x = SetOzonAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOzonAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOzonAuth) ProtoMessage() {}

func (x *SetOzonAuth) ProtoReflect() protoreflect.Message {
	mi := &file_ozon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOzonAuth.ProtoReflect.Descriptor instead.
func (*SetOzonAuth) Descriptor() ([]byte, []int) {
	return file_ozon_proto_rawDescGZIP(), []int{1}
}

func (x *SetOzonAuth) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *SetOzonAuth) GetShopAuth() *ShopOzonAuth {
	if x != nil {
		return x.ShopAuth
	}
	return nil
}

var File_ozon_proto protoreflect.FileDescriptor

var file_ozon_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x4f, 0x7a, 0x6f, 0x6e,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x63,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4f, 0x7a, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x12, 0x31, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x4f, 0x7a, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x41,
	0x75, 0x74, 0x68, 0x32, 0xba, 0x04, 0x0a, 0x04, 0x4f, 0x7a, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x07,
	0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x7a, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0d,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x15, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4f, 0x7a, 0x6f, 0x6e,
	0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x79, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x44, 0x61, 0x79, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ozon_proto_rawDescOnce sync.Once
	file_ozon_proto_rawDescData = file_ozon_proto_rawDesc
)

func file_ozon_proto_rawDescGZIP() []byte {
	file_ozon_proto_rawDescOnce.Do(func() {
		file_ozon_proto_rawDescData = protoimpl.X.CompressGZIP(file_ozon_proto_rawDescData)
	})
	return file_ozon_proto_rawDescData
}

var file_ozon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ozon_proto_goTypes = []interface{}{
	(*ShopOzonAuth)(nil),           // 0: cerasus.ShopOzonAuth
	(*SetOzonAuth)(nil),            // 1: cerasus.SetOzonAuth
	(*Auth)(nil),                   // 2: cerasus.Auth
	(*ShopProductListRequest)(nil), // 3: cerasus.ShopProductListRequest
	(*ShopProductRequest)(nil),     // 4: cerasus.ShopProductRequest
	(*BoolReply)(nil),              // 5: cerasus.BoolReply
	(*CountReply)(nil),             // 6: cerasus.CountReply
	(*ShopProductListReply)(nil),   // 7: cerasus.ShopProductListReply
	(*ShopProduct)(nil),            // 8: cerasus.ShopProduct
	(*DaysSalesReply)(nil),         // 9: cerasus.DaysSalesReply
}
var file_ozon_proto_depIdxs = []int32{
	2,  // 0: cerasus.SetOzonAuth.auth:type_name -> cerasus.Auth
	0,  // 1: cerasus.SetOzonAuth.shopAuth:type_name -> cerasus.ShopOzonAuth
	1,  // 2: cerasus.Ozon.SetAuth:input_type -> cerasus.SetOzonAuth
	2,  // 3: cerasus.Ozon.GetAuth:input_type -> cerasus.Auth
	2,  // 4: cerasus.Ozon.ErrorAuth:input_type -> cerasus.Auth
	2,  // 5: cerasus.Ozon.GetUnsortedCount:input_type -> cerasus.Auth
	3,  // 6: cerasus.Ozon.GetUnsortedList:input_type -> cerasus.ShopProductListRequest
	2,  // 7: cerasus.Ozon.GetProductCount:input_type -> cerasus.Auth
	3,  // 8: cerasus.Ozon.GetProductList:input_type -> cerasus.ShopProductListRequest
	4,  // 9: cerasus.Ozon.GetProduct:input_type -> cerasus.ShopProductRequest
	2,  // 10: cerasus.Ozon.GetDaySales:input_type -> cerasus.Auth
	5,  // 11: cerasus.Ozon.SetAuth:output_type -> cerasus.BoolReply
	0,  // 12: cerasus.Ozon.GetAuth:output_type -> cerasus.ShopOzonAuth
	5,  // 13: cerasus.Ozon.ErrorAuth:output_type -> cerasus.BoolReply
	6,  // 14: cerasus.Ozon.GetUnsortedCount:output_type -> cerasus.CountReply
	7,  // 15: cerasus.Ozon.GetUnsortedList:output_type -> cerasus.ShopProductListReply
	6,  // 16: cerasus.Ozon.GetProductCount:output_type -> cerasus.CountReply
	7,  // 17: cerasus.Ozon.GetProductList:output_type -> cerasus.ShopProductListReply
	8,  // 18: cerasus.Ozon.GetProduct:output_type -> cerasus.ShopProduct
	9,  // 19: cerasus.Ozon.GetDaySales:output_type -> cerasus.DaysSalesReply
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_ozon_proto_init() }
func file_ozon_proto_init() {
	if File_ozon_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ozon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopOzonAuth); i {
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
		file_ozon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOzonAuth); i {
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
			RawDescriptor: file_ozon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ozon_proto_goTypes,
		DependencyIndexes: file_ozon_proto_depIdxs,
		MessageInfos:      file_ozon_proto_msgTypes,
	}.Build()
	File_ozon_proto = out.File
	file_ozon_proto_rawDesc = nil
	file_ozon_proto_goTypes = nil
	file_ozon_proto_depIdxs = nil
}
