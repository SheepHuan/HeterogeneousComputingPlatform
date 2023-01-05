// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: register.proto

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

// ReisgterClient is the client API for Reisgter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReisgterClient interface {
	ResgisterDevice(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type reisgterClient struct {
	cc grpc.ClientConnInterface
}

func NewReisgterClient(cc grpc.ClientConnInterface) ReisgterClient {
	return &reisgterClient{cc}
}

func (c *reisgterClient) ResgisterDevice(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/protos.Reisgter/ResgisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReisgterServer is the server API for Reisgter service.
// All implementations must embed UnimplementedReisgterServer
// for forward compatibility
type ReisgterServer interface {
	ResgisterDevice(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedReisgterServer()
}

// UnimplementedReisgterServer must be embedded to have forward compatible implementations.
type UnimplementedReisgterServer struct {
}

func (UnimplementedReisgterServer) ResgisterDevice(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResgisterDevice not implemented")
}
func (UnimplementedReisgterServer) mustEmbedUnimplementedReisgterServer() {}

// UnsafeReisgterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReisgterServer will
// result in compilation errors.
type UnsafeReisgterServer interface {
	mustEmbedUnimplementedReisgterServer()
}

func RegisterReisgterServer(s grpc.ServiceRegistrar, srv ReisgterServer) {
	s.RegisterService(&Reisgter_ServiceDesc, srv)
}

func _Reisgter_ResgisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReisgterServer).ResgisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Reisgter/ResgisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReisgterServer).ResgisterDevice(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reisgter_ServiceDesc is the grpc.ServiceDesc for Reisgter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reisgter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Reisgter",
	HandlerType: (*ReisgterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResgisterDevice",
			Handler:    _Reisgter_ResgisterDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}