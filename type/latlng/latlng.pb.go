// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/type/latlng.proto

package latlng

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// An object representing a latitude/longitude pair. This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
//
// Example of normalization code in Python:
//
//     def NormalizeLongitude(longitude):
//       """Wraps decimal degrees longitude to [-180.0, 180.0]."""
//       q, r = divmod(longitude, 360.0)
//       if r > 180.0 or (r == 180.0 and q <= -1.0):
//         return r - 360.0
//       return r
//
//     def NormalizeLatLng(latitude, longitude):
//       """Wraps decimal degrees latitude and longitude to
//       [-90.0, 90.0] and [-180.0, 180.0], respectively."""
//       r = latitude % 360.0
//       if r <= 90.0:
//         return r, NormalizeLongitude(longitude)
//       elif r >= 270.0:
//         return r - 360, NormalizeLongitude(longitude)
//       else:
//         return 180 - r, NormalizeLongitude(longitude + 180.0)
//
//     assert 180.0 == NormalizeLongitude(180.0)
//     assert -180.0 == NormalizeLongitude(-180.0)
//     assert -179.0 == NormalizeLongitude(181.0)
//     assert (0.0, 0.0) == NormalizeLatLng(360.0, 0.0)
//     assert (0.0, 0.0) == NormalizeLatLng(-360.0, 0.0)
//     assert (85.0, 180.0) == NormalizeLatLng(95.0, 0.0)
//     assert (-85.0, -170.0) == NormalizeLatLng(-95.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(90.0, 10.0)
//     assert (-90.0, -10.0) == NormalizeLatLng(-90.0, -10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(-180.0, 10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(180.0, 10.0)
//     assert (-90.0, 10.0) == NormalizeLatLng(270.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(-270.0, 10.0)
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude            float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatLng) Reset()         { *m = LatLng{} }
func (m *LatLng) String() string { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()    {}
func (*LatLng) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd168e5b33c51d, []int{0}
}

func (m *LatLng) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatLng.Unmarshal(m, b)
}
func (m *LatLng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatLng.Marshal(b, m, deterministic)
}
func (m *LatLng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatLng.Merge(m, src)
}
func (m *LatLng) XXX_Size() int {
	return xxx_messageInfo_LatLng.Size(m)
}
func (m *LatLng) XXX_DiscardUnknown() {
	xxx_messageInfo_LatLng.DiscardUnknown(m)
}

var xxx_messageInfo_LatLng proto.InternalMessageInfo

func (m *LatLng) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *LatLng) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*LatLng)(nil), "argcv.type.LatLng")
}

func init() { proto.RegisterFile("argcv/proto/type/latlng.proto", fileDescriptor_90cd168e5b33c51d) }

var fileDescriptor_90cd168e5b33c51d = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0x4a, 0x4f,
	0x2e, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x49, 0x2c,
	0xc9, 0xc9, 0x4b, 0xd7, 0x03, 0x8b, 0x08, 0x71, 0x81, 0xa5, 0xf5, 0x40, 0x12, 0x4a, 0x4e, 0x5c,
	0x6c, 0x3e, 0x89, 0x25, 0x3e, 0x79, 0xe9, 0x42, 0x52, 0x5c, 0x1c, 0x39, 0x89, 0x25, 0x99, 0x25,
	0xa5, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x8c, 0x41, 0x70, 0xbe, 0x90, 0x0c, 0x17, 0x67,
	0x4e, 0x7e, 0x5e, 0x3a, 0x44, 0x92, 0x09, 0x2c, 0x89, 0x10, 0x70, 0x8a, 0xe2, 0x12, 0x49, 0xce,
	0xcf, 0xd5, 0x83, 0x98, 0x0a, 0xb6, 0x02, 0x6c, 0xb6, 0x13, 0x37, 0xc4, 0xe4, 0x00, 0x90, 0x48,
	0x00, 0x63, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4,
	0x81, 0xe9, 0xf9, 0xba, 0x60, 0x46, 0x62, 0x41, 0x66, 0x31, 0xb2, 0x3b, 0xad, 0x21, 0xd4, 0x0f,
	0x46, 0xc6, 0x24, 0x36, 0xb0, 0x79, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x50, 0x5a, 0xce,
	0xca, 0xd3, 0x00, 0x00, 0x00,
}
