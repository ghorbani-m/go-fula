// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: file.proto

package file

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

type ActionType int32

const (
	ActionType_READ   ActionType = 0
	ActionType_WRITE  ActionType = 1
	ActionType_MKDIR  ActionType = 2
	ActionType_LS     ActionType = 3
	ActionType_DELETE ActionType = 4
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "READ",
		1: "WRITE",
		2: "MKDIR",
		3: "LS",
		4: "DELETE",
	}
	ActionType_value = map[string]int32{
		"READ":   0,
		"WRITE":  1,
		"MKDIR":  2,
		"LS":     3,
		"DELETE": 4,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_file_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_file_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

type EntryType int32

const (
	EntryType_FILE      EntryType = 0
	EntryType_DIRECTORY EntryType = 1
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "FILE",
		1: "DIRECTORY",
	}
	EntryType_value = map[string]int32{
		"FILE":      0,
		"DIRECTORY": 1,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_file_proto_enumTypes[1].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_file_proto_enumTypes[1]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

type FSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DID       string     `protobuf:"bytes,1,opt,name=DID,proto3" json:"DID,omitempty"`
	Path      string     `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Signature []byte     `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Action    ActionType `protobuf:"varint,4,opt,name=action,proto3,enum=file.ActionType" json:"action,omitempty"`
}

func (x *FSRequest) Reset() {
	*x = FSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FSRequest) ProtoMessage() {}

func (x *FSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FSRequest.ProtoReflect.Descriptor instead.
func (*FSRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *FSRequest) GetDID() string {
	if x != nil {
		return x.DID
	}
	return ""
}

func (x *FSRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FSRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *FSRequest) GetAction() ActionType {
	if x != nil {
		return x.Action
	}
	return ActionType_READ
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	D []int32 `protobuf:"varint,1,rep,packed,name=d,proto3" json:"d,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetD() []int32 {
	if x != nil {
		return x.D
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action ActionType `protobuf:"varint,1,opt,name=action,proto3,enum=file.ActionType" json:"action,omitempty"`
	// Types that are assignable to Body:
	//	*Response_File
	//	*Response_Cid
	//	*Response_Entries
	Body isResponse_Body `protobuf_oneof:"body"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetAction() ActionType {
	if x != nil {
		return x.Action
	}
	return ActionType_READ
}

func (m *Response) GetBody() isResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Response) GetFile() *File {
	if x, ok := x.GetBody().(*Response_File); ok {
		return x.File
	}
	return nil
}

func (x *Response) GetCid() string {
	if x, ok := x.GetBody().(*Response_Cid); ok {
		return x.Cid
	}
	return ""
}

func (x *Response) GetEntries() *DirEntries {
	if x, ok := x.GetBody().(*Response_Entries); ok {
		return x.Entries
	}
	return nil
}

type isResponse_Body interface {
	isResponse_Body()
}

type Response_File struct {
	File *File `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type Response_Cid struct {
	Cid string `protobuf:"bytes,3,opt,name=cid,proto3,oneof"`
}

type Response_Entries struct {
	Entries *DirEntries `protobuf:"bytes,4,opt,name=entries,proto3,oneof"`
}

func (*Response_File) isResponse_Body() {}

func (*Response_Cid) isResponse_Body() {}

func (*Response_Entries) isResponse_Body() {}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	D    []int32 `protobuf:"varint,2,rep,packed,name=d,proto3" json:"d,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetD() []int32 {
	if x != nil {
		return x.D
	}
	return nil
}

type DirEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DirEntry `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DirEntries) Reset() {
	*x = DirEntries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirEntries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirEntries) ProtoMessage() {}

func (x *DirEntries) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirEntries.ProtoReflect.Descriptor instead.
func (*DirEntries) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *DirEntries) GetItems() []*DirEntry {
	if x != nil {
		return x.Items
	}
	return nil
}

type DirEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type EntryType `protobuf:"varint,2,opt,name=type,proto3,enum=file.EntryType" json:"type,omitempty"`
	Size int32     `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Cid  string    `protobuf:"bytes,4,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *DirEntry) Reset() {
	*x = DirEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirEntry) ProtoMessage() {}

func (x *DirEntry) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirEntry.ProtoReflect.Descriptor instead.
func (*DirEntry) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{5}
}

func (x *DirEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirEntry) GetType() EntryType {
	if x != nil {
		return x.Type
	}
	return EntryType_FILE
}

func (x *DirEntry) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DirEntry) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x79, 0x0a, 0x09, 0x46, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x44, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a,
	0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x01, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03,
	0x63, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x28, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x01, 0x64, 0x22, 0x32, 0x0a, 0x0a, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x69, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x69, 0x64, 0x2a, 0x40, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4b, 0x44, 0x49, 0x52, 0x10, 0x02,
	0x12, 0x06, 0x0a, 0x02, 0x4c, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x04, 0x2a, 0x24, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44,
	0x49, 0x52, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6c, 0x61, 0x6e, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x6c, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_file_proto_goTypes = []interface{}{
	(ActionType)(0),    // 0: file.ActionType
	(EntryType)(0),     // 1: file.EntryType
	(*FSRequest)(nil),  // 2: file.FSRequest
	(*Payload)(nil),    // 3: file.Payload
	(*Response)(nil),   // 4: file.Response
	(*File)(nil),       // 5: file.File
	(*DirEntries)(nil), // 6: file.DirEntries
	(*DirEntry)(nil),   // 7: file.DirEntry
}
var file_file_proto_depIdxs = []int32{
	0, // 0: file.FSRequest.action:type_name -> file.ActionType
	0, // 1: file.Response.action:type_name -> file.ActionType
	5, // 2: file.Response.file:type_name -> file.File
	6, // 3: file.Response.entries:type_name -> file.DirEntries
	7, // 4: file.DirEntries.items:type_name -> file.DirEntry
	1, // 5: file.DirEntry.type:type_name -> file.EntryType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FSRequest); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirEntries); i {
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
		file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirEntry); i {
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
	file_file_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Response_File)(nil),
		(*Response_Cid)(nil),
		(*Response_Entries)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		EnumInfos:         file_file_proto_enumTypes,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
