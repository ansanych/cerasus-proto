// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ym-apiV3.proto

package cerasus_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ym_apiV3_proto protoreflect.FileDescriptor

var file_ym_apiV3_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x79, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x56, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x1a, 0x0f, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x79, 0x6d,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x56, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9e, 0x04,
	0x0a, 0x06, 0x59, 0x4d, 0x5f, 0x41, 0x50, 0x49, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x33, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73,
	0x12, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x41,
	0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0d, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x70, 0x69, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x41, 0x70,
	0x69, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x33, 0x2e, 0x59, 0x4d, 0x41, 0x70, 0x69, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e,
	0x59, 0x4d, 0x41, 0x70, 0x69, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x70, 0x69, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x41, 0x70, 0x69, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e,
	0x59, 0x4d, 0x41, 0x70, 0x69, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x41,
	0x70, 0x69, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x4c, 0x6f, 0x61,
	0x64, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x13, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59, 0x4d, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x33, 0x2e, 0x59,
	0x4d, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x00, 0x42, 0x23,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73,
	0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ym_apiV3_proto_goTypes = []interface{}{
	(*PingRequest)(nil),      // 0: cerasusV3.PingRequest
	(*YMAuthData)(nil),       // 1: cerasusV3.YMAuthData
	(*YMApiDateRequest)(nil), // 2: cerasusV3.YMApiDateRequest
	(*YMParams)(nil),         // 3: cerasusV3.YMParams
	(*PingReply)(nil),        // 4: cerasusV3.PingReply
	(*StatusReply)(nil),      // 5: cerasusV3.StatusReply
	(*YMApiOrders)(nil),      // 6: cerasusV3.YMApiOrders
	(*YMApiSales)(nil),       // 7: cerasusV3.YMApiSales
	(*YMApiProducts)(nil),    // 8: cerasusV3.YMApiProducts
}
var file_ym_apiV3_proto_depIdxs = []int32{
	0, // 0: cerasusV3.YM_API.Ping:input_type -> cerasusV3.PingRequest
	1, // 1: cerasusV3.YM_API.GetCampaigns:input_type -> cerasusV3.YMAuthData
	2, // 2: cerasusV3.YM_API.LoadApiOrders:input_type -> cerasusV3.YMApiDateRequest
	2, // 3: cerasusV3.YM_API.GetApiOrders:input_type -> cerasusV3.YMApiDateRequest
	2, // 4: cerasusV3.YM_API.LoadApiSales:input_type -> cerasusV3.YMApiDateRequest
	2, // 5: cerasusV3.YM_API.GetApiSales:input_type -> cerasusV3.YMApiDateRequest
	3, // 6: cerasusV3.YM_API.LoadApiProducts:input_type -> cerasusV3.YMParams
	3, // 7: cerasusV3.YM_API.GetApiProducts:input_type -> cerasusV3.YMParams
	4, // 8: cerasusV3.YM_API.Ping:output_type -> cerasusV3.PingReply
	3, // 9: cerasusV3.YM_API.GetCampaigns:output_type -> cerasusV3.YMParams
	5, // 10: cerasusV3.YM_API.LoadApiOrders:output_type -> cerasusV3.StatusReply
	6, // 11: cerasusV3.YM_API.GetApiOrders:output_type -> cerasusV3.YMApiOrders
	5, // 12: cerasusV3.YM_API.LoadApiSales:output_type -> cerasusV3.StatusReply
	7, // 13: cerasusV3.YM_API.GetApiSales:output_type -> cerasusV3.YMApiSales
	5, // 14: cerasusV3.YM_API.LoadApiProducts:output_type -> cerasusV3.StatusReply
	8, // 15: cerasusV3.YM_API.GetApiProducts:output_type -> cerasusV3.YMApiProducts
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ym_apiV3_proto_init() }
func file_ym_apiV3_proto_init() {
	if File_ym_apiV3_proto != nil {
		return
	}
	file_cerasusV3_proto_init()
	file_ym_dataV3_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ym_apiV3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ym_apiV3_proto_goTypes,
		DependencyIndexes: file_ym_apiV3_proto_depIdxs,
	}.Build()
	File_ym_apiV3_proto = out.File
	file_ym_apiV3_proto_rawDesc = nil
	file_ym_apiV3_proto_goTypes = nil
	file_ym_apiV3_proto_depIdxs = nil
}
