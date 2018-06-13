// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/app/manul/manul.proto

package manul // import "github.com/argcv/go-argcvapis/app/manul/manul"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/argcv/go-argcvapis/api/annotations"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManulServiceClient is the client API for ManulService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManulServiceClient interface {
}

type manulServiceClient struct {
	cc *grpc.ClientConn
}

func NewManulServiceClient(cc *grpc.ClientConn) ManulServiceClient {
	return &manulServiceClient{cc}
}

// ManulServiceServer is the server API for ManulService service.
type ManulServiceServer interface {
}

func RegisterManulServiceServer(s *grpc.Server, srv ManulServiceServer) {
	s.RegisterService(&_ManulService_serviceDesc, srv)
}

var _ManulService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "argcv.app.manul.ManulService",
	HandlerType: (*ManulServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "argcv/proto/app/manul/manul.proto",
}

func init() {
	proto.RegisterFile("argcv/proto/app/manul/manul.proto", fileDescriptor_manul_1ca47c5aabbd7909)
}

var fileDescriptor_manul_1ca47c5aabbd7909 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2c, 0x4a, 0x4f,
	0x2e, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0x28, 0xd0, 0xcf, 0x4d, 0xcc, 0x2b,
	0xcd, 0x81, 0x90, 0x7a, 0x60, 0x51, 0x21, 0x7e, 0xb0, 0x12, 0xbd, 0xc4, 0x82, 0x02, 0x3d, 0xb0,
	0xb0, 0x14, 0x9a, 0x9e, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc,
	0x62, 0x88, 0x1e, 0x23, 0x3e, 0x2e, 0x1e, 0x5f, 0x90, 0xda, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4,
	0x54, 0xa7, 0x44, 0x2e, 0xc9, 0xe4, 0xfc, 0x5c, 0x3d, 0x88, 0x49, 0x60, 0x25, 0x08, 0xf3, 0x9c,
	0xb8, 0xc0, 0x4a, 0x03, 0x40, 0xa2, 0x01, 0x8c, 0x51, 0xc6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0x8b, 0xd2, 0xf3, 0x75, 0xc1, 0x8c, 0xc4, 0x82, 0xcc, 0x62,
	0x74, 0x37, 0x5a, 0x83, 0xc9, 0x1f, 0x8c, 0x8c, 0x49, 0x6c, 0x60, 0x63, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf2, 0xa6, 0xb0, 0x4a, 0xd2, 0x00, 0x00, 0x00,
}
