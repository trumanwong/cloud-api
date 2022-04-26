// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/instance/v1/instance.proto

package v1

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

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceClient interface {
	// 查询可以使用的阿里云地域
	ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error)
	// 查询可以使用的镜像资源
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	// 选择实例规格
	ListInstanceTypes(ctx context.Context, in *ListInstanceTypesRequest, opts ...grpc.CallOption) (*ListInstanceTypesResponse, error)
	// 创建实例
	CreateInstances(ctx context.Context, in *CreateInstancesRequest, opts ...grpc.CallOption) (*CreateInstancesResponse, error)
	// 实例列表
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// 开机
	StartInstances(ctx context.Context, in *StartInstancesRequest, opts ...grpc.CallOption) (*StartInstancesResponse, error)
	// 关机
	StopInstances(ctx context.Context, in *StopInstancesRequest, opts ...grpc.CallOption) (*StopInstancesResponse, error)
	// 重启
	RebootInstances(ctx context.Context, in *RebootInstancesRequest, opts ...grpc.CallOption) (*RebootInstancesResponse, error)
	// 删除实例
	DeleteInstances(ctx context.Context, in *DeleteInstancesRequest, opts ...grpc.CallOption) (*DeleteInstancesResponse, error)
}

type instanceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceClient(cc grpc.ClientConnInterface) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) ListRegions(ctx context.Context, in *ListRegionsRequest, opts ...grpc.CallOption) (*ListRegionsResponse, error) {
	out := new(ListRegionsResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/ListRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/ListImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) ListInstanceTypes(ctx context.Context, in *ListInstanceTypesRequest, opts ...grpc.CallOption) (*ListInstanceTypesResponse, error) {
	out := new(ListInstanceTypesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/ListInstanceTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) CreateInstances(ctx context.Context, in *CreateInstancesRequest, opts ...grpc.CallOption) (*CreateInstancesResponse, error) {
	out := new(CreateInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/CreateInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) StartInstances(ctx context.Context, in *StartInstancesRequest, opts ...grpc.CallOption) (*StartInstancesResponse, error) {
	out := new(StartInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/StartInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) StopInstances(ctx context.Context, in *StopInstancesRequest, opts ...grpc.CallOption) (*StopInstancesResponse, error) {
	out := new(StopInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/StopInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) RebootInstances(ctx context.Context, in *RebootInstancesRequest, opts ...grpc.CallOption) (*RebootInstancesResponse, error) {
	out := new(RebootInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/RebootInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) DeleteInstances(ctx context.Context, in *DeleteInstancesRequest, opts ...grpc.CallOption) (*DeleteInstancesResponse, error) {
	out := new(DeleteInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.instance.v1.Instance/DeleteInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServer is the server API for Instance service.
// All implementations must embed UnimplementedInstanceServer
// for forward compatibility
type InstanceServer interface {
	// 查询可以使用的阿里云地域
	ListRegions(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error)
	// 查询可以使用的镜像资源
	ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	// 选择实例规格
	ListInstanceTypes(context.Context, *ListInstanceTypesRequest) (*ListInstanceTypesResponse, error)
	// 创建实例
	CreateInstances(context.Context, *CreateInstancesRequest) (*CreateInstancesResponse, error)
	// 实例列表
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// 开机
	StartInstances(context.Context, *StartInstancesRequest) (*StartInstancesResponse, error)
	// 关机
	StopInstances(context.Context, *StopInstancesRequest) (*StopInstancesResponse, error)
	// 重启
	RebootInstances(context.Context, *RebootInstancesRequest) (*RebootInstancesResponse, error)
	// 删除实例
	DeleteInstances(context.Context, *DeleteInstancesRequest) (*DeleteInstancesResponse, error)
	mustEmbedUnimplementedInstanceServer()
}

// UnimplementedInstanceServer must be embedded to have forward compatible implementations.
type UnimplementedInstanceServer struct {
}

func (UnimplementedInstanceServer) ListRegions(context.Context, *ListRegionsRequest) (*ListRegionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegions not implemented")
}
func (UnimplementedInstanceServer) ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedInstanceServer) ListInstanceTypes(context.Context, *ListInstanceTypesRequest) (*ListInstanceTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstanceTypes not implemented")
}
func (UnimplementedInstanceServer) CreateInstances(context.Context, *CreateInstancesRequest) (*CreateInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstances not implemented")
}
func (UnimplementedInstanceServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedInstanceServer) StartInstances(context.Context, *StartInstancesRequest) (*StartInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartInstances not implemented")
}
func (UnimplementedInstanceServer) StopInstances(context.Context, *StopInstancesRequest) (*StopInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopInstances not implemented")
}
func (UnimplementedInstanceServer) RebootInstances(context.Context, *RebootInstancesRequest) (*RebootInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebootInstances not implemented")
}
func (UnimplementedInstanceServer) DeleteInstances(context.Context, *DeleteInstancesRequest) (*DeleteInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstances not implemented")
}
func (UnimplementedInstanceServer) mustEmbedUnimplementedInstanceServer() {}

// UnsafeInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServer will
// result in compilation errors.
type UnsafeInstanceServer interface {
	mustEmbedUnimplementedInstanceServer()
}

func RegisterInstanceServer(s grpc.ServiceRegistrar, srv InstanceServer) {
	s.RegisterService(&Instance_ServiceDesc, srv)
}

func _Instance_ListRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).ListRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/ListRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).ListRegions(ctx, req.(*ListRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_ListInstanceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).ListInstanceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/ListInstanceTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).ListInstanceTypes(ctx, req.(*ListInstanceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_CreateInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).CreateInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/CreateInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).CreateInstances(ctx, req.(*CreateInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_StartInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).StartInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/StartInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).StartInstances(ctx, req.(*StartInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_StopInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).StopInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/StopInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).StopInstances(ctx, req.(*StopInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_RebootInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebootInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).RebootInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/RebootInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).RebootInstances(ctx, req.(*RebootInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_DeleteInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).DeleteInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.instance.v1.Instance/DeleteInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).DeleteInstances(ctx, req.(*DeleteInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Instance_ServiceDesc is the grpc.ServiceDesc for Instance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Instance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.instance.v1.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRegions",
			Handler:    _Instance_ListRegions_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _Instance_ListImages_Handler,
		},
		{
			MethodName: "ListInstanceTypes",
			Handler:    _Instance_ListInstanceTypes_Handler,
		},
		{
			MethodName: "CreateInstances",
			Handler:    _Instance_CreateInstances_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _Instance_ListInstances_Handler,
		},
		{
			MethodName: "StartInstances",
			Handler:    _Instance_StartInstances_Handler,
		},
		{
			MethodName: "StopInstances",
			Handler:    _Instance_StopInstances_Handler,
		},
		{
			MethodName: "RebootInstances",
			Handler:    _Instance_RebootInstances_Handler,
		},
		{
			MethodName: "DeleteInstances",
			Handler:    _Instance_DeleteInstances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/instance/v1/instance.proto",
}
