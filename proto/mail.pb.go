// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/mail.proto

package mail_proto

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

type EmailEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	ConfirmedAt int64  `protobuf:"varint,3,opt,name=confirmed_at,json=confirmedAt,proto3" json:"confirmed_at,omitempty"`
	OptOut      bool   `protobuf:"varint,4,opt,name=opt_out,json=optOut,proto3" json:"opt_out,omitempty"`
}

func (x *EmailEntry) Reset() {
	*x = EmailEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailEntry) ProtoMessage() {}

func (x *EmailEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailEntry.ProtoReflect.Descriptor instead.
func (*EmailEntry) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{0}
}

func (x *EmailEntry) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EmailEntry) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailEntry) GetConfirmedAt() int64 {
	if x != nil {
		return x.ConfirmedAt
	}
	return 0
}

func (x *EmailEntry) GetOptOut() bool {
	if x != nil {
		return x.OptOut
	}
	return false
}

type CreateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddr string `protobuf:"bytes,1,opt,name=email_addr,json=emailAddr,proto3" json:"email_addr,omitempty"`
}

func (x *CreateEmailRequest) Reset() {
	*x = CreateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailRequest) ProtoMessage() {}

func (x *CreateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailRequest.ProtoReflect.Descriptor instead.
func (*CreateEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmailRequest) GetEmailAddr() string {
	if x != nil {
		return x.EmailAddr
	}
	return ""
}

type GetEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddr string `protobuf:"bytes,1,opt,name=email_addr,json=emailAddr,proto3" json:"email_addr,omitempty"`
}

func (x *GetEmailRequest) Reset() {
	*x = GetEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailRequest) ProtoMessage() {}

func (x *GetEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailRequest.ProtoReflect.Descriptor instead.
func (*GetEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{2}
}

func (x *GetEmailRequest) GetEmailAddr() string {
	if x != nil {
		return x.EmailAddr
	}
	return ""
}

type UpdateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailEntry *EmailEntry `protobuf:"bytes,1,opt,name=email_entry,json=emailEntry,proto3" json:"email_entry,omitempty"`
}

func (x *UpdateEmailRequest) Reset() {
	*x = UpdateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequest) ProtoMessage() {}

func (x *UpdateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEmailRequest) GetEmailEntry() *EmailEntry {
	if x != nil {
		return x.EmailEntry
	}
	return nil
}

type DeleteEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddr string `protobuf:"bytes,1,opt,name=email_addr,json=emailAddr,proto3" json:"email_addr,omitempty"`
}

func (x *DeleteEmailRequest) Reset() {
	*x = DeleteEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmailRequest) ProtoMessage() {}

func (x *DeleteEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmailRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEmailRequest) GetEmailAddr() string {
	if x != nil {
		return x.EmailAddr
	}
	return ""
}

type GetEmailBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetEmailBatchRequest) Reset() {
	*x = GetEmailBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailBatchRequest) ProtoMessage() {}

func (x *GetEmailBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailBatchRequest.ProtoReflect.Descriptor instead.
func (*GetEmailBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{5}
}

func (x *GetEmailBatchRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetEmailBatchRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type EmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailEntry *EmailEntry `protobuf:"bytes,1,opt,name=email_entry,json=emailEntry,proto3,oneof" json:"email_entry,omitempty"`
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{6}
}

func (x *EmailResponse) GetEmailEntry() *EmailEntry {
	if x != nil {
		return x.EmailEntry
	}
	return nil
}

type GetEmailBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailEntries []*EmailEntry `protobuf:"bytes,1,rep,name=email_entries,json=emailEntries,proto3" json:"email_entries,omitempty"`
}

func (x *GetEmailBatchResponse) Reset() {
	*x = GetEmailBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailBatchResponse) ProtoMessage() {}

func (x *GetEmailBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailBatchResponse.ProtoReflect.Descriptor instead.
func (*GetEmailBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_mail_proto_rawDescGZIP(), []int{7}
}

func (x *GetEmailBatchResponse) GetEmailEntries() []*EmailEntry {
	if x != nil {
		return x.EmailEntries
	}
	return nil
}

var File_proto_mail_proto protoreflect.FileDescriptor

var file_proto_mail_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x0a, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x33, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x22, 0x30,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x22, 0x48, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x33, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x22,
	0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x58, 0x0a, 0x0d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x32, 0xe4, 0x02, 0x0a,
	0x12, 0x4d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x6c, 0x69,
	0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mail_proto_rawDescOnce sync.Once
	file_proto_mail_proto_rawDescData = file_proto_mail_proto_rawDesc
)

func file_proto_mail_proto_rawDescGZIP() []byte {
	file_proto_mail_proto_rawDescOnce.Do(func() {
		file_proto_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mail_proto_rawDescData)
	})
	return file_proto_mail_proto_rawDescData
}

var file_proto_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_mail_proto_goTypes = []interface{}{
	(*EmailEntry)(nil),            // 0: proto.EmailEntry
	(*CreateEmailRequest)(nil),    // 1: proto.CreateEmailRequest
	(*GetEmailRequest)(nil),       // 2: proto.GetEmailRequest
	(*UpdateEmailRequest)(nil),    // 3: proto.UpdateEmailRequest
	(*DeleteEmailRequest)(nil),    // 4: proto.DeleteEmailRequest
	(*GetEmailBatchRequest)(nil),  // 5: proto.GetEmailBatchRequest
	(*EmailResponse)(nil),         // 6: proto.EmailResponse
	(*GetEmailBatchResponse)(nil), // 7: proto.GetEmailBatchResponse
}
var file_proto_mail_proto_depIdxs = []int32{
	0, // 0: proto.UpdateEmailRequest.email_entry:type_name -> proto.EmailEntry
	0, // 1: proto.EmailResponse.email_entry:type_name -> proto.EmailEntry
	0, // 2: proto.GetEmailBatchResponse.email_entries:type_name -> proto.EmailEntry
	1, // 3: proto.MailingListService.CreateEmail:input_type -> proto.CreateEmailRequest
	2, // 4: proto.MailingListService.GetEmail:input_type -> proto.GetEmailRequest
	3, // 5: proto.MailingListService.UpdateEmail:input_type -> proto.UpdateEmailRequest
	4, // 6: proto.MailingListService.DeleteEmail:input_type -> proto.DeleteEmailRequest
	5, // 7: proto.MailingListService.GetEmailBatch:input_type -> proto.GetEmailBatchRequest
	6, // 8: proto.MailingListService.CreateEmail:output_type -> proto.EmailResponse
	6, // 9: proto.MailingListService.GetEmail:output_type -> proto.EmailResponse
	6, // 10: proto.MailingListService.UpdateEmail:output_type -> proto.EmailResponse
	6, // 11: proto.MailingListService.DeleteEmail:output_type -> proto.EmailResponse
	7, // 12: proto.MailingListService.GetEmailBatch:output_type -> proto.GetEmailBatchResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_mail_proto_init() }
func file_proto_mail_proto_init() {
	if File_proto_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailEntry); i {
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
		file_proto_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailRequest); i {
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
		file_proto_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailRequest); i {
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
		file_proto_mail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailRequest); i {
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
		file_proto_mail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmailRequest); i {
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
		file_proto_mail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailBatchRequest); i {
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
		file_proto_mail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailResponse); i {
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
		file_proto_mail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailBatchResponse); i {
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
	file_proto_mail_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mail_proto_goTypes,
		DependencyIndexes: file_proto_mail_proto_depIdxs,
		MessageInfos:      file_proto_mail_proto_msgTypes,
	}.Build()
	File_proto_mail_proto = out.File
	file_proto_mail_proto_rawDesc = nil
	file_proto_mail_proto_goTypes = nil
	file_proto_mail_proto_depIdxs = nil
}
