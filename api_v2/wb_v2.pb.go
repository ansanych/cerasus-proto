// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: wb_v2.proto

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

var File_wb_v2_proto protoreflect.FileDescriptor

var file_wb_v2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x62, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc1, 0x06, 0x0a, 0x02, 0x57,
	0x42, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12,
	0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65,
	0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73,
	0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x23,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73,
	0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_wb_v2_proto_goTypes = []interface{}{
	(*PingRequest)(nil),      // 0: cerasusV2.PingRequest
	(*Auth)(nil),             // 1: cerasusV2.Auth
	(*LineGraphRequest)(nil), // 2: cerasusV2.LineGraphRequest
	(*RequestByIDs)(nil),     // 3: cerasusV2.RequestByIDs
	(*PingReply)(nil),        // 4: cerasusV2.PingReply
	(*AppShopData)(nil),      // 5: cerasusV2.AppShopData
	(*ShopWidget)(nil),       // 6: cerasusV2.ShopWidget
	(*LineGraph)(nil),        // 7: cerasusV2.LineGraph
	(*Count)(nil),            // 8: cerasusV2.Count
	(*WeekGraphic)(nil),      // 9: cerasusV2.WeekGraphic
	(*RoundGraphic)(nil),     // 10: cerasusV2.RoundGraphic
	(*OrderLeaders)(nil),     // 11: cerasusV2.OrderLeaders
	(*ShopProductList)(nil),  // 12: cerasusV2.ShopProductList
}
var file_wb_v2_proto_depIdxs = []int32{
	0,  // 0: cerasusV2.WB.Ping:input_type -> cerasusV2.PingRequest
	1,  // 1: cerasusV2.WB.GetAppData:input_type -> cerasusV2.Auth
	1,  // 2: cerasusV2.WB.GetShopWidget:input_type -> cerasusV2.Auth
	2,  // 3: cerasusV2.WB.GetMainGraphic:input_type -> cerasusV2.LineGraphRequest
	1,  // 4: cerasusV2.WB.GetProductsCount:input_type -> cerasusV2.Auth
	1,  // 5: cerasusV2.WB.GetProductsCountUnsorted:input_type -> cerasusV2.Auth
	1,  // 6: cerasusV2.WB.GetFlowGraphicData:input_type -> cerasusV2.Auth
	1,  // 7: cerasusV2.WB.GetMarginGraphicData:input_type -> cerasusV2.Auth
	2,  // 8: cerasusV2.WB.GetWeekGraphic:input_type -> cerasusV2.LineGraphRequest
	1,  // 9: cerasusV2.WB.GetPayRoundGraphic:input_type -> cerasusV2.Auth
	1,  // 10: cerasusV2.WB.GetCountRoundGraphic:input_type -> cerasusV2.Auth
	1,  // 11: cerasusV2.WB.GetOrderLeaders:input_type -> cerasusV2.Auth
	3,  // 12: cerasusV2.WB.GetShopProducts:input_type -> cerasusV2.RequestByIDs
	4,  // 13: cerasusV2.WB.Ping:output_type -> cerasusV2.PingReply
	5,  // 14: cerasusV2.WB.GetAppData:output_type -> cerasusV2.AppShopData
	6,  // 15: cerasusV2.WB.GetShopWidget:output_type -> cerasusV2.ShopWidget
	7,  // 16: cerasusV2.WB.GetMainGraphic:output_type -> cerasusV2.LineGraph
	8,  // 17: cerasusV2.WB.GetProductsCount:output_type -> cerasusV2.Count
	8,  // 18: cerasusV2.WB.GetProductsCountUnsorted:output_type -> cerasusV2.Count
	8,  // 19: cerasusV2.WB.GetFlowGraphicData:output_type -> cerasusV2.Count
	8,  // 20: cerasusV2.WB.GetMarginGraphicData:output_type -> cerasusV2.Count
	9,  // 21: cerasusV2.WB.GetWeekGraphic:output_type -> cerasusV2.WeekGraphic
	10, // 22: cerasusV2.WB.GetPayRoundGraphic:output_type -> cerasusV2.RoundGraphic
	10, // 23: cerasusV2.WB.GetCountRoundGraphic:output_type -> cerasusV2.RoundGraphic
	11, // 24: cerasusV2.WB.GetOrderLeaders:output_type -> cerasusV2.OrderLeaders
	12, // 25: cerasusV2.WB.GetShopProducts:output_type -> cerasusV2.ShopProductList
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_wb_v2_proto_init() }
func file_wb_v2_proto_init() {
	if File_wb_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wb_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wb_v2_proto_goTypes,
		DependencyIndexes: file_wb_v2_proto_depIdxs,
	}.Build()
	File_wb_v2_proto = out.File
	file_wb_v2_proto_rawDesc = nil
	file_wb_v2_proto_goTypes = nil
	file_wb_v2_proto_depIdxs = nil
}
