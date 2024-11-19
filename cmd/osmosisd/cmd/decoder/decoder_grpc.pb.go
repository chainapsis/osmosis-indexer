// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: decoder.proto

package decoder

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
	CosmosDecoder_Decode_FullMethodName = "/decoder.CosmosDecoder/Decode"
)

// CosmosDecoderClient is the client API for CosmosDecoder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CosmosDecoderClient interface {
	Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeResponse, error)
}

type cosmosDecoderClient struct {
	cc grpc.ClientConnInterface
}

func NewCosmosDecoderClient(cc grpc.ClientConnInterface) CosmosDecoderClient {
	return &cosmosDecoderClient{cc}
}

func (c *cosmosDecoderClient) Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeResponse, error) {
	out := new(DecodeResponse)
	err := c.cc.Invoke(ctx, CosmosDecoder_Decode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CosmosDecoderServer is the server API for CosmosDecoder service.
// All implementations must embed UnimplementedCosmosDecoderServer
// for forward compatibility
type CosmosDecoderServer interface {
	Decode(context.Context, *DecodeRequest) (*DecodeResponse, error)
	mustEmbedUnimplementedCosmosDecoderServer()
}

// UnimplementedCosmosDecoderServer must be embedded to have forward compatible implementations.
type UnimplementedCosmosDecoderServer struct {
}

func (UnimplementedCosmosDecoderServer) Decode(context.Context, *DecodeRequest) (*DecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}
func (UnimplementedCosmosDecoderServer) mustEmbedUnimplementedCosmosDecoderServer() {}

// UnsafeCosmosDecoderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CosmosDecoderServer will
// result in compilation errors.
type UnsafeCosmosDecoderServer interface {
	mustEmbedUnimplementedCosmosDecoderServer()
}

func RegisterCosmosDecoderServer(s grpc.ServiceRegistrar, srv CosmosDecoderServer) {
	s.RegisterService(&CosmosDecoder_ServiceDesc, srv)
}

func _CosmosDecoder_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosmosDecoderServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CosmosDecoder_Decode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosmosDecoderServer).Decode(ctx, req.(*DecodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CosmosDecoder_ServiceDesc is the grpc.ServiceDesc for CosmosDecoder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CosmosDecoder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "decoder.CosmosDecoder",
	HandlerType: (*CosmosDecoderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decode",
			Handler:    _CosmosDecoder_Decode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decoder.proto",
}
