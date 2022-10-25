// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: swissknife.proto

package swissknife

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

type ConvHexa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConvHexa) Reset() {
	*x = ConvHexa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swissknife_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvHexa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvHexa) ProtoMessage() {}

func (x *ConvHexa) ProtoReflect() protoreflect.Message {
	mi := &file_swissknife_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvHexa.ProtoReflect.Descriptor instead.
func (*ConvHexa) Descriptor() ([]byte, []int) {
	return file_swissknife_proto_rawDescGZIP(), []int{0}
}

type ConvBase64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConvBase64) Reset() {
	*x = ConvBase64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swissknife_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvBase64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvBase64) ProtoMessage() {}

func (x *ConvBase64) ProtoReflect() protoreflect.Message {
	mi := &file_swissknife_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvBase64.ProtoReflect.Descriptor instead.
func (*ConvBase64) Descriptor() ([]byte, []int) {
	return file_swissknife_proto_rawDescGZIP(), []int{1}
}

type ConvHexa_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *ConvHexa_Request) Reset() {
	*x = ConvHexa_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swissknife_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvHexa_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvHexa_Request) ProtoMessage() {}

func (x *ConvHexa_Request) ProtoReflect() protoreflect.Message {
	mi := &file_swissknife_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvHexa_Request.ProtoReflect.Descriptor instead.
func (*ConvHexa_Request) Descriptor() ([]byte, []int) {
	return file_swissknife_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConvHexa_Request) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type ConvHexa_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ConvHexa_Response) Reset() {
	*x = ConvHexa_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swissknife_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvHexa_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvHexa_Response) ProtoMessage() {}

func (x *ConvHexa_Response) ProtoReflect() protoreflect.Message {
	mi := &file_swissknife_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvHexa_Response.ProtoReflect.Descriptor instead.
func (*ConvHexa_Response) Descriptor() ([]byte, []int) {
	return file_swissknife_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ConvHexa_Response) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type ConvBase64_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *ConvBase64_Request) Reset() {
	*x = ConvBase64_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swissknife_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvBase64_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvBase64_Request) ProtoMessage() {}

func (x *ConvBase64_Request) ProtoReflect() protoreflect.Message {
	mi := &file_swissknife_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvBase64_Request.ProtoReflect.Descriptor instead.
func (*ConvBase64_Request) Descriptor() ([]byte, []int) {
	return file_swissknife_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ConvBase64_Request) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type ConvBase64_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ConvBase64_Response) Reset() {
	*x = ConvBase64_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_swissknife_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvBase64_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvBase64_Response) ProtoMessage() {}

func (x *ConvBase64_Response) ProtoReflect() protoreflect.Message {
	mi := &file_swissknife_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvBase64_Response.ProtoReflect.Descriptor instead.
func (*ConvBase64_Response) Descriptor() ([]byte, []int) {
	return file_swissknife_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ConvBase64_Response) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_swissknife_proto protoreflect.FileDescriptor

var file_swissknife_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x77, 0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x26, 0x70, 0x6d, 0x67, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x73, 0x77, 0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x22, 0x4f, 0x0a, 0x08, 0x43, 0x6f,
	0x6e, 0x76, 0x48, 0x65, 0x78, 0x61, 0x1a, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x51, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x76, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x1a, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x22, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x9d,
	0x02, 0x0a, 0x0d, 0x53, 0x77, 0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x53, 0x76, 0x63,
	0x12, 0x81, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x76, 0x48, 0x65, 0x78, 0x61, 0x12, 0x38, 0x2e,
	0x70, 0x6d, 0x67, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x77, 0x69, 0x73,
	0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x48, 0x65, 0x78, 0x61, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x6d, 0x67, 0x5f, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x6b, 0x69, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x77, 0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x48, 0x65, 0x78, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x76, 0x42, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x12, 0x3a, 0x2e, 0x70, 0x6d, 0x67, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x77, 0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3b, 0x2e, 0x70, 0x6d, 0x67, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x77,
	0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x42, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x73, 0x77, 0x69, 0x73, 0x73, 0x6b, 0x6e, 0x69, 0x66, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_swissknife_proto_rawDescOnce sync.Once
	file_swissknife_proto_rawDescData = file_swissknife_proto_rawDesc
)

func file_swissknife_proto_rawDescGZIP() []byte {
	file_swissknife_proto_rawDescOnce.Do(func() {
		file_swissknife_proto_rawDescData = protoimpl.X.CompressGZIP(file_swissknife_proto_rawDescData)
	})
	return file_swissknife_proto_rawDescData
}

var file_swissknife_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_swissknife_proto_goTypes = []interface{}{
	(*ConvHexa)(nil),            // 0: pmg_tools.adapterkit_module_swissknife.ConvHexa
	(*ConvBase64)(nil),          // 1: pmg_tools.adapterkit_module_swissknife.ConvBase64
	(*ConvHexa_Request)(nil),    // 2: pmg_tools.adapterkit_module_swissknife.ConvHexa.Request
	(*ConvHexa_Response)(nil),   // 3: pmg_tools.adapterkit_module_swissknife.ConvHexa.Response
	(*ConvBase64_Request)(nil),  // 4: pmg_tools.adapterkit_module_swissknife.ConvBase64.Request
	(*ConvBase64_Response)(nil), // 5: pmg_tools.adapterkit_module_swissknife.ConvBase64.Response
}
var file_swissknife_proto_depIdxs = []int32{
	2, // 0: pmg_tools.adapterkit_module_swissknife.SwissknifeSvc.ConvHexa:input_type -> pmg_tools.adapterkit_module_swissknife.ConvHexa.Request
	4, // 1: pmg_tools.adapterkit_module_swissknife.SwissknifeSvc.ConvBase64:input_type -> pmg_tools.adapterkit_module_swissknife.ConvBase64.Request
	3, // 2: pmg_tools.adapterkit_module_swissknife.SwissknifeSvc.ConvHexa:output_type -> pmg_tools.adapterkit_module_swissknife.ConvHexa.Response
	5, // 3: pmg_tools.adapterkit_module_swissknife.SwissknifeSvc.ConvBase64:output_type -> pmg_tools.adapterkit_module_swissknife.ConvBase64.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_swissknife_proto_init() }
func file_swissknife_proto_init() {
	if File_swissknife_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_swissknife_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvHexa); i {
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
		file_swissknife_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvBase64); i {
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
		file_swissknife_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvHexa_Request); i {
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
		file_swissknife_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvHexa_Response); i {
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
		file_swissknife_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvBase64_Request); i {
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
		file_swissknife_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvBase64_Response); i {
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
			RawDescriptor: file_swissknife_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_swissknife_proto_goTypes,
		DependencyIndexes: file_swissknife_proto_depIdxs,
		MessageInfos:      file_swissknife_proto_msgTypes,
	}.Build()
	File_swissknife_proto = out.File
	file_swissknife_proto_rawDesc = nil
	file_swissknife_proto_goTypes = nil
	file_swissknife_proto_depIdxs = nil
}