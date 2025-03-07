// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: board_v2.proto

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

type ReQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth   *Auth  `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	ID     int64  `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Msv    string `protobuf:"bytes,4,opt,name=msv,proto3" json:"msv,omitempty"`
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ReQueueRequest) Reset() {
	*x = ReQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReQueueRequest) ProtoMessage() {}

func (x *ReQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReQueueRequest.ProtoReflect.Descriptor instead.
func (*ReQueueRequest) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{0}
}

func (x *ReQueueRequest) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *ReQueueRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReQueueRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ReQueueRequest) GetMsv() string {
	if x != nil {
		return x.Msv
	}
	return ""
}

func (x *ReQueueRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type QueuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth      *Auth `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	CompanyID int64 `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
}

func (x *QueuesRequest) Reset() {
	*x = QueuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueuesRequest) ProtoMessage() {}

func (x *QueuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueuesRequest.ProtoReflect.Descriptor instead.
func (*QueuesRequest) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{1}
}

func (x *QueuesRequest) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *QueuesRequest) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

type QueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth       *Auth  `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	CompanyID  int64  `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Msv        string `protobuf:"bytes,4,opt,name=msv,proto3" json:"msv,omitempty"`
	GetAllJobs bool   `protobuf:"varint,5,opt,name=getAllJobs,proto3" json:"getAllJobs,omitempty"`
}

func (x *QueueRequest) Reset() {
	*x = QueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueRequest) ProtoMessage() {}

func (x *QueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueRequest.ProtoReflect.Descriptor instead.
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{2}
}

func (x *QueueRequest) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *QueueRequest) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *QueueRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueueRequest) GetMsv() string {
	if x != nil {
		return x.Msv
	}
	return ""
}

func (x *QueueRequest) GetGetAllJobs() bool {
	if x != nil {
		return x.GetAllJobs
	}
	return false
}

type QueueParamsSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *QueueParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Code   string       `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Msv    string       `protobuf:"bytes,3,opt,name=msv,proto3" json:"msv,omitempty"`
}

func (x *QueueParamsSet) Reset() {
	*x = QueueParamsSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueParamsSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueParamsSet) ProtoMessage() {}

func (x *QueueParamsSet) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueParamsSet.ProtoReflect.Descriptor instead.
func (*QueueParamsSet) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{3}
}

func (x *QueueParamsSet) GetParams() *QueueParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *QueueParamsSet) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueueParamsSet) GetMsv() string {
	if x != nil {
		return x.Msv
	}
	return ""
}

type QueueParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period        int32 `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	AutoReset     bool  `protobuf:"varint,2,opt,name=autoReset,proto3" json:"autoReset,omitempty"`
	AutoResetTime int32 `protobuf:"varint,3,opt,name=autoResetTime,proto3" json:"autoResetTime,omitempty"`
	WarningTime   int32 `protobuf:"varint,4,opt,name=warningTime,proto3" json:"warningTime,omitempty"`
}

func (x *QueueParams) Reset() {
	*x = QueueParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueParams) ProtoMessage() {}

func (x *QueueParams) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueParams.ProtoReflect.Descriptor instead.
func (*QueueParams) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{4}
}

func (x *QueueParams) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *QueueParams) GetAutoReset() bool {
	if x != nil {
		return x.AutoReset
	}
	return false
}

func (x *QueueParams) GetAutoResetTime() int32 {
	if x != nil {
		return x.AutoResetTime
	}
	return 0
}

func (x *QueueParams) GetWarningTime() int32 {
	if x != nil {
		return x.WarningTime
	}
	return 0
}

type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Code        string       `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Msv         string       `protobuf:"bytes,3,opt,name=msv,proto3" json:"msv,omitempty"`
	Processing  int32        `protobuf:"varint,4,opt,name=processing,proto3" json:"processing,omitempty"`
	Errors      bool         `protobuf:"varint,5,opt,name=errors,proto3" json:"errors,omitempty"`
	Warning     bool         `protobuf:"varint,6,opt,name=warning,proto3" json:"warning,omitempty"`
	LastDT      string       `protobuf:"bytes,7,opt,name=lastDT,proto3" json:"lastDT,omitempty"`
	Jobs        []*QueueJob  `protobuf:"bytes,8,rep,name=jobs,proto3" json:"jobs,omitempty"`
	JobsError   []*QueueJob  `protobuf:"bytes,9,rep,name=jobsError,proto3" json:"jobsError,omitempty"`
	JobsWarning []*QueueJob  `protobuf:"bytes,10,rep,name=jobsWarning,proto3" json:"jobsWarning,omitempty"`
	Params      *QueueParams `protobuf:"bytes,11,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{5}
}

func (x *Queue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Queue) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Queue) GetMsv() string {
	if x != nil {
		return x.Msv
	}
	return ""
}

func (x *Queue) GetProcessing() int32 {
	if x != nil {
		return x.Processing
	}
	return 0
}

func (x *Queue) GetErrors() bool {
	if x != nil {
		return x.Errors
	}
	return false
}

func (x *Queue) GetWarning() bool {
	if x != nil {
		return x.Warning
	}
	return false
}

func (x *Queue) GetLastDT() string {
	if x != nil {
		return x.LastDT
	}
	return ""
}

func (x *Queue) GetJobs() []*QueueJob {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *Queue) GetJobsError() []*QueueJob {
	if x != nil {
		return x.JobsError
	}
	return nil
}

func (x *Queue) GetJobsWarning() []*QueueJob {
	if x != nil {
		return x.JobsWarning
	}
	return nil
}

func (x *Queue) GetParams() *QueueParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type Queues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Queue `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Queues) Reset() {
	*x = Queues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queues) ProtoMessage() {}

func (x *Queues) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queues.ProtoReflect.Descriptor instead.
func (*Queues) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{6}
}

func (x *Queues) GetData() []*Queue {
	if x != nil {
		return x.Data
	}
	return nil
}

type QueueJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyID int64          `protobuf:"varint,2,opt,name=companyID,proto3" json:"companyID,omitempty"`
	Shop      string         `protobuf:"bytes,3,opt,name=shop,proto3" json:"shop,omitempty"`
	Status    int32          `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Error     bool           `protobuf:"varint,5,opt,name=error,proto3" json:"error,omitempty"`
	Created   string         `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated   string         `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	Logs      []*QueueJobLog `protobuf:"bytes,8,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *QueueJob) Reset() {
	*x = QueueJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueJob) ProtoMessage() {}

func (x *QueueJob) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueJob.ProtoReflect.Descriptor instead.
func (*QueueJob) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{7}
}

func (x *QueueJob) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueueJob) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

func (x *QueueJob) GetShop() string {
	if x != nil {
		return x.Shop
	}
	return ""
}

func (x *QueueJob) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *QueueJob) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *QueueJob) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *QueueJob) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *QueueJob) GetLogs() []*QueueJobLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

type QueueJobLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime string `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *QueueJobLog) Reset() {
	*x = QueueJobLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_board_v2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueJobLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueJobLog) ProtoMessage() {}

func (x *QueueJobLog) ProtoReflect() protoreflect.Message {
	mi := &file_board_v2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueJobLog.ProtoReflect.Descriptor instead.
func (*QueueJobLog) Descriptor() ([]byte, []int) {
	return file_board_v2_proto_rawDescGZIP(), []int{8}
}

func (x *QueueJobLog) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *QueueJobLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_board_v2_proto protoreflect.FileDescriptor

var file_board_v2_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x1a, 0x10, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x76,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x97, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x76, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x62,
	0x73, 0x22, 0x66, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x53, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x76, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x76, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf0, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x76, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x54, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x54, 0x12, 0x27, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73,
	0x12, 0x31, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x73, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x73, 0x57, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73,
	0x75, 0x73, 0x56, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x0b, 0x6a,
	0x6f, 0x62, 0x73, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xda, 0x01, 0x0a, 0x08, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x4c, 0x6f,
	0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x43, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfa, 0x02, 0x0a,
	0x05, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x65,
	0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73,
	0x56, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x52, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x19,
	0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61,
	0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x65, 0x72,
	0x61, 0x73, 0x75, 0x73, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x56,
	0x32, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x73, 0x61, 0x6e, 0x79, 0x63, 0x68,
	0x2f, 0x63, 0x65, 0x72, 0x61, 0x73, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_board_v2_proto_rawDescOnce sync.Once
	file_board_v2_proto_rawDescData = file_board_v2_proto_rawDesc
)

func file_board_v2_proto_rawDescGZIP() []byte {
	file_board_v2_proto_rawDescOnce.Do(func() {
		file_board_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_board_v2_proto_rawDescData)
	})
	return file_board_v2_proto_rawDescData
}

var file_board_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_board_v2_proto_goTypes = []interface{}{
	(*ReQueueRequest)(nil), // 0: cerasusV2.ReQueueRequest
	(*QueuesRequest)(nil),  // 1: cerasusV2.QueuesRequest
	(*QueueRequest)(nil),   // 2: cerasusV2.QueueRequest
	(*QueueParamsSet)(nil), // 3: cerasusV2.QueueParamsSet
	(*QueueParams)(nil),    // 4: cerasusV2.QueueParams
	(*Queue)(nil),          // 5: cerasusV2.Queue
	(*Queues)(nil),         // 6: cerasusV2.Queues
	(*QueueJob)(nil),       // 7: cerasusV2.QueueJob
	(*QueueJobLog)(nil),    // 8: cerasusV2.QueueJobLog
	(*Auth)(nil),           // 9: cerasusV2.Auth
	(*PingRequest)(nil),    // 10: cerasusV2.PingRequest
	(*RequestByDates)(nil), // 11: cerasusV2.RequestByDates
	(*PingReply)(nil),      // 12: cerasusV2.PingReply
	(*StatusReply)(nil),    // 13: cerasusV2.StatusReply
	(*Count)(nil),          // 14: cerasusV2.Count
}
var file_board_v2_proto_depIdxs = []int32{
	9,  // 0: cerasusV2.ReQueueRequest.auth:type_name -> cerasusV2.Auth
	9,  // 1: cerasusV2.QueuesRequest.auth:type_name -> cerasusV2.Auth
	9,  // 2: cerasusV2.QueueRequest.auth:type_name -> cerasusV2.Auth
	4,  // 3: cerasusV2.QueueParamsSet.params:type_name -> cerasusV2.QueueParams
	7,  // 4: cerasusV2.Queue.jobs:type_name -> cerasusV2.QueueJob
	7,  // 5: cerasusV2.Queue.jobsError:type_name -> cerasusV2.QueueJob
	7,  // 6: cerasusV2.Queue.jobsWarning:type_name -> cerasusV2.QueueJob
	4,  // 7: cerasusV2.Queue.params:type_name -> cerasusV2.QueueParams
	5,  // 8: cerasusV2.Queues.data:type_name -> cerasusV2.Queue
	8,  // 9: cerasusV2.QueueJob.logs:type_name -> cerasusV2.QueueJobLog
	10, // 10: cerasusV2.Board.Ping:input_type -> cerasusV2.PingRequest
	1,  // 11: cerasusV2.Board.GetQueues:input_type -> cerasusV2.QueuesRequest
	2,  // 12: cerasusV2.Board.GetQueue:input_type -> cerasusV2.QueueRequest
	0,  // 13: cerasusV2.Board.ReQueue:input_type -> cerasusV2.ReQueueRequest
	3,  // 14: cerasusV2.Board.SetQueueParams:input_type -> cerasusV2.QueueParamsSet
	11, // 15: cerasusV2.Board.GetLogsCount:input_type -> cerasusV2.RequestByDates
	12, // 16: cerasusV2.Board.Ping:output_type -> cerasusV2.PingReply
	6,  // 17: cerasusV2.Board.GetQueues:output_type -> cerasusV2.Queues
	5,  // 18: cerasusV2.Board.GetQueue:output_type -> cerasusV2.Queue
	13, // 19: cerasusV2.Board.ReQueue:output_type -> cerasusV2.StatusReply
	13, // 20: cerasusV2.Board.SetQueueParams:output_type -> cerasusV2.StatusReply
	14, // 21: cerasusV2.Board.GetLogsCount:output_type -> cerasusV2.Count
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_board_v2_proto_init() }
func file_board_v2_proto_init() {
	if File_board_v2_proto != nil {
		return
	}
	file_cerasus_v2_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_board_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReQueueRequest); i {
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
		file_board_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueuesRequest); i {
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
		file_board_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueRequest); i {
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
		file_board_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueParamsSet); i {
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
		file_board_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueParams); i {
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
		file_board_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
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
		file_board_v2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queues); i {
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
		file_board_v2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueJob); i {
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
		file_board_v2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueJobLog); i {
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
			RawDescriptor: file_board_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_board_v2_proto_goTypes,
		DependencyIndexes: file_board_v2_proto_depIdxs,
		MessageInfos:      file_board_v2_proto_msgTypes,
	}.Build()
	File_board_v2_proto = out.File
	file_board_v2_proto_rawDesc = nil
	file_board_v2_proto_goTypes = nil
	file_board_v2_proto_depIdxs = nil
}
