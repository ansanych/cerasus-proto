// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: products_v2.proto

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

type RequestByShopIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *Auth   `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ID   []int64 `protobuf:"varint,2,rep,packed,name=ID,proto3" json:"ID,omitempty"`
	Shop string  `protobuf:"bytes,3,opt,name=shop,proto3" json:"shop,omitempty"`
}

func (x *RequestByShopIDs) Reset() {
	*x = RequestByShopIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_products_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestByShopIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestByShopIDs) ProtoMessage() {}

func (x *RequestByShopIDs) ProtoReflect() protoreflect.Message {
	mi := &file_products_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestByShopIDs.ProtoReflect.Descriptor instead.
func (*RequestByShopIDs) Descriptor() ([]byte, []int) {
	return file_products_v2_proto_rawDescGZIP(), []int{0}
}

func (x *RequestByShopIDs) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *RequestByShopIDs) GetID() []int64 {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *RequestByShopIDs) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

var File_products_v2_proto protoreflect.FileDescriptor

var file_products_v2_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x68, 0x6f,
	0x70, 0x49, 0x44, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x32, 0x94, 0x02,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x56, 0x32, 0x12, 0x36, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x12,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x42, 0x79, 0x49, 0x44, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x73,
	0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x73, 0x1a, 0x13, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_products_v2_proto_rawDescOnce sync.Once
	file_products_v2_proto_rawDescData = file_products_v2_proto_rawDesc
)

func file_products_v2_proto_rawDescGZIP() []byte {
	file_products_v2_proto_rawDescOnce.Do(func() {
		file_products_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_products_v2_proto_rawDescData)
	})
	return file_products_v2_proto_rawDescData
}

var file_products_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_products_v2_proto_goTypes = []interface{}{
	(*RequestByShopIDs)(nil), // 0: cerasusV2.RequestByShopIDs
	(*Auth)(nil),             // 1: cerasusV2.Auth
	(*PingRequest)(nil),      // 2: cerasusV2.PingRequest
	(*RequestByID)(nil),      // 3: cerasusV2.RequestByID
	(*RequestByIDs)(nil),     // 4: cerasusV2.RequestByIDs
	(*PingReply)(nil),        // 5: cerasusV2.PingReply
	(*Product)(nil),          // 6: cerasusV2.Product
	(*Products)(nil),         // 7: cerasusV2.Products
}
var file_products_v2_proto_depIdxs = []int32{
	1, // 0: cerasusV2.RequestByShopIDs.auth:type_name -> cerasusV2.Auth
	2, // 1: cerasusV2.ProductsV2.Ping:input_type -> cerasusV2.PingRequest
	3, // 2: cerasusV2.ProductsV2.GetProductByID:input_type -> cerasusV2.RequestByID
	4, // 3: cerasusV2.ProductsV2.GetProductsByIDs:input_type -> cerasusV2.RequestByIDs
	0, // 4: cerasusV2.ProductsV2.GetProductsByShopIDs:input_type -> cerasusV2.RequestByShopIDs
	5, // 5: cerasusV2.ProductsV2.Ping:output_type -> cerasusV2.PingReply
	6, // 6: cerasusV2.ProductsV2.GetProductByID:output_type -> cerasusV2.Product
	7, // 7: cerasusV2.ProductsV2.GetProductsByIDs:output_type -> cerasusV2.Products
	7, // 8: cerasusV2.ProductsV2.GetProductsByShopIDs:output_type -> cerasusV2.Products
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_products_v2_proto_init() }
func file_products_v2_proto_init() {
	if File_products_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_products_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestByShopIDs); i {
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
			RawDescriptor: file_products_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_products_v2_proto_goTypes,
		DependencyIndexes: file_products_v2_proto_depIdxs,
		MessageInfos:      file_products_v2_proto_msgTypes,
	}.Build()
	File_products_v2_proto = out.File
	file_products_v2_proto_rawDesc = nil
	file_products_v2_proto_goTypes = nil
	file_products_v2_proto_depIdxs = nil
}
