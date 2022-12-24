// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: math_message.proto

package stub

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

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/controllergrpc.MathService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
// All implementations must embed UnimplementedMathServiceServer
// for forward compatibility
type MathServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	mustEmbedUnimplementedMathServiceServer()
}

// UnimplementedMathServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (UnimplementedMathServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedMathServiceServer) mustEmbedUnimplementedMathServiceServer() {}

// UnsafeMathServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServiceServer will
// result in compilation errors.
type UnsafeMathServiceServer interface {
	mustEmbedUnimplementedMathServiceServer()
}

func RegisterMathServiceServer(s grpc.ServiceRegistrar, srv MathServiceServer) {
	s.RegisterService(&MathService_ServiceDesc, srv)
}

func _MathService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controllergrpc.MathService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MathService_ServiceDesc is the grpc.ServiceDesc for MathService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MathService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controllergrpc.MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _MathService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "math_message.proto",
}