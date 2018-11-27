// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/api/label.proto

package label

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

type LabelDescriptor_ValueType int32

const (
	LabelDescriptor_STRING LabelDescriptor_ValueType = 0
	LabelDescriptor_BOOL   LabelDescriptor_ValueType = 1
	LabelDescriptor_INT64  LabelDescriptor_ValueType = 2
)

var LabelDescriptor_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
}

var LabelDescriptor_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
}

func (x LabelDescriptor_ValueType) String() string {
	return proto.EnumName(LabelDescriptor_ValueType_name, int32(x))
}

func (LabelDescriptor_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac2bbc4ccf6ec7ba, []int{0, 0}
}

type LabelDescriptor struct {
	Key                  string                    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueType            LabelDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=argcv.api.LabelDescriptor_ValueType" json:"value_type,omitempty"`
	Description          string                    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LabelDescriptor) Reset()         { *m = LabelDescriptor{} }
func (m *LabelDescriptor) String() string { return proto.CompactTextString(m) }
func (*LabelDescriptor) ProtoMessage()    {}
func (*LabelDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac2bbc4ccf6ec7ba, []int{0}
}

func (m *LabelDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelDescriptor.Unmarshal(m, b)
}
func (m *LabelDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelDescriptor.Marshal(b, m, deterministic)
}
func (m *LabelDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelDescriptor.Merge(m, src)
}
func (m *LabelDescriptor) XXX_Size() int {
	return xxx_messageInfo_LabelDescriptor.Size(m)
}
func (m *LabelDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_LabelDescriptor proto.InternalMessageInfo

func (m *LabelDescriptor) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelDescriptor) GetValueType() LabelDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return LabelDescriptor_STRING
}

func (m *LabelDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("argcv.api.LabelDescriptor_ValueType", LabelDescriptor_ValueType_name, LabelDescriptor_ValueType_value)
	proto.RegisterType((*LabelDescriptor)(nil), "argcv.api.LabelDescriptor")
}

func init() { proto.RegisterFile("argcv/proto/api/label.proto", fileDescriptor_ac2bbc4ccf6ec7ba) }

var fileDescriptor_ac2bbc4ccf6ec7ba = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2c, 0x4a, 0x4f,
	0x2e, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x49, 0x4c, 0x4a,
	0xcd, 0xd1, 0x03, 0xf3, 0x85, 0x38, 0xc1, 0x92, 0x7a, 0x89, 0x05, 0x99, 0x4a, 0x3b, 0x18, 0xb9,
	0xf8, 0x7d, 0x40, 0x52, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x42, 0x02,
	0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x33,
	0x17, 0x57, 0x59, 0x62, 0x4e, 0x69, 0x6a, 0x7c, 0x49, 0x65, 0x41, 0xaa, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x9f, 0x91, 0x8a, 0x1e, 0xdc, 0x14, 0x3d, 0x34, 0x13, 0xf4, 0xc2, 0x40, 0x8a, 0x43, 0x2a,
	0x0b, 0x52, 0x83, 0x38, 0xcb, 0x60, 0x4c, 0x21, 0x05, 0x2e, 0xee, 0x14, 0xa8, 0x92, 0xcc, 0xfc,
	0x3c, 0x09, 0x66, 0xb0, 0xf1, 0xc8, 0x42, 0x4a, 0x3a, 0x5c, 0x9c, 0x70, 0x9d, 0x42, 0x5c, 0x5c,
	0x6c, 0xc1, 0x21, 0x41, 0x9e, 0x7e, 0xee, 0x02, 0x0c, 0x42, 0x1c, 0x5c, 0x2c, 0x4e, 0xfe, 0xfe,
	0x3e, 0x02, 0x8c, 0x42, 0x9c, 0x5c, 0xac, 0x9e, 0x7e, 0x21, 0x66, 0x26, 0x02, 0x4c, 0x4e, 0xa1,
	0x5c, 0xc2, 0xc9, 0xf9, 0xb9, 0x50, 0x57, 0x80, 0x3d, 0x06, 0x72, 0x8b, 0x13, 0x17, 0xd8, 0x31,
	0x01, 0x20, 0x7e, 0x00, 0x63, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x3e, 0x24, 0x48, 0xd2, 0xf3, 0x75, 0xc1, 0x8c, 0xc4, 0x82, 0xcc, 0x62, 0x44, 0xc8, 0x58,
	0x83, 0xc9, 0x1f, 0x8c, 0x8c, 0x49, 0x6c, 0x60, 0xa3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x36, 0xec, 0x23, 0x55, 0x42, 0x01, 0x00, 0x00,
}
