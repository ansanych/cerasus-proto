// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: oz_v2.proto

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

var File_oz_v2_proto protoreflect.FileDescriptor

var file_oz_v2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x7a, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe4, 0x08, 0x0a, 0x02, 0x4f,
	0x5a, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61,
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
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x73, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x73, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22,
	0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_oz_v2_proto_goTypes = []interface{}{
	(*PingRequest)(nil),      // 0: cerasusV2.PingRequest
	(*Auth)(nil),             // 1: cerasusV2.Auth
	(*LineGraphRequest)(nil), // 2: cerasusV2.LineGraphRequest
	(*RequestByIDs)(nil),     // 3: cerasusV2.RequestByIDs
	(*RequestByDates)(nil),   // 4: cerasusV2.RequestByDates
	(*PingReply)(nil),        // 5: cerasusV2.PingReply
	(*AppShopData)(nil),      // 6: cerasusV2.AppShopData
	(*ShopWidget)(nil),       // 7: cerasusV2.ShopWidget
	(*LineGraph)(nil),        // 8: cerasusV2.LineGraph
	(*Count)(nil),            // 9: cerasusV2.Count
	(*WeekGraphic)(nil),      // 10: cerasusV2.WeekGraphic
	(*RoundGraphic)(nil),     // 11: cerasusV2.RoundGraphic
	(*OrderLeaders)(nil),     // 12: cerasusV2.OrderLeaders
	(*ShopProductList)(nil),  // 13: cerasusV2.ShopProductList
	(*LineGraphics)(nil),     // 14: cerasusV2.LineGraphics
	(*Sales)(nil),            // 15: cerasusV2.Sales
	(*ProductWidget)(nil),    // 16: cerasusV2.ProductWidget
}
var file_oz_v2_proto_depIdxs = []int32{
	0,  // 0: cerasusV2.OZ.Ping:input_type -> cerasusV2.PingRequest
	1,  // 1: cerasusV2.OZ.GetAppData:input_type -> cerasusV2.Auth
	1,  // 2: cerasusV2.OZ.GetShopWidget:input_type -> cerasusV2.Auth
	2,  // 3: cerasusV2.OZ.GetMainGraphic:input_type -> cerasusV2.LineGraphRequest
	1,  // 4: cerasusV2.OZ.GetProductsCount:input_type -> cerasusV2.Auth
	1,  // 5: cerasusV2.OZ.GetProductsCountUnsorted:input_type -> cerasusV2.Auth
	1,  // 6: cerasusV2.OZ.GetFlowGraphicData:input_type -> cerasusV2.Auth
	1,  // 7: cerasusV2.OZ.GetMarginGraphicData:input_type -> cerasusV2.Auth
	2,  // 8: cerasusV2.OZ.GetWeekGraphic:input_type -> cerasusV2.LineGraphRequest
	1,  // 9: cerasusV2.OZ.GetPayRoundGraphic:input_type -> cerasusV2.Auth
	1,  // 10: cerasusV2.OZ.GetCountRoundGraphic:input_type -> cerasusV2.Auth
	1,  // 11: cerasusV2.OZ.GetOrderLeaders:input_type -> cerasusV2.Auth
	3,  // 12: cerasusV2.OZ.GetShopProducts:input_type -> cerasusV2.RequestByIDs
	4,  // 13: cerasusV2.OZ.GetProductGraphics:input_type -> cerasusV2.RequestByDates
	4,  // 14: cerasusV2.OZ.GetSales:input_type -> cerasusV2.RequestByDates
	4,  // 15: cerasusV2.OZ.GetProductWidget:input_type -> cerasusV2.RequestByDates
	4,  // 16: cerasusV2.OZ.GetProductWidgetOrders:input_type -> cerasusV2.RequestByDates
	5,  // 17: cerasusV2.OZ.Ping:output_type -> cerasusV2.PingReply
	6,  // 18: cerasusV2.OZ.GetAppData:output_type -> cerasusV2.AppShopData
	7,  // 19: cerasusV2.OZ.GetShopWidget:output_type -> cerasusV2.ShopWidget
	8,  // 20: cerasusV2.OZ.GetMainGraphic:output_type -> cerasusV2.LineGraph
	9,  // 21: cerasusV2.OZ.GetProductsCount:output_type -> cerasusV2.Count
	9,  // 22: cerasusV2.OZ.GetProductsCountUnsorted:output_type -> cerasusV2.Count
	9,  // 23: cerasusV2.OZ.GetFlowGraphicData:output_type -> cerasusV2.Count
	9,  // 24: cerasusV2.OZ.GetMarginGraphicData:output_type -> cerasusV2.Count
	10, // 25: cerasusV2.OZ.GetWeekGraphic:output_type -> cerasusV2.WeekGraphic
	11, // 26: cerasusV2.OZ.GetPayRoundGraphic:output_type -> cerasusV2.RoundGraphic
	11, // 27: cerasusV2.OZ.GetCountRoundGraphic:output_type -> cerasusV2.RoundGraphic
	12, // 28: cerasusV2.OZ.GetOrderLeaders:output_type -> cerasusV2.OrderLeaders
	13, // 29: cerasusV2.OZ.GetShopProducts:output_type -> cerasusV2.ShopProductList
	14, // 30: cerasusV2.OZ.GetProductGraphics:output_type -> cerasusV2.LineGraphics
	15, // 31: cerasusV2.OZ.GetSales:output_type -> cerasusV2.Sales
	16, // 32: cerasusV2.OZ.GetProductWidget:output_type -> cerasusV2.ProductWidget
	16, // 33: cerasusV2.OZ.GetProductWidgetOrders:output_type -> cerasusV2.ProductWidget
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_oz_v2_proto_init() }
func file_oz_v2_proto_init() {
	if File_oz_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oz_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oz_v2_proto_goTypes,
		DependencyIndexes: file_oz_v2_proto_depIdxs,
	}.Build()
	File_oz_v2_proto = out.File
	file_oz_v2_proto_rawDesc = nil
	file_oz_v2_proto_goTypes = nil
	file_oz_v2_proto_depIdxs = nil
}
