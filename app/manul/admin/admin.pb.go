// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/app/manul/admin.proto

package admin // import "github.com/argcv/go-argcvapis/app/manul/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A simple struct to define authorization info
type Auth struct {
	// Required: project id
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Required: user name
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Required: secret
	Secret               string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_admin_7ab2574a208b2a0b, []int{0}
}
func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (dst *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(dst, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Auth) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Auth) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func init() {
	proto.RegisterType((*Auth)(nil), "argcv.app.manul.Auth")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
}

type adminServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdminServiceClient(cc *grpc.ClientConn) AdminServiceClient {
	return &adminServiceClient{cc}
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "argcv.app.manul.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "argcv/proto/app/manul/admin.proto",
}

func init() {
	proto.RegisterFile("argcv/proto/app/manul/admin.proto", fileDescriptor_admin_7ab2574a208b2a0b)
}

var fileDescriptor_admin_7ab2574a208b2a0b = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0x4a, 0x4f,
	0x2e, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0x28, 0xd0, 0xcf, 0x4d, 0xcc, 0x2b,
	0xcd, 0xd1, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x03, 0x8b, 0x0a, 0xf1, 0x83, 0x95, 0xe8, 0x25,
	0x16, 0x14, 0xe8, 0x81, 0x25, 0x95, 0x7c, 0xb8, 0x58, 0x1c, 0x4b, 0x4b, 0x32, 0x84, 0x24, 0xb8,
	0xd8, 0x0b, 0x8a, 0xf2, 0xb3, 0x52, 0x93, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60,
	0x5c, 0x21, 0x21, 0x2e, 0x96, 0xd2, 0xe2, 0xd4, 0x22, 0x09, 0x26, 0xb0, 0x30, 0x98, 0x2d, 0x24,
	0xc6, 0xc5, 0x56, 0x9c, 0x9a, 0x5c, 0x94, 0x5a, 0x22, 0xc1, 0x0c, 0x16, 0x85, 0xf2, 0x8c, 0xf8,
	0xb8, 0x78, 0x1c, 0x41, 0xb6, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x3a, 0x25, 0x72, 0x49,
	0x26, 0xe7, 0xe7, 0xea, 0x41, 0x2c, 0x05, 0xbb, 0x00, 0x61, 0xb5, 0x13, 0x17, 0x58, 0x69, 0x00,
	0x48, 0x34, 0x80, 0x31, 0xca, 0x38, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0xe2, 0x8f, 0xf4, 0x7c, 0x5d, 0x30, 0x23, 0xb1, 0x20, 0xb3, 0x18, 0xdd, 0x3b, 0xd6, 0x60,
	0xf2, 0x07, 0x23, 0x63, 0x12, 0x1b, 0xd8, 0x58, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86,
	0xc2, 0x66, 0xd9, 0xfd, 0x00, 0x00, 0x00,
}