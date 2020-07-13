// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: calculator.proto

package calculatorpb

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

type PrimeCompose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNo  int32 `protobuf:"varint,1,opt,name=firstNo,proto3" json:"firstNo,omitempty"`
	SecondNo int32 `protobuf:"varint,2,opt,name=secondNo,proto3" json:"secondNo,omitempty"`
}

func (x *PrimeCompose) Reset() {
	*x = PrimeCompose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeCompose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeCompose) ProtoMessage() {}

func (x *PrimeCompose) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeCompose.ProtoReflect.Descriptor instead.
func (*PrimeCompose) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeCompose) GetFirstNo() int32 {
	if x != nil {
		return x.FirstNo
	}
	return 0
}

func (x *PrimeCompose) GetSecondNo() int32 {
	if x != nil {
		return x.SecondNo
	}
	return 0
}

type PrimeCompositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Primecompose *PrimeCompose `protobuf:"bytes,1,opt,name=primecompose,proto3" json:"primecompose,omitempty"`
}

func (x *PrimeCompositionRequest) Reset() {
	*x = PrimeCompositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeCompositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeCompositionRequest) ProtoMessage() {}

func (x *PrimeCompositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeCompositionRequest.ProtoReflect.Descriptor instead.
func (*PrimeCompositionRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeCompositionRequest) GetPrimecompose() *PrimeCompose {
	if x != nil {
		return x.Primecompose
	}
	return nil
}

type PrimeCompositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeCompositionResponse) Reset() {
	*x = PrimeCompositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeCompositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeCompositionResponse) ProtoMessage() {}

func (x *PrimeCompositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeCompositionResponse.ProtoReflect.Descriptor instead.
func (*PrimeCompositionResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeCompositionResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x44,
	0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x4e, 0x6f, 0x22, 0x57, 0x0a, 0x17, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3c, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52,
	0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x32, 0x0a,
	0x18, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x7c, 0x0a, 0x17, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x10,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x19, 0x5a, 0x17, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_calculator_proto_goTypes = []interface{}{
	(*PrimeCompose)(nil),             // 0: calculator.PrimeCompose
	(*PrimeCompositionRequest)(nil),  // 1: calculator.PrimeCompositionRequest
	(*PrimeCompositionResponse)(nil), // 2: calculator.PrimeCompositionResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.PrimeCompositionRequest.primecompose:type_name -> calculator.PrimeCompose
	1, // 1: calculator.PrimeCompositionService.PrimeComposition:input_type -> calculator.PrimeCompositionRequest
	2, // 2: calculator.PrimeCompositionService.PrimeComposition:output_type -> calculator.PrimeCompositionResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeCompose); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeCompositionRequest); i {
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
		file_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeCompositionResponse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PrimeCompositionServiceClient is the client API for PrimeCompositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimeCompositionServiceClient interface {
	PrimeComposition(ctx context.Context, in *PrimeCompositionRequest, opts ...grpc.CallOption) (PrimeCompositionService_PrimeCompositionClient, error)
}

type primeCompositionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrimeCompositionServiceClient(cc grpc.ClientConnInterface) PrimeCompositionServiceClient {
	return &primeCompositionServiceClient{cc}
}

func (c *primeCompositionServiceClient) PrimeComposition(ctx context.Context, in *PrimeCompositionRequest, opts ...grpc.CallOption) (PrimeCompositionService_PrimeCompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeCompositionService_serviceDesc.Streams[0], "/calculator.PrimeCompositionService/PrimeComposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeCompositionServicePrimeCompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeCompositionService_PrimeCompositionClient interface {
	Recv() (*PrimeCompositionResponse, error)
	grpc.ClientStream
}

type primeCompositionServicePrimeCompositionClient struct {
	grpc.ClientStream
}

func (x *primeCompositionServicePrimeCompositionClient) Recv() (*PrimeCompositionResponse, error) {
	m := new(PrimeCompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PrimeCompositionServiceServer is the server API for PrimeCompositionService service.
type PrimeCompositionServiceServer interface {
	PrimeComposition(*PrimeCompositionRequest, PrimeCompositionService_PrimeCompositionServer) error
}

// UnimplementedPrimeCompositionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrimeCompositionServiceServer struct {
}

func (*UnimplementedPrimeCompositionServiceServer) PrimeComposition(*PrimeCompositionRequest, PrimeCompositionService_PrimeCompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeComposition not implemented")
}

func RegisterPrimeCompositionServiceServer(s *grpc.Server, srv PrimeCompositionServiceServer) {
	s.RegisterService(&_PrimeCompositionService_serviceDesc, srv)
}

func _PrimeCompositionService_PrimeComposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeCompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeCompositionServiceServer).PrimeComposition(m, &primeCompositionServicePrimeCompositionServer{stream})
}

type PrimeCompositionService_PrimeCompositionServer interface {
	Send(*PrimeCompositionResponse) error
	grpc.ServerStream
}

type primeCompositionServicePrimeCompositionServer struct {
	grpc.ServerStream
}

func (x *primeCompositionServicePrimeCompositionServer) Send(m *PrimeCompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PrimeCompositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.PrimeCompositionService",
	HandlerType: (*PrimeCompositionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeComposition",
			Handler:       _PrimeCompositionService_PrimeComposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}