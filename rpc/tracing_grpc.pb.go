// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: tracing.proto

package rpc

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

// TerwayTracingClient is the client API for TerwayTracing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TerwayTracingClient interface {
	GetResourceTypes(ctx context.Context, in *Placeholder, opts ...grpc.CallOption) (*ResourcesTypesReply, error)
	GetResources(ctx context.Context, in *ResourceTypeRequest, opts ...grpc.CallOption) (*ResourcesNamesReply, error)
	GetResourceConfig(ctx context.Context, in *ResourceTypeNameRequest, opts ...grpc.CallOption) (*ResourceConfigReply, error)
	GetResourceTrace(ctx context.Context, in *ResourceTypeNameRequest, opts ...grpc.CallOption) (*ResourceTraceReply, error)
	ResourceExecute(ctx context.Context, in *ResourceExecuteRequest, opts ...grpc.CallOption) (TerwayTracing_ResourceExecuteClient, error)
	GetResourceMapping(ctx context.Context, in *Placeholder, opts ...grpc.CallOption) (*PodResourceMappingReply, error)
}

type terwayTracingClient struct {
	cc grpc.ClientConnInterface
}

func NewTerwayTracingClient(cc grpc.ClientConnInterface) TerwayTracingClient {
	return &terwayTracingClient{cc}
}

func (c *terwayTracingClient) GetResourceTypes(ctx context.Context, in *Placeholder, opts ...grpc.CallOption) (*ResourcesTypesReply, error) {
	out := new(ResourcesTypesReply)
	err := c.cc.Invoke(ctx, "/rpc.TerwayTracing/GetResourceTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terwayTracingClient) GetResources(ctx context.Context, in *ResourceTypeRequest, opts ...grpc.CallOption) (*ResourcesNamesReply, error) {
	out := new(ResourcesNamesReply)
	err := c.cc.Invoke(ctx, "/rpc.TerwayTracing/GetResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terwayTracingClient) GetResourceConfig(ctx context.Context, in *ResourceTypeNameRequest, opts ...grpc.CallOption) (*ResourceConfigReply, error) {
	out := new(ResourceConfigReply)
	err := c.cc.Invoke(ctx, "/rpc.TerwayTracing/GetResourceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terwayTracingClient) GetResourceTrace(ctx context.Context, in *ResourceTypeNameRequest, opts ...grpc.CallOption) (*ResourceTraceReply, error) {
	out := new(ResourceTraceReply)
	err := c.cc.Invoke(ctx, "/rpc.TerwayTracing/GetResourceTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terwayTracingClient) ResourceExecute(ctx context.Context, in *ResourceExecuteRequest, opts ...grpc.CallOption) (TerwayTracing_ResourceExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &TerwayTracing_ServiceDesc.Streams[0], "/rpc.TerwayTracing/ResourceExecute", opts...)
	if err != nil {
		return nil, err
	}
	x := &terwayTracingResourceExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TerwayTracing_ResourceExecuteClient interface {
	Recv() (*ResourceExecuteReply, error)
	grpc.ClientStream
}

type terwayTracingResourceExecuteClient struct {
	grpc.ClientStream
}

func (x *terwayTracingResourceExecuteClient) Recv() (*ResourceExecuteReply, error) {
	m := new(ResourceExecuteReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *terwayTracingClient) GetResourceMapping(ctx context.Context, in *Placeholder, opts ...grpc.CallOption) (*PodResourceMappingReply, error) {
	out := new(PodResourceMappingReply)
	err := c.cc.Invoke(ctx, "/rpc.TerwayTracing/GetResourceMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TerwayTracingServer is the server API for TerwayTracing service.
// All implementations must embed UnimplementedTerwayTracingServer
// for forward compatibility
type TerwayTracingServer interface {
	GetResourceTypes(context.Context, *Placeholder) (*ResourcesTypesReply, error)
	GetResources(context.Context, *ResourceTypeRequest) (*ResourcesNamesReply, error)
	GetResourceConfig(context.Context, *ResourceTypeNameRequest) (*ResourceConfigReply, error)
	GetResourceTrace(context.Context, *ResourceTypeNameRequest) (*ResourceTraceReply, error)
	ResourceExecute(*ResourceExecuteRequest, TerwayTracing_ResourceExecuteServer) error
	GetResourceMapping(context.Context, *Placeholder) (*PodResourceMappingReply, error)
	mustEmbedUnimplementedTerwayTracingServer()
}

// UnimplementedTerwayTracingServer must be embedded to have forward compatible implementations.
type UnimplementedTerwayTracingServer struct {
}

func (UnimplementedTerwayTracingServer) GetResourceTypes(context.Context, *Placeholder) (*ResourcesTypesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceTypes not implemented")
}
func (UnimplementedTerwayTracingServer) GetResources(context.Context, *ResourceTypeRequest) (*ResourcesNamesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedTerwayTracingServer) GetResourceConfig(context.Context, *ResourceTypeNameRequest) (*ResourceConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceConfig not implemented")
}
func (UnimplementedTerwayTracingServer) GetResourceTrace(context.Context, *ResourceTypeNameRequest) (*ResourceTraceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceTrace not implemented")
}
func (UnimplementedTerwayTracingServer) ResourceExecute(*ResourceExecuteRequest, TerwayTracing_ResourceExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method ResourceExecute not implemented")
}
func (UnimplementedTerwayTracingServer) GetResourceMapping(context.Context, *Placeholder) (*PodResourceMappingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceMapping not implemented")
}
func (UnimplementedTerwayTracingServer) mustEmbedUnimplementedTerwayTracingServer() {}

// UnsafeTerwayTracingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TerwayTracingServer will
// result in compilation errors.
type UnsafeTerwayTracingServer interface {
	mustEmbedUnimplementedTerwayTracingServer()
}

func RegisterTerwayTracingServer(s grpc.ServiceRegistrar, srv TerwayTracingServer) {
	s.RegisterService(&TerwayTracing_ServiceDesc, srv)
}

func _TerwayTracing_GetResourceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Placeholder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerwayTracingServer).GetResourceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TerwayTracing/GetResourceTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerwayTracingServer).GetResourceTypes(ctx, req.(*Placeholder))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerwayTracing_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerwayTracingServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TerwayTracing/GetResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerwayTracingServer).GetResources(ctx, req.(*ResourceTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerwayTracing_GetResourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerwayTracingServer).GetResourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TerwayTracing/GetResourceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerwayTracingServer).GetResourceConfig(ctx, req.(*ResourceTypeNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerwayTracing_GetResourceTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceTypeNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerwayTracingServer).GetResourceTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TerwayTracing/GetResourceTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerwayTracingServer).GetResourceTrace(ctx, req.(*ResourceTypeNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerwayTracing_ResourceExecute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResourceExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TerwayTracingServer).ResourceExecute(m, &terwayTracingResourceExecuteServer{stream})
}

type TerwayTracing_ResourceExecuteServer interface {
	Send(*ResourceExecuteReply) error
	grpc.ServerStream
}

type terwayTracingResourceExecuteServer struct {
	grpc.ServerStream
}

func (x *terwayTracingResourceExecuteServer) Send(m *ResourceExecuteReply) error {
	return x.ServerStream.SendMsg(m)
}

func _TerwayTracing_GetResourceMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Placeholder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerwayTracingServer).GetResourceMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TerwayTracing/GetResourceMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerwayTracingServer).GetResourceMapping(ctx, req.(*Placeholder))
	}
	return interceptor(ctx, in, info, handler)
}

// TerwayTracing_ServiceDesc is the grpc.ServiceDesc for TerwayTracing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TerwayTracing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.TerwayTracing",
	HandlerType: (*TerwayTracingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResourceTypes",
			Handler:    _TerwayTracing_GetResourceTypes_Handler,
		},
		{
			MethodName: "GetResources",
			Handler:    _TerwayTracing_GetResources_Handler,
		},
		{
			MethodName: "GetResourceConfig",
			Handler:    _TerwayTracing_GetResourceConfig_Handler,
		},
		{
			MethodName: "GetResourceTrace",
			Handler:    _TerwayTracing_GetResourceTrace_Handler,
		},
		{
			MethodName: "GetResourceMapping",
			Handler:    _TerwayTracing_GetResourceMapping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ResourceExecute",
			Handler:       _TerwayTracing_ResourceExecute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tracing.proto",
}
