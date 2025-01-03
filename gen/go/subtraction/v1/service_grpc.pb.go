// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: subtraction/v1/service.proto

package subtractionpb

import (
	context "context"
	factorial "example/gen/go/x/factorial"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SubtractionService_Subtract_FullMethodName  = "/subtraction.v1.SubtractionService/Subtract"
	SubtractionService_Factorial_FullMethodName = "/subtraction.v1.SubtractionService/Factorial"
)

// SubtractionServiceClient is the client API for SubtractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubtractionServiceClient interface {
	Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error)
	Factorial(ctx context.Context, in *factorial.FactorialRequest, opts ...grpc.CallOption) (*factorial.FactorialResponse, error)
}

type subtractionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubtractionServiceClient(cc grpc.ClientConnInterface) SubtractionServiceClient {
	return &subtractionServiceClient{cc}
}

func (c *subtractionServiceClient) Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error) {
	out := new(SubtractResponse)
	err := c.cc.Invoke(ctx, SubtractionService_Subtract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subtractionServiceClient) Factorial(ctx context.Context, in *factorial.FactorialRequest, opts ...grpc.CallOption) (*factorial.FactorialResponse, error) {
	out := new(factorial.FactorialResponse)
	err := c.cc.Invoke(ctx, SubtractionService_Factorial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubtractionServiceServer is the server API for SubtractionService service.
// All implementations should embed UnimplementedSubtractionServiceServer
// for forward compatibility
type SubtractionServiceServer interface {
	Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error)
	Factorial(context.Context, *factorial.FactorialRequest) (*factorial.FactorialResponse, error)
}

// UnimplementedSubtractionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubtractionServiceServer struct {
}

func (UnimplementedSubtractionServiceServer) Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedSubtractionServiceServer) Factorial(context.Context, *factorial.FactorialRequest) (*factorial.FactorialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Factorial not implemented")
}

// UnsafeSubtractionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubtractionServiceServer will
// result in compilation errors.
type UnsafeSubtractionServiceServer interface {
	mustEmbedUnimplementedSubtractionServiceServer()
}

func RegisterSubtractionServiceServer(s grpc.ServiceRegistrar, srv SubtractionServiceServer) {
	s.RegisterService(&SubtractionService_ServiceDesc, srv)
}

func _SubtractionService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubtractionServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubtractionService_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubtractionServiceServer).Subtract(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubtractionService_Factorial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(factorial.FactorialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubtractionServiceServer).Factorial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubtractionService_Factorial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubtractionServiceServer).Factorial(ctx, req.(*factorial.FactorialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubtractionService_ServiceDesc is the grpc.ServiceDesc for SubtractionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubtractionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subtraction.v1.SubtractionService",
	HandlerType: (*SubtractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subtract",
			Handler:    _SubtractionService_Subtract_Handler,
		},
		{
			MethodName: "Factorial",
			Handler:    _SubtractionService_Factorial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subtraction/v1/service.proto",
}
