// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Runtime/Events.Processing/PublicFilters.proto

package processing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "go.dolittle.io/chaostests/contracts/protobuf"
	services "go.dolittle.io/chaostests/contracts/services"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PublicFiltersRegistrationRequest struct {
	CallContext          *services.ReverseCallArgumentsContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	FilterId             *protobuf.Uuid                        `protobuf:"bytes,2,opt,name=filterId,proto3" json:"filterId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PublicFiltersRegistrationRequest) Reset()         { *m = PublicFiltersRegistrationRequest{} }
func (m *PublicFiltersRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*PublicFiltersRegistrationRequest) ProtoMessage()    {}
func (*PublicFiltersRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa415f8b859ec2b0, []int{0}
}

func (m *PublicFiltersRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicFiltersRegistrationRequest.Unmarshal(m, b)
}
func (m *PublicFiltersRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicFiltersRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *PublicFiltersRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicFiltersRegistrationRequest.Merge(m, src)
}
func (m *PublicFiltersRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_PublicFiltersRegistrationRequest.Size(m)
}
func (m *PublicFiltersRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicFiltersRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublicFiltersRegistrationRequest proto.InternalMessageInfo

func (m *PublicFiltersRegistrationRequest) GetCallContext() *services.ReverseCallArgumentsContext {
	if m != nil {
		return m.CallContext
	}
	return nil
}

func (m *PublicFiltersRegistrationRequest) GetFilterId() *protobuf.Uuid {
	if m != nil {
		return m.FilterId
	}
	return nil
}

type PublicFiltersClientToRuntimeMessage struct {
	// Types that are valid to be assigned to Message:
	//	*PublicFiltersClientToRuntimeMessage_RegistrationRequest
	//	*PublicFiltersClientToRuntimeMessage_FilterResult
	Message              isPublicFiltersClientToRuntimeMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *PublicFiltersClientToRuntimeMessage) Reset()         { *m = PublicFiltersClientToRuntimeMessage{} }
func (m *PublicFiltersClientToRuntimeMessage) String() string { return proto.CompactTextString(m) }
func (*PublicFiltersClientToRuntimeMessage) ProtoMessage()    {}
func (*PublicFiltersClientToRuntimeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa415f8b859ec2b0, []int{1}
}

func (m *PublicFiltersClientToRuntimeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicFiltersClientToRuntimeMessage.Unmarshal(m, b)
}
func (m *PublicFiltersClientToRuntimeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicFiltersClientToRuntimeMessage.Marshal(b, m, deterministic)
}
func (m *PublicFiltersClientToRuntimeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicFiltersClientToRuntimeMessage.Merge(m, src)
}
func (m *PublicFiltersClientToRuntimeMessage) XXX_Size() int {
	return xxx_messageInfo_PublicFiltersClientToRuntimeMessage.Size(m)
}
func (m *PublicFiltersClientToRuntimeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicFiltersClientToRuntimeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PublicFiltersClientToRuntimeMessage proto.InternalMessageInfo

type isPublicFiltersClientToRuntimeMessage_Message interface {
	isPublicFiltersClientToRuntimeMessage_Message()
}

type PublicFiltersClientToRuntimeMessage_RegistrationRequest struct {
	RegistrationRequest *PublicFiltersRegistrationRequest `protobuf:"bytes,1,opt,name=registrationRequest,proto3,oneof"`
}

type PublicFiltersClientToRuntimeMessage_FilterResult struct {
	FilterResult *FilterResponse `protobuf:"bytes,2,opt,name=filterResult,proto3,oneof"`
}

func (*PublicFiltersClientToRuntimeMessage_RegistrationRequest) isPublicFiltersClientToRuntimeMessage_Message() {
}

func (*PublicFiltersClientToRuntimeMessage_FilterResult) isPublicFiltersClientToRuntimeMessage_Message() {
}

func (m *PublicFiltersClientToRuntimeMessage) GetMessage() isPublicFiltersClientToRuntimeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PublicFiltersClientToRuntimeMessage) GetRegistrationRequest() *PublicFiltersRegistrationRequest {
	if x, ok := m.GetMessage().(*PublicFiltersClientToRuntimeMessage_RegistrationRequest); ok {
		return x.RegistrationRequest
	}
	return nil
}

func (m *PublicFiltersClientToRuntimeMessage) GetFilterResult() *FilterResponse {
	if x, ok := m.GetMessage().(*PublicFiltersClientToRuntimeMessage_FilterResult); ok {
		return x.FilterResult
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PublicFiltersClientToRuntimeMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PublicFiltersClientToRuntimeMessage_RegistrationRequest)(nil),
		(*PublicFiltersClientToRuntimeMessage_FilterResult)(nil),
	}
}

func init() {
	proto.RegisterType((*PublicFiltersRegistrationRequest)(nil), "dolittle.runtime.events.processing.PublicFiltersRegistrationRequest")
	proto.RegisterType((*PublicFiltersClientToRuntimeMessage)(nil), "dolittle.runtime.events.processing.PublicFiltersClientToRuntimeMessage")
}

func init() {
	proto.RegisterFile("Runtime/Events.Processing/PublicFilters.proto", fileDescriptor_aa415f8b859ec2b0)
}

var fileDescriptor_aa415f8b859ec2b0 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x18, 0x84, 0x37, 0x3d, 0xb0, 0xe0, 0x85, 0x4b, 0x38, 0xb0, 0xea, 0xa9, 0x2a, 0x07, 0xf6, 0x00,
	0x36, 0xea, 0x9e, 0x91, 0x60, 0xb3, 0x94, 0xe5, 0x80, 0x14, 0x99, 0x45, 0x42, 0xdc, 0x52, 0xe7,
	0x6f, 0xb0, 0xe4, 0xda, 0xc1, 0xff, 0xef, 0xaa, 0x0f, 0xc0, 0x43, 0xf0, 0x0a, 0xf4, 0x29, 0x51,
	0x1b, 0xb7, 0x69, 0xd4, 0xa2, 0x6e, 0xcf, 0xf9, 0x67, 0xe6, 0x9b, 0x89, 0xd9, 0x1b, 0x19, 0x2c,
	0xe9, 0x19, 0x88, 0x8f, 0x73, 0xb0, 0x84, 0x3c, 0xf7, 0x4e, 0x01, 0xa2, 0xb6, 0x95, 0xc8, 0xc3,
	0xc4, 0x68, 0x35, 0xd6, 0x86, 0xc0, 0x23, 0xaf, 0xbd, 0x23, 0x97, 0x0e, 0x4b, 0x67, 0x34, 0x91,
	0x01, 0xee, 0x1b, 0x1d, 0x87, 0x46, 0x57, 0x6f, 0x75, 0xfd, 0xc1, 0x38, 0xd8, 0xb2, 0x98, 0x81,
	0xa5, 0xc2, 0xa0, 0xc8, 0x57, 0xba, 0x49, 0x98, 0x8a, 0x6f, 0x41, 0x97, 0x8d, 0x4b, 0x9f, 0x77,
	0x2e, 0xbe, 0x82, 0x9f, 0x6b, 0x05, 0x28, 0x24, 0xcc, 0xc1, 0x23, 0x64, 0x85, 0x31, 0x99, 0xb3,
	0x04, 0x0b, 0x8a, 0xf7, 0xaf, 0xfe, 0x0f, 0xd9, 0xc1, 0x1b, 0xfe, 0x4d, 0xd8, 0xa0, 0x83, 0x2d,
	0xa1, 0xd2, 0x48, 0xbe, 0x20, 0xed, 0xac, 0x84, 0x5f, 0x01, 0x90, 0xd2, 0x9c, 0x5d, 0xa8, 0x36,
	0xe2, 0x32, 0x19, 0x24, 0x57, 0x17, 0x23, 0xce, 0xb7, 0xcd, 0x30, 0xf2, 0xf0, 0x1d, 0x9e, 0x0f,
	0xbe, 0x0a, 0x2b, 0x5e, 0x8c, 0x2a, 0xb9, 0x6b, 0x91, 0x5e, 0xb3, 0xc7, 0xd3, 0x75, 0xde, 0xe7,
	0xf2, 0xb2, 0xb7, 0xb6, 0x7b, 0xd1, 0xda, 0xd5, 0x71, 0x00, 0xbe, 0x1a, 0x40, 0x6e, 0x0f, 0x87,
	0xbf, 0x7b, 0xec, 0x65, 0x87, 0x35, 0x33, 0x1a, 0x2c, 0xdd, 0xbb, 0x58, 0xf6, 0x0b, 0x20, 0x16,
	0x15, 0xa4, 0x0b, 0xf6, 0xdc, 0xef, 0xb7, 0x88, 0xd8, 0xb7, 0xfc, 0xf8, 0x0f, 0xe1, 0xc7, 0x16,
	0xb9, 0x3b, 0x93, 0x87, 0x22, 0xd2, 0xef, 0xec, 0x69, 0x43, 0x2b, 0x01, 0x83, 0xa1, 0x58, 0x6d,
	0xf4, 0x90, 0xc8, 0xf1, 0x46, 0x57, 0x3b, 0x8b, 0x70, 0x77, 0x26, 0x3b, 0x4e, 0x37, 0x4f, 0xd8,
	0x79, 0xac, 0x37, 0x5a, 0x26, 0xec, 0x59, 0x07, 0x30, 0xfd, 0x93, 0xb0, 0xf3, 0xcc, 0x59, 0x0b,
	0x8a, 0xd2, 0x4f, 0x27, 0xf7, 0x3b, 0xbc, 0x62, 0xff, 0xfd, 0x09, 0xd4, 0xcd, 0xf7, 0x7b, 0xd7,
	0x58, 0x45, 0x87, 0xab, 0xe4, 0x6d, 0x72, 0x63, 0x7f, 0xbc, 0xab, 0x5c, 0xeb, 0xa4, 0x9d, 0x50,
	0x3f, 0x0b, 0x87, 0x04, 0x48, 0x28, 0x94, 0xb3, 0xe4, 0x0b, 0x45, 0x28, 0x62, 0x82, 0x68, 0x12,
	0x44, 0x9b, 0xb0, 0xec, 0xbd, 0xbe, 0xdd, 0x88, 0x63, 0x0c, 0xdf, 0x7b, 0xd3, 0x3c, 0xdb, 0x18,
	0x4d, 0x1e, 0xad, 0x1f, 0xcf, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x71, 0xac, 0x78,
	0xa6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PublicFiltersClient is the client API for PublicFilters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublicFiltersClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (PublicFilters_ConnectClient, error)
}

type publicFiltersClient struct {
	cc *grpc.ClientConn
}

func NewPublicFiltersClient(cc *grpc.ClientConn) PublicFiltersClient {
	return &publicFiltersClient{cc}
}

func (c *publicFiltersClient) Connect(ctx context.Context, opts ...grpc.CallOption) (PublicFilters_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PublicFilters_serviceDesc.Streams[0], "/dolittle.runtime.events.processing.PublicFilters/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &publicFiltersConnectClient{stream}
	return x, nil
}

type PublicFilters_ConnectClient interface {
	Send(*PublicFiltersClientToRuntimeMessage) error
	Recv() (*FilterRuntimeToClientMessage, error)
	grpc.ClientStream
}

type publicFiltersConnectClient struct {
	grpc.ClientStream
}

func (x *publicFiltersConnectClient) Send(m *PublicFiltersClientToRuntimeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *publicFiltersConnectClient) Recv() (*FilterRuntimeToClientMessage, error) {
	m := new(FilterRuntimeToClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublicFiltersServer is the server API for PublicFilters service.
type PublicFiltersServer interface {
	Connect(PublicFilters_ConnectServer) error
}

// UnimplementedPublicFiltersServer can be embedded to have forward compatible implementations.
type UnimplementedPublicFiltersServer struct {
}

func (*UnimplementedPublicFiltersServer) Connect(srv PublicFilters_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterPublicFiltersServer(s *grpc.Server, srv PublicFiltersServer) {
	s.RegisterService(&_PublicFilters_serviceDesc, srv)
}

func _PublicFilters_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PublicFiltersServer).Connect(&publicFiltersConnectServer{stream})
}

type PublicFilters_ConnectServer interface {
	Send(*FilterRuntimeToClientMessage) error
	Recv() (*PublicFiltersClientToRuntimeMessage, error)
	grpc.ServerStream
}

type publicFiltersConnectServer struct {
	grpc.ServerStream
}

func (x *publicFiltersConnectServer) Send(m *FilterRuntimeToClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *publicFiltersConnectServer) Recv() (*PublicFiltersClientToRuntimeMessage, error) {
	m := new(PublicFiltersClientToRuntimeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PublicFilters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.events.processing.PublicFilters",
	HandlerType: (*PublicFiltersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _PublicFilters_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Runtime/Events.Processing/PublicFilters.proto",
}
