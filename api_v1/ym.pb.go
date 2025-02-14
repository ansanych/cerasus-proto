// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ym.proto

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

type ShopYMAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	BusinessID int32  `protobuf:"varint,2,opt,name=businessID,proto3" json:"businessID,omitempty"`
	Error      bool   `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ShopYMAuth) Reset() {
	*x = ShopYMAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopYMAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopYMAuth) ProtoMessage() {}

func (x *ShopYMAuth) ProtoReflect() protoreflect.Message {
	mi := &file_ym_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopYMAuth.ProtoReflect.Descriptor instead.
func (*ShopYMAuth) Descriptor() ([]byte, []int) {
	return file_ym_proto_rawDescGZIP(), []int{0}
}

func (x *ShopYMAuth) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ShopYMAuth) GetBusinessID() int32 {
	if x != nil {
		return x.BusinessID
	}
	return 0
}

func (x *ShopYMAuth) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type SetYMAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth     *Auth       `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ShopAuth *ShopYMAuth `protobuf:"bytes,2,opt,name=shopAuth,proto3" json:"shopAuth,omitempty"`
}

func (x *SetYMAuth) Reset() {
	*x = SetYMAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetYMAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetYMAuth) ProtoMessage() {}

func (x *SetYMAuth) ProtoReflect() protoreflect.Message {
	mi := &file_ym_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetYMAuth.ProtoReflect.Descriptor instead.
func (*SetYMAuth) Descriptor() ([]byte, []int) {
	return file_ym_proto_rawDescGZIP(), []int{1}
}

func (x *SetYMAuth) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *SetYMAuth) GetShopAuth() *ShopYMAuth {
	if x != nil {
		return x.ShopAuth
	}
	return nil
}

type YMCampaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Domain        string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	CampaignID    int64  `protobuf:"varint,3,opt,name=campaignID,proto3" json:"campaignID,omitempty"`
	ClientID      int64  `protobuf:"varint,4,opt,name=clientID,proto3" json:"clientID,omitempty"`
	PlacementType string `protobuf:"bytes,5,opt,name=placementType,proto3" json:"placementType,omitempty"`
	Updated       string `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *YMCampaign) Reset() {
	*x = YMCampaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YMCampaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YMCampaign) ProtoMessage() {}

func (x *YMCampaign) ProtoReflect() protoreflect.Message {
	mi := &file_ym_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YMCampaign.ProtoReflect.Descriptor instead.
func (*YMCampaign) Descriptor() ([]byte, []int) {
	return file_ym_proto_rawDescGZIP(), []int{2}
}

func (x *YMCampaign) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *YMCampaign) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *YMCampaign) GetCampaignID() int64 {
	if x != nil {
		return x.CampaignID
	}
	return 0
}

func (x *YMCampaign) GetClientID() int64 {
	if x != nil {
		return x.ClientID
	}
	return 0
}

func (x *YMCampaign) GetPlacementType() string {
	if x != nil {
		return x.PlacementType
	}
	return ""
}

func (x *YMCampaign) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type ForCounterRequestYM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyID     int64   `protobuf:"varint,1,opt,name=companyID,proto3" json:"companyID,omitempty"`
	ShopProductID []int64 `protobuf:"varint,2,rep,packed,name=shopProductID,proto3" json:"shopProductID,omitempty"`
}

func (x *ForCounterRequestYM) Reset() {
	*x = ForCounterRequestYM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForCounterRequestYM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForCounterRequestYM) ProtoMessage() {}

func (x *ForCounterRequestYM) ProtoReflect() protoreflect.Message {
	mi := &file_ym_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForCounterRequestYM.ProtoReflect.Descriptor instead.
func (*ForCounterRequestYM) Descriptor() ([]byte, []int) {
	return file_ym_proto_rawDescGZIP(), []int{3}
}

func (x *ForCounterRequestYM) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *ForCounterRequestYM) GetShopProductID() []int64 {
	if x != nil {
		return x.ShopProductID
	}
	return nil
}

type ForCounterProductDataYM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopProductID int64  `protobuf:"varint,1,opt,name=shopProductID,proto3" json:"shopProductID,omitempty"`
	OfferID       string `protobuf:"bytes,2,opt,name=offerID,proto3" json:"offerID,omitempty"`
}

func (x *ForCounterProductDataYM) Reset() {
	*x = ForCounterProductDataYM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForCounterProductDataYM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForCounterProductDataYM) ProtoMessage() {}

func (x *ForCounterProductDataYM) ProtoReflect() protoreflect.Message {
	mi := &file_ym_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForCounterProductDataYM.ProtoReflect.Descriptor instead.
func (*ForCounterProductDataYM) Descriptor() ([]byte, []int) {
	return file_ym_proto_rawDescGZIP(), []int{4}
}

func (x *ForCounterProductDataYM) GetShopProductID() int64 {
	if x != nil {
		return x.ShopProductID
	}
	return 0
}

func (x *ForCounterProductDataYM) GetOfferID() string {
	if x != nil {
		return x.OfferID
	}
	return ""
}

type ForCounterReplyYM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopYMAuth *ShopYMAuth                `protobuf:"bytes,1,opt,name=ShopYMAuth,proto3" json:"ShopYMAuth,omitempty"`
	Data       []*ForCounterProductDataYM `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Campaigns  []*YMCampaign              `protobuf:"bytes,3,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
}

func (x *ForCounterReplyYM) Reset() {
	*x = ForCounterReplyYM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ym_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForCounterReplyYM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForCounterReplyYM) ProtoMessage() {}

func (x *ForCounterReplyYM) ProtoReflect() protoreflect.Message {
	mi := &file_ym_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForCounterReplyYM.ProtoReflect.Descriptor instead.
func (*ForCounterReplyYM) Descriptor() ([]byte, []int) {
	return file_ym_proto_rawDescGZIP(), []int{5}
}

func (x *ForCounterReplyYM) GetShopYMAuth() *ShopYMAuth {
	if x != nil {
		return x.ShopYMAuth
	}
	return nil
}

func (x *ForCounterReplyYM) GetData() []*ForCounterProductDataYM {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ForCounterReplyYM) GetCampaigns() []*YMCampaign {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

var File_ym_proto protoreflect.FileDescriptor

var file_ym_proto_rawDesc = []byte{
	0x0a, 0x08, 0x79, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x59, 0x4d, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x59, 0x4d, 0x41, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x08,
	0x73, 0x68, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x59, 0x4d, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x75, 0x74, 0x68, 0x22, 0xb0, 0x01,
	0x0a, 0x0a, 0x59, 0x4d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x59, 0x0a, 0x13, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x59, 0x4d, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x17, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x59, 0x4d, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73,
	0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x49, 0x44, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x46, 0x6f, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x59, 0x4d, 0x12, 0x33, 0x0a, 0x0a,
	0x53, 0x68, 0x6f, 0x70, 0x59, 0x4d, 0x41, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x59,
	0x4d, 0x41, 0x75, 0x74, 0x68, 0x52, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x59, 0x4d, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x59,
	0x4d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x59, 0x4d, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52,
	0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x32, 0xb8, 0x0c, 0x0a, 0x02, 0x59,
	0x4d, 0x12, 0x33, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x59, 0x4d, 0x41, 0x75, 0x74, 0x68,
	0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x59,
	0x4d, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0d,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x18, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53,
	0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x79, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x1a, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x44, 0x61, 0x79, 0x73, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x12, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x12, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4d,
	0x61, 0x69, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x44, 0x6f, 0x6e, 0x75, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x12,
	0x0d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x15,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x44, 0x6f, 0x6e, 0x75, 0x74, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x69, 0x63, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x10, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x59, 0x4d, 0x12, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x59, 0x4d, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x59, 0x4d,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x55, 0x72,
	0x6c, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x46,
	0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x46, 0x6f, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ym_proto_rawDescOnce sync.Once
	file_ym_proto_rawDescData = file_ym_proto_rawDesc
)

func file_ym_proto_rawDescGZIP() []byte {
	file_ym_proto_rawDescOnce.Do(func() {
		file_ym_proto_rawDescData = protoimpl.X.CompressGZIP(file_ym_proto_rawDescData)
	})
	return file_ym_proto_rawDescData
}

var file_ym_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ym_proto_goTypes = []interface{}{
	(*ShopYMAuth)(nil),               // 0: cerasus.ShopYMAuth
	(*SetYMAuth)(nil),                // 1: cerasus.SetYMAuth
	(*YMCampaign)(nil),               // 2: cerasus.YMCampaign
	(*ForCounterRequestYM)(nil),      // 3: cerasus.ForCounterRequestYM
	(*ForCounterProductDataYM)(nil),  // 4: cerasus.ForCounterProductDataYM
	(*ForCounterReplyYM)(nil),        // 5: cerasus.ForCounterReplyYM
	(*Auth)(nil),                     // 6: cerasus.Auth
	(*PingRequest)(nil),              // 7: cerasus.PingRequest
	(*ShopProductListRequest)(nil),   // 8: cerasus.ShopProductListRequest
	(*ShopProductRequest)(nil),       // 9: cerasus.ShopProductRequest
	(*ShopProductUpdateRequest)(nil), // 10: cerasus.ShopProductUpdateRequest
	(*SalesRequest)(nil),             // 11: cerasus.SalesRequest
	(*ShopServiceRequest)(nil),       // 12: cerasus.ShopServiceRequest
	(*SaleDetailsRequest)(nil),       // 13: cerasus.SaleDetailsRequest
	(*ProductSalesRequest)(nil),      // 14: cerasus.ProductSalesRequest
	(*MainGraphicRequest)(nil),       // 15: cerasus.MainGraphicRequest
	(*ImageRequest)(nil),             // 16: cerasus.ImageRequest
	(*IDRequest)(nil),                // 17: cerasus.IDRequest
	(*ForBrandSalesRequest)(nil),     // 18: cerasus.ForBrandSalesRequest
	(*ForBrandOrdersRequest)(nil),    // 19: cerasus.ForBrandOrdersRequest
	(*BoolReply)(nil),                // 20: cerasus.BoolReply
	(*PingReply)(nil),                // 21: cerasus.PingReply
	(*CompanyShopData)(nil),          // 22: cerasus.CompanyShopData
	(*CountReply)(nil),               // 23: cerasus.CountReply
	(*ShopProductListReply)(nil),     // 24: cerasus.ShopProductListReply
	(*ShopProduct)(nil),              // 25: cerasus.ShopProduct
	(*DaysSalesReply)(nil),           // 26: cerasus.DaysSalesReply
	(*SalesReply)(nil),               // 27: cerasus.SalesReply
	(*ShopServiceReply)(nil),         // 28: cerasus.ShopServiceReply
	(*SaleDetailsReply)(nil),         // 29: cerasus.SaleDetailsReply
	(*MainGraphicReply)(nil),         // 30: cerasus.MainGraphicReply
	(*ImageReply)(nil),               // 31: cerasus.ImageReply
	(*DonutGraphic)(nil),             // 32: cerasus.DonutGraphic
	(*WeekGraphics)(nil),             // 33: cerasus.WeekGraphics
	(*ProductShopUrls)(nil),          // 34: cerasus.ProductShopUrls
	(*ForBrandSalesReply)(nil),       // 35: cerasus.ForBrandSalesReply
	(*ForBrandOrdersReply)(nil),      // 36: cerasus.ForBrandOrdersReply
}
var file_ym_proto_depIdxs = []int32{
	6,  // 0: cerasus.SetYMAuth.auth:type_name -> cerasus.Auth
	0,  // 1: cerasus.SetYMAuth.shopAuth:type_name -> cerasus.ShopYMAuth
	0,  // 2: cerasus.ForCounterReplyYM.ShopYMAuth:type_name -> cerasus.ShopYMAuth
	4,  // 3: cerasus.ForCounterReplyYM.data:type_name -> cerasus.ForCounterProductDataYM
	2,  // 4: cerasus.ForCounterReplyYM.campaigns:type_name -> cerasus.YMCampaign
	1,  // 5: cerasus.YM.SetAuth:input_type -> cerasus.SetYMAuth
	6,  // 6: cerasus.YM.GetAuth:input_type -> cerasus.Auth
	6,  // 7: cerasus.YM.ErrorAuth:input_type -> cerasus.Auth
	7,  // 8: cerasus.YM.Ping:input_type -> cerasus.PingRequest
	6,  // 9: cerasus.YM.CheckShopData:input_type -> cerasus.Auth
	6,  // 10: cerasus.YM.GetUnsortedCount:input_type -> cerasus.Auth
	8,  // 11: cerasus.YM.GetUnsortedList:input_type -> cerasus.ShopProductListRequest
	6,  // 12: cerasus.YM.GetProductCount:input_type -> cerasus.Auth
	8,  // 13: cerasus.YM.GetProductList:input_type -> cerasus.ShopProductListRequest
	9,  // 14: cerasus.YM.GetProduct:input_type -> cerasus.ShopProductRequest
	10, // 15: cerasus.YM.UpdateProduct:input_type -> cerasus.ShopProductUpdateRequest
	6,  // 16: cerasus.YM.GetDaySales:input_type -> cerasus.Auth
	11, // 17: cerasus.YM.GetSales:input_type -> cerasus.SalesRequest
	12, // 18: cerasus.YM.GetShopServices:input_type -> cerasus.ShopServiceRequest
	13, // 19: cerasus.YM.GetSaleDetail:input_type -> cerasus.SaleDetailsRequest
	14, // 20: cerasus.YM.GetProductSales:input_type -> cerasus.ProductSalesRequest
	15, // 21: cerasus.YM.GetMainGraphic:input_type -> cerasus.MainGraphicRequest
	16, // 22: cerasus.YM.GetImage:input_type -> cerasus.ImageRequest
	6,  // 23: cerasus.YM.GetDonutGraphics:input_type -> cerasus.Auth
	6,  // 24: cerasus.YM.GetWeekGraphics:input_type -> cerasus.Auth
	3,  // 25: cerasus.YM.ForCounterDataYM:input_type -> cerasus.ForCounterRequestYM
	17, // 26: cerasus.YM.GetProductUrls:input_type -> cerasus.IDRequest
	18, // 27: cerasus.YM.ForBrandSales:input_type -> cerasus.ForBrandSalesRequest
	19, // 28: cerasus.YM.ForBrandOrders:input_type -> cerasus.ForBrandOrdersRequest
	20, // 29: cerasus.YM.SetAuth:output_type -> cerasus.BoolReply
	0,  // 30: cerasus.YM.GetAuth:output_type -> cerasus.ShopYMAuth
	20, // 31: cerasus.YM.ErrorAuth:output_type -> cerasus.BoolReply
	21, // 32: cerasus.YM.Ping:output_type -> cerasus.PingReply
	22, // 33: cerasus.YM.CheckShopData:output_type -> cerasus.CompanyShopData
	23, // 34: cerasus.YM.GetUnsortedCount:output_type -> cerasus.CountReply
	24, // 35: cerasus.YM.GetUnsortedList:output_type -> cerasus.ShopProductListReply
	23, // 36: cerasus.YM.GetProductCount:output_type -> cerasus.CountReply
	24, // 37: cerasus.YM.GetProductList:output_type -> cerasus.ShopProductListReply
	25, // 38: cerasus.YM.GetProduct:output_type -> cerasus.ShopProduct
	20, // 39: cerasus.YM.UpdateProduct:output_type -> cerasus.BoolReply
	26, // 40: cerasus.YM.GetDaySales:output_type -> cerasus.DaysSalesReply
	27, // 41: cerasus.YM.GetSales:output_type -> cerasus.SalesReply
	28, // 42: cerasus.YM.GetShopServices:output_type -> cerasus.ShopServiceReply
	29, // 43: cerasus.YM.GetSaleDetail:output_type -> cerasus.SaleDetailsReply
	27, // 44: cerasus.YM.GetProductSales:output_type -> cerasus.SalesReply
	30, // 45: cerasus.YM.GetMainGraphic:output_type -> cerasus.MainGraphicReply
	31, // 46: cerasus.YM.GetImage:output_type -> cerasus.ImageReply
	32, // 47: cerasus.YM.GetDonutGraphics:output_type -> cerasus.DonutGraphic
	33, // 48: cerasus.YM.GetWeekGraphics:output_type -> cerasus.WeekGraphics
	5,  // 49: cerasus.YM.ForCounterDataYM:output_type -> cerasus.ForCounterReplyYM
	34, // 50: cerasus.YM.GetProductUrls:output_type -> cerasus.ProductShopUrls
	35, // 51: cerasus.YM.ForBrandSales:output_type -> cerasus.ForBrandSalesReply
	36, // 52: cerasus.YM.ForBrandOrders:output_type -> cerasus.ForBrandOrdersReply
	29, // [29:53] is the sub-list for method output_type
	5,  // [5:29] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ym_proto_init() }
func file_ym_proto_init() {
	if File_ym_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ym_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopYMAuth); i {
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
		file_ym_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetYMAuth); i {
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
		file_ym_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YMCampaign); i {
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
		file_ym_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForCounterRequestYM); i {
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
		file_ym_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForCounterProductDataYM); i {
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
		file_ym_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForCounterReplyYM); i {
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
			RawDescriptor: file_ym_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ym_proto_goTypes,
		DependencyIndexes: file_ym_proto_depIdxs,
		MessageInfos:      file_ym_proto_msgTypes,
	}.Build()
	File_ym_proto = out.File
	file_ym_proto_rawDesc = nil
	file_ym_proto_goTypes = nil
	file_ym_proto_depIdxs = nil
}
