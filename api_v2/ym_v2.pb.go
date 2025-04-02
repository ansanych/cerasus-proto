// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ym_v2.proto

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

type YMCounterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string         `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	BusinessID   int64          `protobuf:"varint,2,opt,name=businessID,proto3" json:"businessID,omitempty"`
	Campaigns    []*YMCampaign  `protobuf:"bytes,3,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	ShopProducts []*ShopProduct `protobuf:"bytes,4,rep,name=shopProducts,proto3" json:"shopProducts,omitempty"`
}

func (x *YMCounterParams) Reset() {
	*x = YMCounterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YMCounterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YMCounterParams) ProtoMessage() {}

func (x *YMCounterParams) ProtoReflect() protoreflect.Message {
	mi := &file_ym_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YMCounterParams.ProtoReflect.Descriptor instead.
func (*YMCounterParams) Descriptor() ([]byte, []int) {
	return file_ym_v2_proto_rawDescGZIP(), []int{0}
}

func (x *YMCounterParams) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *YMCounterParams) GetBusinessID() int64 {
	if x != nil {
		return x.BusinessID
	}
	return 0
}

func (x *YMCounterParams) GetCampaigns() []*YMCampaign {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

func (x *YMCounterParams) GetShopProducts() []*ShopProduct {
	if x != nil {
		return x.ShopProducts
	}
	return nil
}

var File_ym_v2_proto protoreflect.FileDescriptor

var file_ym_v2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x79, 0x6d, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0f, 0x59,
	0x4d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x59, 0x4d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x09,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xca, 0x0d, 0x0a, 0x02, 0x59, 0x4d, 0x12, 0x36, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x41, 0x70, 0x70, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69,
	0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69,
	0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69,
	0x63, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x0f, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x17, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f,
	0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x56, 0x32, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x19, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x53, 0x61, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x55, 0x6e,
	0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x1a, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x11, 0x53, 0x65, 0x74,
	0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1f,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a,
	0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x59, 0x4d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ym_v2_proto_rawDescOnce sync.Once
	file_ym_v2_proto_rawDescData = file_ym_v2_proto_rawDesc
)

func file_ym_v2_proto_rawDescGZIP() []byte {
	file_ym_v2_proto_rawDescOnce.Do(func() {
		file_ym_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_ym_v2_proto_rawDescData)
	})
	return file_ym_v2_proto_rawDescData
}

var file_ym_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ym_v2_proto_goTypes = []interface{}{
	(*YMCounterParams)(nil),      // 0: cerasusV2.YMCounterParams
	(*YMCampaign)(nil),           // 1: cerasusV2.YMCampaign
	(*ShopProduct)(nil),          // 2: cerasusV2.ShopProduct
	(*PingRequest)(nil),          // 3: cerasusV2.PingRequest
	(*Auth)(nil),                 // 4: cerasusV2.Auth
	(*LineGraphRequest)(nil),     // 5: cerasusV2.LineGraphRequest
	(*RequestByIDs)(nil),         // 6: cerasusV2.RequestByIDs
	(*RequestByID)(nil),          // 7: cerasusV2.RequestByID
	(*RequestByDates)(nil),       // 8: cerasusV2.RequestByDates
	(*SaleRequest)(nil),          // 9: cerasusV2.SaleRequest
	(*ShopProductUrlSetter)(nil), // 10: cerasusV2.ShopProductUrlSetter
	(*ImageRequest)(nil),         // 11: cerasusV2.ImageRequest
	(*SearchRequest)(nil),        // 12: cerasusV2.SearchRequest
	(*OrdersRequest)(nil),        // 13: cerasusV2.OrdersRequest
	(*PingReply)(nil),            // 14: cerasusV2.PingReply
	(*AppShopData)(nil),          // 15: cerasusV2.AppShopData
	(*ShopData)(nil),             // 16: cerasusV2.ShopData
	(*ShopWidget)(nil),           // 17: cerasusV2.ShopWidget
	(*LineGraph)(nil),            // 18: cerasusV2.LineGraph
	(*Count)(nil),                // 19: cerasusV2.Count
	(*WeekGraphic)(nil),          // 20: cerasusV2.WeekGraphic
	(*RoundGraphic)(nil),         // 21: cerasusV2.RoundGraphic
	(*OrderLeaders)(nil),         // 22: cerasusV2.OrderLeaders
	(*ShopProductList)(nil),      // 23: cerasusV2.ShopProductList
	(*LineGraphics)(nil),         // 24: cerasusV2.LineGraphics
	(*Sales)(nil),                // 25: cerasusV2.Sales
	(*ProductWidgets)(nil),       // 26: cerasusV2.ProductWidgets
	(*Sale)(nil),                 // 27: cerasusV2.Sale
	(*StatusReply)(nil),          // 28: cerasusV2.StatusReply
	(*ImageReply)(nil),           // 29: cerasusV2.ImageReply
	(*Orders)(nil),               // 30: cerasusV2.Orders
}
var file_ym_v2_proto_depIdxs = []int32{
	1,  // 0: cerasusV2.YMCounterParams.campaigns:type_name -> cerasusV2.YMCampaign
	2,  // 1: cerasusV2.YMCounterParams.shopProducts:type_name -> cerasusV2.ShopProduct
	3,  // 2: cerasusV2.YM.Ping:input_type -> cerasusV2.PingRequest
	4,  // 3: cerasusV2.YM.GetAppData:input_type -> cerasusV2.Auth
	4,  // 4: cerasusV2.YM.GetShopData:input_type -> cerasusV2.Auth
	4,  // 5: cerasusV2.YM.GetShopWidget:input_type -> cerasusV2.Auth
	5,  // 6: cerasusV2.YM.GetMainGraphic:input_type -> cerasusV2.LineGraphRequest
	4,  // 7: cerasusV2.YM.GetProductsCount:input_type -> cerasusV2.Auth
	4,  // 8: cerasusV2.YM.GetProductsCountUnsorted:input_type -> cerasusV2.Auth
	4,  // 9: cerasusV2.YM.GetFlowGraphicData:input_type -> cerasusV2.Auth
	4,  // 10: cerasusV2.YM.GetMarginGraphicData:input_type -> cerasusV2.Auth
	5,  // 11: cerasusV2.YM.GetWeekGraphic:input_type -> cerasusV2.LineGraphRequest
	4,  // 12: cerasusV2.YM.GetPayRoundGraphic:input_type -> cerasusV2.Auth
	4,  // 13: cerasusV2.YM.GetCountRoundGraphic:input_type -> cerasusV2.Auth
	4,  // 14: cerasusV2.YM.GetOrderLeaders:input_type -> cerasusV2.Auth
	6,  // 15: cerasusV2.YM.GetShopProducts:input_type -> cerasusV2.RequestByIDs
	7,  // 16: cerasusV2.YM.GetShopProduct:input_type -> cerasusV2.RequestByID
	8,  // 17: cerasusV2.YM.GetProductGraphics:input_type -> cerasusV2.RequestByDates
	8,  // 18: cerasusV2.YM.GetSales:input_type -> cerasusV2.RequestByDates
	8,  // 19: cerasusV2.YM.GetProductWidget:input_type -> cerasusV2.RequestByDates
	8,  // 20: cerasusV2.YM.GetProductWidgetOrders:input_type -> cerasusV2.RequestByDates
	9,  // 21: cerasusV2.YM.GetSale:input_type -> cerasusV2.SaleRequest
	4,  // 22: cerasusV2.YM.GetProductsUnsortedList:input_type -> cerasusV2.Auth
	10, // 23: cerasusV2.YM.SetShopProductUrl:input_type -> cerasusV2.ShopProductUrlSetter
	6,  // 24: cerasusV2.YM.GetCounterParams:input_type -> cerasusV2.RequestByIDs
	11, // 25: cerasusV2.YM.GetImage:input_type -> cerasusV2.ImageRequest
	12, // 26: cerasusV2.YM.GetShopProductByCode:input_type -> cerasusV2.SearchRequest
	13, // 27: cerasusV2.YM.GetOrdersForBrand:input_type -> cerasusV2.OrdersRequest
	14, // 28: cerasusV2.YM.Ping:output_type -> cerasusV2.PingReply
	15, // 29: cerasusV2.YM.GetAppData:output_type -> cerasusV2.AppShopData
	16, // 30: cerasusV2.YM.GetShopData:output_type -> cerasusV2.ShopData
	17, // 31: cerasusV2.YM.GetShopWidget:output_type -> cerasusV2.ShopWidget
	18, // 32: cerasusV2.YM.GetMainGraphic:output_type -> cerasusV2.LineGraph
	19, // 33: cerasusV2.YM.GetProductsCount:output_type -> cerasusV2.Count
	19, // 34: cerasusV2.YM.GetProductsCountUnsorted:output_type -> cerasusV2.Count
	19, // 35: cerasusV2.YM.GetFlowGraphicData:output_type -> cerasusV2.Count
	19, // 36: cerasusV2.YM.GetMarginGraphicData:output_type -> cerasusV2.Count
	20, // 37: cerasusV2.YM.GetWeekGraphic:output_type -> cerasusV2.WeekGraphic
	21, // 38: cerasusV2.YM.GetPayRoundGraphic:output_type -> cerasusV2.RoundGraphic
	21, // 39: cerasusV2.YM.GetCountRoundGraphic:output_type -> cerasusV2.RoundGraphic
	22, // 40: cerasusV2.YM.GetOrderLeaders:output_type -> cerasusV2.OrderLeaders
	23, // 41: cerasusV2.YM.GetShopProducts:output_type -> cerasusV2.ShopProductList
	2,  // 42: cerasusV2.YM.GetShopProduct:output_type -> cerasusV2.ShopProduct
	24, // 43: cerasusV2.YM.GetProductGraphics:output_type -> cerasusV2.LineGraphics
	25, // 44: cerasusV2.YM.GetSales:output_type -> cerasusV2.Sales
	26, // 45: cerasusV2.YM.GetProductWidget:output_type -> cerasusV2.ProductWidgets
	26, // 46: cerasusV2.YM.GetProductWidgetOrders:output_type -> cerasusV2.ProductWidgets
	27, // 47: cerasusV2.YM.GetSale:output_type -> cerasusV2.Sale
	23, // 48: cerasusV2.YM.GetProductsUnsortedList:output_type -> cerasusV2.ShopProductList
	28, // 49: cerasusV2.YM.SetShopProductUrl:output_type -> cerasusV2.StatusReply
	0,  // 50: cerasusV2.YM.GetCounterParams:output_type -> cerasusV2.YMCounterParams
	29, // 51: cerasusV2.YM.GetImage:output_type -> cerasusV2.ImageReply
	2,  // 52: cerasusV2.YM.GetShopProductByCode:output_type -> cerasusV2.ShopProduct
	30, // 53: cerasusV2.YM.GetOrdersForBrand:output_type -> cerasusV2.Orders
	28, // [28:54] is the sub-list for method output_type
	2,  // [2:28] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_ym_v2_proto_init() }
func file_ym_v2_proto_init() {
	if File_ym_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ym_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YMCounterParams); i {
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
			RawDescriptor: file_ym_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ym_v2_proto_goTypes,
		DependencyIndexes: file_ym_v2_proto_depIdxs,
		MessageInfos:      file_ym_v2_proto_msgTypes,
	}.Build()
	File_ym_v2_proto = out.File
	file_ym_v2_proto_rawDesc = nil
	file_ym_v2_proto_goTypes = nil
	file_ym_v2_proto_depIdxs = nil
}
