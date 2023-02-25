// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: service/v1/service.proto

package servicev1

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

// MessageAPIServiceClient is the client API for MessageAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageAPIServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error)
}

type messageAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageAPIServiceClient(cc grpc.ClientConnInterface) MessageAPIServiceClient {
	return &messageAPIServiceClient{cc}
}

func (c *messageAPIServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/service.v1.MessageAPIService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageAPIServiceClient) Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error) {
	out := new(ReverseResponse)
	err := c.cc.Invoke(ctx, "/service.v1.MessageAPIService/Reverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageAPIServiceServer is the server API for MessageAPIService service.
// All implementations should embed UnimplementedMessageAPIServiceServer
// for forward compatibility
type MessageAPIServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error)
}

// UnimplementedMessageAPIServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMessageAPIServiceServer struct {
}

func (UnimplementedMessageAPIServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedMessageAPIServiceServer) Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reverse not implemented")
}

// UnsafeMessageAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageAPIServiceServer will
// result in compilation errors.
type UnsafeMessageAPIServiceServer interface {
	mustEmbedUnimplementedMessageAPIServiceServer()
}

func RegisterMessageAPIServiceServer(s grpc.ServiceRegistrar, srv MessageAPIServiceServer) {
	s.RegisterService(&MessageAPIService_ServiceDesc, srv)
}

func _MessageAPIService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageAPIServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.MessageAPIService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageAPIServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageAPIService_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageAPIServiceServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.MessageAPIService/Reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageAPIServiceServer).Reverse(ctx, req.(*ReverseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageAPIService_ServiceDesc is the grpc.ServiceDesc for MessageAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.v1.MessageAPIService",
	HandlerType: (*MessageAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _MessageAPIService_Echo_Handler,
		},
		{
			MethodName: "Reverse",
			Handler:    _MessageAPIService_Reverse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/v1/service.proto",
}
