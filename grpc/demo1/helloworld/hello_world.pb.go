// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/razeencheng/demo-go/grpc/demo1/helloworld/hello_world.proto

package helloworld // import "github.com/razeencheng/demo-go/grpc/demo1/helloworld"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type HelloWorldRequest struct {
	Greeting             string            `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Infos                map[string]string `protobuf:"bytes,2,rep,name=infos,proto3" json:"infos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_5bbefe74e8f48f56, []int{0}
}
func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldRequest.Unmarshal(m, b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
}
func (dst *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(dst, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return xxx_messageInfo_HelloWorldRequest.Size(m)
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *HelloWorldRequest) GetInfos() map[string]string {
	if m != nil {
		return m.Infos
	}
	return nil
}

type HelloWorldResponse struct {
	Reply                string     `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	Details              []*any.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_world_5bbefe74e8f48f56, []int{1}
}
func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloWorldResponse.Unmarshal(m, b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
}
func (dst *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(dst, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return xxx_messageInfo_HelloWorldResponse.Size(m)
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *HelloWorldResponse) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "helloworld.HelloWorldRequest")
	proto.RegisterMapType((map[string]string)(nil), "helloworld.HelloWorldRequest.InfosEntry")
	proto.RegisterType((*HelloWorldResponse)(nil), "helloworld.HelloWorldResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloWorldServiceClient is the client API for HelloWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloWorldServiceClient interface {
	SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type helloWorldServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldServiceClient(cc *grpc.ClientConn) HelloWorldServiceClient {
	return &helloWorldServiceClient{cc}
}

func (c *helloWorldServiceClient) SayHelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorldService/SayHelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServiceServer is the server API for HelloWorldService service.
type HelloWorldServiceServer interface {
	SayHelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
}

func RegisterHelloWorldServiceServer(s *grpc.Server, srv HelloWorldServiceServer) {
	s.RegisterService(&_HelloWorldService_serviceDesc, srv)
}

func _HelloWorldService_SayHelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorldService/SayHelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServiceServer).SayHelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorldService",
	HandlerType: (*HelloWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloWorld",
			Handler:    _HelloWorldService_SayHelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/razeencheng/demo-go/grpc/demo1/helloworld/hello_world.proto",
}

func init() {
	proto.RegisterFile("github.com/razeencheng/demo-go/grpc/demo1/helloworld/hello_world.proto", fileDescriptor_hello_world_5bbefe74e8f48f56)
}

var fileDescriptor_hello_world_5bbefe74e8f48f56 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0xd2, 0xd2, 0x4f, 0x1d, 0x11, 0x74, 0xe9, 0xa1, 0x04, 0x94, 0xd2, 0x53, 0x2f, 0xee,
	0x62, 0x15, 0x29, 0x1e, 0x04, 0x05, 0x45, 0x6f, 0x92, 0x1e, 0x84, 0x5e, 0x24, 0x4d, 0xa7, 0xdb,
	0xe0, 0x76, 0x27, 0x6e, 0x36, 0x95, 0xf8, 0x8f, 0xfc, 0x97, 0xb2, 0x59, 0x63, 0x0a, 0xa2, 0x07,
	0x0f, 0x81, 0x79, 0x33, 0xef, 0xcd, 0x7b, 0x99, 0x85, 0x5b, 0x99, 0xda, 0x65, 0x31, 0xe3, 0x09,
	0xad, 0x84, 0x89, 0xdf, 0x10, 0x75, 0xb2, 0x44, 0x2d, 0xc5, 0x1c, 0x57, 0x74, 0x2c, 0x49, 0x48,
	0x93, 0x25, 0x15, 0x38, 0x11, 0x4b, 0x54, 0x8a, 0x5e, 0xc9, 0xa8, 0xb9, 0x2f, 0x9f, 0xaa, 0x9a,
	0x67, 0x86, 0x2c, 0x31, 0x68, 0xa6, 0xa1, 0xd8, 0xd8, 0x29, 0x49, 0xc5, 0x5a, 0x8a, 0x8a, 0x34,
	0x2b, 0x16, 0x22, 0xb3, 0x65, 0x86, 0xb9, 0x88, 0x75, 0xe9, 0x3e, 0x2f, 0x1e, 0xbc, 0x07, 0x70,
	0x70, 0xe7, 0xf4, 0x8f, 0x4e, 0x1f, 0xe1, 0x4b, 0x81, 0xb9, 0x65, 0x21, 0x6c, 0x4b, 0x83, 0x68,
	0x53, 0x2d, 0x7b, 0x41, 0x3f, 0x18, 0xee, 0x44, 0x5f, 0x98, 0x5d, 0x42, 0x27, 0xd5, 0x0b, 0xca,
	0x7b, 0xad, 0x7e, 0x7b, 0xb8, 0x3b, 0x1a, 0xf2, 0xc6, 0x9e, 0x7f, 0xdb, 0xc4, 0xef, 0x1d, 0xf5,
	0x46, 0x5b, 0x53, 0x46, 0x5e, 0x16, 0x8e, 0x01, 0x9a, 0x26, 0xdb, 0x87, 0xf6, 0x33, 0x96, 0x9f,
	0x26, 0xae, 0x64, 0x5d, 0xe8, 0xac, 0x63, 0x55, 0x60, 0xaf, 0x55, 0xf5, 0x3c, 0xb8, 0x68, 0x8d,
	0x83, 0xc1, 0x14, 0xd8, 0xa6, 0x41, 0x9e, 0x91, 0xce, 0xd1, 0xf1, 0x0d, 0x66, 0xaa, 0xde, 0xe1,
	0x01, 0xe3, 0xb0, 0x35, 0x47, 0x1b, 0xa7, 0xaa, 0xce, 0xd9, 0xe5, 0x92, 0x48, 0x2a, 0xe4, 0xf5,
	0x3d, 0xf8, 0x95, 0x2e, 0xa3, 0x9a, 0x34, 0xc2, 0xcd, 0x33, 0x4c, 0xd0, 0xac, 0xd3, 0x04, 0xd9,
	0x03, 0xec, 0x4d, 0xe2, 0xb2, 0xe9, 0xb3, 0xc3, 0x5f, 0x7f, 0x36, 0x3c, 0xfa, 0x69, 0xec, 0xa3,
	0x0e, 0xfe, 0x5d, 0x9f, 0x4f, 0xcf, 0xfe, 0xf2, 0xea, 0xb3, 0xff, 0x55, 0xea, 0xd3, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x38, 0x6b, 0xfd, 0x41, 0x34, 0x02, 0x00, 0x00,
}
