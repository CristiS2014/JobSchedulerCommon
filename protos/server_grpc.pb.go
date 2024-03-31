// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/server.proto

package protos

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

// ServerAPIClient is the client API for ServerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerAPIClient interface {
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error)
}

type serverAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewServerAPIClient(cc grpc.ClientConnInterface) ServerAPIClient {
	return &serverAPIClient{cc}
}

func (c *serverAPIClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error) {
	out := new(GetJobResponse)
	err := c.cc.Invoke(ctx, "/ServerAPI/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerAPIServer is the server API for ServerAPI service.
// All implementations must embed UnimplementedServerAPIServer
// for forward compatibility
type ServerAPIServer interface {
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)
	mustEmbedUnimplementedServerAPIServer()
}

// UnimplementedServerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedServerAPIServer struct {
}

func (UnimplementedServerAPIServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedServerAPIServer) mustEmbedUnimplementedServerAPIServer() {}

// UnsafeServerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerAPIServer will
// result in compilation errors.
type UnsafeServerAPIServer interface {
	mustEmbedUnimplementedServerAPIServer()
}

func RegisterServerAPIServer(s grpc.ServiceRegistrar, srv ServerAPIServer) {
	s.RegisterService(&ServerAPI_ServiceDesc, srv)
}

func _ServerAPI_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerAPIServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerAPI/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerAPIServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerAPI_ServiceDesc is the grpc.ServiceDesc for ServerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServerAPI",
	HandlerType: (*ServerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _ServerAPI_GetJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/server.proto",
}
