// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: sales.proto

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

type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *Auth   `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ShopCode string  `protobuf:"bytes,2,opt,name=shopCode,proto3" json:"shopCode,omitempty"`
	Sales    []*Sale `protobuf:"bytes,3,rep,name=sales,proto3" json:"sales,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *CalcRequest) GetShopCode() string {
	if x != nil {
		return x.ShopCode
	}
	return ""
}

func (x *CalcRequest) GetSales() []*Sale {
	if x != nil {
		return x.Sales
	}
	return nil
}

type CalcReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales []*Sale `protobuf:"bytes,1,rep,name=sales,proto3" json:"sales,omitempty"`
}

func (x *CalcReply) Reset() {
	*x = CalcReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcReply) ProtoMessage() {}

func (x *CalcReply) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcReply.ProtoReflect.Descriptor instead.
func (*CalcReply) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{1}
}

func (x *CalcReply) GetSales() []*Sale {
	if x != nil {
		return x.Sales
	}
	return nil
}

var File_sales_proto protoreflect.FileDescriptor

var file_sales_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c,
	0x65, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53,
	0x61, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x32, 0x92, 0x01, 0x0a, 0x05, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4d, 0x61, 0x69,
	0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e,
	0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sales_proto_rawDescOnce sync.Once
	file_sales_proto_rawDescData = file_sales_proto_rawDesc
)

func file_sales_proto_rawDescGZIP() []byte {
	file_sales_proto_rawDescOnce.Do(func() {
		file_sales_proto_rawDescData = protoimpl.X.CompressGZIP(file_sales_proto_rawDescData)
	})
	return file_sales_proto_rawDescData
}

var file_sales_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sales_proto_goTypes = []interface{}{
	(*CalcRequest)(nil),        // 0: cerasus.CalcRequest
	(*CalcReply)(nil),          // 1: cerasus.CalcReply
	(*Auth)(nil),               // 2: cerasus.Auth
	(*Sale)(nil),               // 3: cerasus.Sale
	(*MainGraphicRequest)(nil), // 4: cerasus.MainGraphicRequest
	(*MainGraphicReply)(nil),   // 5: cerasus.MainGraphicReply
}
var file_sales_proto_depIdxs = []int32{
	2, // 0: cerasus.CalcRequest.auth:type_name -> cerasus.Auth
	3, // 1: cerasus.CalcRequest.sales:type_name -> cerasus.Sale
	3, // 2: cerasus.CalcReply.sales:type_name -> cerasus.Sale
	0, // 3: cerasus.Sales.CalculatedSales:input_type -> cerasus.CalcRequest
	4, // 4: cerasus.Sales.GetMainGraphic:input_type -> cerasus.MainGraphicRequest
	1, // 5: cerasus.Sales.CalculatedSales:output_type -> cerasus.CalcReply
	5, // 6: cerasus.Sales.GetMainGraphic:output_type -> cerasus.MainGraphicReply
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sales_proto_init() }
func file_sales_proto_init() {
	if File_sales_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sales_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
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
		file_sales_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcReply); i {
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
			RawDescriptor: file_sales_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sales_proto_goTypes,
		DependencyIndexes: file_sales_proto_depIdxs,
		MessageInfos:      file_sales_proto_msgTypes,
	}.Build()
	File_sales_proto = out.File
	file_sales_proto_rawDesc = nil
	file_sales_proto_goTypes = nil
	file_sales_proto_depIdxs = nil
}
