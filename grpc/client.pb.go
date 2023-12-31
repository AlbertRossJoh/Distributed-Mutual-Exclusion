// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: grpc/client.proto

package proto

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

type State int32

const (
	State_HELD     State = 0
	State_WANTED   State = 1
	State_RELEASED State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "HELD",
		1: "WANTED",
		2: "RELEASED",
	}
	State_value = map[string]int32{
		"HELD":     0,
		"WANTED":   1,
		"RELEASED": 2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_client_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_grpc_client_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_grpc_client_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LamportTs int64  `protobuf:"varint,1,opt,name=lamportTs,proto3" json:"lamportTs,omitempty"`
	State     State  `protobuf:"varint,2,opt,name=state,proto3,enum=proto.State" json:"state,omitempty"`
	Id        string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_client_proto_msgTypes[0]
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
	return file_grpc_client_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetLamportTs() int64 {
	if x != nil {
		return x.LamportTs
	}
	return 0
}

func (x *Request) GetState() State {
	if x != nil {
		return x.State
	}
	return State_HELD
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_client_proto_msgTypes[1]
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
	return file_grpc_client_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_grpc_client_proto protoreflect.FileDescriptor

var file_grpc_client_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x2b, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x4c, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x57, 0x41, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x4c, 0x45, 0x41, 0x53, 0x45, 0x44, 0x10, 0x02, 0x32, 0x69, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x61, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_client_proto_rawDescOnce sync.Once
	file_grpc_client_proto_rawDescData = file_grpc_client_proto_rawDesc
)

func file_grpc_client_proto_rawDescGZIP() []byte {
	file_grpc_client_proto_rawDescOnce.Do(func() {
		file_grpc_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_client_proto_rawDescData)
	})
	return file_grpc_client_proto_rawDescData
}

var file_grpc_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_client_proto_goTypes = []interface{}{
	(State)(0),       // 0: proto.State
	(*Request)(nil),  // 1: proto.Request
	(*Response)(nil), // 2: proto.Response
}
var file_grpc_client_proto_depIdxs = []int32{
	0, // 0: proto.Request.state:type_name -> proto.State
	1, // 1: proto.ClientService.MakeRequest:input_type -> proto.Request
	1, // 2: proto.ClientService.Reply:input_type -> proto.Request
	2, // 3: proto.ClientService.MakeRequest:output_type -> proto.Response
	2, // 4: proto.ClientService.Reply:output_type -> proto.Response
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_client_proto_init() }
func file_grpc_client_proto_init() {
	if File_grpc_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_client_proto_goTypes,
		DependencyIndexes: file_grpc_client_proto_depIdxs,
		EnumInfos:         file_grpc_client_proto_enumTypes,
		MessageInfos:      file_grpc_client_proto_msgTypes,
	}.Build()
	File_grpc_client_proto = out.File
	file_grpc_client_proto_rawDesc = nil
	file_grpc_client_proto_goTypes = nil
	file_grpc_client_proto_depIdxs = nil
}
