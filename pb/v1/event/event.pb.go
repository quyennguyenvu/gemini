// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package event

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

type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eca54a60a7c08ce8, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Event) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
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
	return fileDescriptor_event_eca54a60a7c08ce8, []int{1}
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
	Events               []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eca54a60a7c08ce8, []int{2}
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

func (m *ListResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type StoreRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_eca54a60a7c08ce8, []int{3}
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

func (m *StoreRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *StoreRequest) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
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
	return fileDescriptor_event_eca54a60a7c08ce8, []int{4}
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
	proto.RegisterType((*Event)(nil), "event.Event")
	proto.RegisterType((*ListRequest)(nil), "event.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "event.ListResponse")
	proto.RegisterType((*StoreRequest)(nil), "event.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "event.StoreResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/event.EventService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/event.EventService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _EventService_List_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _EventService_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_event_eca54a60a7c08ce8) }

var fileDescriptor_event_eca54a60a7c08ce8 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0xe9, 0x9f, 0x9d, 0xc4, 0xa2, 0x63, 0x0f, 0xa1, 0x28, 0x94, 0xc5, 0x43, 0xf1, 0xd0,
	0x60, 0x3d, 0xe9, 0x4d, 0x41, 0x11, 0xf5, 0x94, 0x3e, 0x80, 0x6c, 0x93, 0xa1, 0x2c, 0x4a, 0x36,
	0x66, 0xb7, 0x95, 0x5e, 0x7d, 0x05, 0xaf, 0xbe, 0x95, 0xaf, 0xe0, 0x83, 0x48, 0x27, 0x5b, 0x6d,
	0xc1, 0x5e, 0xbc, 0xcd, 0x7c, 0x33, 0xf3, 0xed, 0xf7, 0x7d, 0x2c, 0x04, 0x34, 0xa7, 0xdc, 0x0e,
	0x8b, 0x52, 0x5b, 0x8d, 0x0d, 0x6e, 0x7a, 0x87, 0x53, 0xad, 0xa7, 0xcf, 0x14, 0xcb, 0x42, 0xc5,
	0x32, 0xcf, 0xb5, 0x95, 0x56, 0xe9, 0xdc, 0x54, 0x4b, 0xe2, 0x0e, 0x1a, 0xd7, 0xcb, 0x35, 0xec,
	0x80, 0xaf, 0xb2, 0xc8, 0xeb, 0x7b, 0x83, 0x76, 0xe2, 0xab, 0x0c, 0x11, 0xea, 0x13, 0x9d, 0x2d,
	0x22, 0x9f, 0x11, 0xae, 0xf1, 0x08, 0x20, 0x2d, 0x49, 0x5a, 0xca, 0x1e, 0xa5, 0x8d, 0x6a, 0x3c,
	0x69, 0x3b, 0xe4, 0xd2, 0x8a, 0x73, 0x08, 0x1e, 0x94, 0xb1, 0x09, 0xbd, 0xcc, 0xc8, 0x58, 0xdc,
	0x83, 0x9a, 0x2c, 0x94, 0xa3, 0x5c, 0x96, 0x18, 0x41, 0xeb, 0x89, 0x16, 0xaf, 0xba, 0xcc, 0x1c,
	0xed, 0xaa, 0x15, 0x37, 0x10, 0x56, 0xa7, 0xa6, 0xd0, 0xb9, 0xa1, 0x3f, 0x6e, 0x8f, 0xa1, 0xc9,
	0x7e, 0x4c, 0xe4, 0xf7, 0x6b, 0x83, 0x60, 0x14, 0x0e, 0x2b, 0xaf, 0xac, 0x3e, 0x71, 0x33, 0x31,
	0x86, 0x70, 0x6c, 0x75, 0x49, 0xdb, 0x35, 0xfc, 0xc3, 0xd7, 0x3d, 0xec, 0x3a, 0xd2, 0xad, 0xea,
	0xaa, 0xf4, 0xfc, 0x9f, 0xf4, 0x22, 0x68, 0x99, 0x59, 0x9a, 0x92, 0x31, 0x4c, 0xb7, 0x93, 0xac,
	0xda, 0xd1, 0x87, 0x07, 0x21, 0x6b, 0x1e, 0x53, 0x39, 0x57, 0x29, 0xe1, 0x15, 0xd4, 0x97, 0xd6,
	0x11, 0x9d, 0xa1, 0xb5, 0x08, 0x7b, 0x07, 0x1b, 0x58, 0xf5, 0xba, 0xd8, 0x7f, 0xfb, 0xfc, 0x7a,
	0xf7, 0x03, 0x6c, 0xc7, 0xf3, 0xd3, 0x98, 0xe7, 0x78, 0x0b, 0x0d, 0x56, 0x88, 0xab, 0x83, 0xf5,
	0x10, 0x7a, 0xdd, 0x4d, 0xd0, 0xd1, 0x74, 0x99, 0xa6, 0x23, 0x7e, 0x69, 0x2e, 0xbc, 0x93, 0x49,
	0x93, 0xbf, 0xc5, 0xd9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0xe9, 0xc6, 0x15, 0x4a, 0x02,
	0x00, 0x00,
}
