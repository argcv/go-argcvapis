// Code generated by protoc-gen-go. DO NOT EDIT.
// source: argcv/proto/longrunning/operations.proto

package longrunning // import "github.com/argcv/go-argcvapis/longrunning"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/argcv/go-argcvapis/api/annotations"
import status "github.com/argcv/go-argcvapis/status/status"
import any "github.com/golang/protobuf/ptypes/any"
import empty "github.com/golang/protobuf/ptypes/empty"

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

// This resource represents a long-running operation that is the result of a
// network API call.
type Operation struct {
	// The server-assigned name, which is only unique within the same service that
	// originally returns it. If you use the default HTTP mapping, the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Service-specific metadata associated with the operation.  It typically
	// contains progress information and common metadata such as create time.
	// Some services might not provide such metadata.  Any method that returns a
	// long-running operation should document the metadata type, if any.
	Metadata *any.Any `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If true, the operation is completed, and either `error` or `response` is
	// available.
	Done bool `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	// The operation result, which can be either an `error` or a valid `response`.
	// If `done` == `false`, neither `error` nor `response` is set.
	// If `done` == `true`, exactly one of `error` or `response` is set.
	//
	// Types that are valid to be assigned to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result               isOperation_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_26d6b580d5a1337d, []int{0}
}
func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (dst *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(dst, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}
type Operation_Response struct {
	Response *any.Any `protobuf:"bytes,5,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result()    {}
func (*Operation_Response) isOperation_Result() {}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Operation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Operation) GetMetadata() *any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Operation) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Operation) GetError() *status.Status {
	if x, ok := m.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (m *Operation) GetResponse() *any.Any {
	if x, ok := m.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *Operation_Response:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Result has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 4: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(status.Status)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Error{msg}
		return true, err
	case 5: // result.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.Result = &Operation_Response{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// result
	switch x := m.Result.(type) {
	case *Operation_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_Response:
		s := proto.Size(x.Response)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Operations.GetOperation][argcv.longrunning.Operations.GetOperation].
type GetOperationRequest struct {
	// The name of the operation resource.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationRequest) Reset()         { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()    {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_26d6b580d5a1337d, []int{1}
}
func (m *GetOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationRequest.Unmarshal(m, b)
}
func (m *GetOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationRequest.Marshal(b, m, deterministic)
}
func (dst *GetOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationRequest.Merge(dst, src)
}
func (m *GetOperationRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationRequest.Size(m)
}
func (m *GetOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationRequest proto.InternalMessageInfo

func (m *GetOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.ListOperations][argcv.longrunning.Operations.ListOperations].
type ListOperationsRequest struct {
	// The name of the operation collection.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The standard list filter.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The standard list page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The standard list page token.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationsRequest) Reset()         { *m = ListOperationsRequest{} }
func (m *ListOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListOperationsRequest) ProtoMessage()    {}
func (*ListOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_26d6b580d5a1337d, []int{2}
}
func (m *ListOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationsRequest.Unmarshal(m, b)
}
func (m *ListOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationsRequest.Marshal(b, m, deterministic)
}
func (dst *ListOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationsRequest.Merge(dst, src)
}
func (m *ListOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListOperationsRequest.Size(m)
}
func (m *ListOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationsRequest proto.InternalMessageInfo

func (m *ListOperationsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListOperationsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListOperationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListOperationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for [Operations.ListOperations][argcv.longrunning.Operations.ListOperations].
type ListOperationsResponse struct {
	// A list of operations that matches the specified filter in the request.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// The standard List next-page token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationsResponse) Reset()         { *m = ListOperationsResponse{} }
func (m *ListOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListOperationsResponse) ProtoMessage()    {}
func (*ListOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_26d6b580d5a1337d, []int{3}
}
func (m *ListOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationsResponse.Unmarshal(m, b)
}
func (m *ListOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationsResponse.Marshal(b, m, deterministic)
}
func (dst *ListOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationsResponse.Merge(dst, src)
}
func (m *ListOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListOperationsResponse.Size(m)
}
func (m *ListOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationsResponse proto.InternalMessageInfo

func (m *ListOperationsResponse) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListOperationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request message for [Operations.CancelOperation][argcv.longrunning.Operations.CancelOperation].
type CancelOperationRequest struct {
	// The name of the operation resource to be cancelled.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOperationRequest) Reset()         { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()    {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_26d6b580d5a1337d, []int{4}
}
func (m *CancelOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOperationRequest.Unmarshal(m, b)
}
func (m *CancelOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOperationRequest.Marshal(b, m, deterministic)
}
func (dst *CancelOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOperationRequest.Merge(dst, src)
}
func (m *CancelOperationRequest) XXX_Size() int {
	return xxx_messageInfo_CancelOperationRequest.Size(m)
}
func (m *CancelOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOperationRequest proto.InternalMessageInfo

func (m *CancelOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for [Operations.DeleteOperation][argcv.longrunning.Operations.DeleteOperation].
type DeleteOperationRequest struct {
	// The name of the operation resource to be deleted.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOperationRequest) Reset()         { *m = DeleteOperationRequest{} }
func (m *DeleteOperationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOperationRequest) ProtoMessage()    {}
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_26d6b580d5a1337d, []int{5}
}
func (m *DeleteOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOperationRequest.Unmarshal(m, b)
}
func (m *DeleteOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOperationRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOperationRequest.Merge(dst, src)
}
func (m *DeleteOperationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOperationRequest.Size(m)
}
func (m *DeleteOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOperationRequest proto.InternalMessageInfo

func (m *DeleteOperationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Operation)(nil), "argcv.longrunning.Operation")
	proto.RegisterType((*GetOperationRequest)(nil), "argcv.longrunning.GetOperationRequest")
	proto.RegisterType((*ListOperationsRequest)(nil), "argcv.longrunning.ListOperationsRequest")
	proto.RegisterType((*ListOperationsResponse)(nil), "argcv.longrunning.ListOperationsResponse")
	proto.RegisterType((*CancelOperationRequest)(nil), "argcv.longrunning.CancelOperationRequest")
	proto.RegisterType((*DeleteOperationRequest)(nil), "argcv.longrunning.DeleteOperationRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperationsClient is the client API for Operations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationsClient interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	//
	// NOTE: the `name` binding below allows API services to override the binding
	// to use different resource name schemes, such as `users/*/operations`.
	ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `argcv.error.Code.UNIMPLEMENTED`.
	DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `argcv.error.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][argcv.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][argcv.longrunning.Operation.error] value with a [argcv.status.Status.code][argcv.status.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error) {
	out := new(ListOperationsResponse)
	err := c.cc.Invoke(ctx, "/argcv.longrunning.Operations/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/argcv.longrunning.Operations/GetOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/argcv.longrunning.Operations/DeleteOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/argcv.longrunning.Operations/CancelOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServer is the server API for Operations service.
type OperationsServer interface {
	// Lists operations that match the specified filter in the request. If the
	// server doesn't support this method, it returns `UNIMPLEMENTED`.
	//
	// NOTE: the `name` binding below allows API services to override the binding
	// to use different resource name schemes, such as `users/*/operations`.
	ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error)
	// Gets the latest state of a long-running operation.  Clients can use this
	// method to poll the operation result at intervals as recommended by the API
	// service.
	GetOperation(context.Context, *GetOperationRequest) (*Operation, error)
	// Deletes a long-running operation. This method indicates that the client is
	// no longer interested in the operation result. It does not cancel the
	// operation. If the server doesn't support this method, it returns
	// `argcv.error.Code.UNIMPLEMENTED`.
	DeleteOperation(context.Context, *DeleteOperationRequest) (*empty.Empty, error)
	// Starts asynchronous cancellation on a long-running operation.  The server
	// makes a best effort to cancel the operation, but success is not
	// guaranteed.  If the server doesn't support this method, it returns
	// `argcv.error.Code.UNIMPLEMENTED`.  Clients can use
	// [Operations.GetOperation][argcv.longrunning.Operations.GetOperation] or
	// other methods to check whether the cancellation succeeded or whether the
	// operation completed despite cancellation. On successful cancellation,
	// the operation is not deleted; instead, it becomes an operation with
	// an [Operation.error][argcv.longrunning.Operation.error] value with a [argcv.status.Status.code][argcv.status.Status.code] of 1,
	// corresponding to `Code.CANCELLED`.
	CancelOperation(context.Context, *CancelOperationRequest) (*empty.Empty, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/argcv.longrunning.Operations/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).ListOperations(ctx, req.(*ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_GetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).GetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/argcv.longrunning.Operations/GetOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).GetOperation(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DeleteOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DeleteOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/argcv.longrunning.Operations/DeleteOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DeleteOperation(ctx, req.(*DeleteOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/argcv.longrunning.Operations/CancelOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).CancelOperation(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "argcv.longrunning.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOperations",
			Handler:    _Operations_ListOperations_Handler,
		},
		{
			MethodName: "GetOperation",
			Handler:    _Operations_GetOperation_Handler,
		},
		{
			MethodName: "DeleteOperation",
			Handler:    _Operations_DeleteOperation_Handler,
		},
		{
			MethodName: "CancelOperation",
			Handler:    _Operations_CancelOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "argcv/proto/longrunning/operations.proto",
}

func init() {
	proto.RegisterFile("argcv/proto/longrunning/operations.proto", fileDescriptor_operations_26d6b580d5a1337d)
}

var fileDescriptor_operations_26d6b580d5a1337d = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0xb4, 0x49, 0x94, 0xdc, 0x7e, 0x1f, 0x11, 0x03, 0x0d, 0xc6, 0xa1, 0x6a, 0xb0, 0x50,
	0xe5, 0x46, 0xc1, 0x86, 0x20, 0x36, 0x05, 0x16, 0x04, 0x10, 0x5d, 0x20, 0x11, 0xb9, 0xac, 0xd8,
	0x54, 0x93, 0xf4, 0xd6, 0x18, 0x9c, 0x19, 0x63, 0x8f, 0x23, 0x5a, 0xc4, 0x8f, 0x10, 0xe2, 0x05,
	0x78, 0x32, 0xc4, 0x2b, 0xf0, 0x10, 0x2c, 0x91, 0xc7, 0x4e, 0x6c, 0x52, 0x07, 0x65, 0xe5, 0x9b,
	0xb9, 0x67, 0xce, 0x99, 0x73, 0xe6, 0x66, 0xc0, 0x64, 0xa1, 0x3b, 0x9e, 0xda, 0x41, 0x28, 0xa4,
	0xb0, 0x7d, 0xc1, 0xdd, 0x30, 0xe6, 0xdc, 0xe3, 0xae, 0x2d, 0x02, 0x0c, 0x99, 0xf4, 0x04, 0x8f,
	0x2c, 0xd5, 0xa4, 0x17, 0x15, 0xd2, 0x2a, 0x60, 0xf4, 0xeb, 0xc5, 0xcd, 0x2c, 0xf0, 0x6c, 0xc6,
	0xb9, 0x90, 0xc5, 0x5d, 0xfa, 0x4e, 0x11, 0x12, 0x49, 0x26, 0xe3, 0x28, 0xfb, 0x64, 0x80, 0xab,
	0xae, 0x10, 0xae, 0x8f, 0x29, 0x62, 0x14, 0x9f, 0xd8, 0x8c, 0x9f, 0x66, 0xad, 0xf6, 0x62, 0x0b,
	0x27, 0x81, 0xcc, 0x9a, 0xc6, 0x0f, 0x02, 0x8d, 0xe7, 0xb3, 0x33, 0x52, 0x0a, 0x15, 0xce, 0x26,
	0xa8, 0x91, 0x0e, 0x31, 0x1b, 0x8e, 0xaa, 0xe9, 0x2d, 0xa8, 0x4f, 0x50, 0xb2, 0x63, 0x26, 0x99,
	0xb6, 0xde, 0x21, 0xe6, 0x66, 0xff, 0xb2, 0x95, 0x32, 0x5a, 0x33, 0x46, 0xeb, 0x21, 0x3f, 0x75,
	0xe6, 0xa8, 0x84, 0xe5, 0x58, 0x70, 0xd4, 0x36, 0x3a, 0xc4, 0xac, 0x3b, 0xaa, 0xa6, 0x3d, 0xa8,
	0x62, 0x18, 0x8a, 0x50, 0xab, 0x64, 0x14, 0x69, 0x0c, 0x99, 0x87, 0x43, 0xf5, 0x39, 0x58, 0x73,
	0x52, 0x10, 0xed, 0x43, 0x3d, 0xc4, 0x28, 0x10, 0x3c, 0x42, 0xad, 0xba, 0x5c, 0xf3, 0x60, 0xcd,
	0x99, 0xe3, 0x06, 0x75, 0xa8, 0x85, 0x18, 0xc5, 0xbe, 0x34, 0xf6, 0xe0, 0xd2, 0x53, 0x94, 0x73,
	0x57, 0x0e, 0xbe, 0x8d, 0x31, 0x92, 0x65, 0xe6, 0x8c, 0x4f, 0xb0, 0xf5, 0xcc, 0x8b, 0x72, 0x6c,
	0xb4, 0x08, 0xae, 0x14, 0x92, 0x68, 0x41, 0xed, 0xc4, 0xf3, 0x25, 0x86, 0x19, 0x45, 0xf6, 0x8b,
	0xb6, 0xa1, 0x11, 0x30, 0x17, 0x8f, 0x22, 0xef, 0x0c, 0x55, 0x44, 0x55, 0xa7, 0x9e, 0x2c, 0x1c,
	0x7a, 0x67, 0x48, 0xb7, 0x01, 0x54, 0x53, 0x8a, 0x37, 0xc8, 0x55, 0x24, 0x0d, 0x47, 0xc1, 0x5f,
	0x24, 0x0b, 0xc6, 0x47, 0x68, 0x2d, 0x1e, 0x20, 0xf5, 0x43, 0xef, 0x03, 0xe4, 0xc3, 0xa3, 0x91,
	0xce, 0x86, 0xb9, 0xd9, 0xbf, 0x66, 0x9d, 0x9b, 0x1e, 0x2b, 0xf7, 0x59, 0xc0, 0xd3, 0x5d, 0x68,
	0x72, 0x7c, 0x27, 0x8f, 0x0a, 0xda, 0xeb, 0x4a, 0xfb, 0xff, 0x64, 0x79, 0x38, 0xd7, 0xef, 0x41,
	0xeb, 0x11, 0xe3, 0x63, 0xf4, 0x57, 0x8a, 0xab, 0x07, 0xad, 0xc7, 0xe8, 0xa3, 0xc4, 0x55, 0xd0,
	0xfd, 0x6f, 0x15, 0x80, 0xdc, 0x18, 0xfd, 0x4a, 0xe0, 0xc2, 0xdf, 0x5e, 0xa9, 0x59, 0xe2, 0xa7,
	0xf4, 0x3e, 0xf4, 0xbd, 0x15, 0x90, 0x69, 0x70, 0xc6, 0xf6, 0x97, 0x9f, 0xbf, 0xbe, 0xaf, 0x5f,
	0xa1, 0x5b, 0xf6, 0xf4, 0xb6, 0xfd, 0x3e, 0x39, 0xc9, 0x83, 0x3c, 0x98, 0x0f, 0x74, 0x0a, 0xff,
	0x15, 0xa7, 0x83, 0xee, 0x96, 0x30, 0x97, 0x8c, 0x8f, 0xfe, 0xcf, 0xec, 0x8d, 0x8e, 0x12, 0xd5,
	0xa9, 0x56, 0x26, 0x6a, 0x77, 0xbb, 0x89, 0x6e, 0x73, 0x21, 0x3b, 0x5a, 0x66, 0xaa, 0x3c, 0x5f,
	0xbd, 0x75, 0x6e, 0xfe, 0x9f, 0x24, 0xff, 0xe2, 0x99, 0x6e, 0x77, 0xb9, 0xee, 0x67, 0x02, 0xcd,
	0x85, 0x2b, 0x2e, 0x15, 0x2e, 0x1f, 0x83, 0xa5, 0xc2, 0x5d, 0x25, 0x7c, 0xc3, 0xd8, 0x59, 0x26,
	0xbc, 0x3f, 0x56, 0x84, 0xfb, 0xa4, 0x3b, 0x78, 0x0d, 0xed, 0xb1, 0x98, 0x64, 0x9a, 0x8a, 0xa7,
	0xa8, 0x3c, 0x68, 0xe6, 0x97, 0x38, 0x4c, 0x9a, 0x43, 0xf2, 0xf2, 0xae, 0xeb, 0xc9, 0x57, 0xf1,
	0xc8, 0x1a, 0x8b, 0x89, 0x9d, 0x3e, 0x7d, 0xae, 0xb8, 0xa9, 0x0a, 0x16, 0x78, 0x51, 0xf1, 0x85,
	0xbd, 0x57, 0xa8, 0x7f, 0x13, 0x32, 0xaa, 0x29, 0xfe, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0b, 0x20, 0xca, 0xfd, 0x92, 0x05, 0x00, 0x00,
}
