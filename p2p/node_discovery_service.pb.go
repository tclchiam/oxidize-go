// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_discovery_service.proto

/*
Package p2p is a generated protocol buffer package.

It is generated from these files:
	node_discovery_service.proto

It has these top-level messages:
	PingRequest
	PingResponse
	VersionRequest
	VersionResponse
*/
package p2p

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

type PingRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type VersionRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type VersionResponse struct {
	LatestHash       []byte  `protobuf:"bytes,1,req,name=latestHash" json:"latestHash,omitempty"`
	LatestIndex      *uint64 `protobuf:"varint,2,req,name=latestIndex" json:"latestIndex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *VersionResponse) GetLatestHash() []byte {
	if m != nil {
		return m.LatestHash
	}
	return nil
}

func (m *VersionResponse) GetLatestIndex() uint64 {
	if m != nil && m.LatestIndex != nil {
		return *m.LatestIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "p2p.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "p2p.PingResponse")
	proto.RegisterType((*VersionRequest)(nil), "p2p.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "p2p.VersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DiscoveryService service

type DiscoveryServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type discoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryServiceClient(cc *grpc.ClientConn) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/p2p.DiscoveryService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/p2p.DiscoveryService/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DiscoveryService service

type DiscoveryServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterDiscoveryServiceServer(s *grpc.Server, srv DiscoveryServiceServer) {
	s.RegisterService(&_DiscoveryService_serviceDesc, srv)
}

func _DiscoveryService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.DiscoveryService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p2p.DiscoveryService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "p2p.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DiscoveryService_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _DiscoveryService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_discovery_service.proto",
}

func init() { proto.RegisterFile("node_discovery_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xcb, 0x4f, 0x49,
	0x8d, 0x4f, 0xc9, 0x2c, 0x4e, 0xce, 0x2f, 0x4b, 0x2d, 0xaa, 0x8c, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x30, 0x2a, 0x50, 0xe2, 0xe5,
	0xe2, 0x0e, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0xe2, 0xe3, 0xe2,
	0x81, 0x70, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x04, 0xb8, 0xf8, 0xc2, 0x52, 0x8b, 0x8a,
	0x33, 0xf3, 0xf3, 0x60, 0x2a, 0x82, 0xb9, 0xf8, 0xe1, 0x22, 0x10, 0x45, 0x42, 0x72, 0x5c, 0x5c,
	0x39, 0x89, 0x25, 0xa9, 0xc5, 0x25, 0x1e, 0x89, 0xc5, 0x19, 0x12, 0x8c, 0x0a, 0x4c, 0x1a, 0x3c,
	0x41, 0x48, 0x22, 0x42, 0x0a, 0x5c, 0xdc, 0x10, 0x9e, 0x67, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0x93,
	0x02, 0x93, 0x06, 0x4b, 0x10, 0xb2, 0x90, 0x51, 0x29, 0x97, 0x80, 0x0b, 0xcc, 0x95, 0xc1, 0x10,
	0x47, 0x0a, 0x69, 0x73, 0xb1, 0x80, 0x9c, 0x22, 0x24, 0xa0, 0x57, 0x60, 0x54, 0xa0, 0x87, 0xe4,
	0x48, 0x29, 0x41, 0x24, 0x11, 0xa8, 0x13, 0x4c, 0xb8, 0xd8, 0xa1, 0xae, 0x12, 0x12, 0x06, 0xcb,
	0xa2, 0xba, 0x5a, 0x4a, 0x04, 0x55, 0x10, 0xa2, 0xcb, 0x89, 0x35, 0x0a, 0x14, 0x06, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x01, 0xd9, 0x60, 0x27, 0x01, 0x00, 0x00,
}