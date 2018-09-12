// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/app/manul/file.proto

package file // import "github.com/argcv/go-argcvapis/app/manul/file"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Generatic File
type File struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string          `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Size                 uint64          `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Meta                 *_struct.Struct `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 []byte          `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_532d6b4e8d3fb72b, []int{0}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *File) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetMeta() *_struct.Struct {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Files struct {
	Meta                 *_struct.Struct `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data                 []*File         `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Files) Reset()         { *m = Files{} }
func (m *Files) String() string { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()    {}
func (*Files) Descriptor() ([]byte, []int) {
	return fileDescriptor_532d6b4e8d3fb72b, []int{1}
}
func (m *Files) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Files.Unmarshal(m, b)
}
func (m *Files) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Files.Marshal(b, m, deterministic)
}
func (m *Files) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Files.Merge(m, src)
}
func (m *Files) XXX_Size() int {
	return xxx_messageInfo_Files.Size(m)
}
func (m *Files) XXX_DiscardUnknown() {
	xxx_messageInfo_Files.DiscardUnknown(m)
}

var xxx_messageInfo_Files proto.InternalMessageInfo

func (m *Files) GetMeta() *_struct.Struct {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Files) GetData() []*File {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "argcv.app.manul.File")
	proto.RegisterType((*Files)(nil), "argcv.app.manul.Files")
}

func init() { proto.RegisterFile("argcv/proto/app/manul/file.proto", fileDescriptor_532d6b4e8d3fb72b) }

var fileDescriptor_532d6b4e8d3fb72b = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4b, 0x33, 0x31,
	0x10, 0xc5, 0x99, 0x76, 0xfb, 0x41, 0xd3, 0x0f, 0x84, 0x05, 0x71, 0x15, 0x0f, 0xa1, 0xa7, 0x88,
	0x98, 0x60, 0x3d, 0x7a, 0xeb, 0xc1, 0x73, 0x59, 0x6f, 0x1e, 0x94, 0xd9, 0x35, 0x4d, 0x03, 0x9b,
	0x26, 0xec, 0x66, 0x3d, 0x78, 0xf5, 0x9f, 0xf6, 0x28, 0x99, 0x54, 0x85, 0x9e, 0xbc, 0x84, 0xc7,
	0xcb, 0xcb, 0xef, 0x65, 0x86, 0x71, 0xec, 0x4d, 0xfb, 0xa6, 0x42, 0xef, 0xa3, 0x57, 0x18, 0x82,
	0x72, 0xb8, 0x1f, 0x3b, 0xb5, 0xb5, 0x9d, 0x96, 0x64, 0x96, 0x27, 0x94, 0x90, 0x18, 0x82, 0xa4,
	0xbb, 0x8b, 0x4b, 0xe3, 0xbd, 0xe9, 0x74, 0x7e, 0xd3, 0x8c, 0x5b, 0x35, 0xc4, 0x7e, 0x6c, 0x63,
	0x8e, 0x2f, 0x3f, 0x80, 0x15, 0x0f, 0xb6, 0xd3, 0x65, 0xc9, 0x8a, 0x3d, 0x3a, 0x5d, 0x01, 0x07,
	0x31, 0xaf, 0x49, 0x27, 0x2f, 0x60, 0xdc, 0x55, 0x93, 0xec, 0x25, 0x9d, 0xbc, 0xc1, 0xbe, 0xeb,
	0x6a, 0xca, 0x41, 0x14, 0x35, 0xe9, 0xf2, 0x9a, 0x15, 0x4e, 0x47, 0xac, 0x0a, 0x0e, 0x62, 0xb1,
	0x3a, 0x93, 0xb9, 0x51, 0x7e, 0x37, 0xca, 0x47, 0x6a, 0xac, 0x29, 0x94, 0x00, 0xaf, 0x18, 0xb1,
	0x9a, 0x71, 0x10, 0xff, 0x6b, 0xd2, 0xcb, 0x17, 0x36, 0x4b, 0x9f, 0x18, 0x7e, 0x48, 0xf0, 0x17,
	0xd2, 0xd5, 0x81, 0x34, 0xe1, 0x53, 0xb1, 0x58, 0x9d, 0xca, 0xa3, 0xc9, 0x65, 0x42, 0xe6, 0x82,
	0xf5, 0x33, 0x3b, 0x6f, 0xbd, 0x3b, 0x24, 0x88, 0xf6, 0x9b, 0x5b, 0xcf, 0x53, 0x70, 0x93, 0xcc,
	0x0d, 0x3c, 0xdd, 0x1a, 0x1b, 0x77, 0x63, 0x23, 0x5b, 0xef, 0x54, 0x5e, 0xb6, 0xf1, 0x37, 0x24,
	0x30, 0xd8, 0xe1, 0x68, 0xe7, 0xf7, 0xe9, 0xf8, 0x04, 0x68, 0xfe, 0x11, 0xf3, 0xee, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0xe6, 0x0e, 0xdf, 0xa0, 0x01, 0x00, 0x00,
}
