// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: fight.proto

package fight

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Hero_Gender int32

const (
	Hero_MALE   Hero_Gender = 0
	Hero_FEMALE Hero_Gender = 1
)

// Enum value maps for Hero_Gender.
var (
	Hero_Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Hero_Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Hero_Gender) Enum() *Hero_Gender {
	p := new(Hero_Gender)
	*p = x
	return p
}

func (x Hero_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hero_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_fight_proto_enumTypes[0].Descriptor()
}

func (Hero_Gender) Type() protoreflect.EnumType {
	return &file_fight_proto_enumTypes[0]
}

func (x Hero_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hero_Gender.Descriptor instead.
func (Hero_Gender) EnumDescriptor() ([]byte, []int) {
	return file_fight_proto_rawDescGZIP(), []int{1, 0}
}

type ListHerosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// support: specify hero names
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListHerosRequest) Reset() {
	*x = ListHerosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHerosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHerosRequest) ProtoMessage() {}

func (x *ListHerosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHerosRequest.ProtoReflect.Descriptor instead.
func (*ListHerosRequest) Descriptor() ([]byte, []int) {
	return file_fight_proto_rawDescGZIP(), []int{0}
}

func (x *ListHerosRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age    uint32      `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Gender Hero_Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=fight.Hero_Gender" json:"gender,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_fight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_fight_proto_rawDescGZIP(), []int{1}
}

func (x *Hero) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hero) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Hero) GetGender() Hero_Gender {
	if x != nil {
		return x.Gender
	}
	return Hero_MALE
}

var File_fight_proto protoreflect.FileDescriptor

var file_fight_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x78,
	0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x66,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x32, 0x3c, 0x0a, 0x05, 0x46, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x33, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x73, 0x12, 0x17,
	0x2e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x2e,
	0x48, 0x65, 0x72, 0x6f, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x66, 0x69, 0x67, 0x68,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fight_proto_rawDescOnce sync.Once
	file_fight_proto_rawDescData = file_fight_proto_rawDesc
)

func file_fight_proto_rawDescGZIP() []byte {
	file_fight_proto_rawDescOnce.Do(func() {
		file_fight_proto_rawDescData = protoimpl.X.CompressGZIP(file_fight_proto_rawDescData)
	})
	return file_fight_proto_rawDescData
}

var file_fight_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fight_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fight_proto_goTypes = []interface{}{
	(Hero_Gender)(0),         // 0: fight.Hero.Gender
	(*ListHerosRequest)(nil), // 1: fight.ListHerosRequest
	(*Hero)(nil),             // 2: fight.Hero
}
var file_fight_proto_depIdxs = []int32{
	0, // 0: fight.Hero.gender:type_name -> fight.Hero.Gender
	1, // 1: fight.Fight.ListHeros:input_type -> fight.ListHerosRequest
	2, // 2: fight.Fight.ListHeros:output_type -> fight.Hero
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fight_proto_init() }
func file_fight_proto_init() {
	if File_fight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHerosRequest); i {
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
		file_fight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
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
			RawDescriptor: file_fight_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fight_proto_goTypes,
		DependencyIndexes: file_fight_proto_depIdxs,
		EnumInfos:         file_fight_proto_enumTypes,
		MessageInfos:      file_fight_proto_msgTypes,
	}.Build()
	File_fight_proto = out.File
	file_fight_proto_rawDesc = nil
	file_fight_proto_goTypes = nil
	file_fight_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FightClient is the client API for Fight service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FightClient interface {
	ListHeros(ctx context.Context, in *ListHerosRequest, opts ...grpc.CallOption) (Fight_ListHerosClient, error)
}

type fightClient struct {
	cc grpc.ClientConnInterface
}

func NewFightClient(cc grpc.ClientConnInterface) FightClient {
	return &fightClient{cc}
}

func (c *fightClient) ListHeros(ctx context.Context, in *ListHerosRequest, opts ...grpc.CallOption) (Fight_ListHerosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fight_serviceDesc.Streams[0], "/fight.Fight/ListHeros", opts...)
	if err != nil {
		return nil, err
	}
	x := &fightListHerosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fight_ListHerosClient interface {
	Recv() (*Hero, error)
	grpc.ClientStream
}

type fightListHerosClient struct {
	grpc.ClientStream
}

func (x *fightListHerosClient) Recv() (*Hero, error) {
	m := new(Hero)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FightServer is the server API for Fight service.
type FightServer interface {
	ListHeros(*ListHerosRequest, Fight_ListHerosServer) error
}

// UnimplementedFightServer can be embedded to have forward compatible implementations.
type UnimplementedFightServer struct {
}

func (*UnimplementedFightServer) ListHeros(*ListHerosRequest, Fight_ListHerosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListHeros not implemented")
}

func RegisterFightServer(s *grpc.Server, srv FightServer) {
	s.RegisterService(&_Fight_serviceDesc, srv)
}

func _Fight_ListHeros_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListHerosRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FightServer).ListHeros(m, &fightListHerosServer{stream})
}

type Fight_ListHerosServer interface {
	Send(*Hero) error
	grpc.ServerStream
}

type fightListHerosServer struct {
	grpc.ServerStream
}

func (x *fightListHerosServer) Send(m *Hero) error {
	return x.ServerStream.SendMsg(m)
}

var _Fight_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fight.Fight",
	HandlerType: (*FightServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListHeros",
			Handler:       _Fight_ListHeros_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fight.proto",
}