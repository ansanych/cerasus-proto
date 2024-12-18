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

type ForCounterRequestOZ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyID     int64   `protobuf:"varint,1,opt,name=companyID,proto3" json:"companyID,omitempty"`
	ShopProductID []int64 `protobuf:"varint,2,rep,packed,name=shopProductID,proto3" json:"shopProductID,omitempty"`
}

func (x *ForCounterRequestOZ) Reset() {
	*x = ForCounterRequestOZ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForCounterRequestOZ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForCounterRequestOZ) ProtoMessage() {}

func (x *ForCounterRequestOZ) ProtoReflect() protoreflect.Message {
	mi := &file_ozon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForCounterRequestOZ.ProtoReflect.Descriptor instead.
func (*ForCounterRequestOZ) Descriptor() ([]byte, []int) {
	return file_ozon_proto_rawDescGZIP(), []int{2}
}

func (x *ForCounterRequestOZ) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *ForCounterRequestOZ) GetShopProductID() []int64 {
	if x != nil {
		return x.ShopProductID
	}
	return nil
}

type ForCounterProductDataOZ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopProductID int64 `protobuf:"varint,1,opt,name=shopProductID,proto3" json:"shopProductID,omitempty"`
	OzonProductID int64 `protobuf:"varint,2,opt,name=ozonProductID,proto3" json:"ozonProductID,omitempty"`
}

func (x *ForCounterProductDataOZ) Reset() {
	*x = ForCounterProductDataOZ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForCounterProductDataOZ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForCounterProductDataOZ) ProtoMessage() {}

func (x *ForCounterProductDataOZ) ProtoReflect() protoreflect.Message {
	mi := &file_ozon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForCounterProductDataOZ.ProtoReflect.Descriptor instead.
func (*ForCounterProductDataOZ) Descriptor() ([]byte, []int) {
	return file_ozon_proto_rawDescGZIP(), []int{3}
}

func (x *ForCounterProductDataOZ) GetShopProductID() int64 {
	if x != nil {
		return x.ShopProductID
	}
	return 0
}

func (x *ForCounterProductDataOZ) GetOzonProductID() int64 {
	if x != nil {
		return x.OzonProductID
	}
	return 0
}

type CompanyWarehouseOZ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	WarehouseID int64  `protobuf:"varint,3,opt,name=warehouseID,proto3" json:"warehouseID,omitempty"`
	Active      bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Updated     string `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *CompanyWarehouseOZ) Reset() {
	*x = CompanyWarehouseOZ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyWarehouseOZ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyWarehouseOZ) ProtoMessage() {}

func (x *CompanyWarehouseOZ) ProtoReflect() protoreflect.Message {
	mi := &file_ozon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyWarehouseOZ.ProtoReflect.Descriptor instead.
func (*CompanyWarehouseOZ) Descriptor() ([]byte, []int) {
	return file_ozon_proto_rawDescGZIP(), []int{4}
}

func (x *CompanyWarehouseOZ) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CompanyWarehouseOZ) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanyWarehouseOZ) GetWarehouseID() int64 {
	if x != nil {
		return x.WarehouseID
	}
	return 0
}

func (x *CompanyWarehouseOZ) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *CompanyWarehouseOZ) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type ForCounterReplyOZ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopOzonAuth *ShopOzonAuth              `protobuf:"bytes,1,opt,name=shopOzonAuth,proto3" json:"shopOzonAuth,omitempty"`
	Data         []*ForCounterProductDataOZ `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Warehouses   []*CompanyWarehouseOZ      `protobuf:"bytes,3,rep,name=warehouses,proto3" json:"warehouses,omitempty"`
}

func (x *ForCounterReplyOZ) Reset() {
	*x = ForCounterReplyOZ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ozon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForCounterReplyOZ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForCounterReplyOZ) ProtoMessage() {}

func (x *ForCounterReplyOZ) ProtoReflect() protoreflect.Message {
	mi := &file_ozon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForCounterReplyOZ.ProtoReflect.Descriptor instead.
func (*ForCounterReplyOZ) Descriptor() ([]byte, []int) {
	return file_ozon_proto_rawDescGZIP(), []int{5}
}

func (x *ForCounterReplyOZ) GetShopOzonAuth() *ShopOzonAuth {
	if x != nil {
		return x.ShopOzonAuth
	}
	return nil
}

func (x *ForCounterReplyOZ) GetData() []*ForCounterProductDataOZ {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ForCounterReplyOZ) GetWarehouses() []*CompanyWarehouseOZ {
	if x != nil {
		return x.Warehouses
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
	0x75, 0x74, 0x68, 0x22, 0x59, 0x0a, 0x13, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x5a, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x70,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0d, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x65,
	0x0a, 0x17, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x5a, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x24, 0x0a, 0x0d, 0x6f, 0x7a, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x7a, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x8c, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x5a, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4f, 0x5a, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x68,
	0x6f, 0x70, 0x4f, 0x7a, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x4f,
	0x7a, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x4f, 0x7a, 0x6f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x46, 0x6f,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x4f, 0x5a, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0a, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4f, 0x5a, 0x52, 0x0a, 0x77, 0x61,
	0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x32, 0xbe, 0x0c, 0x0a, 0x04, 0x4f, 0x7a, 0x6f,
	0x6e, 0x12, 0x35, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x7a, 0x6f, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x4f, 0x7a, 0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x6e,
	0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x79,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x44,
	0x61, 0x79, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53,
	0x61, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a,
	0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x6e, 0x75, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x12,
	0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x15,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x44, 0x6f, 0x6e, 0x75, 0x74, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x4f, 0x5a, 0x12, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4f, 0x5a, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4f, 0x5a,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x68,
	0x6f, 0x70, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68,
	0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_ozon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ozon_proto_goTypes = []interface{}{
	(*ShopOzonAuth)(nil),             // 0: cerasus.ShopOzonAuth
	(*SetOzonAuth)(nil),              // 1: cerasus.SetOzonAuth
	(*ForCounterRequestOZ)(nil),      // 2: cerasus.ForCounterRequestOZ
	(*ForCounterProductDataOZ)(nil),  // 3: cerasus.ForCounterProductDataOZ
	(*CompanyWarehouseOZ)(nil),       // 4: cerasus.CompanyWarehouseOZ
	(*ForCounterReplyOZ)(nil),        // 5: cerasus.ForCounterReplyOZ
	(*Auth)(nil),                     // 6: cerasus.Auth
	(*ShopProductListRequest)(nil),   // 7: cerasus.ShopProductListRequest
	(*ShopProductRequest)(nil),       // 8: cerasus.ShopProductRequest
	(*ShopProductUpdateRequest)(nil), // 9: cerasus.ShopProductUpdateRequest
	(*SalesRequest)(nil),             // 10: cerasus.SalesRequest
	(*ShopServiceRequest)(nil),       // 11: cerasus.ShopServiceRequest
	(*SaleDetailsRequest)(nil),       // 12: cerasus.SaleDetailsRequest
	(*ProductSalesRequest)(nil),      // 13: cerasus.ProductSalesRequest
	(*MainGraphicRequest)(nil),       // 14: cerasus.MainGraphicRequest
	(*ImageRequest)(nil),             // 15: cerasus.ImageRequest
	(*PingRequest)(nil),              // 16: cerasus.PingRequest
	(*IDRequest)(nil),                // 17: cerasus.IDRequest
	(*ForBrandSalesRequest)(nil),     // 18: cerasus.ForBrandSalesRequest
	(*ForBrandOrdersRequest)(nil),    // 19: cerasus.ForBrandOrdersRequest
	(*BoolReply)(nil),                // 20: cerasus.BoolReply
	(*CountReply)(nil),               // 21: cerasus.CountReply
	(*ShopProductListReply)(nil),     // 22: cerasus.ShopProductListReply
	(*ShopProduct)(nil),              // 23: cerasus.ShopProduct
	(*DaysSalesReply)(nil),           // 24: cerasus.DaysSalesReply
	(*SalesReply)(nil),               // 25: cerasus.SalesReply
	(*ShopServiceReply)(nil),         // 26: cerasus.ShopServiceReply
	(*SaleDetailsReply)(nil),         // 27: cerasus.SaleDetailsReply
	(*MainGraphicReply)(nil),         // 28: cerasus.MainGraphicReply
	(*ImageReply)(nil),               // 29: cerasus.ImageReply
	(*CompanyShopData)(nil),          // 30: cerasus.CompanyShopData
	(*DonutGraphic)(nil),             // 31: cerasus.DonutGraphic
	(*WeekGraphics)(nil),             // 32: cerasus.WeekGraphics
	(*PingReply)(nil),                // 33: cerasus.PingReply
	(*ProductShopUrls)(nil),          // 34: cerasus.ProductShopUrls
	(*ForBrandSalesReply)(nil),       // 35: cerasus.ForBrandSalesReply
	(*ForBrandOrdersReply)(nil),      // 36: cerasus.ForBrandOrdersReply
}
var file_ozon_proto_depIdxs = []int32{
	6,  // 0: cerasus.SetOzonAuth.auth:type_name -> cerasus.Auth
	0,  // 1: cerasus.SetOzonAuth.shopAuth:type_name -> cerasus.ShopOzonAuth
	0,  // 2: cerasus.ForCounterReplyOZ.shopOzonAuth:type_name -> cerasus.ShopOzonAuth
	3,  // 3: cerasus.ForCounterReplyOZ.data:type_name -> cerasus.ForCounterProductDataOZ
	4,  // 4: cerasus.ForCounterReplyOZ.warehouses:type_name -> cerasus.CompanyWarehouseOZ
	1,  // 5: cerasus.Ozon.SetAuth:input_type -> cerasus.SetOzonAuth
	6,  // 6: cerasus.Ozon.GetAuth:input_type -> cerasus.Auth
	6,  // 7: cerasus.Ozon.ErrorAuth:input_type -> cerasus.Auth
	6,  // 8: cerasus.Ozon.GetUnsortedCount:input_type -> cerasus.Auth
	7,  // 9: cerasus.Ozon.GetUnsortedList:input_type -> cerasus.ShopProductListRequest
	6,  // 10: cerasus.Ozon.GetProductCount:input_type -> cerasus.Auth
	7,  // 11: cerasus.Ozon.GetProductList:input_type -> cerasus.ShopProductListRequest
	8,  // 12: cerasus.Ozon.GetProduct:input_type -> cerasus.ShopProductRequest
	9,  // 13: cerasus.Ozon.UpdateProduct:input_type -> cerasus.ShopProductUpdateRequest
	6,  // 14: cerasus.Ozon.GetDaySales:input_type -> cerasus.Auth
	10, // 15: cerasus.Ozon.GetSales:input_type -> cerasus.SalesRequest
	11, // 16: cerasus.Ozon.GetShopServices:input_type -> cerasus.ShopServiceRequest
	12, // 17: cerasus.Ozon.GetSaleDetail:input_type -> cerasus.SaleDetailsRequest
	13, // 18: cerasus.Ozon.GetProductSales:input_type -> cerasus.ProductSalesRequest
	14, // 19: cerasus.Ozon.GetMainGraphic:input_type -> cerasus.MainGraphicRequest
	15, // 20: cerasus.Ozon.GetImage:input_type -> cerasus.ImageRequest
	6,  // 21: cerasus.Ozon.CheckShopData:input_type -> cerasus.Auth
	6,  // 22: cerasus.Ozon.GetDonutGraphics:input_type -> cerasus.Auth
	6,  // 23: cerasus.Ozon.GetWeekGraphics:input_type -> cerasus.Auth
	2,  // 24: cerasus.Ozon.ForCounterDataOZ:input_type -> cerasus.ForCounterRequestOZ
	16, // 25: cerasus.Ozon.Ping:input_type -> cerasus.PingRequest
	17, // 26: cerasus.Ozon.GetProductUrls:input_type -> cerasus.IDRequest
	18, // 27: cerasus.Ozon.ForBrandSales:input_type -> cerasus.ForBrandSalesRequest
	19, // 28: cerasus.Ozon.ForBrandOrders:input_type -> cerasus.ForBrandOrdersRequest
	20, // 29: cerasus.Ozon.SetAuth:output_type -> cerasus.BoolReply
	0,  // 30: cerasus.Ozon.GetAuth:output_type -> cerasus.ShopOzonAuth
	20, // 31: cerasus.Ozon.ErrorAuth:output_type -> cerasus.BoolReply
	21, // 32: cerasus.Ozon.GetUnsortedCount:output_type -> cerasus.CountReply
	22, // 33: cerasus.Ozon.GetUnsortedList:output_type -> cerasus.ShopProductListReply
	21, // 34: cerasus.Ozon.GetProductCount:output_type -> cerasus.CountReply
	22, // 35: cerasus.Ozon.GetProductList:output_type -> cerasus.ShopProductListReply
	23, // 36: cerasus.Ozon.GetProduct:output_type -> cerasus.ShopProduct
	20, // 37: cerasus.Ozon.UpdateProduct:output_type -> cerasus.BoolReply
	24, // 38: cerasus.Ozon.GetDaySales:output_type -> cerasus.DaysSalesReply
	25, // 39: cerasus.Ozon.GetSales:output_type -> cerasus.SalesReply
	26, // 40: cerasus.Ozon.GetShopServices:output_type -> cerasus.ShopServiceReply
	27, // 41: cerasus.Ozon.GetSaleDetail:output_type -> cerasus.SaleDetailsReply
	25, // 42: cerasus.Ozon.GetProductSales:output_type -> cerasus.SalesReply
	28, // 43: cerasus.Ozon.GetMainGraphic:output_type -> cerasus.MainGraphicReply
	29, // 44: cerasus.Ozon.GetImage:output_type -> cerasus.ImageReply
	30, // 45: cerasus.Ozon.CheckShopData:output_type -> cerasus.CompanyShopData
	31, // 46: cerasus.Ozon.GetDonutGraphics:output_type -> cerasus.DonutGraphic
	32, // 47: cerasus.Ozon.GetWeekGraphics:output_type -> cerasus.WeekGraphics
	5,  // 48: cerasus.Ozon.ForCounterDataOZ:output_type -> cerasus.ForCounterReplyOZ
	33, // 49: cerasus.Ozon.Ping:output_type -> cerasus.PingReply
	34, // 50: cerasus.Ozon.GetProductUrls:output_type -> cerasus.ProductShopUrls
	35, // 51: cerasus.Ozon.ForBrandSales:output_type -> cerasus.ForBrandSalesReply
	36, // 52: cerasus.Ozon.ForBrandOrders:output_type -> cerasus.ForBrandOrdersReply
	29, // [29:53] is the sub-list for method output_type
	5,  // [5:29] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
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
		file_ozon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForCounterRequestOZ); i {
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
		file_ozon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForCounterProductDataOZ); i {
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
		file_ozon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyWarehouseOZ); i {
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
		file_ozon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForCounterReplyOZ); i {
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
			NumMessages:   6,
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
