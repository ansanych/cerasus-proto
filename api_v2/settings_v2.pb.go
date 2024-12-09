// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: settings_v2.proto

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

type AppServiceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string        `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Status    ServiceStatus `protobuf:"varint,2,opt,name=status,proto3,enum=cerasusV2.ServiceStatus" json:"status,omitempty"`
	StartDate string        `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   string        `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Limit     int32         `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Access    bool          `protobuf:"varint,6,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *AppServiceData) Reset() {
	*x = AppServiceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppServiceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppServiceData) ProtoMessage() {}

func (x *AppServiceData) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppServiceData.ProtoReflect.Descriptor instead.
func (*AppServiceData) Descriptor() ([]byte, []int) {
	return file_settings_v2_proto_rawDescGZIP(), []int{0}
}

func (x *AppServiceData) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AppServiceData) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_DISABLED
}

func (x *AppServiceData) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *AppServiceData) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *AppServiceData) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AppServiceData) GetAccess() bool {
	if x != nil {
		return x.Access
	}
	return false
}

type UserAppData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shops    []*AppShopData    `protobuf:"bytes,1,rep,name=shops,proto3" json:"shops,omitempty"`
	Services []*AppServiceData `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *UserAppData) Reset() {
	*x = UserAppData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAppData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAppData) ProtoMessage() {}

func (x *UserAppData) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAppData.ProtoReflect.Descriptor instead.
func (*UserAppData) Descriptor() ([]byte, []int) {
	return file_settings_v2_proto_rawDescGZIP(), []int{1}
}

func (x *UserAppData) GetShops() []*AppShopData {
	if x != nil {
		return x.Shops
	}
	return nil
}

func (x *UserAppData) GetServices() []*AppServiceData {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_settings_v2_proto protoreflect.FileDescriptor

var file_settings_v2_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x72, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x68, 0x6f,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x12, 0x35, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x32, 0xc4, 0x02, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x16, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a,
	0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a,
	0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63,
	0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_settings_v2_proto_rawDescOnce sync.Once
	file_settings_v2_proto_rawDescData = file_settings_v2_proto_rawDesc
)

func file_settings_v2_proto_rawDescGZIP() []byte {
	file_settings_v2_proto_rawDescOnce.Do(func() {
		file_settings_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_v2_proto_rawDescData)
	})
	return file_settings_v2_proto_rawDescData
}

var file_settings_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_settings_v2_proto_goTypes = []interface{}{
	(*AppServiceData)(nil),   // 0: cerasusV2.AppServiceData
	(*UserAppData)(nil),      // 1: cerasusV2.UserAppData
	(ServiceStatus)(0),       // 2: cerasusV2.ServiceStatus
	(*AppShopData)(nil),      // 3: cerasusV2.AppShopData
	(*PingRequest)(nil),      // 4: cerasusV2.PingRequest
	(*Auth)(nil),             // 5: cerasusV2.Auth
	(*LineGraphRequest)(nil), // 6: cerasusV2.LineGraphRequest
	(*PingReply)(nil),        // 7: cerasusV2.PingReply
	(*LineGraph)(nil),        // 8: cerasusV2.LineGraph
	(*RoundGraphic)(nil),     // 9: cerasusV2.RoundGraphic
}
var file_settings_v2_proto_depIdxs = []int32{
	2, // 0: cerasusV2.AppServiceData.status:type_name -> cerasusV2.ServiceStatus
	3, // 1: cerasusV2.UserAppData.shops:type_name -> cerasusV2.AppShopData
	0, // 2: cerasusV2.UserAppData.services:type_name -> cerasusV2.AppServiceData
	4, // 3: cerasusV2.Settings.Ping:input_type -> cerasusV2.PingRequest
	5, // 4: cerasusV2.Settings.GetUserAppData:input_type -> cerasusV2.Auth
	6, // 5: cerasusV2.Settings.GetMainGraphic:input_type -> cerasusV2.LineGraphRequest
	5, // 6: cerasusV2.Settings.GetFlowGraphic:input_type -> cerasusV2.Auth
	5, // 7: cerasusV2.Settings.GetMarginGraphic:input_type -> cerasusV2.Auth
	7, // 8: cerasusV2.Settings.Ping:output_type -> cerasusV2.PingReply
	1, // 9: cerasusV2.Settings.GetUserAppData:output_type -> cerasusV2.UserAppData
	8, // 10: cerasusV2.Settings.GetMainGraphic:output_type -> cerasusV2.LineGraph
	9, // 11: cerasusV2.Settings.GetFlowGraphic:output_type -> cerasusV2.RoundGraphic
	9, // 12: cerasusV2.Settings.GetMarginGraphic:output_type -> cerasusV2.RoundGraphic
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_settings_v2_proto_init() }
func file_settings_v2_proto_init() {
	if File_settings_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_settings_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppServiceData); i {
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
		file_settings_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAppData); i {
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
			RawDescriptor: file_settings_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_settings_v2_proto_goTypes,
		DependencyIndexes: file_settings_v2_proto_depIdxs,
		MessageInfos:      file_settings_v2_proto_msgTypes,
	}.Build()
	File_settings_v2_proto = out.File
	file_settings_v2_proto_rawDesc = nil
	file_settings_v2_proto_goTypes = nil
	file_settings_v2_proto_depIdxs = nil
}
