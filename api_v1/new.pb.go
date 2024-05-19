// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: new.proto

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

type AllNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Newsline bool  `protobuf:"varint,2,opt,name=newsline,proto3" json:"newsline,omitempty"`
}

func (x *AllNewsRequest) Reset() {
	*x = AllNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllNewsRequest) ProtoMessage() {}

func (x *AllNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllNewsRequest.ProtoReflect.Descriptor instead.
func (*AllNewsRequest) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{0}
}

func (x *AllNewsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AllNewsRequest) GetNewsline() bool {
	if x != nil {
		return x.Newsline
	}
	return false
}

type OneNewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *OneNewsRequest) Reset() {
	*x = OneNewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneNewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneNewsRequest) ProtoMessage() {}

func (x *OneNewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneNewsRequest.ProtoReflect.Descriptor instead.
func (*OneNewsRequest) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{1}
}

func (x *OneNewsRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type OneNews struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title          string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Preview        string `protobuf:"bytes,3,opt,name=preview,proto3" json:"preview,omitempty"`
	TextShort      string `protobuf:"bytes,4,opt,name=text_short,json=textShort,proto3" json:"text_short,omitempty"`
	TextFull       string `protobuf:"bytes,5,opt,name=text_full,json=textFull,proto3" json:"text_full,omitempty"`
	Active         int64  `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	Newsline       int64  `protobuf:"varint,7,opt,name=newsline,proto3" json:"newsline,omitempty"`
	SeoTitle       string `protobuf:"bytes,8,opt,name=seo_title,json=seoTitle,proto3" json:"seo_title,omitempty"`
	SeoDiscription string `protobuf:"bytes,9,opt,name=seo_discription,json=seoDiscription,proto3" json:"seo_discription,omitempty"`
	SeoKeywords    string `protobuf:"bytes,10,opt,name=seo_keywords,json=seoKeywords,proto3" json:"seo_keywords,omitempty"`
}

func (x *OneNews) Reset() {
	*x = OneNews{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneNews) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneNews) ProtoMessage() {}

func (x *OneNews) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneNews.ProtoReflect.Descriptor instead.
func (*OneNews) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{2}
}

func (x *OneNews) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *OneNews) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OneNews) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *OneNews) GetTextShort() string {
	if x != nil {
		return x.TextShort
	}
	return ""
}

func (x *OneNews) GetTextFull() string {
	if x != nil {
		return x.TextFull
	}
	return ""
}

func (x *OneNews) GetActive() int64 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *OneNews) GetNewsline() int64 {
	if x != nil {
		return x.Newsline
	}
	return 0
}

func (x *OneNews) GetSeoTitle() string {
	if x != nil {
		return x.SeoTitle
	}
	return ""
}

func (x *OneNews) GetSeoDiscription() string {
	if x != nil {
		return x.SeoDiscription
	}
	return ""
}

func (x *OneNews) GetSeoKeywords() string {
	if x != nil {
		return x.SeoKeywords
	}
	return ""
}

type OneNewsShort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Preview   string `protobuf:"bytes,3,opt,name=preview,proto3" json:"preview,omitempty"`
	TextShort string `protobuf:"bytes,4,opt,name=text_short,json=textShort,proto3" json:"text_short,omitempty"`
}

func (x *OneNewsShort) Reset() {
	*x = OneNewsShort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneNewsShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneNewsShort) ProtoMessage() {}

func (x *OneNewsShort) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneNewsShort.ProtoReflect.Descriptor instead.
func (*OneNewsShort) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{3}
}

func (x *OneNewsShort) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *OneNewsShort) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *OneNewsShort) GetPreview() string {
	if x != nil {
		return x.Preview
	}
	return ""
}

func (x *OneNewsShort) GetTextShort() string {
	if x != nil {
		return x.TextShort
	}
	return ""
}

type NewsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News       []*OneNewsShort `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty"`
	Pagination *Pagination     `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *NewsList) Reset() {
	*x = NewsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_new_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsList) ProtoMessage() {}

func (x *NewsList) ProtoReflect() protoreflect.Message {
	mi := &file_new_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsList.ProtoReflect.Descriptor instead.
func (*NewsList) Descriptor() ([]byte, []int) {
	return file_new_proto_rawDescGZIP(), []int{4}
}

func (x *NewsList) GetNews() []*OneNewsShort {
	if x != nil {
		return x.News
	}
	return nil
}

func (x *NewsList) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_new_proto protoreflect.FileDescriptor

var file_new_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x1a, 0x0d, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0xa2, 0x02, 0x0a, 0x07, 0x4f, 0x6e, 0x65, 0x4e,
	0x65, 0x77, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x73, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6f, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6f, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x6f, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6f, 0x44, 0x69,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x6f,
	0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x6f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x6d, 0x0a, 0x0c,
	0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x6a, 0x0a, 0x08, 0x4e,
	0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x6e, 0x65,
	0x77, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x93, 0x03, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73,
	0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x17,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4f, 0x6e,
	0x65, 0x4e, 0x65, 0x77, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e,
	0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x10, 0x2e,
	0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x4f, 0x6e, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x1a,
	0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75,
	0x73, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61,
	0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_new_proto_rawDescOnce sync.Once
	file_new_proto_rawDescData = file_new_proto_rawDesc
)

func file_new_proto_rawDescGZIP() []byte {
	file_new_proto_rawDescOnce.Do(func() {
		file_new_proto_rawDescData = protoimpl.X.CompressGZIP(file_new_proto_rawDescData)
	})
	return file_new_proto_rawDescData
}

var file_new_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_new_proto_goTypes = []interface{}{
	(*AllNewsRequest)(nil), // 0: cerasus.AllNewsRequest
	(*OneNewsRequest)(nil), // 1: cerasus.OneNewsRequest
	(*OneNews)(nil),        // 2: cerasus.OneNews
	(*OneNewsShort)(nil),   // 3: cerasus.OneNewsShort
	(*NewsList)(nil),       // 4: cerasus.NewsList
	(*Pagination)(nil),     // 5: cerasus.Pagination
	(*InsertReply)(nil),    // 6: cerasus.InsertReply
	(*ImageRequest)(nil),   // 7: cerasus.ImageRequest
	(*PingRequest)(nil),    // 8: cerasus.PingRequest
	(*BoolReply)(nil),      // 9: cerasus.BoolReply
	(*ImageReply)(nil),     // 10: cerasus.ImageReply
	(*PingReply)(nil),      // 11: cerasus.PingReply
}
var file_new_proto_depIdxs = []int32{
	3,  // 0: cerasus.NewsList.news:type_name -> cerasus.OneNewsShort
	5,  // 1: cerasus.NewsList.pagination:type_name -> cerasus.Pagination
	0,  // 2: cerasus.News.GetAllNews:input_type -> cerasus.AllNewsRequest
	1,  // 3: cerasus.News.GetOneNews:input_type -> cerasus.OneNewsRequest
	2,  // 4: cerasus.News.CreateNews:input_type -> cerasus.OneNews
	2,  // 5: cerasus.News.UpdateNews:input_type -> cerasus.OneNews
	6,  // 6: cerasus.News.DeleteNews:input_type -> cerasus.InsertReply
	7,  // 7: cerasus.News.GetImage:input_type -> cerasus.ImageRequest
	8,  // 8: cerasus.News.Ping:input_type -> cerasus.PingRequest
	4,  // 9: cerasus.News.GetAllNews:output_type -> cerasus.NewsList
	2,  // 10: cerasus.News.GetOneNews:output_type -> cerasus.OneNews
	6,  // 11: cerasus.News.CreateNews:output_type -> cerasus.InsertReply
	9,  // 12: cerasus.News.UpdateNews:output_type -> cerasus.BoolReply
	9,  // 13: cerasus.News.DeleteNews:output_type -> cerasus.BoolReply
	10, // 14: cerasus.News.GetImage:output_type -> cerasus.ImageReply
	11, // 15: cerasus.News.Ping:output_type -> cerasus.PingReply
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_new_proto_init() }
func file_new_proto_init() {
	if File_new_proto != nil {
		return
	}
	file_cerasus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_new_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllNewsRequest); i {
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
		file_new_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneNewsRequest); i {
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
		file_new_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneNews); i {
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
		file_new_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneNewsShort); i {
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
		file_new_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsList); i {
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
			RawDescriptor: file_new_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_new_proto_goTypes,
		DependencyIndexes: file_new_proto_depIdxs,
		MessageInfos:      file_new_proto_msgTypes,
	}.Build()
	File_new_proto = out.File
	file_new_proto_rawDesc = nil
	file_new_proto_goTypes = nil
	file_new_proto_depIdxs = nil
}
