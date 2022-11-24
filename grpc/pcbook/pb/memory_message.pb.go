// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: grpc/pcbook/proto/memory_message.proto

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

type Memory_Unit int32

const (
	Memory_Unknown  Memory_Unit = 0
	Memory_Bit      Memory_Unit = 1
	Memory_Byte     Memory_Unit = 2
	Memory_Kilobyte Memory_Unit = 3
	Memory_Megabyte Memory_Unit = 4
	Memory_Gigabyte Memory_Unit = 5
	Memory_Terabyte Memory_Unit = 6
)

// Enum value maps for Memory_Unit.
var (
	Memory_Unit_name = map[int32]string{
		0: "Unknown",
		1: "Bit",
		2: "Byte",
		3: "Kilobyte",
		4: "Megabyte",
		5: "Gigabyte",
		6: "Terabyte",
	}
	Memory_Unit_value = map[string]int32{
		"Unknown":  0,
		"Bit":      1,
		"Byte":     2,
		"Kilobyte": 3,
		"Megabyte": 4,
		"Gigabyte": 5,
		"Terabyte": 6,
	}
)

func (x Memory_Unit) Enum() *Memory_Unit {
	p := new(Memory_Unit)
	*p = x
	return p
}

func (x Memory_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Memory_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_pcbook_proto_memory_message_proto_enumTypes[0].Descriptor()
}

func (Memory_Unit) Type() protoreflect.EnumType {
	return &file_grpc_pcbook_proto_memory_message_proto_enumTypes[0]
}

func (x Memory_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Memory_Unit.Descriptor instead.
func (Memory_Unit) EnumDescriptor() ([]byte, []int) {
	return file_grpc_pcbook_proto_memory_message_proto_rawDescGZIP(), []int{0, 0}
}

type Memory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64      `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit  Memory_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=pb.Memory_Unit" json:"unit,omitempty"`
}

func (x *Memory) Reset() {
	*x = Memory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pcbook_proto_memory_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memory) ProtoMessage() {}

func (x *Memory) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pcbook_proto_memory_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memory.ProtoReflect.Descriptor instead.
func (*Memory) Descriptor() ([]byte, []int) {
	return file_grpc_pcbook_proto_memory_message_proto_rawDescGZIP(), []int{0}
}

func (x *Memory) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Memory) GetUnit() Memory_Unit {
	if x != nil {
		return x.Unit
	}
	return Memory_Unknown
}

var File_grpc_pcbook_proto_memory_message_proto protoreflect.FileDescriptor

var file_grpc_pcbook_proto_memory_message_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x63, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa3, 0x01, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x22, 0x5e, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x69, 0x74, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x42, 0x79, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x69,
	0x6c, 0x6f, 0x62, 0x79, 0x74, 0x65, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x65, 0x67, 0x61,
	0x62, 0x79, 0x74, 0x65, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x69, 0x67, 0x61, 0x62, 0x79,
	0x74, 0x65, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x61, 0x62, 0x79, 0x74, 0x65,
	0x10, 0x06, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x63, 0x62,
	0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_pcbook_proto_memory_message_proto_rawDescOnce sync.Once
	file_grpc_pcbook_proto_memory_message_proto_rawDescData = file_grpc_pcbook_proto_memory_message_proto_rawDesc
)

func file_grpc_pcbook_proto_memory_message_proto_rawDescGZIP() []byte {
	file_grpc_pcbook_proto_memory_message_proto_rawDescOnce.Do(func() {
		file_grpc_pcbook_proto_memory_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_pcbook_proto_memory_message_proto_rawDescData)
	})
	return file_grpc_pcbook_proto_memory_message_proto_rawDescData
}

var file_grpc_pcbook_proto_memory_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_pcbook_proto_memory_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_pcbook_proto_memory_message_proto_goTypes = []interface{}{
	(Memory_Unit)(0), // 0: pb.Memory.Unit
	(*Memory)(nil),   // 1: pb.Memory
}
var file_grpc_pcbook_proto_memory_message_proto_depIdxs = []int32{
	0, // 0: pb.Memory.unit:type_name -> pb.Memory.Unit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_pcbook_proto_memory_message_proto_init() }
func file_grpc_pcbook_proto_memory_message_proto_init() {
	if File_grpc_pcbook_proto_memory_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_pcbook_proto_memory_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memory); i {
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
			RawDescriptor: file_grpc_pcbook_proto_memory_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_pcbook_proto_memory_message_proto_goTypes,
		DependencyIndexes: file_grpc_pcbook_proto_memory_message_proto_depIdxs,
		EnumInfos:         file_grpc_pcbook_proto_memory_message_proto_enumTypes,
		MessageInfos:      file_grpc_pcbook_proto_memory_message_proto_msgTypes,
	}.Build()
	File_grpc_pcbook_proto_memory_message_proto = out.File
	file_grpc_pcbook_proto_memory_message_proto_rawDesc = nil
	file_grpc_pcbook_proto_memory_message_proto_goTypes = nil
	file_grpc_pcbook_proto_memory_message_proto_depIdxs = nil
}