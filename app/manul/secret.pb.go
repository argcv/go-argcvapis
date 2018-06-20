// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/app/manul/secret.proto

package manul // import "github.com/argcv/go-argcvapis/app/manul"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/argcv/go-argcvapis/api/annotations"
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
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Secret               string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_63915116ce98f41e, []int{0}
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

func (m *Secret) GetProjectId() string {
	if m != nil {
		return m.ProjectId
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
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// leave it empty for global
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Types that are valid to be assigned to Origin:
	//	*UpdateSecretRequest_GlobalSecret
	//	*UpdateSecretRequest_CurrentSecret
	//	*UpdateSecretRequest_TempToken
	Origin               isUpdateSecretRequest_Origin `protobuf_oneof:"origin"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UpdateSecretRequest) Reset()         { *m = UpdateSecretRequest{} }
func (m *UpdateSecretRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSecretRequest) ProtoMessage()    {}
func (*UpdateSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_63915116ce98f41e, []int{1}
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

type isUpdateSecretRequest_Origin interface {
	isUpdateSecretRequest_Origin()
}

type UpdateSecretRequest_GlobalSecret struct {
	GlobalSecret string `protobuf:"bytes,3,opt,name=global_secret,json=globalSecret,proto3,oneof"`
}
type UpdateSecretRequest_CurrentSecret struct {
	CurrentSecret string `protobuf:"bytes,4,opt,name=current_secret,json=currentSecret,proto3,oneof"`
}
type UpdateSecretRequest_TempToken struct {
	TempToken string `protobuf:"bytes,5,opt,name=temp_token,json=tempToken,proto3,oneof"`
}

func (*UpdateSecretRequest_GlobalSecret) isUpdateSecretRequest_Origin()  {}
func (*UpdateSecretRequest_CurrentSecret) isUpdateSecretRequest_Origin() {}
func (*UpdateSecretRequest_TempToken) isUpdateSecretRequest_Origin()     {}

func (m *UpdateSecretRequest) GetOrigin() isUpdateSecretRequest_Origin {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *UpdateSecretRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateSecretRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *UpdateSecretRequest) GetGlobalSecret() string {
	if x, ok := m.GetOrigin().(*UpdateSecretRequest_GlobalSecret); ok {
		return x.GlobalSecret
	}
	return ""
}

func (m *UpdateSecretRequest) GetCurrentSecret() string {
	if x, ok := m.GetOrigin().(*UpdateSecretRequest_CurrentSecret); ok {
		return x.CurrentSecret
	}
	return ""
}

func (m *UpdateSecretRequest) GetTempToken() string {
	if x, ok := m.GetOrigin().(*UpdateSecretRequest_TempToken); ok {
		return x.TempToken
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpdateSecretRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpdateSecretRequest_OneofMarshaler, _UpdateSecretRequest_OneofUnmarshaler, _UpdateSecretRequest_OneofSizer, []interface{}{
		(*UpdateSecretRequest_GlobalSecret)(nil),
		(*UpdateSecretRequest_CurrentSecret)(nil),
		(*UpdateSecretRequest_TempToken)(nil),
	}
}

func _UpdateSecretRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpdateSecretRequest)
	// origin
	switch x := m.Origin.(type) {
	case *UpdateSecretRequest_GlobalSecret:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.GlobalSecret)
	case *UpdateSecretRequest_CurrentSecret:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.CurrentSecret)
	case *UpdateSecretRequest_TempToken:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TempToken)
	case nil:
	default:
		return fmt.Errorf("UpdateSecretRequest.Origin has unexpected type %T", x)
	}
	return nil
}

func _UpdateSecretRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpdateSecretRequest)
	switch tag {
	case 3: // origin.global_secret
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Origin = &UpdateSecretRequest_GlobalSecret{x}
		return true, err
	case 4: // origin.current_secret
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Origin = &UpdateSecretRequest_CurrentSecret{x}
		return true, err
	case 5: // origin.temp_token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Origin = &UpdateSecretRequest_TempToken{x}
		return true, err
	default:
		return false, nil
	}
}

func _UpdateSecretRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpdateSecretRequest)
	// origin
	switch x := m.Origin.(type) {
	case *UpdateSecretRequest_GlobalSecret:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.GlobalSecret)))
		n += len(x.GlobalSecret)
	case *UpdateSecretRequest_CurrentSecret:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.CurrentSecret)))
		n += len(x.CurrentSecret)
	case *UpdateSecretRequest_TempToken:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.TempToken)))
		n += len(x.TempToken)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
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
	return fileDescriptor_secret_63915116ce98f41e, []int{2}
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

type isUpdateSecretResponse_Result interface {
	isUpdateSecretResponse_Result()
}

type UpdateSecretResponse_Error struct {
	Error *status.Status `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}
type UpdateSecretResponse_Secret struct {
	Secret *Secret `protobuf:"bytes,4,opt,name=secret,proto3,oneof"`
}

func (*UpdateSecretResponse_Error) isUpdateSecretResponse_Result()  {}
func (*UpdateSecretResponse_Secret) isUpdateSecretResponse_Result() {}

func (m *UpdateSecretResponse) GetResult() isUpdateSecretResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotSecretRequest) Reset()         { *m = ForgotSecretRequest{} }
func (m *ForgotSecretRequest) String() string { return proto.CompactTextString(m) }
func (*ForgotSecretRequest) ProtoMessage()    {}
func (*ForgotSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_secret_63915116ce98f41e, []int{3}
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
	return fileDescriptor_secret_63915116ce98f41e, []int{4}
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

func init() {
	proto.RegisterFile("argcv/proto/app/manul/secret.proto", fileDescriptor_secret_63915116ce98f41e)
}

var fileDescriptor_secret_63915116ce98f41e = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6e, 0x13, 0x3d,
	0x10, 0xce, 0xe6, 0x6f, 0xb7, 0xcd, 0x34, 0xf9, 0x11, 0x4e, 0x44, 0x93, 0x48, 0x28, 0xb0, 0xa2,
	0x02, 0x55, 0x74, 0x57, 0x2d, 0xb7, 0x72, 0xcb, 0x01, 0xa5, 0xb7, 0x6a, 0x03, 0x12, 0x42, 0x48,
	0x91, 0xe3, 0x0c, 0xcb, 0x42, 0x76, 0x6d, 0x6c, 0xef, 0x5e, 0x80, 0x0b, 0x12, 0x4f, 0xc0, 0xb3,
	0xf0, 0x0e, 0xdc, 0x79, 0x05, 0x1e, 0x82, 0x23, 0x5a, 0xdb, 0x21, 0x09, 0xac, 0x44, 0x05, 0x97,
	0x38, 0x33, 0xf3, 0x79, 0xbe, 0xf9, 0x3e, 0xdb, 0x0b, 0x01, 0x95, 0x09, 0x2b, 0x23, 0x21, 0xb9,
	0xe6, 0x11, 0x15, 0x22, 0xca, 0x68, 0x5e, 0x2c, 0x23, 0x85, 0x4c, 0xa2, 0x0e, 0x4d, 0x9a, 0x5c,
	0x33, 0x98, 0x90, 0x0a, 0x11, 0x9a, 0xea, 0xf0, 0xf6, 0xf6, 0xa6, 0x34, 0xa2, 0x79, 0xce, 0x35,
	0xd5, 0x29, 0xcf, 0x95, 0xdd, 0x33, 0x1c, 0x6d, 0x42, 0x94, 0xa6, 0xba, 0x50, 0x6e, 0xb1, 0x80,
	0xe0, 0x29, 0xf8, 0x53, 0x43, 0x42, 0x0e, 0x61, 0xaf, 0x50, 0x28, 0x67, 0xe9, 0xa2, 0xef, 0xdd,
	0xf2, 0xee, 0xb5, 0x62, 0xbf, 0x0a, 0x2f, 0x16, 0xe4, 0x26, 0x80, 0x90, 0xfc, 0x15, 0x32, 0x5d,
	0xd5, 0x9a, 0xa6, 0xd6, 0x72, 0x99, 0x8b, 0x05, 0xb9, 0x01, 0xbe, 0x1d, 0xb3, 0xff, 0x9f, 0xdd,
	0x66, 0xa3, 0xe0, 0x8b, 0x07, 0xdd, 0x27, 0x62, 0x41, 0x35, 0x5a, 0x82, 0x18, 0xdf, 0x14, 0xa8,
	0xfe, 0x9e, 0xe7, 0x08, 0x3a, 0xc9, 0x92, 0xcf, 0xe9, 0x72, 0xb6, 0x49, 0x37, 0x69, 0xc4, 0x6d,
	0x9b, 0x76, 0x32, 0xee, 0xc2, 0xff, 0xac, 0x90, 0x12, 0x73, 0xbd, 0xc2, 0xed, 0x38, 0x5c, 0xc7,
	0xe5, 0x1d, 0x70, 0x04, 0xa0, 0x31, 0x13, 0x33, 0xcd, 0x5f, 0x63, 0xde, 0xdf, 0x75, 0xa0, 0x56,
	0x95, 0x7b, 0x5c, 0xa5, 0xc6, 0xfb, 0xe0, 0x73, 0x99, 0x26, 0x69, 0x1e, 0x7c, 0xf6, 0xa0, 0xb7,
	0x2d, 0x45, 0x09, 0x9e, 0x2b, 0x24, 0x7d, 0xd8, 0x53, 0x05, 0x63, 0xa8, 0x94, 0xd1, 0xb2, 0x1f,
	0xaf, 0xc2, 0xaa, 0x92, 0xa1, 0x52, 0x34, 0x41, 0xa7, 0x64, 0x15, 0x92, 0xfb, 0xb0, 0x8b, 0x52,
	0x72, 0x69, 0xe6, 0x3f, 0x38, 0xeb, 0x85, 0xf6, 0x58, 0xdd, 0xa9, 0x4c, 0xcd, 0x32, 0x69, 0xc4,
	0x16, 0x44, 0x4e, 0x7f, 0xba, 0xbb, 0x63, 0xe0, 0x87, 0xe1, 0x2f, 0xb7, 0x20, 0xb4, 0x23, 0x4d,
	0x1a, 0x2b, 0xe3, 0xab, 0xb9, 0x25, 0xaa, 0x62, 0xa9, 0x83, 0x10, 0xba, 0x8f, 0xb8, 0x4c, 0xb8,
	0xbe, 0xda, 0x09, 0x04, 0x25, 0xf4, 0xb6, 0xf1, 0xff, 0x20, 0xf3, 0xf8, 0x0a, 0x32, 0x9d, 0xc8,
	0xb3, 0x8f, 0x4d, 0xe8, 0x58, 0xca, 0x29, 0xca, 0x32, 0x65, 0x48, 0xde, 0x41, 0x7b, 0xd3, 0x70,
	0x72, 0xe7, 0x37, 0xd9, 0x35, 0x57, 0x6b, 0x78, 0xf4, 0x07, 0x94, 0x95, 0x13, 0x8c, 0x3e, 0x7c,
	0xfd, 0xf6, 0xa9, 0x39, 0x18, 0xf6, 0xa2, 0xf2, 0xd4, 0x3d, 0xb1, 0xe8, 0xad, 0x73, 0xe4, 0xfd,
	0xb9, 0x77, 0x4c, 0x4a, 0x68, 0x6f, 0xfa, 0x50, 0xc3, 0x5e, 0x63, 0x6b, 0x0d, 0x7b, 0x9d, 0x99,
	0xc1, 0xc0, 0xb0, 0x77, 0x83, 0xeb, 0x6b, 0xf6, 0xf3, 0x17, 0x06, 0x38, 0x7e, 0x0e, 0x03, 0xc6,
	0x33, 0xd7, 0xc6, 0xbc, 0xcf, 0x75, 0xb3, 0xf1, 0x81, 0xed, 0x73, 0x59, 0xa5, 0x2f, 0xbd, 0x67,
	0x27, 0x49, 0xaa, 0x5f, 0x16, 0xf3, 0x90, 0xf1, 0x2c, 0xb2, 0x8f, 0x3c, 0xe1, 0x27, 0xe6, 0x0f,
	0x15, 0xa9, 0x5a, 0x7f, 0x43, 0x1e, 0x9a, 0xdf, 0xef, 0x9e, 0x37, 0xf7, 0x4d, 0xc7, 0x07, 0x3f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xf9, 0xa6, 0x0a, 0x6c, 0x04, 0x00, 0x00,
}
