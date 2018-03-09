// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent/agent.proto

/*
Package agent is a generated protocol buffer package.

It is generated from these files:
	agent/agent.proto

It has these top-level messages:
	TunnelRequest
	TunnelResponse
*/
package agent

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

type TunnelRequest struct {
	Dial string `protobuf:"bytes,1,opt,name=dial" json:"dial,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TunnelRequest) Reset()                    { *m = TunnelRequest{} }
func (m *TunnelRequest) String() string            { return proto.CompactTextString(m) }
func (*TunnelRequest) ProtoMessage()               {}
func (*TunnelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TunnelRequest) GetDial() string {
	if m != nil {
		return m.Dial
	}
	return ""
}

func (m *TunnelRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TunnelResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TunnelResponse) Reset()                    { *m = TunnelResponse{} }
func (m *TunnelResponse) String() string            { return proto.CompactTextString(m) }
func (*TunnelResponse) ProtoMessage()               {}
func (*TunnelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TunnelResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TunnelResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TunnelRequest)(nil), "agent.TunnelRequest")
	proto.RegisterType((*TunnelResponse)(nil), "agent.TunnelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Tunnel(ctx context.Context, opts ...grpc.CallOption) (Service_TunnelClient, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Tunnel(ctx context.Context, opts ...grpc.CallOption) (Service_TunnelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Service_serviceDesc.Streams[0], c.cc, "/agent.Service/Tunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceTunnelClient{stream}
	return x, nil
}

type Service_TunnelClient interface {
	Send(*TunnelRequest) error
	Recv() (*TunnelResponse, error)
	grpc.ClientStream
}

type serviceTunnelClient struct {
	grpc.ClientStream
}

func (x *serviceTunnelClient) Send(m *TunnelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceTunnelClient) Recv() (*TunnelResponse, error) {
	m := new(TunnelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Service service

type ServiceServer interface {
	Tunnel(Service_TunnelServer) error
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Tunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).Tunnel(&serviceTunnelServer{stream})
}

type Service_TunnelServer interface {
	Send(*TunnelResponse) error
	Recv() (*TunnelRequest, error)
	grpc.ServerStream
}

type serviceTunnelServer struct {
	grpc.ServerStream
}

func (x *serviceTunnelServer) Send(m *TunnelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceTunnelServer) Recv() (*TunnelRequest, error) {
	m := new(TunnelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tunnel",
			Handler:       _Service_Tunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent/agent.proto",
}

func init() { proto.RegisterFile("agent/agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x39,
	0x17, 0x6f, 0x48, 0x69, 0x5e, 0x5e, 0x6a, 0x4e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
	0x10, 0x17, 0x4b, 0x4a, 0x66, 0x62, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x0d,
	0x16, 0x4b, 0x2c, 0x49, 0x94, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0xac, 0xb8,
	0xf8, 0x60, 0x1a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x8b, 0x8a,
	0xf2, 0x8b, 0xa0, 0x5a, 0x21, 0x1c, 0x6c, 0x7a, 0x8d, 0x5c, 0xb8, 0xd8, 0x83, 0x53, 0x8b, 0xca,
	0x32, 0x93, 0x53, 0x85, 0x2c, 0xb9, 0xd8, 0x20, 0xc6, 0x08, 0x89, 0xe8, 0x41, 0x9c, 0x87, 0xe2,
	0x1c, 0x29, 0x51, 0x34, 0x51, 0x88, 0x5d, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0x8f, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x36, 0x4f, 0x60, 0x84, 0xdd, 0x00, 0x00, 0x00,
}
