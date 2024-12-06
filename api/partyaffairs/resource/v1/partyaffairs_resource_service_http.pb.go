// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/partyaffairs/resource/partyaffairs_resource_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationResourceCreateResource = "/partyaffairs.api.partyaffairs.resource.v1.Resource/CreateResource"
const OperationResourceCreateResourceClassify = "/partyaffairs.api.partyaffairs.resource.v1.Resource/CreateResourceClassify"
const OperationResourceDeleteResource = "/partyaffairs.api.partyaffairs.resource.v1.Resource/DeleteResource"
const OperationResourceDeleteResourceClassify = "/partyaffairs.api.partyaffairs.resource.v1.Resource/DeleteResourceClassify"
const OperationResourceGetResource = "/partyaffairs.api.partyaffairs.resource.v1.Resource/GetResource"
const OperationResourceListResource = "/partyaffairs.api.partyaffairs.resource.v1.Resource/ListResource"
const OperationResourceListResourceClassify = "/partyaffairs.api.partyaffairs.resource.v1.Resource/ListResourceClassify"
const OperationResourceUpdateResource = "/partyaffairs.api.partyaffairs.resource.v1.Resource/UpdateResource"
const OperationResourceUpdateResourceClassify = "/partyaffairs.api.partyaffairs.resource.v1.Resource/UpdateResourceClassify"

type ResourceHTTPServer interface {
	// CreateResource CreateResource 创建咨询信息
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceReply, error)
	// CreateResourceClassify CreateResourceClassify 创建资料分组
	CreateResourceClassify(context.Context, *CreateResourceClassifyRequest) (*CreateResourceClassifyReply, error)
	// DeleteResource DeleteResource 删除咨询信息
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceReply, error)
	// DeleteResourceClassify DeleteResourceClassify 删除资料分组
	DeleteResourceClassify(context.Context, *DeleteResourceClassifyRequest) (*DeleteResourceClassifyReply, error)
	// GetResource GetResource 获取指定的咨询信息
	GetResource(context.Context, *GetResourceRequest) (*GetResourceReply, error)
	// ListResource ListResource 获取咨询信息列表
	ListResource(context.Context, *ListResourceRequest) (*ListResourceReply, error)
	// ListResourceClassify ListResourceClassify 获取资料分组列表
	ListResourceClassify(context.Context, *ListResourceClassifyRequest) (*ListResourceClassifyReply, error)
	// UpdateResource UpdateResource 更新咨询信息
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceReply, error)
	// UpdateResourceClassify UpdateResourceClassify 更新资料分组
	UpdateResourceClassify(context.Context, *UpdateResourceClassifyRequest) (*UpdateResourceClassifyReply, error)
}

func RegisterResourceHTTPServer(s *http.Server, srv ResourceHTTPServer) {
	r := s.Route("/")
	r.GET("/partyaffairs/client/v1/resource/classifies", _Resource_ListResourceClassify0_HTTP_Handler(srv))
	r.GET("/partyaffairs/api/v1/resource/classifies", _Resource_ListResourceClassify1_HTTP_Handler(srv))
	r.POST("/partyaffairs/api/v1/resource/classify", _Resource_CreateResourceClassify0_HTTP_Handler(srv))
	r.PUT("/partyaffairs/api/v1/resource/classify", _Resource_UpdateResourceClassify0_HTTP_Handler(srv))
	r.DELETE("/partyaffairs/api/v1/resource/classify", _Resource_DeleteResourceClassify0_HTTP_Handler(srv))
	r.GET("/partyaffairs/client/v1/resource", _Resource_GetResource0_HTTP_Handler(srv))
	r.GET("/partyaffairs/api/v1/resource", _Resource_GetResource1_HTTP_Handler(srv))
	r.GET("/partyaffairs/client/v1/resources", _Resource_ListResource0_HTTP_Handler(srv))
	r.GET("/partyaffairs/api/v1/resources", _Resource_ListResource1_HTTP_Handler(srv))
	r.POST("/partyaffairs/api/v1/resource", _Resource_CreateResource0_HTTP_Handler(srv))
	r.PUT("/partyaffairs/api/v1/resource", _Resource_UpdateResource0_HTTP_Handler(srv))
	r.DELETE("/partyaffairs/api/v1/resource", _Resource_DeleteResource0_HTTP_Handler(srv))
}

func _Resource_ListResourceClassify0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListResourceClassifyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceListResourceClassify)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListResourceClassify(ctx, req.(*ListResourceClassifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourceClassifyReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_ListResourceClassify1_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListResourceClassifyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceListResourceClassify)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListResourceClassify(ctx, req.(*ListResourceClassifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourceClassifyReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_CreateResourceClassify0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateResourceClassifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceCreateResourceClassify)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateResourceClassify(ctx, req.(*CreateResourceClassifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResourceClassifyReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_UpdateResourceClassify0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateResourceClassifyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceUpdateResourceClassify)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateResourceClassify(ctx, req.(*UpdateResourceClassifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResourceClassifyReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_DeleteResourceClassify0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteResourceClassifyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceDeleteResourceClassify)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteResourceClassify(ctx, req.(*DeleteResourceClassifyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResourceClassifyReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_GetResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceGetResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetResource(ctx, req.(*GetResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_GetResource1_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceGetResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetResource(ctx, req.(*GetResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_ListResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceListResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListResource(ctx, req.(*ListResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_ListResource1_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceListResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListResource(ctx, req.(*ListResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_CreateResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceCreateResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateResource(ctx, req.(*CreateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_UpdateResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateResourceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceUpdateResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateResource(ctx, req.(*UpdateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResourceReply)
		return ctx.Result(200, reply)
	}
}

func _Resource_DeleteResource0_HTTP_Handler(srv ResourceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceDeleteResource)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteResource(ctx, req.(*DeleteResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResourceReply)
		return ctx.Result(200, reply)
	}
}

type ResourceHTTPClient interface {
	CreateResource(ctx context.Context, req *CreateResourceRequest, opts ...http.CallOption) (rsp *CreateResourceReply, err error)
	CreateResourceClassify(ctx context.Context, req *CreateResourceClassifyRequest, opts ...http.CallOption) (rsp *CreateResourceClassifyReply, err error)
	DeleteResource(ctx context.Context, req *DeleteResourceRequest, opts ...http.CallOption) (rsp *DeleteResourceReply, err error)
	DeleteResourceClassify(ctx context.Context, req *DeleteResourceClassifyRequest, opts ...http.CallOption) (rsp *DeleteResourceClassifyReply, err error)
	GetResource(ctx context.Context, req *GetResourceRequest, opts ...http.CallOption) (rsp *GetResourceReply, err error)
	ListResource(ctx context.Context, req *ListResourceRequest, opts ...http.CallOption) (rsp *ListResourceReply, err error)
	ListResourceClassify(ctx context.Context, req *ListResourceClassifyRequest, opts ...http.CallOption) (rsp *ListResourceClassifyReply, err error)
	UpdateResource(ctx context.Context, req *UpdateResourceRequest, opts ...http.CallOption) (rsp *UpdateResourceReply, err error)
	UpdateResourceClassify(ctx context.Context, req *UpdateResourceClassifyRequest, opts ...http.CallOption) (rsp *UpdateResourceClassifyReply, err error)
}

type ResourceHTTPClientImpl struct {
	cc *http.Client
}

func NewResourceHTTPClient(client *http.Client) ResourceHTTPClient {
	return &ResourceHTTPClientImpl{client}
}

func (c *ResourceHTTPClientImpl) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...http.CallOption) (*CreateResourceReply, error) {
	var out CreateResourceReply
	pattern := "/partyaffairs/api/v1/resource"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationResourceCreateResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) CreateResourceClassify(ctx context.Context, in *CreateResourceClassifyRequest, opts ...http.CallOption) (*CreateResourceClassifyReply, error) {
	var out CreateResourceClassifyReply
	pattern := "/partyaffairs/api/v1/resource/classify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationResourceCreateResourceClassify))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...http.CallOption) (*DeleteResourceReply, error) {
	var out DeleteResourceReply
	pattern := "/partyaffairs/api/v1/resource"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceDeleteResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) DeleteResourceClassify(ctx context.Context, in *DeleteResourceClassifyRequest, opts ...http.CallOption) (*DeleteResourceClassifyReply, error) {
	var out DeleteResourceClassifyReply
	pattern := "/partyaffairs/api/v1/resource/classify"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceDeleteResourceClassify))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) GetResource(ctx context.Context, in *GetResourceRequest, opts ...http.CallOption) (*GetResourceReply, error) {
	var out GetResourceReply
	pattern := "/partyaffairs/api/v1/resource"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceGetResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) ListResource(ctx context.Context, in *ListResourceRequest, opts ...http.CallOption) (*ListResourceReply, error) {
	var out ListResourceReply
	pattern := "/partyaffairs/api/v1/resources"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceListResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) ListResourceClassify(ctx context.Context, in *ListResourceClassifyRequest, opts ...http.CallOption) (*ListResourceClassifyReply, error) {
	var out ListResourceClassifyReply
	pattern := "/partyaffairs/api/v1/resource/classifies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationResourceListResourceClassify))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...http.CallOption) (*UpdateResourceReply, error) {
	var out UpdateResourceReply
	pattern := "/partyaffairs/api/v1/resource"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationResourceUpdateResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ResourceHTTPClientImpl) UpdateResourceClassify(ctx context.Context, in *UpdateResourceClassifyRequest, opts ...http.CallOption) (*UpdateResourceClassifyReply, error) {
	var out UpdateResourceClassifyReply
	pattern := "/partyaffairs/api/v1/resource/classify"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationResourceUpdateResourceClassify))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
