// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: all.proto

package __

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

type Pb_CMD int32

const (
	Pb_CMD_USER   Pb_CMD = 0
	Pb_CMD_FRIEND Pb_CMD = 1
	Pb_CMD_CARD   Pb_CMD = 3
)

// Enum value maps for Pb_CMD.
var (
	Pb_CMD_name = map[int32]string{
		0: "CMD_USER",
		1: "CMD_FRIEND",
		3: "CMD_CARD",
	}
	Pb_CMD_value = map[string]int32{
		"CMD_USER":   0,
		"CMD_FRIEND": 1,
		"CMD_CARD":   3,
	}
)

func (x Pb_CMD) Enum() *Pb_CMD {
	p := new(Pb_CMD)
	*p = x
	return p
}

func (x Pb_CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pb_CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_all_proto_enumTypes[0].Descriptor()
}

func (Pb_CMD) Type() protoreflect.EnumType {
	return &file_all_proto_enumTypes[0]
}

func (x Pb_CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pb_CMD.Descriptor instead.
func (Pb_CMD) EnumDescriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 0}
}

type Pb_ACT int32

const (
	Pb_ACT_USER_INFO    Pb_ACT = 0
	Pb_ACT_USER_CREATE  Pb_ACT = 1
	Pb_ACT_USER_ADD_EXP Pb_ACT = 2
	Pb_ACT_ACCOUNT_INFO Pb_ACT = 4
)

// Enum value maps for Pb_ACT.
var (
	Pb_ACT_name = map[int32]string{
		0: "ACT_USER_INFO",
		1: "ACT_USER_CREATE",
		2: "ACT_USER_ADD_EXP",
		4: "ACT_ACCOUNT_INFO",
	}
	Pb_ACT_value = map[string]int32{
		"ACT_USER_INFO":    0,
		"ACT_USER_CREATE":  1,
		"ACT_USER_ADD_EXP": 2,
		"ACT_ACCOUNT_INFO": 4,
	}
)

func (x Pb_ACT) Enum() *Pb_ACT {
	p := new(Pb_ACT)
	*p = x
	return p
}

func (x Pb_ACT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pb_ACT) Descriptor() protoreflect.EnumDescriptor {
	return file_all_proto_enumTypes[1].Descriptor()
}

func (Pb_ACT) Type() protoreflect.EnumType {
	return &file_all_proto_enumTypes[1]
}

func (x Pb_ACT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pb_ACT.Descriptor instead.
func (Pb_ACT) EnumDescriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 1}
}

type Pb struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb) Reset() {
	*x = Pb{}
	mi := &file_all_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb) ProtoMessage() {}

func (x *Pb) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb.ProtoReflect.Descriptor instead.
func (*Pb) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0}
}

type Pb_RqstAccountInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *Pb_AccountInfo        `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
	Index         *int32                 `protobuf:"varint,2,opt,name=index,proto3,oneof" json:"index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_RqstAccountInfo) Reset() {
	*x = Pb_RqstAccountInfo{}
	mi := &file_all_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_RqstAccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RqstAccountInfo) ProtoMessage() {}

func (x *Pb_RqstAccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RqstAccountInfo.ProtoReflect.Descriptor instead.
func (*Pb_RqstAccountInfo) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Pb_RqstAccountInfo) GetInfo() *Pb_AccountInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Pb_RqstAccountInfo) GetIndex() int32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type Pb_RespAccountInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id            int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_RespAccountInfo) Reset() {
	*x = Pb_RespAccountInfo{}
	mi := &file_all_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_RespAccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RespAccountInfo) ProtoMessage() {}

func (x *Pb_RespAccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RespAccountInfo.ProtoReflect.Descriptor instead.
func (*Pb_RespAccountInfo) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Pb_RespAccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pb_RespAccountInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Pb_AccountInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Head          string                 `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_AccountInfo) Reset() {
	*x = Pb_AccountInfo{}
	mi := &file_all_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_AccountInfo) ProtoMessage() {}

func (x *Pb_AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_AccountInfo.ProtoReflect.Descriptor instead.
func (*Pb_AccountInfo) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Pb_AccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pb_AccountInfo) GetHead() string {
	if x != nil {
		return x.Head
	}
	return ""
}

type Pb_RqstUserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_RqstUserInfo) Reset() {
	*x = Pb_RqstUserInfo{}
	mi := &file_all_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_RqstUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RqstUserInfo) ProtoMessage() {}

func (x *Pb_RqstUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RqstUserInfo.ProtoReflect.Descriptor instead.
func (*Pb_RqstUserInfo) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 3}
}

type Pb_RespUserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id            int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_RespUserInfo) Reset() {
	*x = Pb_RespUserInfo{}
	mi := &file_all_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_RespUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RespUserInfo) ProtoMessage() {}

func (x *Pb_RespUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RespUserInfo.ProtoReflect.Descriptor instead.
func (*Pb_RespUserInfo) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Pb_RespUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pb_RespUserInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Pb_RqstCreateUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id            int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_RqstCreateUser) Reset() {
	*x = Pb_RqstCreateUser{}
	mi := &file_all_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_RqstCreateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RqstCreateUser) ProtoMessage() {}

func (x *Pb_RqstCreateUser) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RqstCreateUser.ProtoReflect.Descriptor instead.
func (*Pb_RqstCreateUser) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Pb_RqstCreateUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pb_RqstCreateUser) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Pb_RespCreateUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id            int32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pb_RespCreateUser) Reset() {
	*x = Pb_RespCreateUser{}
	mi := &file_all_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pb_RespCreateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RespCreateUser) ProtoMessage() {}

func (x *Pb_RespCreateUser) ProtoReflect() protoreflect.Message {
	mi := &file_all_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RespCreateUser.ProtoReflect.Descriptor instead.
func (*Pb_RespCreateUser) Descriptor() ([]byte, []int) {
	return file_all_proto_rawDescGZIP(), []int{0, 6}
}

func (x *Pb_RespCreateUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pb_RespCreateUser) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_all_proto protoreflect.FileDescriptor

var file_all_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x04, 0x0a, 0x02,
	0x70, 0x62, 0x1a, 0x69, 0x0a, 0x0f, 0x52, 0x71, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x35, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x1a, 0x35, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x1a, 0x0e, 0x0a, 0x0c, 0x52,
	0x71, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x32, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x1a,
	0x34, 0x0a, 0x0e, 0x52, 0x71, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x34, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x03, 0x43,
	0x4d, 0x44, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4d, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4d, 0x44, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4d, 0x44, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x03, 0x22, 0x59,
	0x0a, 0x03, 0x41, 0x43, 0x54, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x43, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x58,
	0x50, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x43, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x04, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_all_proto_rawDescOnce sync.Once
	file_all_proto_rawDescData = file_all_proto_rawDesc
)

func file_all_proto_rawDescGZIP() []byte {
	file_all_proto_rawDescOnce.Do(func() {
		file_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_all_proto_rawDescData)
	})
	return file_all_proto_rawDescData
}

var file_all_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_all_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_all_proto_goTypes = []any{
	(Pb_CMD)(0),                // 0: pb.CMD
	(Pb_ACT)(0),                // 1: pb.ACT
	(*Pb)(nil),                 // 2: pb
	(*Pb_RqstAccountInfo)(nil), // 3: pb.RqstAccountInfo
	(*Pb_RespAccountInfo)(nil), // 4: pb.RespAccountInfo
	(*Pb_AccountInfo)(nil),     // 5: pb.AccountInfo
	(*Pb_RqstUserInfo)(nil),    // 6: pb.RqstUserInfo
	(*Pb_RespUserInfo)(nil),    // 7: pb.RespUserInfo
	(*Pb_RqstCreateUser)(nil),  // 8: pb.RqstCreateUser
	(*Pb_RespCreateUser)(nil),  // 9: pb.RespCreateUser
}
var file_all_proto_depIdxs = []int32{
	5, // 0: pb.RqstAccountInfo.info:type_name -> pb.AccountInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_all_proto_init() }
func file_all_proto_init() {
	if File_all_proto != nil {
		return
	}
	file_all_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_all_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_all_proto_goTypes,
		DependencyIndexes: file_all_proto_depIdxs,
		EnumInfos:         file_all_proto_enumTypes,
		MessageInfos:      file_all_proto_msgTypes,
	}.Build()
	File_all_proto = out.File
	file_all_proto_rawDesc = nil
	file_all_proto_goTypes = nil
	file_all_proto_depIdxs = nil
}
