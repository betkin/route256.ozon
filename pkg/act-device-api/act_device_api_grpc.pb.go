// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package act_device_api

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

// ActDeviceApiServiceClient is the client API for ActDeviceApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActDeviceApiServiceClient interface {
	// CreateDeviceV1 - Create a device
	CreateDeviceV1(ctx context.Context, in *CreateDeviceV1Request, opts ...grpc.CallOption) (*CreateDeviceV1Response, error)
	// DescribeDeviceV1 - Describe a device
	DescribeDeviceV1(ctx context.Context, in *DescribeDeviceV1Request, opts ...grpc.CallOption) (*DescribeDeviceV1Response, error)
	// LogDeviceV1 - List of device events
	LogDeviceV1(ctx context.Context, in *LogDeviceV1Request, opts ...grpc.CallOption) (*LogDeviceV1Response, error)
	// ListDevicesV1 - List of devices
	ListDevicesV1(ctx context.Context, in *ListDevicesV1Request, opts ...grpc.CallOption) (*ListDevicesV1Response, error)
	// UpdateDeviceV1 - Update a device
	UpdateDeviceV1(ctx context.Context, in *UpdateDeviceV1Request, opts ...grpc.CallOption) (*UpdateDeviceV1Response, error)
	// RemoveDeviceV1 - Remove a device
	RemoveDeviceV1(ctx context.Context, in *RemoveDeviceV1Request, opts ...grpc.CallOption) (*RemoveDeviceV1Response, error)
}

type actDeviceApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActDeviceApiServiceClient(cc grpc.ClientConnInterface) ActDeviceApiServiceClient {
	return &actDeviceApiServiceClient{cc}
}

func (c *actDeviceApiServiceClient) CreateDeviceV1(ctx context.Context, in *CreateDeviceV1Request, opts ...grpc.CallOption) (*CreateDeviceV1Response, error) {
	out := new(CreateDeviceV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.act_device_api.v1.ActDeviceApiService/CreateDeviceV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actDeviceApiServiceClient) DescribeDeviceV1(ctx context.Context, in *DescribeDeviceV1Request, opts ...grpc.CallOption) (*DescribeDeviceV1Response, error) {
	out := new(DescribeDeviceV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.act_device_api.v1.ActDeviceApiService/DescribeDeviceV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actDeviceApiServiceClient) LogDeviceV1(ctx context.Context, in *LogDeviceV1Request, opts ...grpc.CallOption) (*LogDeviceV1Response, error) {
	out := new(LogDeviceV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.act_device_api.v1.ActDeviceApiService/LogDeviceV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actDeviceApiServiceClient) ListDevicesV1(ctx context.Context, in *ListDevicesV1Request, opts ...grpc.CallOption) (*ListDevicesV1Response, error) {
	out := new(ListDevicesV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.act_device_api.v1.ActDeviceApiService/ListDevicesV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actDeviceApiServiceClient) UpdateDeviceV1(ctx context.Context, in *UpdateDeviceV1Request, opts ...grpc.CallOption) (*UpdateDeviceV1Response, error) {
	out := new(UpdateDeviceV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.act_device_api.v1.ActDeviceApiService/UpdateDeviceV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actDeviceApiServiceClient) RemoveDeviceV1(ctx context.Context, in *RemoveDeviceV1Request, opts ...grpc.CallOption) (*RemoveDeviceV1Response, error) {
	out := new(RemoveDeviceV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.act_device_api.v1.ActDeviceApiService/RemoveDeviceV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActDeviceApiServiceServer is the server API for ActDeviceApiService service.
// All implementations must embed UnimplementedActDeviceApiServiceServer
// for forward compatibility
type ActDeviceApiServiceServer interface {
	// CreateDeviceV1 - Create a device
	CreateDeviceV1(context.Context, *CreateDeviceV1Request) (*CreateDeviceV1Response, error)
	// DescribeDeviceV1 - Describe a device
	DescribeDeviceV1(context.Context, *DescribeDeviceV1Request) (*DescribeDeviceV1Response, error)
	// LogDeviceV1 - List of device events
	LogDeviceV1(context.Context, *LogDeviceV1Request) (*LogDeviceV1Response, error)
	// ListDevicesV1 - List of devices
	ListDevicesV1(context.Context, *ListDevicesV1Request) (*ListDevicesV1Response, error)
	// UpdateDeviceV1 - Update a device
	UpdateDeviceV1(context.Context, *UpdateDeviceV1Request) (*UpdateDeviceV1Response, error)
	// RemoveDeviceV1 - Remove a device
	RemoveDeviceV1(context.Context, *RemoveDeviceV1Request) (*RemoveDeviceV1Response, error)
	mustEmbedUnimplementedActDeviceApiServiceServer()
}

// UnimplementedActDeviceApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActDeviceApiServiceServer struct {
}

func (UnimplementedActDeviceApiServiceServer) CreateDeviceV1(context.Context, *CreateDeviceV1Request) (*CreateDeviceV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceV1 not implemented")
}
func (UnimplementedActDeviceApiServiceServer) DescribeDeviceV1(context.Context, *DescribeDeviceV1Request) (*DescribeDeviceV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeDeviceV1 not implemented")
}
func (UnimplementedActDeviceApiServiceServer) LogDeviceV1(context.Context, *LogDeviceV1Request) (*LogDeviceV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogDeviceV1 not implemented")
}
func (UnimplementedActDeviceApiServiceServer) ListDevicesV1(context.Context, *ListDevicesV1Request) (*ListDevicesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevicesV1 not implemented")
}
func (UnimplementedActDeviceApiServiceServer) UpdateDeviceV1(context.Context, *UpdateDeviceV1Request) (*UpdateDeviceV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceV1 not implemented")
}
func (UnimplementedActDeviceApiServiceServer) RemoveDeviceV1(context.Context, *RemoveDeviceV1Request) (*RemoveDeviceV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDeviceV1 not implemented")
}
func (UnimplementedActDeviceApiServiceServer) mustEmbedUnimplementedActDeviceApiServiceServer() {}

// UnsafeActDeviceApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActDeviceApiServiceServer will
// result in compilation errors.
type UnsafeActDeviceApiServiceServer interface {
	mustEmbedUnimplementedActDeviceApiServiceServer()
}

func RegisterActDeviceApiServiceServer(s grpc.ServiceRegistrar, srv ActDeviceApiServiceServer) {
	s.RegisterService(&ActDeviceApiService_ServiceDesc, srv)
}

func _ActDeviceApiService_CreateDeviceV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActDeviceApiServiceServer).CreateDeviceV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.act_device_api.v1.ActDeviceApiService/CreateDeviceV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActDeviceApiServiceServer).CreateDeviceV1(ctx, req.(*CreateDeviceV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActDeviceApiService_DescribeDeviceV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDeviceV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActDeviceApiServiceServer).DescribeDeviceV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.act_device_api.v1.ActDeviceApiService/DescribeDeviceV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActDeviceApiServiceServer).DescribeDeviceV1(ctx, req.(*DescribeDeviceV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActDeviceApiService_LogDeviceV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogDeviceV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActDeviceApiServiceServer).LogDeviceV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.act_device_api.v1.ActDeviceApiService/LogDeviceV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActDeviceApiServiceServer).LogDeviceV1(ctx, req.(*LogDeviceV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActDeviceApiService_ListDevicesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActDeviceApiServiceServer).ListDevicesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.act_device_api.v1.ActDeviceApiService/ListDevicesV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActDeviceApiServiceServer).ListDevicesV1(ctx, req.(*ListDevicesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActDeviceApiService_UpdateDeviceV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActDeviceApiServiceServer).UpdateDeviceV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.act_device_api.v1.ActDeviceApiService/UpdateDeviceV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActDeviceApiServiceServer).UpdateDeviceV1(ctx, req.(*UpdateDeviceV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActDeviceApiService_RemoveDeviceV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActDeviceApiServiceServer).RemoveDeviceV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.act_device_api.v1.ActDeviceApiService/RemoveDeviceV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActDeviceApiServiceServer).RemoveDeviceV1(ctx, req.(*RemoveDeviceV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ActDeviceApiService_ServiceDesc is the grpc.ServiceDesc for ActDeviceApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActDeviceApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.act_device_api.v1.ActDeviceApiService",
	HandlerType: (*ActDeviceApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeviceV1",
			Handler:    _ActDeviceApiService_CreateDeviceV1_Handler,
		},
		{
			MethodName: "DescribeDeviceV1",
			Handler:    _ActDeviceApiService_DescribeDeviceV1_Handler,
		},
		{
			MethodName: "LogDeviceV1",
			Handler:    _ActDeviceApiService_LogDeviceV1_Handler,
		},
		{
			MethodName: "ListDevicesV1",
			Handler:    _ActDeviceApiService_ListDevicesV1_Handler,
		},
		{
			MethodName: "UpdateDeviceV1",
			Handler:    _ActDeviceApiService_UpdateDeviceV1_Handler,
		},
		{
			MethodName: "RemoveDeviceV1",
			Handler:    _ActDeviceApiService_RemoveDeviceV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/act_device_api/v1/act_device_api.proto",
}
