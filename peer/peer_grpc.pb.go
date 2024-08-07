// sets the .proto file syntax version

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: peer.proto

// sets the protobuf package name (i.e. definitions namespace)

package peer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PeerService_ControlStream_FullMethodName = "/border0.peer.v1.PeerService/ControlStream"
)

// PeerServiceClient is the client API for PeerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerServiceClient interface {
	ControlStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ControlStreamRequest, ControlStreamResponse], error)
}

type peerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerServiceClient(cc grpc.ClientConnInterface) PeerServiceClient {
	return &peerServiceClient{cc}
}

func (c *peerServiceClient) ControlStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ControlStreamRequest, ControlStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PeerService_ServiceDesc.Streams[0], PeerService_ControlStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ControlStreamRequest, ControlStreamResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PeerService_ControlStreamClient = grpc.BidiStreamingClient[ControlStreamRequest, ControlStreamResponse]

// PeerServiceServer is the server API for PeerService service.
// All implementations must embed UnimplementedPeerServiceServer
// for forward compatibility.
type PeerServiceServer interface {
	ControlStream(grpc.BidiStreamingServer[ControlStreamRequest, ControlStreamResponse]) error
	mustEmbedUnimplementedPeerServiceServer()
}

// UnimplementedPeerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPeerServiceServer struct{}

func (UnimplementedPeerServiceServer) ControlStream(grpc.BidiStreamingServer[ControlStreamRequest, ControlStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ControlStream not implemented")
}
func (UnimplementedPeerServiceServer) mustEmbedUnimplementedPeerServiceServer() {}
func (UnimplementedPeerServiceServer) testEmbeddedByValue()                     {}

// UnsafePeerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerServiceServer will
// result in compilation errors.
type UnsafePeerServiceServer interface {
	mustEmbedUnimplementedPeerServiceServer()
}

func RegisterPeerServiceServer(s grpc.ServiceRegistrar, srv PeerServiceServer) {
	// If the following call pancis, it indicates UnimplementedPeerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PeerService_ServiceDesc, srv)
}

func _PeerService_ControlStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerServiceServer).ControlStream(&grpc.GenericServerStream[ControlStreamRequest, ControlStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PeerService_ControlStreamServer = grpc.BidiStreamingServer[ControlStreamRequest, ControlStreamResponse]

// PeerService_ServiceDesc is the grpc.ServiceDesc for PeerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "border0.peer.v1.PeerService",
	HandlerType: (*PeerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ControlStream",
			Handler:       _PeerService_ControlStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer.proto",
}
