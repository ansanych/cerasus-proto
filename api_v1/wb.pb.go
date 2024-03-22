// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: wb.proto

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

type ShopWBAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Error bool   `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ShopWBAuth) Reset() {
	*x = ShopWBAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopWBAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopWBAuth) ProtoMessage() {}

func (x *ShopWBAuth) ProtoReflect() protoreflect.Message {
	mi := &file_wb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopWBAuth.ProtoReflect.Descriptor instead.
func (*ShopWBAuth) Descriptor() ([]byte, []int) {
	return file_wb_proto_rawDescGZIP(), []int{0}
}

func (x *ShopWBAuth) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ShopWBAuth) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type SetWBAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *Auth       `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ShopAuth *ShopWBAuth `protobuf:"bytes,2,opt,name=shopAuth,proto3" json:"shopAuth,omitempty"`
}

func (x *SetWBAuth) Reset() {
	*x = SetWBAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetWBAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetWBAuth) ProtoMessage() {}

func (x *SetWBAuth) ProtoReflect() protoreflect.Message {
	mi := &file_wb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetWBAuth.ProtoReflect.Descriptor instead.
func (*SetWBAuth) Descriptor() ([]byte, []int) {
	return file_wb_proto_rawDescGZIP(), []int{1}
}

func (x *SetWBAuth) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *SetWBAuth) GetShopAuth() *ShopWBAuth {
	if x != nil {
		return x.ShopAuth
	}
	return nil
}

var File_wb_proto protoreflect.FileDescriptor

var file_wb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x77, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x42, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x57, 0x42, 0x41, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x08,
	0x73, 0x68, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x42, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x32, 0x90, 0x08,
	0x0a, 0x02, 0x57, 0x42, 0x12, 0x33, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x57, 0x42, 0x41,
	0x75, 0x74, 0x68, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x57, 0x42, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a,
	0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x79, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x44, 0x61, 0x79, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x12, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x6f,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_proto_rawDescOnce sync.Once
	file_wb_proto_rawDescData = file_wb_proto_rawDesc
)

func file_wb_proto_rawDescGZIP() []byte {
	file_wb_proto_rawDescOnce.Do(func() {
		file_wb_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_proto_rawDescData)
	})
	return file_wb_proto_rawDescData
}

var file_wb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wb_proto_goTypes = []interface{}{
	(*ShopWBAuth)(nil),             // 0: cerasus.ShopWBAuth
	(*SetWBAuth)(nil),              // 1: cerasus.SetWBAuth
	(*Auth)(nil),                   // 2: cerasus.Auth
	(*ShopProductListRequest)(nil), // 3: cerasus.ShopProductListRequest
	(*ShopProductRequest)(nil),     // 4: cerasus.ShopProductRequest
	(*SalesRequest)(nil),           // 5: cerasus.SalesRequest
	(*ShopServiceRequest)(nil),     // 6: cerasus.ShopServiceRequest
	(*SaleDetailsRequest)(nil),     // 7: cerasus.SaleDetailsRequest
	(*ProductSalesRequest)(nil),    // 8: cerasus.ProductSalesRequest
	(*MainGraphicRequest)(nil),     // 9: cerasus.MainGraphicRequest
	(*ImageRequest)(nil),           // 10: cerasus.ImageRequest
	(*BoolReply)(nil),              // 11: cerasus.BoolReply
	(*CountReply)(nil),             // 12: cerasus.CountReply
	(*ShopProductListReply)(nil),   // 13: cerasus.ShopProductListReply
	(*ShopProduct)(nil),            // 14: cerasus.ShopProduct
	(*DaysSalesReply)(nil),         // 15: cerasus.DaysSalesReply
	(*SalesReply)(nil),             // 16: cerasus.SalesReply
	(*ShopServiceReply)(nil),       // 17: cerasus.ShopServiceReply
	(*SaleDetailsReply)(nil),       // 18: cerasus.SaleDetailsReply
	(*MainGraphicReply)(nil),       // 19: cerasus.MainGraphicReply
	(*ImageReply)(nil),             // 20: cerasus.ImageReply
	(*CompanyShopData)(nil),        // 21: cerasus.CompanyShopData
}
var file_wb_proto_depIdxs = []int32{
	2,  // 0: cerasus.SetWBAuth.auth:type_name -> cerasus.Auth
	0,  // 1: cerasus.SetWBAuth.shopAuth:type_name -> cerasus.ShopWBAuth
	1,  // 2: cerasus.WB.SetAuth:input_type -> cerasus.SetWBAuth
	2,  // 3: cerasus.WB.GetAuth:input_type -> cerasus.Auth
	2,  // 4: cerasus.WB.ErrorAuth:input_type -> cerasus.Auth
	2,  // 5: cerasus.WB.GetUnsortedCount:input_type -> cerasus.Auth
	3,  // 6: cerasus.WB.GetUnsortedList:input_type -> cerasus.ShopProductListRequest
	2,  // 7: cerasus.WB.GetProductCount:input_type -> cerasus.Auth
	3,  // 8: cerasus.WB.GetProductList:input_type -> cerasus.ShopProductListRequest
	4,  // 9: cerasus.WB.GetProduct:input_type -> cerasus.ShopProductRequest
	2,  // 10: cerasus.WB.GetDaySales:input_type -> cerasus.Auth
	5,  // 11: cerasus.WB.GetSales:input_type -> cerasus.SalesRequest
	6,  // 12: cerasus.WB.GetShopServices:input_type -> cerasus.ShopServiceRequest
	7,  // 13: cerasus.WB.GetSaleDetail:input_type -> cerasus.SaleDetailsRequest
	8,  // 14: cerasus.WB.GetProductSales:input_type -> cerasus.ProductSalesRequest
	9,  // 15: cerasus.WB.GetMainGraphic:input_type -> cerasus.MainGraphicRequest
	10, // 16: cerasus.WB.GetImage:input_type -> cerasus.ImageRequest
	2,  // 17: cerasus.WB.CheckShopData:input_type -> cerasus.Auth
	11, // 18: cerasus.WB.SetAuth:output_type -> cerasus.BoolReply
	0,  // 19: cerasus.WB.GetAuth:output_type -> cerasus.ShopWBAuth
	11, // 20: cerasus.WB.ErrorAuth:output_type -> cerasus.BoolReply
	12, // 21: cerasus.WB.GetUnsortedCount:output_type -> cerasus.CountReply
	13, // 22: cerasus.WB.GetUnsortedList:output_type -> cerasus.ShopProductListReply
	12, // 23: cerasus.WB.GetProductCount:output_type -> cerasus.CountReply
	13, // 24: cerasus.WB.GetProductList:output_type -> cerasus.ShopProductListReply
	14, // 25: cerasus.WB.GetProduct:output_type -> cerasus.ShopProduct
	15, // 26: cerasus.WB.GetDaySales:output_type -> cerasus.DaysSalesReply
	16, // 27: cerasus.WB.GetSales:output_type -> cerasus.SalesReply
	17, // 28: cerasus.WB.GetShopServices:output_type -> cerasus.ShopServiceReply
	18, // 29: cerasus.WB.GetSaleDetail:output_type -> cerasus.SaleDetailsReply
	16, // 30: cerasus.WB.GetProductSales:output_type -> cerasus.SalesReply
	19, // 31: cerasus.WB.GetMainGraphic:output_type -> cerasus.MainGraphicReply
	20, // 32: cerasus.WB.GetImage:output_type -> cerasus.ImageReply
	21, // 33: cerasus.WB.CheckShopData:output_type -> cerasus.CompanyShopData
	18, // [18:34] is the sub-list for method output_type
	2,  // [2:18] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_wb_proto_init() }
func file_wb_proto_init() {
	if File_wb_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopWBAuth); i {
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
		file_wb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetWBAuth); i {
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
			RawDescriptor: file_wb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wb_proto_goTypes,
		DependencyIndexes: file_wb_proto_depIdxs,
		MessageInfos:      file_wb_proto_msgTypes,
	}.Build()
	File_wb_proto = out.File
	file_wb_proto_rawDesc = nil
	file_wb_proto_goTypes = nil
	file_wb_proto_depIdxs = nil
}
