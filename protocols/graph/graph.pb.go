// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: fula/protocols/graph/graph.proto

package graph

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query          string         `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	VariableValues *_struct.Value `protobuf:"bytes,2,opt,name=variable_values,json=variableValues,proto3" json:"variable_values,omitempty"`
	OperationName  string         `protobuf:"bytes,3,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	Subscribe      bool           `protobuf:"varint,4,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fula_protocols_graph_graph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_fula_protocols_graph_graph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_fula_protocols_graph_graph_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Request) GetVariableValues() *_struct.Value {
	if x != nil {
		return x.VariableValues
	}
	return nil
}

func (x *Request) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

func (x *Request) GetSubscribe() bool {
	if x != nil {
		return x.Subscribe
	}
	return false
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to DataOrError:
	//	*Result_Data
	//	*Result_Error
	DataOrError isResult_DataOrError `protobuf_oneof:"data_or_error"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fula_protocols_graph_graph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_fula_protocols_graph_graph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_fula_protocols_graph_graph_proto_rawDescGZIP(), []int{1}
}

func (m *Result) GetDataOrError() isResult_DataOrError {
	if m != nil {
		return m.DataOrError
	}
	return nil
}

func (x *Result) GetData() *_struct.Value {
	if x, ok := x.GetDataOrError().(*Result_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Result) GetError() *_struct.Value {
	if x, ok := x.GetDataOrError().(*Result_Error); ok {
		return x.Error
	}
	return nil
}

type isResult_DataOrError interface {
	isResult_DataOrError()
}

type Result_Data struct {
	Data *_struct.Value `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type Result_Error struct {
	Error *_struct.Value `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*Result_Data) isResult_DataOrError() {}

func (*Result_Error) isResult_DataOrError() {}

var File_fula_protocols_graph_graph_proto protoreflect.FileDescriptor

var file_fula_protocols_graph_graph_proto_rawDesc = []byte{
	0x0a, 0x20, 0x66, 0x75, 0x6c, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x0f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22,
	0x77, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6f, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fula_protocols_graph_graph_proto_rawDescOnce sync.Once
	file_fula_protocols_graph_graph_proto_rawDescData = file_fula_protocols_graph_graph_proto_rawDesc
)

func file_fula_protocols_graph_graph_proto_rawDescGZIP() []byte {
	file_fula_protocols_graph_graph_proto_rawDescOnce.Do(func() {
		file_fula_protocols_graph_graph_proto_rawDescData = protoimpl.X.CompressGZIP(file_fula_protocols_graph_graph_proto_rawDescData)
	})
	return file_fula_protocols_graph_graph_proto_rawDescData
}

var file_fula_protocols_graph_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fula_protocols_graph_graph_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: graph.Request
	(*Result)(nil),        // 1: graph.Result
	(*_struct.Value)(nil), // 2: google.protobuf.Value
}
var file_fula_protocols_graph_graph_proto_depIdxs = []int32{
	2, // 0: graph.Request.variable_values:type_name -> google.protobuf.Value
	2, // 1: graph.Result.data:type_name -> google.protobuf.Value
	2, // 2: graph.Result.error:type_name -> google.protobuf.Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fula_protocols_graph_graph_proto_init() }
func file_fula_protocols_graph_graph_proto_init() {
	if File_fula_protocols_graph_graph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fula_protocols_graph_graph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_fula_protocols_graph_graph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
	file_fula_protocols_graph_graph_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Result_Data)(nil),
		(*Result_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fula_protocols_graph_graph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fula_protocols_graph_graph_proto_goTypes,
		DependencyIndexes: file_fula_protocols_graph_graph_proto_depIdxs,
		MessageInfos:      file_fula_protocols_graph_graph_proto_msgTypes,
	}.Build()
	File_fula_protocols_graph_graph_proto = out.File
	file_fula_protocols_graph_graph_proto_rawDesc = nil
	file_fula_protocols_graph_graph_proto_goTypes = nil
	file_fula_protocols_graph_graph_proto_depIdxs = nil
}