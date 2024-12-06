// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/partyaffairs/banner/partyaffairs_banner_service.proto

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
	Banner_ListBanner_FullMethodName   = "/partyaffairs.api.partyaffairs.banner.v1.Banner/ListBanner"
	Banner_CreateBanner_FullMethodName = "/partyaffairs.api.partyaffairs.banner.v1.Banner/CreateBanner"
	Banner_UpdateBanner_FullMethodName = "/partyaffairs.api.partyaffairs.banner.v1.Banner/UpdateBanner"
	Banner_DeleteBanner_FullMethodName = "/partyaffairs.api.partyaffairs.banner.v1.Banner/DeleteBanner"
)

// BannerClient is the client API for Banner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BannerClient interface {
	// ListBanner 获取轮播图信息列表
	ListBanner(ctx context.Context, in *ListBannerRequest, opts ...grpc.CallOption) (*ListBannerReply, error)
	// CreateBanner 创建轮播图信息
	CreateBanner(ctx context.Context, in *CreateBannerRequest, opts ...grpc.CallOption) (*CreateBannerReply, error)
	// UpdateBanner 更新轮播图信息
	UpdateBanner(ctx context.Context, in *UpdateBannerRequest, opts ...grpc.CallOption) (*UpdateBannerReply, error)
	// DeleteBanner 删除轮播图信息
	DeleteBanner(ctx context.Context, in *DeleteBannerRequest, opts ...grpc.CallOption) (*DeleteBannerReply, error)
}

type bannerClient struct {
	cc grpc.ClientConnInterface
}

func NewBannerClient(cc grpc.ClientConnInterface) BannerClient {
	return &bannerClient{cc}
}

func (c *bannerClient) ListBanner(ctx context.Context, in *ListBannerRequest, opts ...grpc.CallOption) (*ListBannerReply, error) {
	out := new(ListBannerReply)
	err := c.cc.Invoke(ctx, Banner_ListBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerClient) CreateBanner(ctx context.Context, in *CreateBannerRequest, opts ...grpc.CallOption) (*CreateBannerReply, error) {
	out := new(CreateBannerReply)
	err := c.cc.Invoke(ctx, Banner_CreateBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerClient) UpdateBanner(ctx context.Context, in *UpdateBannerRequest, opts ...grpc.CallOption) (*UpdateBannerReply, error) {
	out := new(UpdateBannerReply)
	err := c.cc.Invoke(ctx, Banner_UpdateBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerClient) DeleteBanner(ctx context.Context, in *DeleteBannerRequest, opts ...grpc.CallOption) (*DeleteBannerReply, error) {
	out := new(DeleteBannerReply)
	err := c.cc.Invoke(ctx, Banner_DeleteBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BannerServer is the server API for Banner service.
// All implementations must embed UnimplementedBannerServer
// for forward compatibility
type BannerServer interface {
	// ListBanner 获取轮播图信息列表
	ListBanner(context.Context, *ListBannerRequest) (*ListBannerReply, error)
	// CreateBanner 创建轮播图信息
	CreateBanner(context.Context, *CreateBannerRequest) (*CreateBannerReply, error)
	// UpdateBanner 更新轮播图信息
	UpdateBanner(context.Context, *UpdateBannerRequest) (*UpdateBannerReply, error)
	// DeleteBanner 删除轮播图信息
	DeleteBanner(context.Context, *DeleteBannerRequest) (*DeleteBannerReply, error)
	mustEmbedUnimplementedBannerServer()
}

// UnimplementedBannerServer must be embedded to have forward compatible implementations.
type UnimplementedBannerServer struct {
}

func (UnimplementedBannerServer) ListBanner(context.Context, *ListBannerRequest) (*ListBannerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBanner not implemented")
}
func (UnimplementedBannerServer) CreateBanner(context.Context, *CreateBannerRequest) (*CreateBannerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanner not implemented")
}
func (UnimplementedBannerServer) UpdateBanner(context.Context, *UpdateBannerRequest) (*UpdateBannerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
func (UnimplementedBannerServer) DeleteBanner(context.Context, *DeleteBannerRequest) (*DeleteBannerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (UnimplementedBannerServer) mustEmbedUnimplementedBannerServer() {}

// UnsafeBannerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BannerServer will
// result in compilation errors.
type UnsafeBannerServer interface {
	mustEmbedUnimplementedBannerServer()
}

func RegisterBannerServer(s grpc.ServiceRegistrar, srv BannerServer) {
	s.RegisterService(&Banner_ServiceDesc, srv)
}

func _Banner_ListBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServer).ListBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banner_ListBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServer).ListBanner(ctx, req.(*ListBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banner_CreateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServer).CreateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banner_CreateBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServer).CreateBanner(ctx, req.(*CreateBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banner_UpdateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServer).UpdateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banner_UpdateBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServer).UpdateBanner(ctx, req.(*UpdateBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banner_DeleteBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServer).DeleteBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banner_DeleteBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServer).DeleteBanner(ctx, req.(*DeleteBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Banner_ServiceDesc is the grpc.ServiceDesc for Banner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Banner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "partyaffairs.api.partyaffairs.banner.v1.Banner",
	HandlerType: (*BannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBanner",
			Handler:    _Banner_ListBanner_Handler,
		},
		{
			MethodName: "CreateBanner",
			Handler:    _Banner_CreateBanner_Handler,
		},
		{
			MethodName: "UpdateBanner",
			Handler:    _Banner_UpdateBanner_Handler,
		},
		{
			MethodName: "DeleteBanner",
			Handler:    _Banner_DeleteBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/partyaffairs/banner/partyaffairs_banner_service.proto",
}
