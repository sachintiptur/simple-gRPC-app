// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: proto/env.proto

package proto

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

// GetterClient is the client API for Getter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetterClient interface {
	GetEnvVariable(ctx context.Context, in *EnvRequest, opts ...grpc.CallOption) (*EnvResponse, error)
}

type getterClient struct {
	cc grpc.ClientConnInterface
}

func NewGetterClient(cc grpc.ClientConnInterface) GetterClient {
	return &getterClient{cc}
}

func (c *getterClient) GetEnvVariable(ctx context.Context, in *EnvRequest, opts ...grpc.CallOption) (*EnvResponse, error) {
	out := new(EnvResponse)
	err := c.cc.Invoke(ctx, "/env.Getter/GetEnvVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetterServer is the server API for Getter service.
// All implementations must embed UnimplementedGetterServer
// for forward compatibility
type GetterServer interface {
	GetEnvVariable(context.Context, *EnvRequest) (*EnvResponse, error)
	mustEmbedUnimplementedGetterServer()
}

// UnimplementedGetterServer must be embedded to have forward compatible implementations.
type UnimplementedGetterServer struct {
}

func (UnimplementedGetterServer) GetEnvVariable(context.Context, *EnvRequest) (*EnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnvVariable not implemented")
}
func (UnimplementedGetterServer) mustEmbedUnimplementedGetterServer() {}

// UnsafeGetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetterServer will
// result in compilation errors.
type UnsafeGetterServer interface {
	mustEmbedUnimplementedGetterServer()
}

func RegisterGetterServer(s grpc.ServiceRegistrar, srv GetterServer) {
	s.RegisterService(&Getter_ServiceDesc, srv)
}

func _Getter_GetEnvVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetterServer).GetEnvVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/env.Getter/GetEnvVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetterServer).GetEnvVariable(ctx, req.(*EnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Getter_ServiceDesc is the grpc.ServiceDesc for Getter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Getter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "env.Getter",
	HandlerType: (*GetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEnvVariable",
			Handler:    _Getter_GetEnvVariable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/env.proto",
}