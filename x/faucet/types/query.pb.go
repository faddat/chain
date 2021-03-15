// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: faucet/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("faucet/query.proto", fileDescriptor_32e01ab1e3e8ff22) }

var fileDescriptor_32e01ab1e3e8ff22 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0x31, 0x8e, 0xc2, 0x30,
	0x10, 0x85, 0xe1, 0xa4, 0xd8, 0x5d, 0x29, 0x65, 0xaa, 0x55, 0xb4, 0xf2, 0x01, 0x16, 0xc9, 0xa3,
	0xc0, 0x0d, 0x90, 0x38, 0x00, 0x2d, 0xdd, 0xd8, 0x18, 0xc7, 0x22, 0xf1, 0x98, 0xd8, 0x41, 0xe4,
	0x16, 0x1c, 0x8b, 0x32, 0x25, 0x25, 0x4a, 0x2e, 0x82, 0x88, 0x29, 0xe8, 0xbf, 0xf7, 0xf4, 0x67,
	0xf9, 0x01, 0x3b, 0xa9, 0x02, 0x9c, 0x3a, 0xd5, 0xf6, 0xdc, 0xb5, 0x14, 0x28, 0xff, 0xf5, 0x68,
	0xf7, 0xa2, 0x26, 0x79, 0x34, 0xc4, 0x65, 0x85, 0xc6, 0xf2, 0xa8, 0x8a, 0x3f, 0x4d, 0xa4, 0x6b,
	0x05, 0xe8, 0x0c, 0xa0, 0xb5, 0x14, 0x30, 0x18, 0xb2, 0x3e, 0xee, 0x8a, 0x7f, 0x49, 0xbe, 0x21,
	0x0f, 0x02, 0xbd, 0x8a, 0x87, 0x70, 0x2e, 0x85, 0x0a, 0x58, 0x82, 0x43, 0x6d, 0xec, 0x8c, 0xa3,
	0x5d, 0xfe, 0x64, 0x5f, 0xdb, 0x97, 0x58, 0x6f, 0x6e, 0x23, 0x4b, 0x87, 0x91, 0xa5, 0x8f, 0x91,
	0xa5, 0xd7, 0x89, 0x25, 0xc3, 0xc4, 0x92, 0xfb, 0xc4, 0x92, 0xdd, 0x42, 0x9b, 0x50, 0x75, 0x82,
	0x4b, 0x6a, 0xe0, 0xa3, 0x08, 0xe6, 0x22, 0xb8, 0xc0, 0xbb, 0x3c, 0xf4, 0x4e, 0x79, 0xf1, 0x3d,
	0xdf, 0xae, 0x9e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xa4, 0x5c, 0xed, 0xd0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sandblockio.chain.faucet.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "faucet/query.proto",
}