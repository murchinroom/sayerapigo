// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: muvtuber/sayer/v1/sayer.proto

// Sayer is the TTS (Text To Speech) service. It is used to convert text to
// speech

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

const (
	SayerService_Say_FullMethodName = "/muvtuber.sayer.v1.SayerService/Say"
)

// SayerServiceClient is the client API for SayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SayerServiceClient interface {
	// Say converts text to speech and returns the audio file.
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type sayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSayerServiceClient(cc grpc.ClientConnInterface) SayerServiceClient {
	return &sayerServiceClient{cc}
}

func (c *sayerServiceClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, SayerService_Say_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayerServiceServer is the server API for SayerService service.
// All implementations must embed UnimplementedSayerServiceServer
// for forward compatibility
type SayerServiceServer interface {
	// Say converts text to speech and returns the audio file.
	Say(context.Context, *SayRequest) (*SayResponse, error)
	mustEmbedUnimplementedSayerServiceServer()
}

// UnimplementedSayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSayerServiceServer struct {
}

func (UnimplementedSayerServiceServer) Say(context.Context, *SayRequest) (*SayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}
func (UnimplementedSayerServiceServer) mustEmbedUnimplementedSayerServiceServer() {}

// UnsafeSayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SayerServiceServer will
// result in compilation errors.
type UnsafeSayerServiceServer interface {
	mustEmbedUnimplementedSayerServiceServer()
}

func RegisterSayerServiceServer(s grpc.ServiceRegistrar, srv SayerServiceServer) {
	s.RegisterService(&SayerService_ServiceDesc, srv)
}

func _SayerService_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayerServiceServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SayerService_Say_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayerServiceServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SayerService_ServiceDesc is the grpc.ServiceDesc for SayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "muvtuber.sayer.v1.SayerService",
	HandlerType: (*SayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _SayerService_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "muvtuber/sayer/v1/sayer.proto",
}
