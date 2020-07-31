// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: annotation.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FIELD_BIND int32

const (
	FIELD_BIND_FROM_DEFAULT FIELD_BIND = 0
	FIELD_BIND_FROM_CONTEXT FIELD_BIND = 1
	FIELD_BIND_FROM_QUERY   FIELD_BIND = 2
	FIELD_BIND_FROM_HEADER  FIELD_BIND = 3
	FIELD_BIND_FROM_PARAMS  FIELD_BIND = 4
)

// Enum value maps for FIELD_BIND.
var (
	FIELD_BIND_name = map[int32]string{
		0: "FROM_DEFAULT",
		1: "FROM_CONTEXT",
		2: "FROM_QUERY",
		3: "FROM_HEADER",
		4: "FROM_PARAMS",
	}
	FIELD_BIND_value = map[string]int32{
		"FROM_DEFAULT": 0,
		"FROM_CONTEXT": 1,
		"FROM_QUERY":   2,
		"FROM_HEADER":  3,
		"FROM_PARAMS":  4,
	}
)

func (x FIELD_BIND) Enum() *FIELD_BIND {
	p := new(FIELD_BIND)
	*p = x
	return p
}

func (x FIELD_BIND) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FIELD_BIND) Descriptor() protoreflect.EnumDescriptor {
	return file_annotation_proto_enumTypes[0].Descriptor()
}

func (FIELD_BIND) Type() protoreflect.EnumType {
	return &file_annotation_proto_enumTypes[0]
}

func (x FIELD_BIND) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FIELD_BIND.Descriptor instead.
func (FIELD_BIND) EnumDescriptor() ([]byte, []int) {
	return file_annotation_proto_rawDescGZIP(), []int{0}
}

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Pattern:
	//	*Http_Post
	//	*Http_Get
	//	*Http_Delete
	//	*Http_Put
	//	*Http_Patch
	//	*Http_Option
	Pattern isHttp_Pattern `protobuf_oneof:"pattern"`
	Use     []string       `protobuf:"bytes,7,rep,name=use,proto3" json:"use,omitempty"`
	Timeout int32          `protobuf:"varint,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Handler string         `protobuf:"bytes,9,opt,name=handler,proto3" json:"handler,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_annotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_annotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_annotation_proto_rawDescGZIP(), []int{0}
}

func (m *Http) GetPattern() isHttp_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (x *Http) GetPost() string {
	if x, ok := x.GetPattern().(*Http_Post); ok {
		return x.Post
	}
	return ""
}

func (x *Http) GetGet() string {
	if x, ok := x.GetPattern().(*Http_Get); ok {
		return x.Get
	}
	return ""
}

func (x *Http) GetDelete() string {
	if x, ok := x.GetPattern().(*Http_Delete); ok {
		return x.Delete
	}
	return ""
}

func (x *Http) GetPut() string {
	if x, ok := x.GetPattern().(*Http_Put); ok {
		return x.Put
	}
	return ""
}

func (x *Http) GetPatch() string {
	if x, ok := x.GetPattern().(*Http_Patch); ok {
		return x.Patch
	}
	return ""
}

func (x *Http) GetOption() string {
	if x, ok := x.GetPattern().(*Http_Option); ok {
		return x.Option
	}
	return ""
}

func (x *Http) GetUse() []string {
	if x != nil {
		return x.Use
	}
	return nil
}

func (x *Http) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Http) GetHandler() string {
	if x != nil {
		return x.Handler
	}
	return ""
}

type isHttp_Pattern interface {
	isHttp_Pattern()
}

type Http_Post struct {
	Post string `protobuf:"bytes,1,opt,name=post,proto3,oneof"`
}

type Http_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type Http_Delete struct {
	Delete string `protobuf:"bytes,3,opt,name=delete,proto3,oneof"`
}

type Http_Put struct {
	Put string `protobuf:"bytes,4,opt,name=put,proto3,oneof"`
}

type Http_Patch struct {
	Patch string `protobuf:"bytes,5,opt,name=patch,proto3,oneof"`
}

type Http_Option struct {
	Option string `protobuf:"bytes,6,opt,name=option,proto3,oneof"`
}

func (*Http_Post) isHttp_Pattern() {}

func (*Http_Get) isHttp_Pattern() {}

func (*Http_Delete) isHttp_Pattern() {}

func (*Http_Put) isHttp_Pattern() {}

func (*Http_Patch) isHttp_Pattern() {}

func (*Http_Option) isHttp_Pattern() {}

var file_annotation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Http)(nil),
		Field:         3110202,
		Name:          "gapi.http",
		Tag:           "bytes,3110202,opt,name=http",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         4110202,
		Name:          "gapi.server",
		Tag:           "bytes,4110202,opt,name=server",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         4110204,
		Name:          "gapi.default_handler",
		Tag:           "bytes,4110204,opt,name=default_handler",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         4110205,
		Name:          "gapi.default_timeout",
		Tag:           "varint,4110205,opt,name=default_timeout",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         4110206,
		Name:          "gapi.path_prefix",
		Tag:           "bytes,4110206,opt,name=path_prefix",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         5110202,
		Name:          "gapi.flat",
		Tag:           "varint,5110202,opt,name=flat",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         6110202,
		Name:          "gapi.alias",
		Tag:           "bytes,6110202,opt,name=alias",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         6110203,
		Name:          "gapi.omit_empty",
		Tag:           "varint,6110203,opt,name=omit_empty",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         6110204,
		Name:          "gapi.raw_data",
		Tag:           "varint,6110204,opt,name=raw_data",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         6110206,
		Name:          "gapi.from_context",
		Tag:           "varint,6110206,opt,name=from_context",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         6110207,
		Name:          "gapi.validate",
		Tag:           "varint,6110207,opt,name=validate",
		Filename:      "annotation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FIELD_BIND)(nil),
		Field:         6110209,
		Name:          "gapi.bind",
		Tag:           "varint,6110209,opt,name=bind,enum=gapi.FIELD_BIND",
		Filename:      "annotation.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional gapi.Http http = 3110202;
	E_Http = &file_annotation_proto_extTypes[0]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// optional string server = 4110202;
	E_Server = &file_annotation_proto_extTypes[1]
	// optional string default_handler = 4110204;
	E_DefaultHandler = &file_annotation_proto_extTypes[2]
	// optional int32 default_timeout = 4110205;
	E_DefaultTimeout = &file_annotation_proto_extTypes[3]
	// optional string path_prefix = 4110206;
	E_PathPrefix = &file_annotation_proto_extTypes[4]
)

// Extension fields to descriptor.MessageOptions.
var (
	// optional bool flat = 5110202;
	E_Flat = &file_annotation_proto_extTypes[5]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional string alias = 6110202;
	E_Alias = &file_annotation_proto_extTypes[6]
	// optional bool omit_empty = 6110203;
	E_OmitEmpty = &file_annotation_proto_extTypes[7]
	// optional bool raw_data = 6110204;
	E_RawData = &file_annotation_proto_extTypes[8]
	// optional bool from_context = 6110206;
	E_FromContext = &file_annotation_proto_extTypes[9]
	// optional bool validate = 6110207;
	E_Validate = &file_annotation_proto_extTypes[10]
	// optional gapi.FIELD_BIND bind = 6110209;
	E_Bind = &file_annotation_proto_extTypes[11]
)

var File_annotation_proto protoreflect.FileDescriptor

var file_annotation_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x04, 0x48,
	0x74, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x67, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x70, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x73, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x2a, 0x62,
	0x0a, 0x0a, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x12, 0x10, 0x0a, 0x0c,
	0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x53,
	0x10, 0x04, 0x3a, 0x41, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xba, 0xea, 0xbd, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x3a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xfa, 0xee, 0xfa, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x3a, 0x4b, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfc, 0xee, 0xfa, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x3a, 0x4b,
	0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xfd, 0xee, 0xfa, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x3a, 0x43, 0x0a, 0x0b, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0xee, 0xfa, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x3a, 0x36, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xba, 0xf3, 0xb7, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x74, 0x3a, 0x36, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xfa, 0xf7, 0xf4, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x3a, 0x3f, 0x0a, 0x0a, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfb, 0xf7,
	0xf4, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x3a, 0x3b, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfc, 0xf7, 0xf4,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x43,
	0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfe, 0xf7,
	0xf4, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x3a, 0x3c, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff,
	0xf7, 0xf4, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x46, 0x0a, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x81, 0xf8, 0xf4, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x42,
	0x49, 0x4e, 0x44, 0x52, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x69, 0x64, 0x75, 0x6f, 0x6b, 0x65,
	0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_annotation_proto_rawDescOnce sync.Once
	file_annotation_proto_rawDescData = file_annotation_proto_rawDesc
)

func file_annotation_proto_rawDescGZIP() []byte {
	file_annotation_proto_rawDescOnce.Do(func() {
		file_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_annotation_proto_rawDescData)
	})
	return file_annotation_proto_rawDescData
}

var file_annotation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_annotation_proto_goTypes = []interface{}{
	(FIELD_BIND)(0),                   // 0: gapi.FIELD_BIND
	(*Http)(nil),                      // 1: gapi.Http
	(*descriptor.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
	(*descriptor.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
	(*descriptor.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 5: google.protobuf.FieldOptions
}
var file_annotation_proto_depIdxs = []int32{
	2,  // 0: gapi.http:extendee -> google.protobuf.MethodOptions
	3,  // 1: gapi.server:extendee -> google.protobuf.ServiceOptions
	3,  // 2: gapi.default_handler:extendee -> google.protobuf.ServiceOptions
	3,  // 3: gapi.default_timeout:extendee -> google.protobuf.ServiceOptions
	3,  // 4: gapi.path_prefix:extendee -> google.protobuf.ServiceOptions
	4,  // 5: gapi.flat:extendee -> google.protobuf.MessageOptions
	5,  // 6: gapi.alias:extendee -> google.protobuf.FieldOptions
	5,  // 7: gapi.omit_empty:extendee -> google.protobuf.FieldOptions
	5,  // 8: gapi.raw_data:extendee -> google.protobuf.FieldOptions
	5,  // 9: gapi.from_context:extendee -> google.protobuf.FieldOptions
	5,  // 10: gapi.validate:extendee -> google.protobuf.FieldOptions
	5,  // 11: gapi.bind:extendee -> google.protobuf.FieldOptions
	1,  // 12: gapi.http:type_name -> gapi.Http
	0,  // 13: gapi.bind:type_name -> gapi.FIELD_BIND
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	12, // [12:14] is the sub-list for extension type_name
	0,  // [0:12] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_annotation_proto_init() }
func file_annotation_proto_init() {
	if File_annotation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_annotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
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
	file_annotation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Http_Post)(nil),
		(*Http_Get)(nil),
		(*Http_Delete)(nil),
		(*Http_Put)(nil),
		(*Http_Patch)(nil),
		(*Http_Option)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_annotation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 12,
			NumServices:   0,
		},
		GoTypes:           file_annotation_proto_goTypes,
		DependencyIndexes: file_annotation_proto_depIdxs,
		EnumInfos:         file_annotation_proto_enumTypes,
		MessageInfos:      file_annotation_proto_msgTypes,
		ExtensionInfos:    file_annotation_proto_extTypes,
	}.Build()
	File_annotation_proto = out.File
	file_annotation_proto_rawDesc = nil
	file_annotation_proto_goTypes = nil
	file_annotation_proto_depIdxs = nil
}
