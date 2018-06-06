// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/api/http.proto

package http // import "github.com/argcv/go-argcvapis/api/http"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Http struct {
	Rules                []*HttpRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Http) Reset()         { *m = Http{} }
func (m *Http) String() string { return proto.CompactTextString(m) }
func (*Http) ProtoMessage()    {}
func (*Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_dc4fa9b2ac2486f1, []int{0}
}
func (m *Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http.Unmarshal(m, b)
}
func (m *Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http.Marshal(b, m, deterministic)
}
func (dst *Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http.Merge(dst, src)
}
func (m *Http) XXX_Size() int {
	return xxx_messageInfo_Http.Size(m)
}
func (m *Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Http proto.InternalMessageInfo

func (m *Http) GetRules() []*HttpRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type HttpRule struct {
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern              isHttpRule_Pattern `protobuf_oneof:"pattern"`
	Body                 string             `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	AdditionalBindings   []*HttpRule        `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings,proto3" json:"additional_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_dc4fa9b2ac2486f1, []int{1}
}
func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (dst *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(dst, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}
type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}
type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}
type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}
type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}
type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern()    {}
func (*HttpRule_Put) isHttpRule_Pattern()    {}
func (*HttpRule_Post) isHttpRule_Pattern()   {}
func (*HttpRule_Delete) isHttpRule_Pattern() {}
func (*HttpRule_Patch) isHttpRule_Pattern()  {}
func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HttpRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HttpRule_OneofMarshaler, _HttpRule_OneofUnmarshaler, _HttpRule_OneofSizer, []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

func _HttpRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Get)
	case *HttpRule_Put:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Put)
	case *HttpRule_Post:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Post)
	case *HttpRule_Delete:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Delete)
	case *HttpRule_Patch:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Patch)
	case *HttpRule_Custom:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HttpRule.Pattern has unexpected type %T", x)
	}
	return nil
}

func _HttpRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HttpRule)
	switch tag {
	case 2: // pattern.get
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Get{x}
		return true, err
	case 3: // pattern.put
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Put{x}
		return true, err
	case 4: // pattern.post
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Post{x}
		return true, err
	case 5: // pattern.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Delete{x}
		return true, err
	case 6: // pattern.patch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Pattern = &HttpRule_Patch{x}
		return true, err
	case 8: // pattern.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CustomHttpPattern)
		err := b.DecodeMessage(msg)
		m.Pattern = &HttpRule_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HttpRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HttpRule)
	// pattern
	switch x := m.Pattern.(type) {
	case *HttpRule_Get:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Get)))
		n += len(x.Get)
	case *HttpRule_Put:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Put)))
		n += len(x.Put)
	case *HttpRule_Post:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Post)))
		n += len(x.Post)
	case *HttpRule_Delete:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Delete)))
		n += len(x.Delete)
	case *HttpRule_Patch:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Patch)))
		n += len(x.Patch)
	case *HttpRule_Custom:
		s := proto.Size(x.Custom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CustomHttpPattern struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_dc4fa9b2ac2486f1, []int{2}
}
func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (dst *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(dst, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Http)(nil), "argcv.api.Http")
	proto.RegisterType((*HttpRule)(nil), "argcv.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "argcv.api.CustomHttpPattern")
}

func init() { proto.RegisterFile("argcv/proto/api/http.proto", fileDescriptor_http_dc4fa9b2ac2486f1) }

var fileDescriptor_http_dc4fa9b2ac2486f1 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x7e, 0xd3, 0xa6, 0x69, 0x33, 0x3d, 0xbd, 0x53, 0x91, 0xa5, 0x78, 0x28, 0x3d, 0x55, 0xc4,
	0x04, 0x2b, 0x78, 0xe9, 0x2d, 0x7a, 0xe8, 0xb1, 0xe4, 0xe8, 0x45, 0x36, 0xc9, 0x92, 0x2c, 0xa6,
	0xd9, 0x25, 0x99, 0x08, 0xfe, 0x2e, 0xff, 0x9c, 0x47, 0xd9, 0xcd, 0xb6, 0x0a, 0xe2, 0x25, 0xcc,
	0xf3, 0xc1, 0x3c, 0x93, 0x99, 0x85, 0x25, 0x6f, 0xcb, 0xfc, 0x2d, 0xd6, 0xad, 0x22, 0x15, 0x73,
	0x2d, 0xe3, 0x8a, 0x48, 0x47, 0x16, 0x62, 0x68, 0xb5, 0x88, 0x6b, 0xb9, 0xbe, 0x03, 0x7f, 0x4f,
	0xa4, 0xf1, 0x1a, 0x26, 0x6d, 0x5f, 0x8b, 0x8e, 0x79, 0xab, 0xf1, 0x66, 0xbe, 0x5d, 0x44, 0x67,
	0x4b, 0x64, 0xf4, 0xb4, 0xaf, 0x45, 0x3a, 0x38, 0xd6, 0x1f, 0x23, 0x98, 0x9d, 0x38, 0x5c, 0xc2,
	0xac, 0x13, 0xb5, 0xc8, 0x49, 0xb5, 0xcc, 0x5b, 0x79, 0x9b, 0x30, 0x3d, 0x63, 0x44, 0x18, 0x97,
	0x82, 0xd8, 0xc8, 0xd0, 0xfb, 0x7f, 0xa9, 0x01, 0x86, 0xd3, 0x3d, 0xb1, 0xf1, 0x89, 0xd3, 0x3d,
	0xe1, 0x05, 0xf8, 0x5a, 0x75, 0xc4, 0x7c, 0x47, 0x5a, 0x84, 0x0c, 0x82, 0x42, 0xd4, 0x82, 0x04,
	0x9b, 0x38, 0xde, 0x61, 0xbc, 0x84, 0x89, 0xe6, 0x94, 0x57, 0x2c, 0x70, 0xc2, 0x00, 0xf1, 0x01,
	0x82, 0xbc, 0xef, 0x48, 0x1d, 0xd9, 0x6c, 0xe5, 0x6d, 0xe6, 0xdb, 0xab, 0x1f, 0x3f, 0xf1, 0x68,
	0x05, 0x33, 0xf6, 0x81, 0x13, 0x89, 0xb6, 0x31, 0xfd, 0x06, 0x37, 0x22, 0xf8, 0x99, 0x2a, 0xde,
	0xd9, 0xd4, 0xce, 0x6f, 0x6b, 0x7c, 0x82, 0x05, 0x2f, 0x0a, 0x49, 0x52, 0x35, 0xbc, 0x7e, 0xc9,
	0x64, 0x53, 0xc8, 0xa6, 0xec, 0xd8, 0xfc, 0xef, 0xed, 0xe0, 0xb7, 0x3f, 0x71, 0xf6, 0x24, 0x84,
	0xa9, 0x1e, 0xe2, 0xd6, 0x3b, 0xf8, 0xff, 0x6b, 0x06, 0x93, 0xfc, 0x2a, 0x9b, 0xc2, 0x6d, 0xce,
	0xd6, 0x86, 0xd3, 0x9c, 0xaa, 0x61, 0x6d, 0xa9, 0xad, 0x93, 0x14, 0x16, 0xb9, 0x3a, 0xba, 0x54,
	0x7b, 0x43, 0x93, 0x9d, 0x84, 0xb6, 0x97, 0x81, 0x07, 0xef, 0xf9, 0xa6, 0x94, 0x54, 0xf5, 0x59,
	0x94, 0xab, 0x63, 0x3c, 0xdc, 0xbe, 0x54, 0xb7, 0xb6, 0xe0, 0x5a, 0x76, 0xe7, 0x27, 0xb0, 0x33,
	0x9f, 0x4f, 0xcf, 0xcb, 0x02, 0xdb, 0xe7, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x95, 0x05, 0x04,
	0x73, 0x29, 0x02, 0x00, 0x00,
}
