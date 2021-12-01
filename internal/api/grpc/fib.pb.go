// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: fib.proto

package api

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

type Segment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	Stop  uint64 `protobuf:"varint,2,opt,name=stop,proto3" json:"stop,omitempty"`
}

func (x *Segment) Reset() {
	*x = Segment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Segment) ProtoMessage() {}

func (x *Segment) ProtoReflect() protoreflect.Message {
	mi := &file_fib_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Segment.ProtoReflect.Descriptor instead.
func (*Segment) Descriptor() ([]byte, []int) {
	return file_fib_proto_rawDescGZIP(), []int{0}
}

func (x *Segment) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Segment) GetStop() uint64 {
	if x != nil {
		return x.Stop
	}
	return 0
}

type FibonacciSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numbers []uint64 `protobuf:"varint,1,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *FibonacciSlice) Reset() {
	*x = FibonacciSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciSlice) ProtoMessage() {}

func (x *FibonacciSlice) ProtoReflect() protoreflect.Message {
	mi := &file_fib_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciSlice.ProtoReflect.Descriptor instead.
func (*FibonacciSlice) Descriptor() ([]byte, []int) {
	return file_fib_proto_rawDescGZIP(), []int{1}
}

func (x *FibonacciSlice) GetNumbers() []uint64 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

var File_fib_proto protoreflect.FileDescriptor

var file_fib_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x33, 0x0a, 0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x22, 0x2a, 0x0a, 0x0e, 0x46, 0x69,
	0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32, 0x49, 0x0a, 0x09, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61,
	0x63, 0x63, 0x69, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x62, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x22,
	0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6e, 0x61, 0x36, 0x39, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x69, 0x62, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fib_proto_rawDescOnce sync.Once
	file_fib_proto_rawDescData = file_fib_proto_rawDesc
)

func file_fib_proto_rawDescGZIP() []byte {
	file_fib_proto_rawDescOnce.Do(func() {
		file_fib_proto_rawDescData = protoimpl.X.CompressGZIP(file_fib_proto_rawDescData)
	})
	return file_fib_proto_rawDescData
}

var file_fib_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fib_proto_goTypes = []interface{}{
	(*Segment)(nil),        // 0: grpc_api.Segment
	(*FibonacciSlice)(nil), // 1: grpc_api.FibonacciSlice
}
var file_fib_proto_depIdxs = []int32{
	0, // 0: grpc_api.Fibonacci.GetFibSlice:input_type -> grpc_api.Segment
	1, // 1: grpc_api.Fibonacci.GetFibSlice:output_type -> grpc_api.FibonacciSlice
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fib_proto_init() }
func file_fib_proto_init() {
	if File_fib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Segment); i {
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
		file_fib_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciSlice); i {
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
			RawDescriptor: file_fib_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fib_proto_goTypes,
		DependencyIndexes: file_fib_proto_depIdxs,
		MessageInfos:      file_fib_proto_msgTypes,
	}.Build()
	File_fib_proto = out.File
	file_fib_proto_rawDesc = nil
	file_fib_proto_goTypes = nil
	file_fib_proto_depIdxs = nil
}
