// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Fundamentals/Execution/ExecutionContext.proto

package execution

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	protobuf "go.dolittle.io/chaostests/contracts/protobuf"
	security "go.dolittle.io/chaostests/contracts/security"
	versioning "go.dolittle.io/chaostests/contracts/versioning"
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

type ExecutionContext struct {
	MicroserviceId       *protobuf.Uuid      `protobuf:"bytes,1,opt,name=microserviceId,proto3" json:"microserviceId,omitempty"`
	TenantId             *protobuf.Uuid      `protobuf:"bytes,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	Version              *versioning.Version `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	CorrelationId        *protobuf.Uuid      `protobuf:"bytes,4,opt,name=correlationId,proto3" json:"correlationId,omitempty"`
	Claims               []*security.Claim   `protobuf:"bytes,5,rep,name=claims,proto3" json:"claims,omitempty"`
	Environment          string              `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExecutionContext) Reset()         { *m = ExecutionContext{} }
func (m *ExecutionContext) String() string { return proto.CompactTextString(m) }
func (*ExecutionContext) ProtoMessage()    {}
func (*ExecutionContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_442a9ca5768d6252, []int{0}
}

func (m *ExecutionContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionContext.Unmarshal(m, b)
}
func (m *ExecutionContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionContext.Marshal(b, m, deterministic)
}
func (m *ExecutionContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionContext.Merge(m, src)
}
func (m *ExecutionContext) XXX_Size() int {
	return xxx_messageInfo_ExecutionContext.Size(m)
}
func (m *ExecutionContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionContext.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionContext proto.InternalMessageInfo

func (m *ExecutionContext) GetMicroserviceId() *protobuf.Uuid {
	if m != nil {
		return m.MicroserviceId
	}
	return nil
}

func (m *ExecutionContext) GetTenantId() *protobuf.Uuid {
	if m != nil {
		return m.TenantId
	}
	return nil
}

func (m *ExecutionContext) GetVersion() *versioning.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ExecutionContext) GetCorrelationId() *protobuf.Uuid {
	if m != nil {
		return m.CorrelationId
	}
	return nil
}

func (m *ExecutionContext) GetClaims() []*security.Claim {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *ExecutionContext) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecutionContext)(nil), "dolittle.execution.ExecutionContext")
}

func init() {
	proto.RegisterFile("Fundamentals/Execution/ExecutionContext.proto", fileDescriptor_442a9ca5768d6252)
}

var fileDescriptor_442a9ca5768d6252 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x69, 0xfb, 0xbe, 0x55, 0xb7, 0x28, 0xb2, 0x17, 0x97, 0xd2, 0x43, 0x14, 0x84, 0x5e,
	0xba, 0x91, 0x16, 0xbc, 0x89, 0x60, 0x55, 0xe8, 0x45, 0x24, 0xa2, 0x07, 0x6f, 0xdb, 0xcd, 0x58,
	0x17, 0xd2, 0x9d, 0xb2, 0x3b, 0x09, 0xf5, 0x2b, 0xf9, 0x75, 0xfc, 0x42, 0x92, 0x36, 0x7f, 0x4c,
	0x0f, 0xbd, 0x2d, 0xe1, 0xf7, 0x4c, 0x7e, 0xf3, 0x0c, 0x1b, 0x3d, 0xa6, 0x36, 0x56, 0x4b, 0xb0,
	0xa4, 0x12, 0x1f, 0x3e, 0xac, 0x41, 0xa7, 0x64, 0xd0, 0xd6, 0xaf, 0x29, 0x5a, 0x82, 0x35, 0xc9,
	0x95, 0x43, 0x42, 0xce, 0x63, 0x4c, 0x0c, 0x51, 0x02, 0x12, 0x4a, 0xa0, 0x1f, 0x34, 0x46, 0x3c,
	0xe7, 0xdc, 0x3c, 0xfd, 0x08, 0x5f, 0x53, 0x13, 0x6f, 0x53, 0xfd, 0xf3, 0x06, 0xf1, 0x02, 0x3a,
	0x75, 0x86, 0xbe, 0xc2, 0x69, 0xa2, 0xcc, 0xb2, 0x40, 0x2e, 0x1b, 0xc8, 0x1b, 0x38, 0x6f, 0xd0,
	0x1a, 0xbb, 0x28, 0x9f, 0x5b, 0xec, 0xe2, 0xa7, 0xcd, 0x4e, 0x77, 0xd5, 0xf8, 0x2d, 0x3b, 0x59,
	0x1a, 0xed, 0xd0, 0x83, 0xcb, 0x8c, 0x86, 0x59, 0x2c, 0x5a, 0x41, 0x6b, 0xd8, 0x1b, 0x9f, 0xc9,
	0xca, 0x76, 0x55, 0x58, 0xc9, 0xdc, 0x2a, 0xda, 0xc1, 0xf9, 0x84, 0x1d, 0x12, 0x58, 0x65, 0x69,
	0x16, 0x8b, 0xf6, 0xfe, 0x68, 0x05, 0xf2, 0x6b, 0x76, 0x90, 0x6d, 0xdd, 0x44, 0x67, 0x93, 0x19,
	0xd4, 0x99, 0xac, 0xf2, 0x97, 0x85, 0x7f, 0x54, 0xc2, 0xfc, 0x86, 0x1d, 0x6b, 0x74, 0x0e, 0x12,
	0x95, 0xef, 0x30, 0x8b, 0xc5, 0xbf, 0xfd, 0x7f, 0x6c, 0xd2, 0xfc, 0x8a, 0x75, 0x75, 0xde, 0x9b,
	0x17, 0xff, 0x83, 0xce, 0xb0, 0x37, 0x16, 0x75, 0xce, 0x17, 0xc5, 0xca, 0x4d, 0xb1, 0x51, 0xc1,
	0xf1, 0x80, 0xf5, 0xc0, 0x66, 0xc6, 0xa1, 0xcd, 0xeb, 0x15, 0xdd, 0xa0, 0x35, 0x3c, 0x8a, 0xfe,
	0x7e, 0xba, 0x7b, 0x7a, 0x1f, 0x2d, 0xb0, 0x9e, 0x63, 0x30, 0xd4, 0x9f, 0x0a, 0x3d, 0x81, 0x27,
	0x1f, 0x6a, 0xb4, 0xe4, 0x94, 0x26, 0x1f, 0x56, 0x27, 0xff, 0x6e, 0x0f, 0xee, 0x4b, 0xb8, 0xba,
	0x86, 0x9c, 0x96, 0xe0, 0xbc, 0xbb, 0xd9, 0x60, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x48,
	0x66, 0xe2, 0x5d, 0x02, 0x00, 0x00,
}
