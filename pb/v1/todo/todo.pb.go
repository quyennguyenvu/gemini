// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ToDo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reminder             string   `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToDo) Reset()         { *m = ToDo{} }
func (m *ToDo) String() string { return proto.CompactTextString(m) }
func (*ToDo) ProtoMessage()    {}
func (*ToDo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_e99543ea72ad93ec, []int{0}
}
func (m *ToDo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDo.Unmarshal(m, b)
}
func (m *ToDo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDo.Marshal(b, m, deterministic)
}
func (dst *ToDo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDo.Merge(dst, src)
}
func (m *ToDo) XXX_Size() int {
	return xxx_messageInfo_ToDo.Size(m)
}
func (m *ToDo) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDo.DiscardUnknown(m)
}

var xxx_messageInfo_ToDo proto.InternalMessageInfo

func (m *ToDo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToDo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDo) GetReminder() string {
	if m != nil {
		return m.Reminder
	}
	return ""
}

type ListRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Keyword              string   `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_e99543ea72ad93ec, []int{1}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type ListResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ToDos                []*ToDo  `protobuf:"bytes,2,rep,name=toDos,proto3" json:"toDos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_e99543ea72ad93ec, []int{2}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListResponse) GetToDos() []*ToDo {
	if m != nil {
		return m.ToDos
	}
	return nil
}

type StoreRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reminder             string   `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_e99543ea72ad93ec, []int{3}
}
func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (dst *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(dst, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *StoreRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StoreRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *StoreRequest) GetReminder() string {
	if m != nil {
		return m.Reminder
	}
	return ""
}

type StoreResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_e99543ea72ad93ec, []int{4}
}
func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (dst *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(dst, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *StoreResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoreResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ToDo)(nil), "todo.ToDo")
	proto.RegisterType((*ListRequest)(nil), "todo.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "todo.ListResponse")
	proto.RegisterType((*StoreRequest)(nil), "todo.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "todo.StoreResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/todo.Todo/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/todo.Todo/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Todo_List_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Todo_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_todo_e99543ea72ad93ec) }

var fileDescriptor_todo_e99543ea72ad93ec = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x55, 0xd2, 0xf4, 0xb5, 0xef, 0xb6, 0x45, 0xe5, 0x96, 0xc1, 0x8a, 0x18, 0xa2, 0x4c, 0x15,
	0x43, 0x23, 0xca, 0x04, 0x0b, 0x12, 0x2a, 0x13, 0x4c, 0x81, 0x1f, 0x08, 0xb5, 0x55, 0x19, 0x4a,
	0x6e, 0xb0, 0x5d, 0x10, 0x2b, 0x1f, 0xc0, 0xc2, 0xa7, 0xf1, 0x0b, 0x7c, 0x08, 0xb2, 0x9d, 0x46,
	0x41, 0xa2, 0x23, 0x9b, 0xcf, 0x39, 0xf6, 0x39, 0xf7, 0x9e, 0x04, 0xc0, 0x10, 0xa7, 0x59, 0xa5,
	0xc8, 0x10, 0x46, 0xf6, 0x1c, 0x1f, 0xae, 0x88, 0x56, 0x6b, 0x91, 0x15, 0x95, 0xcc, 0x8a, 0xb2,
	0x24, 0x53, 0x18, 0x49, 0xa5, 0xf6, 0x77, 0xd2, 0x7b, 0x88, 0x6e, 0x69, 0x41, 0xb8, 0x07, 0xa1,
	0xe4, 0x2c, 0x48, 0x82, 0xe9, 0xff, 0x3c, 0x94, 0x1c, 0x0f, 0xa0, 0x6b, 0xa4, 0x59, 0x0b, 0x16,
	0x3a, 0xca, 0x03, 0x4c, 0x60, 0xc0, 0x85, 0x5e, 0x2a, 0x59, 0x59, 0x0f, 0xd6, 0x71, 0x5a, 0x9b,
	0xc2, 0x18, 0xfa, 0x4a, 0x3c, 0xca, 0x92, 0x0b, 0xc5, 0x22, 0x27, 0x37, 0x38, 0x3d, 0x85, 0xc1,
	0xb5, 0xd4, 0x26, 0x17, 0x4f, 0x1b, 0xa1, 0x0d, 0x8e, 0xa1, 0x53, 0x54, 0xb2, 0xce, 0xb4, 0x47,
	0x64, 0xd0, 0x7b, 0x10, 0xaf, 0x2f, 0xa4, 0x78, 0x1d, 0xbb, 0x85, 0xe9, 0x05, 0x0c, 0xfd, 0x53,
	0x5d, 0x51, 0xa9, 0xc5, 0x2f, 0x6f, 0x13, 0xe8, 0x1a, 0x5a, 0x90, 0x66, 0x61, 0xd2, 0x99, 0x0e,
	0xe6, 0x30, 0x73, 0x45, 0xd8, 0xdd, 0x72, 0x2f, 0xa4, 0x06, 0x86, 0x37, 0x86, 0x94, 0xd8, 0x9d,
	0xff, 0x17, 0x4b, 0x5f, 0xc1, 0xa8, 0x4e, 0xdd, 0x39, 0xba, 0xef, 0x3e, 0x6c, 0xba, 0x67, 0xd0,
	0xd3, 0x9b, 0xe5, 0x52, 0x68, 0xed, 0xc2, 0xfa, 0xf9, 0x16, 0xce, 0xdf, 0x03, 0xfb, 0xb9, 0x38,
	0xe1, 0x39, 0x44, 0xb6, 0x0f, 0xdc, 0xf7, 0x6b, 0xb6, 0x6a, 0x8d, 0xb1, 0x4d, 0xf9, 0xcc, 0x74,
	0xfc, 0xf6, 0xf9, 0xf5, 0x11, 0x02, 0xf6, 0xb3, 0xe7, 0xe3, 0xcc, 0xca, 0x78, 0x09, 0x5d, 0x37,
	0x16, 0xd6, 0xd7, 0xdb, 0xcd, 0xc4, 0x93, 0x1f, 0x5c, 0xed, 0x31, 0x71, 0x1e, 0xa3, 0xb4, 0xf1,
	0x38, 0x0b, 0x8e, 0xee, 0xfe, 0xb9, 0xbf, 0xe8, 0xe4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x28, 0x90,
	0xf6, 0xc7, 0x77, 0x02, 0x00, 0x00,
}
