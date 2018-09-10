// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/app/manul/secret.proto

package secret // import "github.com/argcv/go-argcvapis/app/manul/secret"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/argcv/go-argcvapis/api/annotations"
import user "github.com/argcv/go-argcvapis/app/manul/user"
import status "github.com/argcv/go-argcvapis/status/status"

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

// A simple struct to define secretorization info
type Secret struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e5759c6d0aac59, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (dst *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(dst, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Secret) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type UpdateSecretRequest struct {
	// by user id
	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	CurrentSecret string `protobuf:"bytes,3,opt,name=current_secret,json=currentSecret,proto3" json:"current_secret,omitempty"`
	TempToken     string `protobuf:"bytes,4,opt,name=temp_token,json=tempToken,proto3" json:"temp_token,omitempty"`
	// Optional: if contains auth info
	Auth                 *user.AuthToken `protobuf:"bytes,5,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateSecretRequest) Reset()         { *m = UpdateSecretRequest{} }
func (m *UpdateSecretRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSecretRequest) ProtoMessage()    {}
func (*UpdateSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e5759c6d0aac59, []int{1}
}
func (m *UpdateSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSecretRequest.Unmarshal(m, b)
}
func (m *UpdateSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSecretRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSecretRequest.Merge(dst, src)
}
func (m *UpdateSecretRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSecretRequest.Size(m)
}
func (m *UpdateSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSecretRequest proto.InternalMessageInfo

func (m *UpdateSecretRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateSecretRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdateSecretRequest) GetCurrentSecret() string {
	if m != nil {
		return m.CurrentSecret
	}
	return ""
}

func (m *UpdateSecretRequest) GetTempToken() string {
	if m != nil {
		return m.TempToken
	}
	return ""
}

func (m *UpdateSecretRequest) GetAuth() *user.AuthToken {
	if m != nil {
		return m.Auth
	}
	return nil
}

type UpdateSecretResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The operation result, which can be either an `error` or a valid `secrets`.
	// If `success` == `false`, neither `error` nor `secrets` is set.
	// If `success` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*UpdateSecretResponse_Error
	//	*UpdateSecretResponse_Secret
	Result               isUpdateSecretResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UpdateSecretResponse) Reset()         { *m = UpdateSecretResponse{} }
func (m *UpdateSecretResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateSecretResponse) ProtoMessage()    {}
func (*UpdateSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e5759c6d0aac59, []int{2}
}
func (m *UpdateSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSecretResponse.Unmarshal(m, b)
}
func (m *UpdateSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSecretResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSecretResponse.Merge(dst, src)
}
func (m *UpdateSecretResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateSecretResponse.Size(m)
}
func (m *UpdateSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSecretResponse proto.InternalMessageInfo

func (m *UpdateSecretResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateSecretResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type isUpdateSecretResponse_Result interface {
	isUpdateSecretResponse_Result()
}

type UpdateSecretResponse_Error struct {
	Error *status.Status `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

type UpdateSecretResponse_Secret struct {
	Secret *Secret `protobuf:"bytes,4,opt,name=secret,proto3,oneof"`
}

func (*UpdateSecretResponse_Error) isUpdateSecretResponse_Result() {}

func (*UpdateSecretResponse_Secret) isUpdateSecretResponse_Result() {}

func (m *UpdateSecretResponse) GetResult() isUpdateSecretResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UpdateSecretResponse) GetError() *status.Status {
	if x, ok := m.GetResult().(*UpdateSecretResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (m *UpdateSecretResponse) GetSecret() *Secret {
	if x, ok := m.GetResult().(*UpdateSecretResponse_Secret); ok {
		return x.Secret
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpdateSecretResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpdateSecretResponse_OneofMarshaler, _UpdateSecretResponse_OneofUnmarshaler, _UpdateSecretResponse_OneofSizer, []interface{}{
		(*UpdateSecretResponse_Error)(nil),
		(*UpdateSecretResponse_Secret)(nil),
	}
}

func _UpdateSecretResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpdateSecretResponse)
	// result
	switch x := m.Result.(type) {
	case *UpdateSecretResponse_Error:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *UpdateSecretResponse_Secret:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Secret); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UpdateSecretResponse.Result has unexpected type %T", x)
	}
	return nil
}

func _UpdateSecretResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpdateSecretResponse)
	switch tag {
	case 3: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(status.Status)
		err := b.DecodeMessage(msg)
		m.Result = &UpdateSecretResponse_Error{msg}
		return true, err
	case 4: // result.secret
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Secret)
		err := b.DecodeMessage(msg)
		m.Result = &UpdateSecretResponse_Secret{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UpdateSecretResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpdateSecretResponse)
	// result
	switch x := m.Result.(type) {
	case *UpdateSecretResponse_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpdateSecretResponse_Secret:
		s := proto.Size(x.Secret)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ForgotSecretRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotSecretRequest) Reset()         { *m = ForgotSecretRequest{} }
func (m *ForgotSecretRequest) String() string { return proto.CompactTextString(m) }
func (*ForgotSecretRequest) ProtoMessage()    {}
func (*ForgotSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e5759c6d0aac59, []int{3}
}
func (m *ForgotSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotSecretRequest.Unmarshal(m, b)
}
func (m *ForgotSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotSecretRequest.Marshal(b, m, deterministic)
}
func (dst *ForgotSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotSecretRequest.Merge(dst, src)
}
func (m *ForgotSecretRequest) XXX_Size() int {
	return xxx_messageInfo_ForgotSecretRequest.Size(m)
}
func (m *ForgotSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotSecretRequest proto.InternalMessageInfo

func (m *ForgotSecretRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ForgotSecretRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type ForgotSecretResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// error message
	Error                *status.Status `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ForgotSecretResponse) Reset()         { *m = ForgotSecretResponse{} }
func (m *ForgotSecretResponse) String() string { return proto.CompactTextString(m) }
func (*ForgotSecretResponse) ProtoMessage()    {}
func (*ForgotSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e5759c6d0aac59, []int{4}
}
func (m *ForgotSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotSecretResponse.Unmarshal(m, b)
}
func (m *ForgotSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotSecretResponse.Marshal(b, m, deterministic)
}
func (dst *ForgotSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotSecretResponse.Merge(dst, src)
}
func (m *ForgotSecretResponse) XXX_Size() int {
	return xxx_messageInfo_ForgotSecretResponse.Size(m)
}
func (m *ForgotSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotSecretResponse proto.InternalMessageInfo

func (m *ForgotSecretResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ForgotSecretResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ForgotSecretResponse) GetError() *status.Status {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Secret)(nil), "argcv.app.manul.Secret")
	proto.RegisterType((*UpdateSecretRequest)(nil), "argcv.app.manul.UpdateSecretRequest")
	proto.RegisterType((*UpdateSecretResponse)(nil), "argcv.app.manul.UpdateSecretResponse")
	proto.RegisterType((*ForgotSecretRequest)(nil), "argcv.app.manul.ForgotSecretRequest")
	proto.RegisterType((*ForgotSecretResponse)(nil), "argcv.app.manul.ForgotSecretResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretServiceClient interface {
	UpdateSecret(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*UpdateSecretResponse, error)
	ForgotSecret(ctx context.Context, in *ForgotSecretRequest, opts ...grpc.CallOption) (*ForgotSecretResponse, error)
}

type secretServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretServiceClient(cc *grpc.ClientConn) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) UpdateSecret(ctx context.Context, in *UpdateSecretRequest, opts ...grpc.CallOption) (*UpdateSecretResponse, error) {
	out := new(UpdateSecretResponse)
	err := c.cc.Invoke(ctx, "/argcv.app.manul.SecretService/UpdateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ForgotSecret(ctx context.Context, in *ForgotSecretRequest, opts ...grpc.CallOption) (*ForgotSecretResponse, error) {
	out := new(ForgotSecretResponse)
	err := c.cc.Invoke(ctx, "/argcv.app.manul.SecretService/ForgotSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
type SecretServiceServer interface {
	UpdateSecret(context.Context, *UpdateSecretRequest) (*UpdateSecretResponse, error)
	ForgotSecret(context.Context, *ForgotSecretRequest) (*ForgotSecretResponse, error)
}

func RegisterSecretServiceServer(s *grpc.Server, srv SecretServiceServer) {
	s.RegisterService(&_SecretService_serviceDesc, srv)
}

func _SecretService_UpdateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).UpdateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/argcv.app.manul.SecretService/UpdateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).UpdateSecret(ctx, req.(*UpdateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ForgotSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).ForgotSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/argcv.app.manul.SecretService/ForgotSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).ForgotSecret(ctx, req.(*ForgotSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "argcv.app.manul.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSecret",
			Handler:    _SecretService_UpdateSecret_Handler,
		},
		{
			MethodName: "ForgotSecret",
			Handler:    _SecretService_ForgotSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "argcv/proto/app/manul/secret.proto",
}

func init() { proto.RegisterFile("argcv/proto/app/manul/secret.proto", fileDescriptor_99e5759c6d0aac59) }

var fileDescriptor_99e5759c6d0aac59 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0xe4, 0x4b, 0xd3, 0xe4, 0xa6, 0xfd, 0x10, 0x93, 0x88, 0xa6, 0x46, 0x88, 0x60, 0x51,
	0x51, 0x2a, 0xb0, 0xd5, 0x20, 0x16, 0x94, 0x15, 0x59, 0xa0, 0x02, 0x12, 0x8a, 0x12, 0xd8, 0x74,
	0x13, 0x4d, 0x9d, 0x4b, 0x62, 0x51, 0x7b, 0xcc, 0xfc, 0x78, 0x53, 0x75, 0xc3, 0x2b, 0xf0, 0x1c,
	0x2c, 0x11, 0x6b, 0x9e, 0x81, 0x57, 0xe0, 0x21, 0x58, 0x22, 0xcf, 0x4c, 0x68, 0x42, 0xcc, 0x8f,
	0x04, 0xab, 0xf1, 0x3d, 0x73, 0x7c, 0xee, 0xb9, 0x67, 0xec, 0x01, 0x9f, 0x89, 0x69, 0x94, 0x87,
	0x99, 0xe0, 0x8a, 0x87, 0x2c, 0xcb, 0xc2, 0x84, 0xa5, 0xfa, 0x34, 0x94, 0x18, 0x09, 0x54, 0x81,
	0x81, 0xe9, 0x25, 0xc3, 0x09, 0x58, 0x96, 0x05, 0x66, 0xd7, 0xbb, 0xb1, 0xfc, 0x52, 0x1c, 0xb2,
	0x34, 0xe5, 0x8a, 0xa9, 0x98, 0xa7, 0xd2, 0xbe, 0xe3, 0x5d, 0x5f, 0xa4, 0x48, 0xc5, 0x94, 0x96,
	0x6e, 0x71, 0x84, 0x6e, 0x79, 0x63, 0x2d, 0x51, 0x58, 0x86, 0xff, 0x00, 0x6a, 0x23, 0x63, 0x83,
	0x6e, 0xc3, 0x46, 0x81, 0x8f, 0xe3, 0x49, 0x87, 0x74, 0xc9, 0x5e, 0x63, 0x58, 0x2b, 0xca, 0x27,
	0x13, 0x7a, 0x05, 0x6a, 0xd6, 0x69, 0xa7, 0x62, 0x71, 0x5b, 0xf9, 0x9f, 0x08, 0xb4, 0x5e, 0x66,
	0x13, 0xa6, 0xd0, 0x2a, 0x0c, 0xf1, 0x8d, 0x46, 0xf9, 0x0b, 0xa1, 0xab, 0xd0, 0x30, 0x1b, 0x29,
	0x4b, 0xd0, 0x69, 0xd5, 0x0b, 0xe0, 0x39, 0x4b, 0x90, 0xee, 0xc2, 0xff, 0x91, 0x16, 0x02, 0x53,
	0x35, 0x76, 0xdd, 0xfe, 0x33, 0x8c, 0x2d, 0x87, 0x3a, 0x97, 0xd7, 0x00, 0x14, 0x26, 0xd9, 0x58,
	0xf1, 0xd7, 0x98, 0x76, 0xaa, 0x86, 0xd2, 0x28, 0x90, 0x17, 0x05, 0x40, 0x03, 0xa8, 0x32, 0xad,
	0x66, 0x9d, 0xf5, 0x2e, 0xd9, 0x6b, 0xf6, 0xbc, 0xe0, 0x87, 0x50, 0x83, 0x47, 0x5a, 0xcd, 0x0c,
	0x73, 0x68, 0x78, 0xfe, 0x07, 0x02, 0xed, 0xe5, 0x19, 0x64, 0xc6, 0x53, 0x89, 0xb4, 0x03, 0x1b,
	0x52, 0x47, 0x11, 0x4a, 0x69, 0x86, 0xa8, 0x0f, 0xe7, 0x65, 0xb1, 0x93, 0xa0, 0x94, 0x6c, 0x3a,
	0x9f, 0x61, 0x5e, 0xd2, 0x3b, 0xb0, 0x8e, 0x42, 0x70, 0x61, 0x9c, 0x37, 0x7b, 0x6d, 0xd7, 0xdd,
	0x9d, 0xc8, 0xc8, 0x2c, 0x47, 0x6b, 0x43, 0x4b, 0xa2, 0x07, 0xdf, 0x63, 0xad, 0x1a, 0xfa, 0xf6,
	0x8a, 0x59, 0x6b, 0xe9, 0x68, 0x6d, 0x9e, 0x78, 0xbf, 0x0e, 0x35, 0x81, 0x52, 0x9f, 0x2a, 0xff,
	0x19, 0xb4, 0x1e, 0x73, 0x31, 0xe5, 0xea, 0x1f, 0x44, 0xef, 0xe7, 0xd0, 0x5e, 0x16, 0xfb, 0x8b,
	0x0c, 0xf6, 0xff, 0x20, 0x03, 0x97, 0x40, 0xef, 0x63, 0x05, 0xb6, 0x6c, 0xcb, 0x11, 0x8a, 0x3c,
	0x8e, 0x90, 0xbe, 0x27, 0xb0, 0xb9, 0x78, 0x1c, 0xf4, 0xe6, 0x4a, 0x28, 0x25, 0x5f, 0x9c, 0xb7,
	0xfb, 0x1b, 0x96, 0x9d, 0xc7, 0x1f, 0xbc, 0xfd, 0xfc, 0xe5, 0x5d, 0xe5, 0xa9, 0xe7, 0x85, 0xf9,
	0x81, 0xfb, 0xf9, 0x0e, 0x05, 0x4a, 0x54, 0xe1, 0x99, 0x4b, 0xed, 0xfc, 0x90, 0xec, 0x1f, 0xdf,
	0xa6, 0xb7, 0x7e, 0x4e, 0x08, 0xcf, 0x2e, 0x3e, 0xbf, 0x73, 0x9a, 0xc3, 0xe6, 0x62, 0x72, 0x25,
	0x76, 0x4b, 0x4e, 0xa9, 0xc4, 0x6e, 0x59, 0xfc, 0xfe, 0x8e, 0xb1, 0xdb, 0xf2, 0x2f, 0x2f, 0xb8,
	0x79, 0x65, 0x88, 0xfd, 0x09, 0xec, 0x44, 0x3c, 0x71, 0x32, 0xe6, 0x47, 0xbe, 0x10, 0xeb, 0x37,
	0xad, 0xce, 0xa0, 0x80, 0x07, 0xe4, 0xf8, 0xfe, 0x34, 0x56, 0x33, 0x7d, 0x12, 0x44, 0x3c, 0x09,
	0xed, 0x75, 0x30, 0xe5, 0x77, 0xcd, 0x03, 0xcb, 0x62, 0xb9, 0x72, 0x1d, 0x3d, 0xb4, 0xcb, 0x57,
	0x42, 0x4e, 0x6a, 0x46, 0xf9, 0xde, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x5b, 0xd2, 0xb4,
	0xbf, 0x04, 0x00, 0x00,
}
