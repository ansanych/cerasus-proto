// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: cerasus_v2.proto

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

// Status
type Status int32

const (
	Status_UNKNOWN Status = 0
	Status_OK      Status = 1
	Status_ERROR   Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
		2: "ERROR",
	}
	Status_value = map[string]int32{
		"UNKNOWN": 0,
		"OK":      1,
		"ERROR":   2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_cerasus_v2_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_cerasus_v2_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{0}
}

// Service status
type ServiceStatus int32

const (
	ServiceStatus_DISABLED ServiceStatus = 0
	ServiceStatus_ENABLED  ServiceStatus = 1
	ServiceStatus_BID      ServiceStatus = 2
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "DISABLED",
		1: "ENABLED",
		2: "BID",
	}
	ServiceStatus_value = map[string]int32{
		"DISABLED": 0,
		"ENABLED":  1,
		"BID":      2,
	}
)

func (x ServiceStatus) Enum() *ServiceStatus {
	p := new(ServiceStatus)
	*p = x
	return p
}

func (x ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cerasus_v2_proto_enumTypes[1].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_cerasus_v2_proto_enumTypes[1]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{1}
}

// Ping
type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=cerasusV2.Status" json:"status,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{2}
}

func (x *StatusReply) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

// Auth
type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenValidation bool    `protobuf:"varint,1,opt,name=tokenValidation,proto3" json:"tokenValidation,omitempty"`
	UserID          int64   `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	CompanyID       int64   `protobuf:"varint,3,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Roles           []*Role `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{3}
}

func (x *Auth) GetTokenValidation() bool {
	if x != nil {
		return x.TokenValidation
	}
	return false
}

func (x *Auth) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Auth) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *Auth) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Brand bool   `protobuf:"varint,3,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{4}
}

func (x *Company) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetBrand() bool {
	if x != nil {
		return x.Brand
	}
	return false
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code  string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{5}
}

func (x *Role) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Role) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	VerifyEmail bool     `protobuf:"varint,4,opt,name=verifyEmail,proto3" json:"verifyEmail,omitempty"`
	Phone       string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	VerifyPhone bool     `protobuf:"varint,6,opt,name=verifyPhone,proto3" json:"verifyPhone,omitempty"`
	Company     *Company `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
	Roles       []*Role  `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
	Active      bool     `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetVerifyEmail() bool {
	if x != nil {
		return x.VerifyEmail
	}
	return false
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetVerifyPhone() bool {
	if x != nil {
		return x.VerifyPhone
	}
	return false
}

func (x *User) GetCompany() *Company {
	if x != nil {
		return x.Company
	}
	return nil
}

func (x *User) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *User) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

// Shops
type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Logo   string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	Url    string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Active bool   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{7}
}

func (x *Shop) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Shop) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Shop) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Shop) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Shop) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Shop) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type AppShopData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shop     *Shop `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	Auth     bool  `protobuf:"varint,2,opt,name=auth,proto3" json:"auth,omitempty"`
	Products bool  `protobuf:"varint,3,opt,name=products,proto3" json:"products,omitempty"`
	Sales    bool  `protobuf:"varint,4,opt,name=sales,proto3" json:"sales,omitempty"`
	Orders   bool  `protobuf:"varint,5,opt,name=orders,proto3" json:"orders,omitempty"`
	Days     bool  `protobuf:"varint,6,opt,name=days,proto3" json:"days,omitempty"`
	Active   bool  `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	Access   bool  `protobuf:"varint,8,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *AppShopData) Reset() {
	*x = AppShopData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppShopData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppShopData) ProtoMessage() {}

func (x *AppShopData) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppShopData.ProtoReflect.Descriptor instead.
func (*AppShopData) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{8}
}

func (x *AppShopData) GetShop() *Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

func (x *AppShopData) GetAuth() bool {
	if x != nil {
		return x.Auth
	}
	return false
}

func (x *AppShopData) GetProducts() bool {
	if x != nil {
		return x.Products
	}
	return false
}

func (x *AppShopData) GetSales() bool {
	if x != nil {
		return x.Sales
	}
	return false
}

func (x *AppShopData) GetOrders() bool {
	if x != nil {
		return x.Orders
	}
	return false
}

func (x *AppShopData) GetDays() bool {
	if x != nil {
		return x.Days
	}
	return false
}

func (x *AppShopData) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *AppShopData) GetAccess() bool {
	if x != nil {
		return x.Access
	}
	return false
}

type ShopWidgetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	SalesCount  int32  `protobuf:"varint,2,opt,name=salesCount,proto3" json:"salesCount,omitempty"`
	SalesSum    int32  `protobuf:"varint,3,opt,name=salesSum,proto3" json:"salesSum,omitempty"`
	OrdersCount int32  `protobuf:"varint,4,opt,name=ordersCount,proto3" json:"ordersCount,omitempty"`
	OrdersSum   int32  `protobuf:"varint,5,opt,name=ordersSum,proto3" json:"ordersSum,omitempty"`
}

func (x *ShopWidgetData) Reset() {
	*x = ShopWidgetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopWidgetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopWidgetData) ProtoMessage() {}

func (x *ShopWidgetData) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopWidgetData.ProtoReflect.Descriptor instead.
func (*ShopWidgetData) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{9}
}

func (x *ShopWidgetData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ShopWidgetData) GetSalesCount() int32 {
	if x != nil {
		return x.SalesCount
	}
	return 0
}

func (x *ShopWidgetData) GetSalesSum() int32 {
	if x != nil {
		return x.SalesSum
	}
	return 0
}

func (x *ShopWidgetData) GetOrdersCount() int32 {
	if x != nil {
		return x.OrdersCount
	}
	return 0
}

func (x *ShopWidgetData) GetOrdersSum() int32 {
	if x != nil {
		return x.OrdersSum
	}
	return 0
}

type ShopWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ShopWidgetData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ShopWidget) Reset() {
	*x = ShopWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cerasus_v2_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopWidget) ProtoMessage() {}

func (x *ShopWidget) ProtoReflect() protoreflect.Message {
	mi := &file_cerasus_v2_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopWidget.ProtoReflect.Descriptor instead.
func (*ShopWidget) Descriptor() ([]byte, []int) {
	return file_cerasus_v2_proto_rawDescGZIP(), []int{10}
}

func (x *ShopWidget) GetData() []*ShopWidgetData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_cerasus_v2_proto protoreflect.FileDescriptor

var file_cerasus_v2_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x22, 0x27, 0x0a,
	0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x38, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x28, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44,
	0x12, 0x25, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x87,
	0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x7e, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f,
	0x67, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70,
	0x53, 0x68, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x61,
	0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0xa2, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x61,
	0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x53, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x53, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x53, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x53, 0x75, 0x6d, 0x22, 0x3b, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x68, 0x6f,
	0x70, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x2a, 0x33, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x44, 0x10, 0x02,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68, 0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cerasus_v2_proto_rawDescOnce sync.Once
	file_cerasus_v2_proto_rawDescData = file_cerasus_v2_proto_rawDesc
)

func file_cerasus_v2_proto_rawDescGZIP() []byte {
	file_cerasus_v2_proto_rawDescOnce.Do(func() {
		file_cerasus_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_cerasus_v2_proto_rawDescData)
	})
	return file_cerasus_v2_proto_rawDescData
}

var file_cerasus_v2_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cerasus_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cerasus_v2_proto_goTypes = []interface{}{
	(Status)(0),            // 0: cerasusV2.Status
	(ServiceStatus)(0),     // 1: cerasusV2.ServiceStatus
	(*PingRequest)(nil),    // 2: cerasusV2.PingRequest
	(*PingReply)(nil),      // 3: cerasusV2.PingReply
	(*StatusReply)(nil),    // 4: cerasusV2.StatusReply
	(*Auth)(nil),           // 5: cerasusV2.Auth
	(*Company)(nil),        // 6: cerasusV2.Company
	(*Role)(nil),           // 7: cerasusV2.Role
	(*User)(nil),           // 8: cerasusV2.User
	(*Shop)(nil),           // 9: cerasusV2.Shop
	(*AppShopData)(nil),    // 10: cerasusV2.AppShopData
	(*ShopWidgetData)(nil), // 11: cerasusV2.ShopWidgetData
	(*ShopWidget)(nil),     // 12: cerasusV2.ShopWidget
}
var file_cerasus_v2_proto_depIdxs = []int32{
	0,  // 0: cerasusV2.StatusReply.status:type_name -> cerasusV2.Status
	7,  // 1: cerasusV2.Auth.roles:type_name -> cerasusV2.Role
	6,  // 2: cerasusV2.User.company:type_name -> cerasusV2.Company
	7,  // 3: cerasusV2.User.roles:type_name -> cerasusV2.Role
	9,  // 4: cerasusV2.AppShopData.shop:type_name -> cerasusV2.Shop
	11, // 5: cerasusV2.ShopWidget.data:type_name -> cerasusV2.ShopWidgetData
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cerasus_v2_proto_init() }
func file_cerasus_v2_proto_init() {
	if File_cerasus_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cerasus_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_cerasus_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReply); i {
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
		file_cerasus_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_cerasus_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_cerasus_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_cerasus_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_cerasus_v2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_cerasus_v2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
		file_cerasus_v2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppShopData); i {
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
		file_cerasus_v2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopWidgetData); i {
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
		file_cerasus_v2_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopWidget); i {
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
			RawDescriptor: file_cerasus_v2_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cerasus_v2_proto_goTypes,
		DependencyIndexes: file_cerasus_v2_proto_depIdxs,
		EnumInfos:         file_cerasus_v2_proto_enumTypes,
		MessageInfos:      file_cerasus_v2_proto_msgTypes,
	}.Build()
	File_cerasus_v2_proto = out.File
	file_cerasus_v2_proto_rawDesc = nil
	file_cerasus_v2_proto_goTypes = nil
	file_cerasus_v2_proto_depIdxs = nil
}
