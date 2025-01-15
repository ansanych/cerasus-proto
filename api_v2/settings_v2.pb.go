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

type MarginLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level Level   `protobuf:"varint,1,opt,name=level,proto3,enum=cerasusV2.Level" json:"level,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MarginLevel) Reset() {
	*x = MarginLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarginLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarginLevel) ProtoMessage() {}

func (x *MarginLevel) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarginLevel.ProtoReflect.Descriptor instead.
func (*MarginLevel) Descriptor() ([]byte, []int) {
	return file_settings_v2_proto_rawDescGZIP(), []int{2}
}

func (x *MarginLevel) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_NOLEVEL
}

func (x *MarginLevel) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type MarginLevelsBrand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand  *Brand         `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Levels []*MarginLevel `protobuf:"bytes,2,rep,name=levels,proto3" json:"levels,omitempty"`
}

func (x *MarginLevelsBrand) Reset() {
	*x = MarginLevelsBrand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarginLevelsBrand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarginLevelsBrand) ProtoMessage() {}

func (x *MarginLevelsBrand) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarginLevelsBrand.ProtoReflect.Descriptor instead.
func (*MarginLevelsBrand) Descriptor() ([]byte, []int) {
	return file_settings_v2_proto_rawDescGZIP(), []int{3}
}

func (x *MarginLevelsBrand) GetBrand() *Brand {
	if x != nil {
		return x.Brand
	}
	return nil
}

func (x *MarginLevelsBrand) GetLevels() []*MarginLevel {
	if x != nil {
		return x.Levels
	}
	return nil
}

type MarginLevels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Levels      []*MarginLevel       `protobuf:"bytes,1,rep,name=levels,proto3" json:"levels,omitempty"`
	BrandLevels []*MarginLevelsBrand `protobuf:"bytes,2,rep,name=brandLevels,proto3" json:"brandLevels,omitempty"`
}

func (x *MarginLevels) Reset() {
	*x = MarginLevels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarginLevels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarginLevels) ProtoMessage() {}

func (x *MarginLevels) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarginLevels.ProtoReflect.Descriptor instead.
func (*MarginLevels) Descriptor() ([]byte, []int) {
	return file_settings_v2_proto_rawDescGZIP(), []int{4}
}

func (x *MarginLevels) GetLevels() []*MarginLevel {
	if x != nil {
		return x.Levels
	}
	return nil
}

func (x *MarginLevels) GetBrandLevels() []*MarginLevelsBrand {
	if x != nil {
		return x.BrandLevels
	}
	return nil
}

type SetGeoPlaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*GeoPlaceData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SetGeoPlaceRequest) Reset() {
	*x = SetGeoPlaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGeoPlaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGeoPlaceRequest) ProtoMessage() {}

func (x *SetGeoPlaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_v2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGeoPlaceRequest.ProtoReflect.Descriptor instead.
func (*SetGeoPlaceRequest) Descriptor() ([]byte, []int) {
	return file_settings_v2_proto_rawDescGZIP(), []int{5}
}

func (x *SetGeoPlaceRequest) GetData() []*GeoPlaceData {
	if x != nil {
		return x.Data
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
	0x63, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x6b, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a,
	0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x22, 0x7e, 0x0a,
	0x0c, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x2e, 0x0a,
	0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x3e, 0x0a,
	0x0b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4d,
	0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x22, 0x41, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x47, 0x65,
	0x6f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xe0, 0x07, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x36, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x1a, 0x11, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x10, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22,
	0x00, 0x12, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74,
	0x65, 0x73, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c,
	0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x54, 0x61, 0x78, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x54, 0x61, 0x78, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4d,
	0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x19,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x53,
	0x65, 0x74, 0x47, 0x65, 0x6f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x65, 0x6f, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_settings_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_settings_v2_proto_goTypes = []interface{}{
	(*AppServiceData)(nil),     // 0: cerasusV2.AppServiceData
	(*UserAppData)(nil),        // 1: cerasusV2.UserAppData
	(*MarginLevel)(nil),        // 2: cerasusV2.MarginLevel
	(*MarginLevelsBrand)(nil),  // 3: cerasusV2.MarginLevelsBrand
	(*MarginLevels)(nil),       // 4: cerasusV2.MarginLevels
	(*SetGeoPlaceRequest)(nil), // 5: cerasusV2.SetGeoPlaceRequest
	(ServiceStatus)(0),         // 6: cerasusV2.ServiceStatus
	(*AppShopData)(nil),        // 7: cerasusV2.AppShopData
	(Level)(0),                 // 8: cerasusV2.Level
	(*Brand)(nil),              // 9: cerasusV2.Brand
	(*GeoPlaceData)(nil),       // 10: cerasusV2.GeoPlaceData
	(*PingRequest)(nil),        // 11: cerasusV2.PingRequest
	(*Auth)(nil),               // 12: cerasusV2.Auth
	(*LineGraphRequest)(nil),   // 13: cerasusV2.LineGraphRequest
	(*RequestByID)(nil),        // 14: cerasusV2.RequestByID
	(*RequestByDates)(nil),     // 15: cerasusV2.RequestByDates
	(*PingReply)(nil),          // 16: cerasusV2.PingReply
	(*LineGraph)(nil),          // 17: cerasusV2.LineGraph
	(*RoundGraphic)(nil),       // 18: cerasusV2.RoundGraphic
	(*WeekGraphic)(nil),        // 19: cerasusV2.WeekGraphic
	(*OrderLeaders)(nil),       // 20: cerasusV2.OrderLeaders
	(*Brands)(nil),             // 21: cerasusV2.Brands
	(*LineGraphics)(nil),       // 22: cerasusV2.LineGraphics
	(*Taxes)(nil),              // 23: cerasusV2.Taxes
	(*ProductWidgets)(nil),     // 24: cerasusV2.ProductWidgets
	(*StatusReply)(nil),        // 25: cerasusV2.StatusReply
}
var file_settings_v2_proto_depIdxs = []int32{
	6,  // 0: cerasusV2.AppServiceData.status:type_name -> cerasusV2.ServiceStatus
	7,  // 1: cerasusV2.UserAppData.shops:type_name -> cerasusV2.AppShopData
	0,  // 2: cerasusV2.UserAppData.services:type_name -> cerasusV2.AppServiceData
	8,  // 3: cerasusV2.MarginLevel.level:type_name -> cerasusV2.Level
	9,  // 4: cerasusV2.MarginLevelsBrand.brand:type_name -> cerasusV2.Brand
	2,  // 5: cerasusV2.MarginLevelsBrand.levels:type_name -> cerasusV2.MarginLevel
	2,  // 6: cerasusV2.MarginLevels.levels:type_name -> cerasusV2.MarginLevel
	3,  // 7: cerasusV2.MarginLevels.brandLevels:type_name -> cerasusV2.MarginLevelsBrand
	10, // 8: cerasusV2.SetGeoPlaceRequest.data:type_name -> cerasusV2.GeoPlaceData
	11, // 9: cerasusV2.Settings.Ping:input_type -> cerasusV2.PingRequest
	12, // 10: cerasusV2.Settings.GetUserAppData:input_type -> cerasusV2.Auth
	13, // 11: cerasusV2.Settings.GetMainGraphic:input_type -> cerasusV2.LineGraphRequest
	12, // 12: cerasusV2.Settings.GetFlowGraphic:input_type -> cerasusV2.Auth
	12, // 13: cerasusV2.Settings.GetMarginGraphic:input_type -> cerasusV2.Auth
	13, // 14: cerasusV2.Settings.GetWeekGraphic:input_type -> cerasusV2.LineGraphRequest
	12, // 15: cerasusV2.Settings.GetOrderLeaders:input_type -> cerasusV2.Auth
	12, // 16: cerasusV2.Settings.GetCompanyBrands:input_type -> cerasusV2.Auth
	14, // 17: cerasusV2.Settings.GetBrand:input_type -> cerasusV2.RequestByID
	15, // 18: cerasusV2.Settings.GetProductGraphics:input_type -> cerasusV2.RequestByDates
	12, // 19: cerasusV2.Settings.GetTaxes:input_type -> cerasusV2.Auth
	12, // 20: cerasusV2.Settings.GetMarginLevels:input_type -> cerasusV2.Auth
	15, // 21: cerasusV2.Settings.GetProductWidget:input_type -> cerasusV2.RequestByDates
	15, // 22: cerasusV2.Settings.GetProductWidgetOrders:input_type -> cerasusV2.RequestByDates
	5,  // 23: cerasusV2.Settings.SetGeoPlace:input_type -> cerasusV2.SetGeoPlaceRequest
	16, // 24: cerasusV2.Settings.Ping:output_type -> cerasusV2.PingReply
	1,  // 25: cerasusV2.Settings.GetUserAppData:output_type -> cerasusV2.UserAppData
	17, // 26: cerasusV2.Settings.GetMainGraphic:output_type -> cerasusV2.LineGraph
	18, // 27: cerasusV2.Settings.GetFlowGraphic:output_type -> cerasusV2.RoundGraphic
	18, // 28: cerasusV2.Settings.GetMarginGraphic:output_type -> cerasusV2.RoundGraphic
	19, // 29: cerasusV2.Settings.GetWeekGraphic:output_type -> cerasusV2.WeekGraphic
	20, // 30: cerasusV2.Settings.GetOrderLeaders:output_type -> cerasusV2.OrderLeaders
	21, // 31: cerasusV2.Settings.GetCompanyBrands:output_type -> cerasusV2.Brands
	9,  // 32: cerasusV2.Settings.GetBrand:output_type -> cerasusV2.Brand
	22, // 33: cerasusV2.Settings.GetProductGraphics:output_type -> cerasusV2.LineGraphics
	23, // 34: cerasusV2.Settings.GetTaxes:output_type -> cerasusV2.Taxes
	4,  // 35: cerasusV2.Settings.GetMarginLevels:output_type -> cerasusV2.MarginLevels
	24, // 36: cerasusV2.Settings.GetProductWidget:output_type -> cerasusV2.ProductWidgets
	24, // 37: cerasusV2.Settings.GetProductWidgetOrders:output_type -> cerasusV2.ProductWidgets
	25, // 38: cerasusV2.Settings.SetGeoPlace:output_type -> cerasusV2.StatusReply
	24, // [24:39] is the sub-list for method output_type
	9,  // [9:24] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
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
		file_settings_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarginLevel); i {
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
		file_settings_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarginLevelsBrand); i {
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
		file_settings_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarginLevels); i {
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
		file_settings_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGeoPlaceRequest); i {
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
			NumMessages:   6,
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
