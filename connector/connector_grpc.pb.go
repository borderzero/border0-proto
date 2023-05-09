// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: connector.proto

package connector

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConnectorServiceClient is the client API for ConnectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type connectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorServiceClient(cc grpc.ClientConnInterface) ConnectorServiceClient {
	return &connectorServiceClient{cc}
}

func (c *connectorServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/connector.ConnectorService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorServiceServer is the server API for ConnectorService service.
// All implementations must embed UnimplementedConnectorServiceServer
// for forward compatibility
type ConnectorServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedConnectorServiceServer()
}

// UnimplementedConnectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorServiceServer struct {
}

func (UnimplementedConnectorServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedConnectorServiceServer) mustEmbedUnimplementedConnectorServiceServer() {}

// UnsafeConnectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorServiceServer will
// result in compilation errors.
type UnsafeConnectorServiceServer interface {
	mustEmbedUnimplementedConnectorServiceServer()
}

func RegisterConnectorServiceServer(s grpc.ServiceRegistrar, srv ConnectorServiceServer) {
	s.RegisterService(&ConnectorService_ServiceDesc, srv)
}

func _ConnectorService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connector.ConnectorService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectorService_ServiceDesc is the grpc.ServiceDesc for ConnectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connector.ConnectorService",
	HandlerType: (*ConnectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ConnectorService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connector.proto",
}
