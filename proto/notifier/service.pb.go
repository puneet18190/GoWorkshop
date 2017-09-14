// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/notifier/service.proto

/*
Package notifier is a generated protocol buffer package.

It is generated from these files:
	proto/notifier/service.proto

It has these top-level messages:
	NewUserRequest
	NewUserResponse
*/
package notifier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type NewUserRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *NewUserRequest) Reset()                    { *m = NewUserRequest{} }
func (m *NewUserRequest) String() string            { return proto.CompactTextString(m) }
func (*NewUserRequest) ProtoMessage()               {}
func (*NewUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type NewUserResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Code    uint64 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *NewUserResponse) Reset()                    { *m = NewUserResponse{} }
func (m *NewUserResponse) String() string            { return proto.CompactTextString(m) }
func (*NewUserResponse) ProtoMessage()               {}
func (*NewUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NewUserResponse) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*NewUserRequest)(nil), "notifier.NewUserRequest")
	proto.RegisterType((*NewUserResponse)(nil), "notifier.NewUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Notifier service

type NotifierClient interface {
	NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error)
}

type notifierClient struct {
	cc *grpc.ClientConn
}

func NewNotifierClient(cc *grpc.ClientConn) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) NewUser(ctx context.Context, in *NewUserRequest, opts ...grpc.CallOption) (*NewUserResponse, error) {
	out := new(NewUserResponse)
	err := grpc.Invoke(ctx, "/notifier.Notifier/NewUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notifier service

type NotifierServer interface {
	NewUser(context.Context, *NewUserRequest) (*NewUserResponse, error)
}

func RegisterNotifierServer(s *grpc.Server, srv NotifierServer) {
	s.RegisterService(&_Notifier_serviceDesc, srv)
}

func _Notifier_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifier/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NewUser(ctx, req.(*NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifier.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUser",
			Handler:    _Notifier_NewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/notifier/service.proto",
}

func init() { proto.RegisterFile("proto/notifier/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x2d, 0xd2, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x03, 0x0b, 0x0b, 0x71, 0xc0, 0xc4, 0x95, 0xd4, 0xb8, 0xf8, 0xfc, 0x52, 0xcb,
	0x43, 0x8b, 0x53, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x53,
	0x73, 0x13, 0x33, 0x73, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x7b, 0x2e,
	0x7e, 0xb8, 0xba, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2,
	0xe2, 0xc4, 0xf4, 0x54, 0xa8, 0x52, 0x18, 0x57, 0x48, 0x88, 0x8b, 0x25, 0x39, 0x3f, 0x25, 0x55,
	0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x08, 0xcc, 0x36, 0xf2, 0xe1, 0xe2, 0xf0, 0x83, 0x5a, 0x2a,
	0xe4, 0xc0, 0xc5, 0x0e, 0x35, 0x4c, 0x48, 0x42, 0x0f, 0xe6, 0x14, 0x3d, 0x54, 0x77, 0x48, 0x49,
	0x62, 0x91, 0x81, 0xd8, 0xac, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x07, 0x09, 0x1f, 0xe7, 0x00, 0x00, 0x00,
}
