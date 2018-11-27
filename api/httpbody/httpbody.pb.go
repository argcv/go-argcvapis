// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/api/httpbody.proto

package httpbody

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HttpBody struct {
	ContentType          string   `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpBody) Reset()         { *m = HttpBody{} }
func (m *HttpBody) String() string { return proto.CompactTextString(m) }
func (*HttpBody) ProtoMessage()    {}
func (*HttpBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_94b7c0231911c425, []int{0}
}

func (m *HttpBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpBody.Unmarshal(m, b)
}
func (m *HttpBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpBody.Marshal(b, m, deterministic)
}
func (m *HttpBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBody.Merge(m, src)
}
func (m *HttpBody) XXX_Size() int {
	return xxx_messageInfo_HttpBody.Size(m)
}
func (m *HttpBody) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBody.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBody proto.InternalMessageInfo

func (m *HttpBody) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HttpBody) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpBody)(nil), "argcv.api.HttpBody")
}

func init() { proto.RegisterFile("argcv/proto/api/httpbody.proto", fileDescriptor_94b7c0231911c425) }

var fileDescriptor_94b7c0231911c425 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2c, 0x4a, 0x4f,
	0x2e, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x28, 0x29, 0x29,
	0x48, 0xca, 0x4f, 0xa9, 0xd4, 0x03, 0x0b, 0x09, 0x71, 0x82, 0xe5, 0xf5, 0x12, 0x0b, 0x32, 0x95,
	0x1c, 0xb9, 0x38, 0x3c, 0x4a, 0x4a, 0x0a, 0x9c, 0xf2, 0x53, 0x2a, 0x85, 0x14, 0xb9, 0x78, 0x92,
	0xf3, 0xf3, 0x4a, 0x52, 0xf3, 0x4a, 0xe2, 0x4b, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xb8, 0xa1, 0x62, 0x21, 0x95, 0x05, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25,
	0x89, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x53, 0x1c, 0x97, 0x70, 0x72, 0x7e,
	0xae, 0x1e, 0xc4, 0x4c, 0xb0, 0x05, 0x20, 0x93, 0x9d, 0x78, 0x61, 0xe6, 0x06, 0x80, 0x84, 0x02,
	0x18, 0xa3, 0x8c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x21, 0x0e,
	0x4c, 0xcf, 0xd7, 0x05, 0x33, 0x12, 0x0b, 0x32, 0x8b, 0x51, 0xdc, 0x69, 0x0d, 0x63, 0xfc, 0x60,
	0x64, 0x4c, 0x62, 0x03, 0x9b, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x8b, 0x94, 0x9c,
	0xd6, 0x00, 0x00, 0x00,
}
