// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/partyaffairs/activity/partyaffairs_activity_service.proto

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

const (
	Activity_GetActivity_FullMethodName    = "/partyaffairs.api.partyaffairs.activity.v1.Activity/GetActivity"
	Activity_ListActivity_FullMethodName   = "/partyaffairs.api.partyaffairs.activity.v1.Activity/ListActivity"
	Activity_CreateActivity_FullMethodName = "/partyaffairs.api.partyaffairs.activity.v1.Activity/CreateActivity"
	Activity_UpdateActivity_FullMethodName = "/partyaffairs.api.partyaffairs.activity.v1.Activity/UpdateActivity"
	Activity_DeleteActivity_FullMethodName = "/partyaffairs.api.partyaffairs.activity.v1.Activity/DeleteActivity"
)

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	// GetActivity 获取指定的活动信息
	GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityReply, error)
	// ListActivity 获取活动信息列表
	ListActivity(ctx context.Context, in *ListActivityRequest, opts ...grpc.CallOption) (*ListActivityReply, error)
	// CreateActivity 创建活动信息
	CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityReply, error)
	// UpdateActivity 更新活动信息
	UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*UpdateActivityReply, error)
	// DeleteActivity 删除活动信息
	DeleteActivity(ctx context.Context, in *DeleteActivityRequest, opts ...grpc.CallOption) (*DeleteActivityReply, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityReply, error) {
	out := new(GetActivityReply)
	err := c.cc.Invoke(ctx, Activity_GetActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) ListActivity(ctx context.Context, in *ListActivityRequest, opts ...grpc.CallOption) (*ListActivityReply, error) {
	out := new(ListActivityReply)
	err := c.cc.Invoke(ctx, Activity_ListActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityReply, error) {
	out := new(CreateActivityReply)
	err := c.cc.Invoke(ctx, Activity_CreateActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*UpdateActivityReply, error) {
	out := new(UpdateActivityReply)
	err := c.cc.Invoke(ctx, Activity_UpdateActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) DeleteActivity(ctx context.Context, in *DeleteActivityRequest, opts ...grpc.CallOption) (*DeleteActivityReply, error) {
	out := new(DeleteActivityReply)
	err := c.cc.Invoke(ctx, Activity_DeleteActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	// GetActivity 获取指定的活动信息
	GetActivity(context.Context, *GetActivityRequest) (*GetActivityReply, error)
	// ListActivity 获取活动信息列表
	ListActivity(context.Context, *ListActivityRequest) (*ListActivityReply, error)
	// CreateActivity 创建活动信息
	CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityReply, error)
	// UpdateActivity 更新活动信息
	UpdateActivity(context.Context, *UpdateActivityRequest) (*UpdateActivityReply, error)
	// DeleteActivity 删除活动信息
	DeleteActivity(context.Context, *DeleteActivityRequest) (*DeleteActivityReply, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) GetActivity(context.Context, *GetActivityRequest) (*GetActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedActivityServer) ListActivity(context.Context, *ListActivityRequest) (*ListActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivity not implemented")
}
func (UnimplementedActivityServer) CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedActivityServer) UpdateActivity(context.Context, *UpdateActivityRequest) (*UpdateActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivity not implemented")
}
func (UnimplementedActivityServer) DeleteActivity(context.Context, *DeleteActivityRequest) (*DeleteActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivity not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_GetActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivity(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_ListActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).ListActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_ListActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).ListActivity(ctx, req.(*ListActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_CreateActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_UpdateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).UpdateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_UpdateActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).UpdateActivity(ctx, req.(*UpdateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_DeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).DeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_DeleteActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).DeleteActivity(ctx, req.(*DeleteActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partyaffairs.api.partyaffairs.activity.v1.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActivity",
			Handler:    _Activity_GetActivity_Handler,
		},
		{
			MethodName: "ListActivity",
			Handler:    _Activity_ListActivity_Handler,
		},
		{
			MethodName: "CreateActivity",
			Handler:    _Activity_CreateActivity_Handler,
		},
		{
			MethodName: "UpdateActivity",
			Handler:    _Activity_UpdateActivity_Handler,
		},
		{
			MethodName: "DeleteActivity",
			Handler:    _Activity_DeleteActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/partyaffairs/activity/partyaffairs_activity_service.proto",
}
