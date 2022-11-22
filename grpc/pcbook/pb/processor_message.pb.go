// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/pcbook/proto/processor_message.proto

package pb

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

type CPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand      string  `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumCores   uint32  `protobuf:"varint,3,opt,name=numCores,proto3" json:"numCores,omitempty"`
	NumThreads uint32  `protobuf:"varint,4,opt,name=numThreads,proto3" json:"numThreads,omitempty"`
	MinGhz     float64 `protobuf:"fixed64,5,opt,name=minGhz,proto3" json:"minGhz,omitempty"`
	MaxGhz     float64 `protobuf:"fixed64,6,opt,name=maxGhz,proto3" json:"maxGhz,omitempty"`
}

func (x *CPU) Reset() {
	*x = CPU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pcbook_proto_processor_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPU) ProtoMessage() {}

func (x *CPU) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pcbook_proto_processor_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
func (*CPU) Descriptor() ([]byte, []int) {
	return file_grpc_pcbook_proto_processor_message_proto_rawDescGZIP(), []int{0}
}

func (x *CPU) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CPU) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CPU) GetNumCores() uint32 {
	if x != nil {
		return x.NumCores
	}
	return 0
}

func (x *CPU) GetNumThreads() uint32 {
	if x != nil {
		return x.NumThreads
	}
	return 0
}

func (x *CPU) GetMinGhz() float64 {
	if x != nil {
		return x.MinGhz
	}
	return 0
}

func (x *CPU) GetMaxGhz() float64 {
	if x != nil {
		return x.MaxGhz
	}
	return 0
}

var File_grpc_pcbook_proto_processor_message_proto protoreflect.FileDescriptor

var file_grpc_pcbook_proto_processor_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x9b, 0x01, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x6e, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x69, 0x6e, 0x47, 0x68, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6d,
	0x69, 0x6e, 0x47, 0x68, 0x7a, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x47, 0x68, 0x7a, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x47, 0x68, 0x7a, 0x42, 0x12, 0x5a,
	0x10, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_pcbook_proto_processor_message_proto_rawDescOnce sync.Once
	file_grpc_pcbook_proto_processor_message_proto_rawDescData = file_grpc_pcbook_proto_processor_message_proto_rawDesc
)

func file_grpc_pcbook_proto_processor_message_proto_rawDescGZIP() []byte {
	file_grpc_pcbook_proto_processor_message_proto_rawDescOnce.Do(func() {
		file_grpc_pcbook_proto_processor_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_pcbook_proto_processor_message_proto_rawDescData)
	})
	return file_grpc_pcbook_proto_processor_message_proto_rawDescData
}

var file_grpc_pcbook_proto_processor_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_pcbook_proto_processor_message_proto_goTypes = []interface{}{
	(*CPU)(nil), // 0: pb.CPU
}
var file_grpc_pcbook_proto_processor_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_pcbook_proto_processor_message_proto_init() }
func file_grpc_pcbook_proto_processor_message_proto_init() {
	if File_grpc_pcbook_proto_processor_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_pcbook_proto_processor_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CPU); i {
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
			RawDescriptor: file_grpc_pcbook_proto_processor_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_pcbook_proto_processor_message_proto_goTypes,
		DependencyIndexes: file_grpc_pcbook_proto_processor_message_proto_depIdxs,
		MessageInfos:      file_grpc_pcbook_proto_processor_message_proto_msgTypes,
	}.Build()
	File_grpc_pcbook_proto_processor_message_proto = out.File
	file_grpc_pcbook_proto_processor_message_proto_rawDesc = nil
	file_grpc_pcbook_proto_processor_message_proto_goTypes = nil
	file_grpc_pcbook_proto_processor_message_proto_depIdxs = nil
}
