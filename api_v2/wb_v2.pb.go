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
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WBCounterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Warehouses   []int64        `protobuf:"varint,2,rep,packed,name=warehouses,proto3" json:"warehouses,omitempty"`
	ShopProducts []*ShopProduct `protobuf:"bytes,3,rep,name=shopProducts,proto3" json:"shopProducts,omitempty"`
}

func (x *WBCounterParams) Reset() {
	*x = WBCounterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wb_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WBCounterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WBCounterParams) ProtoMessage() {}

func (x *WBCounterParams) ProtoReflect() protoreflect.Message {
	mi := &file_wb_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WBCounterParams.ProtoReflect.Descriptor instead.
func (*WBCounterParams) Descriptor() ([]byte, []int) {
	return file_wb_v2_proto_rawDescGZIP(), []int{0}
}

func (x *WBCounterParams) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WBCounterParams) GetWarehouses() []int64 {
	if x != nil {
		return x.Warehouses
	}
	return nil
}

func (x *WBCounterParams) GetShopProducts() []*ShopProduct {
	if x != nil {
		return x.ShopProducts
	}
	return nil
}

var File_wb_v2_proto protoreflect.FileDescriptor

var file_wb_v2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x62, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x57,
	0x42, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x32, 0xca, 0x0d, 0x0a, 0x02, 0x57, 0x42, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x16,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x68,
	0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x6e,
	0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a,
	0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x57,
	0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x73, 0x1a, 0x1a, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x16, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x69, 0x63, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73,
	0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e,
	0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74,
	0x65, 0x73, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x12,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x61, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x57, 0x42,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x18,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61,
	0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wb_v2_proto_rawDescOnce sync.Once
	file_wb_v2_proto_rawDescData = file_wb_v2_proto_rawDesc
)

func file_wb_v2_proto_rawDescGZIP() []byte {
	file_wb_v2_proto_rawDescOnce.Do(func() {
		file_wb_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_wb_v2_proto_rawDescData)
	})
	return file_wb_v2_proto_rawDescData
}

var file_wb_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wb_v2_proto_goTypes = []interface{}{
	(*WBCounterParams)(nil),      // 0: cerasusV2.WBCounterParams
	(*ShopProduct)(nil),          // 1: cerasusV2.ShopProduct
	(*PingRequest)(nil),          // 2: cerasusV2.PingRequest
	(*Auth)(nil),                 // 3: cerasusV2.Auth
	(*LineGraphRequest)(nil),     // 4: cerasusV2.LineGraphRequest
	(*RequestByIDs)(nil),         // 5: cerasusV2.RequestByIDs
	(*RequestByID)(nil),          // 6: cerasusV2.RequestByID
	(*RequestByDates)(nil),       // 7: cerasusV2.RequestByDates
	(*SaleRequest)(nil),          // 8: cerasusV2.SaleRequest
	(*ShopProductUrlSetter)(nil), // 9: cerasusV2.ShopProductUrlSetter
	(*ImageRequest)(nil),         // 10: cerasusV2.ImageRequest
	(*SearchRequest)(nil),        // 11: cerasusV2.SearchRequest
	(*OrdersRequest)(nil),        // 12: cerasusV2.OrdersRequest
	(*PingReply)(nil),            // 13: cerasusV2.PingReply
	(*AppShopData)(nil),          // 14: cerasusV2.AppShopData
	(*ShopData)(nil),             // 15: cerasusV2.ShopData
	(*ShopWidget)(nil),           // 16: cerasusV2.ShopWidget
	(*LineGraph)(nil),            // 17: cerasusV2.LineGraph
	(*Count)(nil),                // 18: cerasusV2.Count
	(*WeekGraphic)(nil),          // 19: cerasusV2.WeekGraphic
	(*RoundGraphic)(nil),         // 20: cerasusV2.RoundGraphic
	(*OrderLeaders)(nil),         // 21: cerasusV2.OrderLeaders
	(*ShopProductList)(nil),      // 22: cerasusV2.ShopProductList
	(*LineGraphics)(nil),         // 23: cerasusV2.LineGraphics
	(*Sales)(nil),                // 24: cerasusV2.Sales
	(*ProductWidgets)(nil),       // 25: cerasusV2.ProductWidgets
	(*Sale)(nil),                 // 26: cerasusV2.Sale
	(*StatusReply)(nil),          // 27: cerasusV2.StatusReply
	(*ImageReply)(nil),           // 28: cerasusV2.ImageReply
	(*Orders)(nil),               // 29: cerasusV2.Orders
}
var file_wb_v2_proto_depIdxs = []int32{
	1,  // 0: cerasusV2.WBCounterParams.shopProducts:type_name -> cerasusV2.ShopProduct
	2,  // 1: cerasusV2.WB.Ping:input_type -> cerasusV2.PingRequest
	3,  // 2: cerasusV2.WB.GetAppData:input_type -> cerasusV2.Auth
	3,  // 3: cerasusV2.WB.GetShopData:input_type -> cerasusV2.Auth
	3,  // 4: cerasusV2.WB.GetShopWidget:input_type -> cerasusV2.Auth
	4,  // 5: cerasusV2.WB.GetMainGraphic:input_type -> cerasusV2.LineGraphRequest
	3,  // 6: cerasusV2.WB.GetProductsCount:input_type -> cerasusV2.Auth
	3,  // 7: cerasusV2.WB.GetProductsCountUnsorted:input_type -> cerasusV2.Auth
	3,  // 8: cerasusV2.WB.GetFlowGraphicData:input_type -> cerasusV2.Auth
	3,  // 9: cerasusV2.WB.GetMarginGraphicData:input_type -> cerasusV2.Auth
	4,  // 10: cerasusV2.WB.GetWeekGraphic:input_type -> cerasusV2.LineGraphRequest
	3,  // 11: cerasusV2.WB.GetPayRoundGraphic:input_type -> cerasusV2.Auth
	3,  // 12: cerasusV2.WB.GetCountRoundGraphic:input_type -> cerasusV2.Auth
	3,  // 13: cerasusV2.WB.GetOrderLeaders:input_type -> cerasusV2.Auth
	5,  // 14: cerasusV2.WB.GetShopProducts:input_type -> cerasusV2.RequestByIDs
	6,  // 15: cerasusV2.WB.GetShopProduct:input_type -> cerasusV2.RequestByID
	7,  // 16: cerasusV2.WB.GetProductGraphics:input_type -> cerasusV2.RequestByDates
	7,  // 17: cerasusV2.WB.GetSales:input_type -> cerasusV2.RequestByDates
	7,  // 18: cerasusV2.WB.GetProductWidget:input_type -> cerasusV2.RequestByDates
	7,  // 19: cerasusV2.WB.GetProductWidgetOrders:input_type -> cerasusV2.RequestByDates
	8,  // 20: cerasusV2.WB.GetSale:input_type -> cerasusV2.SaleRequest
	3,  // 21: cerasusV2.WB.GetProductsUnsortedList:input_type -> cerasusV2.Auth
	9,  // 22: cerasusV2.WB.SetShopProductUrl:input_type -> cerasusV2.ShopProductUrlSetter
	5,  // 23: cerasusV2.WB.GetCounterParams:input_type -> cerasusV2.RequestByIDs
	10, // 24: cerasusV2.WB.GetImage:input_type -> cerasusV2.ImageRequest
	11, // 25: cerasusV2.WB.GetShopProductByCode:input_type -> cerasusV2.SearchRequest
	12, // 26: cerasusV2.WB.GetOrdersForBrand:input_type -> cerasusV2.OrdersRequest
	13, // 27: cerasusV2.WB.Ping:output_type -> cerasusV2.PingReply
	14, // 28: cerasusV2.WB.GetAppData:output_type -> cerasusV2.AppShopData
	15, // 29: cerasusV2.WB.GetShopData:output_type -> cerasusV2.ShopData
	16, // 30: cerasusV2.WB.GetShopWidget:output_type -> cerasusV2.ShopWidget
	17, // 31: cerasusV2.WB.GetMainGraphic:output_type -> cerasusV2.LineGraph
	18, // 32: cerasusV2.WB.GetProductsCount:output_type -> cerasusV2.Count
	18, // 33: cerasusV2.WB.GetProductsCountUnsorted:output_type -> cerasusV2.Count
	18, // 34: cerasusV2.WB.GetFlowGraphicData:output_type -> cerasusV2.Count
	18, // 35: cerasusV2.WB.GetMarginGraphicData:output_type -> cerasusV2.Count
	19, // 36: cerasusV2.WB.GetWeekGraphic:output_type -> cerasusV2.WeekGraphic
	20, // 37: cerasusV2.WB.GetPayRoundGraphic:output_type -> cerasusV2.RoundGraphic
	20, // 38: cerasusV2.WB.GetCountRoundGraphic:output_type -> cerasusV2.RoundGraphic
	21, // 39: cerasusV2.WB.GetOrderLeaders:output_type -> cerasusV2.OrderLeaders
	22, // 40: cerasusV2.WB.GetShopProducts:output_type -> cerasusV2.ShopProductList
	1,  // 41: cerasusV2.WB.GetShopProduct:output_type -> cerasusV2.ShopProduct
	23, // 42: cerasusV2.WB.GetProductGraphics:output_type -> cerasusV2.LineGraphics
	24, // 43: cerasusV2.WB.GetSales:output_type -> cerasusV2.Sales
	25, // 44: cerasusV2.WB.GetProductWidget:output_type -> cerasusV2.ProductWidgets
	25, // 45: cerasusV2.WB.GetProductWidgetOrders:output_type -> cerasusV2.ProductWidgets
	26, // 46: cerasusV2.WB.GetSale:output_type -> cerasusV2.Sale
	22, // 47: cerasusV2.WB.GetProductsUnsortedList:output_type -> cerasusV2.ShopProductList
	27, // 48: cerasusV2.WB.SetShopProductUrl:output_type -> cerasusV2.StatusReply
	0,  // 49: cerasusV2.WB.GetCounterParams:output_type -> cerasusV2.WBCounterParams
	28, // 50: cerasusV2.WB.GetImage:output_type -> cerasusV2.ImageReply
	1,  // 51: cerasusV2.WB.GetShopProductByCode:output_type -> cerasusV2.ShopProduct
	29, // 52: cerasusV2.WB.GetOrdersForBrand:output_type -> cerasusV2.Orders
	27, // [27:53] is the sub-list for method output_type
	1,  // [1:27] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_wb_v2_proto_init() }
func file_wb_v2_proto_init() {
	if File_wb_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wb_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WBCounterParams); i {
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
			RawDescriptor: file_wb_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wb_v2_proto_goTypes,
		DependencyIndexes: file_wb_v2_proto_depIdxs,
		MessageInfos:      file_wb_v2_proto_msgTypes,
	}.Build()
	File_wb_v2_proto = out.File
	file_wb_v2_proto_rawDesc = nil
	file_wb_v2_proto_goTypes = nil
	file_wb_v2_proto_depIdxs = nil
}
