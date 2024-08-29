// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: brand.proto

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

type BProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CompanyID int64          `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Title     string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Comment   string         `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	Article   string         `protobuf:"bytes,5,opt,name=article,proto3" json:"article,omitempty"`
	Active    bool           `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	Updated   string         `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	Preview   string         `protobuf:"bytes,8,opt,name=preview,proto3" json:"preview,omitempty"`
	Price     *BProductPrice `protobuf:"bytes,9,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *BProduct) Reset() {
	*x = BProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BProduct) ProtoMessage() {}

func (x *BProduct) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BProduct.ProtoReflect.Descriptor instead.
func (*BProduct) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{0}
}

func (x *BProduct) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BProduct) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *BProduct) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BProduct) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *BProduct) GetArticle() string {
	if x != nil {
		return x.Article
	}
	return ""
}

func (x *BProduct) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *BProduct) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *BProduct) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *BProduct) GetPrice() *BProductPrice {
	if x != nil {
		return x.Price
	}
	return nil
}

type BProductFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	File    string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	Type    string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Active  bool   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Updated string `protobuf:"bytes,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *BProductFile) Reset() {
	*x = BProductFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BProductFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BProductFile) ProtoMessage() {}

func (x *BProductFile) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BProductFile.ProtoReflect.Descriptor instead.
func (*BProductFile) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{1}
}

func (x *BProductFile) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BProductFile) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *BProductFile) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BProductFile) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *BProductFile) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

type BSeller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64      `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Comment       string     `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Person        string     `protobuf:"bytes,4,opt,name=person,proto3" json:"person,omitempty"`
	Phone         string     `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Email         string     `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	CompaniesLink []*Company `protobuf:"bytes,7,rep,name=companiesLink,proto3" json:"companiesLink,omitempty"`
}

func (x *BSeller) Reset() {
	*x = BSeller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BSeller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BSeller) ProtoMessage() {}

func (x *BSeller) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BSeller.ProtoReflect.Descriptor instead.
func (*BSeller) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{2}
}

func (x *BSeller) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BSeller) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BSeller) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *BSeller) GetPerson() string {
	if x != nil {
		return x.Person
	}
	return ""
}

func (x *BSeller) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BSeller) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BSeller) GetCompaniesLink() []*Company {
	if x != nil {
		return x.CompaniesLink
	}
	return nil
}

type BBrandLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Brand *Brand `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *BBrandLink) Reset() {
	*x = BBrandLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBrandLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBrandLink) ProtoMessage() {}

func (x *BBrandLink) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBrandLink.ProtoReflect.Descriptor instead.
func (*BBrandLink) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{3}
}

func (x *BBrandLink) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *BBrandLink) GetBrand() *Brand {
	if x != nil {
		return x.Brand
	}
	return nil
}

type BrandLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *BrandLinkRequest) Reset() {
	*x = BrandLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandLinkRequest) ProtoMessage() {}

func (x *BrandLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandLinkRequest.ProtoReflect.Descriptor instead.
func (*BrandLinkRequest) Descriptor() ([]byte, []int) {
	return file_brand_proto_rawDescGZIP(), []int{4}
}

func (x *BrandLinkRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_brand_proto protoreflect.FileDescriptor

var file_brand_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x08, 0x42, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x42, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x78, 0x0a, 0x0c, 0x42, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0xc3,
	0x01, 0x0a, 0x07, 0x42, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x36, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x42, 0x0a, 0x0a, 0x42, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x32, 0x7e, 0x0a, 0x06,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x42, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e,
	0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brand_proto_rawDescOnce sync.Once
	file_brand_proto_rawDescData = file_brand_proto_rawDesc
)

func file_brand_proto_rawDescGZIP() []byte {
	file_brand_proto_rawDescOnce.Do(func() {
		file_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_brand_proto_rawDescData)
	})
	return file_brand_proto_rawDescData
}

var file_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_brand_proto_goTypes = []interface{}{
	(*BProduct)(nil),         // 0: cerasus.BProduct
	(*BProductFile)(nil),     // 1: cerasus.BProductFile
	(*BSeller)(nil),          // 2: cerasus.BSeller
	(*BBrandLink)(nil),       // 3: cerasus.BBrandLink
	(*BrandLinkRequest)(nil), // 4: cerasus.BrandLinkRequest
	(*BProductPrice)(nil),    // 5: cerasus.BProductPrice
	(*Company)(nil),          // 6: cerasus.Company
	(*Brand)(nil),            // 7: cerasus.Brand
	(*PingRequest)(nil),      // 8: cerasus.PingRequest
	(*PingReply)(nil),        // 9: cerasus.PingReply
}
var file_brand_proto_depIdxs = []int32{
	5, // 0: cerasus.BProduct.price:type_name -> cerasus.BProductPrice
	6, // 1: cerasus.BSeller.companiesLink:type_name -> cerasus.Company
	7, // 2: cerasus.BBrandLink.brand:type_name -> cerasus.Brand
	8, // 3: cerasus.Brands.Ping:input_type -> cerasus.PingRequest
	4, // 4: cerasus.Brands.GetBrandLink:input_type -> cerasus.BrandLinkRequest
	9, // 5: cerasus.Brands.Ping:output_type -> cerasus.PingReply
	3, // 6: cerasus.Brands.GetBrandLink:output_type -> cerasus.BBrandLink
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_brand_proto_init() }
func file_brand_proto_init() {
	if File_brand_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BProduct); i {
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
		file_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BProductFile); i {
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
		file_brand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BSeller); i {
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
		file_brand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBrandLink); i {
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
		file_brand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandLinkRequest); i {
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
			RawDescriptor: file_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brand_proto_goTypes,
		DependencyIndexes: file_brand_proto_depIdxs,
		MessageInfos:      file_brand_proto_msgTypes,
	}.Build()
	File_brand_proto = out.File
	file_brand_proto_rawDesc = nil
	file_brand_proto_goTypes = nil
	file_brand_proto_depIdxs = nil
}
