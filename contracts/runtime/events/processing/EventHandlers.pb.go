// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Runtime/Events.Processing/EventHandlers.proto

package processing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	artifacts "go.dolittle.io/chaostests/contracts/artifacts"
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

type EventHandlersRegistrationRequest struct {
	CallContext          *services.ReverseCallArgumentsContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	ScopeId              *protobuf.Uuid                        `protobuf:"bytes,2,opt,name=scopeId,proto3" json:"scopeId,omitempty"`
	EventHandlerId       *protobuf.Uuid                        `protobuf:"bytes,3,opt,name=eventHandlerId,proto3" json:"eventHandlerId,omitempty"`
	Types                []*artifacts.Artifact                 `protobuf:"bytes,4,rep,name=types,proto3" json:"types,omitempty"`
	Partitioned          bool                                  `protobuf:"varint,5,opt,name=partitioned,proto3" json:"partitioned,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *EventHandlersRegistrationRequest) Reset()         { *m = EventHandlersRegistrationRequest{} }
func (m *EventHandlersRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*EventHandlersRegistrationRequest) ProtoMessage()    {}
func (*EventHandlersRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b38c310015445, []int{0}
}

func (m *EventHandlersRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandlersRegistrationRequest.Unmarshal(m, b)
}
func (m *EventHandlersRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandlersRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *EventHandlersRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandlersRegistrationRequest.Merge(m, src)
}
func (m *EventHandlersRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_EventHandlersRegistrationRequest.Size(m)
}
func (m *EventHandlersRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandlersRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandlersRegistrationRequest proto.InternalMessageInfo

func (m *EventHandlersRegistrationRequest) GetCallContext() *services.ReverseCallArgumentsContext {
	if m != nil {
		return m.CallContext
	}
	return nil
}

func (m *EventHandlersRegistrationRequest) GetScopeId() *protobuf.Uuid {
	if m != nil {
		return m.ScopeId
	}
	return nil
}

func (m *EventHandlersRegistrationRequest) GetEventHandlerId() *protobuf.Uuid {
	if m != nil {
		return m.EventHandlerId
	}
	return nil
}

func (m *EventHandlersRegistrationRequest) GetTypes() []*artifacts.Artifact {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *EventHandlersRegistrationRequest) GetPartitioned() bool {
	if m != nil {
		return m.Partitioned
	}
	return false
}

type EventHandlerResponse struct {
	CallContext          *services.ReverseCallResponseContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Failure              *ProcessorFailure                    `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *EventHandlerResponse) Reset()         { *m = EventHandlerResponse{} }
func (m *EventHandlerResponse) String() string { return proto.CompactTextString(m) }
func (*EventHandlerResponse) ProtoMessage()    {}
func (*EventHandlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b38c310015445, []int{1}
}

func (m *EventHandlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandlerResponse.Unmarshal(m, b)
}
func (m *EventHandlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandlerResponse.Marshal(b, m, deterministic)
}
func (m *EventHandlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandlerResponse.Merge(m, src)
}
func (m *EventHandlerResponse) XXX_Size() int {
	return xxx_messageInfo_EventHandlerResponse.Size(m)
}
func (m *EventHandlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandlerResponse proto.InternalMessageInfo

func (m *EventHandlerResponse) GetCallContext() *services.ReverseCallResponseContext {
	if m != nil {
		return m.CallContext
	}
	return nil
}

func (m *EventHandlerResponse) GetFailure() *ProcessorFailure {
	if m != nil {
		return m.Failure
	}
	return nil
}

type EventHandlersClientToRuntimeMessage struct {
	// Types that are valid to be assigned to Message:
	//	*EventHandlersClientToRuntimeMessage_RegistrationRequest
	//	*EventHandlersClientToRuntimeMessage_HandleResult
	Message              isEventHandlersClientToRuntimeMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *EventHandlersClientToRuntimeMessage) Reset()         { *m = EventHandlersClientToRuntimeMessage{} }
func (m *EventHandlersClientToRuntimeMessage) String() string { return proto.CompactTextString(m) }
func (*EventHandlersClientToRuntimeMessage) ProtoMessage()    {}
func (*EventHandlersClientToRuntimeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b38c310015445, []int{2}
}

func (m *EventHandlersClientToRuntimeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandlersClientToRuntimeMessage.Unmarshal(m, b)
}
func (m *EventHandlersClientToRuntimeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandlersClientToRuntimeMessage.Marshal(b, m, deterministic)
}
func (m *EventHandlersClientToRuntimeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandlersClientToRuntimeMessage.Merge(m, src)
}
func (m *EventHandlersClientToRuntimeMessage) XXX_Size() int {
	return xxx_messageInfo_EventHandlersClientToRuntimeMessage.Size(m)
}
func (m *EventHandlersClientToRuntimeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandlersClientToRuntimeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandlersClientToRuntimeMessage proto.InternalMessageInfo

type isEventHandlersClientToRuntimeMessage_Message interface {
	isEventHandlersClientToRuntimeMessage_Message()
}

type EventHandlersClientToRuntimeMessage_RegistrationRequest struct {
	RegistrationRequest *EventHandlersRegistrationRequest `protobuf:"bytes,1,opt,name=registrationRequest,proto3,oneof"`
}

type EventHandlersClientToRuntimeMessage_HandleResult struct {
	HandleResult *EventHandlerResponse `protobuf:"bytes,2,opt,name=handleResult,proto3,oneof"`
}

func (*EventHandlersClientToRuntimeMessage_RegistrationRequest) isEventHandlersClientToRuntimeMessage_Message() {
}

func (*EventHandlersClientToRuntimeMessage_HandleResult) isEventHandlersClientToRuntimeMessage_Message() {
}

func (m *EventHandlersClientToRuntimeMessage) GetMessage() isEventHandlersClientToRuntimeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *EventHandlersClientToRuntimeMessage) GetRegistrationRequest() *EventHandlersRegistrationRequest {
	if x, ok := m.GetMessage().(*EventHandlersClientToRuntimeMessage_RegistrationRequest); ok {
		return x.RegistrationRequest
	}
	return nil
}

func (m *EventHandlersClientToRuntimeMessage) GetHandleResult() *EventHandlerResponse {
	if x, ok := m.GetMessage().(*EventHandlersClientToRuntimeMessage_HandleResult); ok {
		return x.HandleResult
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventHandlersClientToRuntimeMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventHandlersClientToRuntimeMessage_RegistrationRequest)(nil),
		(*EventHandlersClientToRuntimeMessage_HandleResult)(nil),
	}
}

type EventHandlerRegistrationResponse struct {
	Failure              *protobuf.Failure `protobuf:"bytes,1,opt,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EventHandlerRegistrationResponse) Reset()         { *m = EventHandlerRegistrationResponse{} }
func (m *EventHandlerRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*EventHandlerRegistrationResponse) ProtoMessage()    {}
func (*EventHandlerRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b38c310015445, []int{3}
}

func (m *EventHandlerRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandlerRegistrationResponse.Unmarshal(m, b)
}
func (m *EventHandlerRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandlerRegistrationResponse.Marshal(b, m, deterministic)
}
func (m *EventHandlerRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandlerRegistrationResponse.Merge(m, src)
}
func (m *EventHandlerRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_EventHandlerRegistrationResponse.Size(m)
}
func (m *EventHandlerRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandlerRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandlerRegistrationResponse proto.InternalMessageInfo

func (m *EventHandlerRegistrationResponse) GetFailure() *protobuf.Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

type HandleEventRequest struct {
	CallContext          *services.ReverseCallRequestContext `protobuf:"bytes,1,opt,name=callContext,proto3" json:"callContext,omitempty"`
	Event                *StreamEvent                        `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	RetryProcessingState *RetryProcessingState               `protobuf:"bytes,3,opt,name=retryProcessingState,proto3" json:"retryProcessingState,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *HandleEventRequest) Reset()         { *m = HandleEventRequest{} }
func (m *HandleEventRequest) String() string { return proto.CompactTextString(m) }
func (*HandleEventRequest) ProtoMessage()    {}
func (*HandleEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b38c310015445, []int{4}
}

func (m *HandleEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleEventRequest.Unmarshal(m, b)
}
func (m *HandleEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleEventRequest.Marshal(b, m, deterministic)
}
func (m *HandleEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleEventRequest.Merge(m, src)
}
func (m *HandleEventRequest) XXX_Size() int {
	return xxx_messageInfo_HandleEventRequest.Size(m)
}
func (m *HandleEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleEventRequest proto.InternalMessageInfo

func (m *HandleEventRequest) GetCallContext() *services.ReverseCallRequestContext {
	if m != nil {
		return m.CallContext
	}
	return nil
}

func (m *HandleEventRequest) GetEvent() *StreamEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *HandleEventRequest) GetRetryProcessingState() *RetryProcessingState {
	if m != nil {
		return m.RetryProcessingState
	}
	return nil
}

type EventHandlerRuntimeToClientMessage struct {
	// Types that are valid to be assigned to Message:
	//	*EventHandlerRuntimeToClientMessage_RegistrationResponse
	//	*EventHandlerRuntimeToClientMessage_HandleRequest
	Message              isEventHandlerRuntimeToClientMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *EventHandlerRuntimeToClientMessage) Reset()         { *m = EventHandlerRuntimeToClientMessage{} }
func (m *EventHandlerRuntimeToClientMessage) String() string { return proto.CompactTextString(m) }
func (*EventHandlerRuntimeToClientMessage) ProtoMessage()    {}
func (*EventHandlerRuntimeToClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c0b38c310015445, []int{5}
}

func (m *EventHandlerRuntimeToClientMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandlerRuntimeToClientMessage.Unmarshal(m, b)
}
func (m *EventHandlerRuntimeToClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandlerRuntimeToClientMessage.Marshal(b, m, deterministic)
}
func (m *EventHandlerRuntimeToClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandlerRuntimeToClientMessage.Merge(m, src)
}
func (m *EventHandlerRuntimeToClientMessage) XXX_Size() int {
	return xxx_messageInfo_EventHandlerRuntimeToClientMessage.Size(m)
}
func (m *EventHandlerRuntimeToClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandlerRuntimeToClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandlerRuntimeToClientMessage proto.InternalMessageInfo

type isEventHandlerRuntimeToClientMessage_Message interface {
	isEventHandlerRuntimeToClientMessage_Message()
}

type EventHandlerRuntimeToClientMessage_RegistrationResponse struct {
	RegistrationResponse *EventHandlerRegistrationResponse `protobuf:"bytes,1,opt,name=registrationResponse,proto3,oneof"`
}

type EventHandlerRuntimeToClientMessage_HandleRequest struct {
	HandleRequest *HandleEventRequest `protobuf:"bytes,2,opt,name=handleRequest,proto3,oneof"`
}

func (*EventHandlerRuntimeToClientMessage_RegistrationResponse) isEventHandlerRuntimeToClientMessage_Message() {
}

func (*EventHandlerRuntimeToClientMessage_HandleRequest) isEventHandlerRuntimeToClientMessage_Message() {
}

func (m *EventHandlerRuntimeToClientMessage) GetMessage() isEventHandlerRuntimeToClientMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *EventHandlerRuntimeToClientMessage) GetRegistrationResponse() *EventHandlerRegistrationResponse {
	if x, ok := m.GetMessage().(*EventHandlerRuntimeToClientMessage_RegistrationResponse); ok {
		return x.RegistrationResponse
	}
	return nil
}

func (m *EventHandlerRuntimeToClientMessage) GetHandleRequest() *HandleEventRequest {
	if x, ok := m.GetMessage().(*EventHandlerRuntimeToClientMessage_HandleRequest); ok {
		return x.HandleRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventHandlerRuntimeToClientMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventHandlerRuntimeToClientMessage_RegistrationResponse)(nil),
		(*EventHandlerRuntimeToClientMessage_HandleRequest)(nil),
	}
}

func init() {
	proto.RegisterType((*EventHandlersRegistrationRequest)(nil), "dolittle.runtime.events.processing.EventHandlersRegistrationRequest")
	proto.RegisterType((*EventHandlerResponse)(nil), "dolittle.runtime.events.processing.EventHandlerResponse")
	proto.RegisterType((*EventHandlersClientToRuntimeMessage)(nil), "dolittle.runtime.events.processing.EventHandlersClientToRuntimeMessage")
	proto.RegisterType((*EventHandlerRegistrationResponse)(nil), "dolittle.runtime.events.processing.EventHandlerRegistrationResponse")
	proto.RegisterType((*HandleEventRequest)(nil), "dolittle.runtime.events.processing.HandleEventRequest")
	proto.RegisterType((*EventHandlerRuntimeToClientMessage)(nil), "dolittle.runtime.events.processing.EventHandlerRuntimeToClientMessage")
}

func init() {
	proto.RegisterFile("Runtime/Events.Processing/EventHandlers.proto", fileDescriptor_5c0b38c310015445)
}

var fileDescriptor_5c0b38c310015445 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0xde, 0x2e, 0xe2, 0xea, 0xac, 0x78, 0x18, 0x49, 0xdc, 0x6c, 0x3c, 0x6c, 0x4a, 0x4c, 0x36,
	0x0a, 0x53, 0x5d, 0x89, 0xf1, 0x62, 0x0c, 0x2c, 0xe0, 0x72, 0x10, 0xc9, 0x80, 0x89, 0xf1, 0x40,
	0x52, 0xda, 0xc7, 0xd2, 0xa4, 0xcc, 0xd4, 0x99, 0x29, 0x01, 0x7f, 0x82, 0xfe, 0x02, 0x2f, 0xfe,
	0x00, 0x4f, 0xde, 0xbc, 0xf9, 0xdb, 0xcc, 0x76, 0x66, 0x68, 0xeb, 0x0e, 0xb0, 0x78, 0x6b, 0xa6,
	0xef, 0xfb, 0xde, 0x7b, 0xdf, 0xf7, 0x4d, 0x8b, 0x56, 0x68, 0xce, 0x54, 0x72, 0x02, 0xc1, 0xe6,
	0x29, 0x30, 0x25, 0xc9, 0xae, 0xe0, 0x11, 0x48, 0x99, 0xb0, 0xb1, 0x3e, 0x19, 0x85, 0x2c, 0x4e,
	0x41, 0x48, 0x92, 0x09, 0xae, 0x38, 0xf6, 0x63, 0x9e, 0x26, 0x4a, 0xa5, 0x40, 0x84, 0xc6, 0x11,
	0xd0, 0xb8, 0xec, 0x02, 0xd7, 0x7d, 0xbc, 0x95, 0xb3, 0x38, 0x3c, 0x01, 0xa6, 0xc2, 0x54, 0x06,
	0x6b, 0x42, 0x25, 0x47, 0x61, 0xa4, 0xca, 0x27, 0x4d, 0xd5, 0xed, 0xd5, 0xca, 0x76, 0x27, 0x67,
	0x87, 0xf9, 0x51, 0xf0, 0x21, 0x4f, 0x62, 0x53, 0xb1, 0xe4, 0xae, 0xd8, 0x0a, 0x93, 0x34, 0x17,
	0x60, 0x8a, 0x48, 0xad, 0x68, 0x0f, 0xc4, 0x69, 0x12, 0x81, 0x0c, 0x28, 0x9c, 0x82, 0x90, 0x30,
	0x0c, 0xd3, 0x74, 0xc8, 0x99, 0x82, 0x33, 0xdb, 0xf6, 0xe9, 0xe5, 0x0b, 0xef, 0x29, 0x01, 0xe1,
	0x49, 0x71, 0x6e, 0x8a, 0x9f, 0x5c, 0x5e, 0x6c, 0x1e, 0xb9, 0x95, 0xc6, 0xff, 0xd3, 0x44, 0xbd,
	0x9a, 0x64, 0x14, 0xc6, 0x89, 0x54, 0x22, 0x54, 0x09, 0x67, 0x14, 0x3e, 0xe7, 0x20, 0x15, 0xde,
	0x45, 0xed, 0xa8, 0x1c, 0xa9, 0xe3, 0xf5, 0xbc, 0x7e, 0x7b, 0x40, 0xc8, 0x85, 0xaa, 0xd2, 0xcc,
	0x4f, 0x2a, 0xf3, 0xaf, 0x89, 0x71, 0x3e, 0xd9, 0x4f, 0x1a, 0x14, 0xad, 0x52, 0xe0, 0xe7, 0xa8,
	0x25, 0x23, 0x9e, 0xc1, 0x76, 0xdc, 0x69, 0x16, 0x6c, 0x0f, 0x4b, 0xb6, 0xcc, 0x48, 0x46, 0x26,
	0xa2, 0x52, 0x5b, 0x87, 0xdf, 0xa0, 0xfb, 0x50, 0x19, 0x74, 0x3b, 0xee, 0xcc, 0x5d, 0x8d, 0xfc,
	0xa7, 0x1c, 0x0f, 0xd0, 0xbc, 0x3a, 0xcf, 0x40, 0x76, 0x6e, 0xf5, 0xe6, 0xfa, 0xed, 0xc1, 0xa3,
	0x12, 0x17, 0x5a, 0xb7, 0x89, 0x75, 0x9b, 0xea, 0x52, 0xdc, 0x43, 0xed, 0x6c, 0xf2, 0x76, 0xa2,
	0x06, 0xc4, 0x9d, 0xf9, 0x9e, 0xd7, 0xbf, 0x43, 0xab, 0x47, 0xfe, 0x6f, 0x0f, 0x2d, 0x56, 0x05,
	0xa4, 0x20, 0x33, 0xce, 0x24, 0xe0, 0xf7, 0x2e, 0xd1, 0x56, 0xae, 0x16, 0xcd, 0x82, 0x9d, 0x9a,
	0xed, 0xa0, 0xd6, 0x91, 0x0e, 0x91, 0xd1, 0x6c, 0x95, 0x5c, 0x9f, 0x6b, 0x72, 0xe1, 0xb8, 0x09,
	0x20, 0xb5, 0x24, 0xfe, 0xd7, 0x26, 0x5a, 0xaa, 0x59, 0x3f, 0x4c, 0x13, 0x60, 0x6a, 0x9f, 0x9b,
	0xf8, 0xbc, 0x03, 0x29, 0xc3, 0x31, 0xe0, 0x33, 0xf4, 0x40, 0x4c, 0x87, 0xc2, 0x2c, 0xb4, 0x31,
	0xcb, 0x0c, 0xd7, 0x05, 0x6c, 0xd4, 0xa0, 0xae, 0x16, 0xf8, 0x00, 0xdd, 0x3b, 0x2e, 0x50, 0x14,
	0x64, 0x9e, 0x2a, 0xb3, 0xf6, 0xab, 0x9b, 0xb6, 0xb4, 0xaa, 0x8e, 0x1a, 0xb4, 0xc6, 0xb7, 0x7e,
	0x17, 0xb5, 0xcc, 0x92, 0xfe, 0xc7, 0xfa, 0x35, 0xa8, 0x0f, 0x69, 0x1c, 0x5d, 0x2d, 0x0d, 0xd0,
	0xcb, 0x77, 0x1d, 0xd1, 0x9b, 0x92, 0xf9, 0x7b, 0x13, 0x61, 0xcd, 0x5a, 0x34, 0xb0, 0xbb, 0xed,
	0xb8, 0xe2, 0xb1, 0x7c, 0x5d, 0x3c, 0x0a, 0xac, 0x33, 0x1d, 0x9b, 0x68, 0xbe, 0x50, 0xc1, 0x88,
	0x14, 0xcc, 0x22, 0x52, 0xe5, 0xd3, 0x41, 0x35, 0x1a, 0xa7, 0x68, 0x51, 0x80, 0x12, 0xe7, 0xe5,
	0x37, 0x63, 0x4f, 0x85, 0x0a, 0xcc, 0x5d, 0x9b, 0x49, 0x7a, 0xea, 0xc0, 0x53, 0x27, 0xab, 0xff,
	0xad, 0x89, 0xfc, 0x9a, 0xec, 0x9a, 0x74, 0x9f, 0xeb, 0x28, 0xda, 0x04, 0x7e, 0x99, 0x0c, 0x35,
	0x6d, 0xc8, 0xff, 0x46, 0xd0, 0x65, 0xee, 0xa8, 0x41, 0x9d, 0x3d, 0xf0, 0x01, 0x5a, 0xb0, 0x99,
	0xd1, 0xb9, 0xd7, 0xfa, 0xbe, 0x9c, 0xa5, 0xe9, 0xb4, 0xed, 0xa3, 0x06, 0xad, 0xd3, 0x55, 0x32,
	0x38, 0xf8, 0xe5, 0xa1, 0x85, 0xda, 0x55, 0xc1, 0x3f, 0x3c, 0xd4, 0x1a, 0x72, 0xc6, 0x20, 0x52,
	0xf8, 0xed, 0x8d, 0x6f, 0x9a, 0xfb, 0x3e, 0x77, 0xb7, 0x6e, 0xac, 0x97, 0xd3, 0x95, 0xbe, 0xf7,
	0xcc, 0x5b, 0x67, 0x9f, 0x5e, 0x8f, 0x79, 0xc9, 0x97, 0xf0, 0x20, 0x3a, 0x0e, 0xb9, 0x54, 0x20,
	0x95, 0x0c, 0x22, 0xce, 0x94, 0x28, 0x7e, 0xa2, 0xa6, 0x4f, 0xa0, 0xfb, 0x04, 0x65, 0x9f, 0x9f,
	0xcd, 0xe5, 0x0d, 0x0b, 0x36, 0x6d, 0xc8, 0xd4, 0x5f, 0x8b, 0x0c, 0x2d, 0xd1, 0xe1, 0xed, 0xe2,
	0x9e, 0xbd, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xff, 0x32, 0xce, 0x01, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventHandlersClient is the client API for EventHandlers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventHandlersClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (EventHandlers_ConnectClient, error)
}

type eventHandlersClient struct {
	cc *grpc.ClientConn
}

func NewEventHandlersClient(cc *grpc.ClientConn) EventHandlersClient {
	return &eventHandlersClient{cc}
}

func (c *eventHandlersClient) Connect(ctx context.Context, opts ...grpc.CallOption) (EventHandlers_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventHandlers_serviceDesc.Streams[0], "/dolittle.runtime.events.processing.EventHandlers/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventHandlersConnectClient{stream}
	return x, nil
}

type EventHandlers_ConnectClient interface {
	Send(*EventHandlersClientToRuntimeMessage) error
	Recv() (*EventHandlerRuntimeToClientMessage, error)
	grpc.ClientStream
}

type eventHandlersConnectClient struct {
	grpc.ClientStream
}

func (x *eventHandlersConnectClient) Send(m *EventHandlersClientToRuntimeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventHandlersConnectClient) Recv() (*EventHandlerRuntimeToClientMessage, error) {
	m := new(EventHandlerRuntimeToClientMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventHandlersServer is the server API for EventHandlers service.
type EventHandlersServer interface {
	Connect(EventHandlers_ConnectServer) error
}

// UnimplementedEventHandlersServer can be embedded to have forward compatible implementations.
type UnimplementedEventHandlersServer struct {
}

func (*UnimplementedEventHandlersServer) Connect(srv EventHandlers_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterEventHandlersServer(s *grpc.Server, srv EventHandlersServer) {
	s.RegisterService(&_EventHandlers_serviceDesc, srv)
}

func _EventHandlers_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventHandlersServer).Connect(&eventHandlersConnectServer{stream})
}

type EventHandlers_ConnectServer interface {
	Send(*EventHandlerRuntimeToClientMessage) error
	Recv() (*EventHandlersClientToRuntimeMessage, error)
	grpc.ServerStream
}

type eventHandlersConnectServer struct {
	grpc.ServerStream
}

func (x *eventHandlersConnectServer) Send(m *EventHandlerRuntimeToClientMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventHandlersConnectServer) Recv() (*EventHandlersClientToRuntimeMessage, error) {
	m := new(EventHandlersClientToRuntimeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventHandlers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dolittle.runtime.events.processing.EventHandlers",
	HandlerType: (*EventHandlersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _EventHandlers_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Runtime/Events.Processing/EventHandlers.proto",
}