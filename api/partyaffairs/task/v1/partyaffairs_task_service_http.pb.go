// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: api/partyaffairs/task/partyaffairs_task_service.proto

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

const OperationTaskCreateTask = "/partyaffairs.api.partyaffairs.task.v1.Task/CreateTask"
const OperationTaskCreateTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/CreateTaskValue"
const OperationTaskDeleteTask = "/partyaffairs.api.partyaffairs.task.v1.Task/DeleteTask"
const OperationTaskDeleteTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/DeleteTaskValue"
const OperationTaskExportTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/ExportTaskValue"
const OperationTaskGetCurTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/GetCurTaskValue"
const OperationTaskGetTask = "/partyaffairs.api.partyaffairs.task.v1.Task/GetTask"
const OperationTaskGetTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/GetTaskValue"
const OperationTaskListClientTask = "/partyaffairs.api.partyaffairs.task.v1.Task/ListClientTask"
const OperationTaskListTask = "/partyaffairs.api.partyaffairs.task.v1.Task/ListTask"
const OperationTaskListTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/ListTaskValue"
const OperationTaskUpdateTask = "/partyaffairs.api.partyaffairs.task.v1.Task/UpdateTask"
const OperationTaskUpdateTaskValue = "/partyaffairs.api.partyaffairs.task.v1.Task/UpdateTaskValue"

type TaskHTTPServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskReply, error)
	CreateTaskValue(context.Context, *CreateTaskValueRequest) (*CreateTaskValueReply, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskReply, error)
	DeleteTaskValue(context.Context, *DeleteTaskValueRequest) (*DeleteTaskValueReply, error)
	ExportTaskValue(context.Context, *ExportTaskValueRequest) (*ExportTaskValueReply, error)
	GetCurTaskValue(context.Context, *GetCurTaskValueRequest) (*GetCurTaskValueReply, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error)
	GetTaskValue(context.Context, *GetTaskValueRequest) (*GetTaskValueReply, error)
	ListClientTask(context.Context, *ListClientTaskRequest) (*ListClientTaskReply, error)
	ListTask(context.Context, *ListTaskRequest) (*ListTaskReply, error)
	ListTaskValue(context.Context, *ListTaskValueRequest) (*ListTaskValueReply, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskReply, error)
	UpdateTaskValue(context.Context, *UpdateTaskValueRequest) (*UpdateTaskValueReply, error)
}

func RegisterTaskHTTPServer(s *http.Server, srv TaskHTTPServer) {
	r := s.Route("/")
	r.GET("/partyaffairs/api/v1/tasks", _Task_ListTask0_HTTP_Handler(srv))
	r.GET("/partyaffairs/client/v1/tasks", _Task_ListClientTask0_HTTP_Handler(srv))
	r.GET("/partyaffairs/client/v1/task", _Task_GetTask0_HTTP_Handler(srv))
	r.GET("/partyaffairs/api/v1/task", _Task_GetTask1_HTTP_Handler(srv))
	r.POST("/partyaffairs/api/v1/task", _Task_CreateTask0_HTTP_Handler(srv))
	r.PUT("/partyaffairs/api/v1/task", _Task_UpdateTask0_HTTP_Handler(srv))
	r.DELETE("/partyaffairs/api/v1/task", _Task_DeleteTask0_HTTP_Handler(srv))
	r.GET("/partyaffairs/api/v1/task/values", _Task_ListTaskValue0_HTTP_Handler(srv))
	r.GET("/partyaffairs/api/v1/task/value", _Task_GetTaskValue0_HTTP_Handler(srv))
	r.POST("/partyaffairs/api/v1/task/values", _Task_ExportTaskValue0_HTTP_Handler(srv))
	r.GET("/partyaffairs/client/v1/task/value", _Task_GetCurTaskValue0_HTTP_Handler(srv))
	r.POST("/partyaffairs/client/v1/task/value", _Task_CreateTaskValue0_HTTP_Handler(srv))
	r.PUT("/partyaffairs/client/v1/task/value", _Task_UpdateTaskValue0_HTTP_Handler(srv))
	r.DELETE("/partyaffairs/api/v1/task/value", _Task_DeleteTaskValue0_HTTP_Handler(srv))
}

func _Task_ListTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskListTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTask(ctx, req.(*ListTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_ListClientTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListClientTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskListClientTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListClientTask(ctx, req.(*ListClientTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListClientTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTask(ctx, req.(*GetTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTask1_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTask(ctx, req.(*GetTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CreateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateTask(ctx, req.(*CreateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_DeleteTask0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskDeleteTask)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteTask(ctx, req.(*DeleteTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTaskReply)
		return ctx.Result(200, reply)
	}
}

func _Task_ListTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTaskValueRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskListTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTaskValue(ctx, req.(*ListTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTaskValueReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTaskValueRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTaskValue(ctx, req.(*GetTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTaskValueReply)
		return ctx.Result(200, reply)
	}
}

func _Task_ExportTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExportTaskValueRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskExportTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ExportTaskValue(ctx, req.(*ExportTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExportTaskValueReply)
		return ctx.Result(200, reply)
	}
}

func _Task_GetCurTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurTaskValueRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskGetCurTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetCurTaskValue(ctx, req.(*GetCurTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurTaskValueReply)
		return ctx.Result(200, reply)
	}
}

func _Task_CreateTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTaskValueRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskCreateTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateTaskValue(ctx, req.(*CreateTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTaskValueReply)
		return ctx.Result(200, reply)
	}
}

func _Task_UpdateTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskValueRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskUpdateTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateTaskValue(ctx, req.(*UpdateTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskValueReply)
		return ctx.Result(200, reply)
	}
}

func _Task_DeleteTaskValue0_HTTP_Handler(srv TaskHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTaskValueRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTaskDeleteTaskValue)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteTaskValue(ctx, req.(*DeleteTaskValueRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTaskValueReply)
		return ctx.Result(200, reply)
	}
}

type TaskHTTPClient interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest, opts ...http.CallOption) (rsp *CreateTaskReply, err error)
	CreateTaskValue(ctx context.Context, req *CreateTaskValueRequest, opts ...http.CallOption) (rsp *CreateTaskValueReply, err error)
	DeleteTask(ctx context.Context, req *DeleteTaskRequest, opts ...http.CallOption) (rsp *DeleteTaskReply, err error)
	DeleteTaskValue(ctx context.Context, req *DeleteTaskValueRequest, opts ...http.CallOption) (rsp *DeleteTaskValueReply, err error)
	ExportTaskValue(ctx context.Context, req *ExportTaskValueRequest, opts ...http.CallOption) (rsp *ExportTaskValueReply, err error)
	GetCurTaskValue(ctx context.Context, req *GetCurTaskValueRequest, opts ...http.CallOption) (rsp *GetCurTaskValueReply, err error)
	GetTask(ctx context.Context, req *GetTaskRequest, opts ...http.CallOption) (rsp *GetTaskReply, err error)
	GetTaskValue(ctx context.Context, req *GetTaskValueRequest, opts ...http.CallOption) (rsp *GetTaskValueReply, err error)
	ListClientTask(ctx context.Context, req *ListClientTaskRequest, opts ...http.CallOption) (rsp *ListClientTaskReply, err error)
	ListTask(ctx context.Context, req *ListTaskRequest, opts ...http.CallOption) (rsp *ListTaskReply, err error)
	ListTaskValue(ctx context.Context, req *ListTaskValueRequest, opts ...http.CallOption) (rsp *ListTaskValueReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *UpdateTaskReply, err error)
	UpdateTaskValue(ctx context.Context, req *UpdateTaskValueRequest, opts ...http.CallOption) (rsp *UpdateTaskValueReply, err error)
}

type TaskHTTPClientImpl struct {
	cc *http.Client
}

func NewTaskHTTPClient(client *http.Client) TaskHTTPClient {
	return &TaskHTTPClientImpl{client}
}

func (c *TaskHTTPClientImpl) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...http.CallOption) (*CreateTaskReply, error) {
	var out CreateTaskReply
	pattern := "/partyaffairs/api/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) CreateTaskValue(ctx context.Context, in *CreateTaskValueRequest, opts ...http.CallOption) (*CreateTaskValueReply, error) {
	var out CreateTaskValueReply
	pattern := "/partyaffairs/client/v1/task/value"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskCreateTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...http.CallOption) (*DeleteTaskReply, error) {
	var out DeleteTaskReply
	pattern := "/partyaffairs/api/v1/task"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskDeleteTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) DeleteTaskValue(ctx context.Context, in *DeleteTaskValueRequest, opts ...http.CallOption) (*DeleteTaskValueReply, error) {
	var out DeleteTaskValueReply
	pattern := "/partyaffairs/api/v1/task/value"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskDeleteTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ExportTaskValue(ctx context.Context, in *ExportTaskValueRequest, opts ...http.CallOption) (*ExportTaskValueReply, error) {
	var out ExportTaskValueReply
	pattern := "/partyaffairs/api/v1/task/values"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskExportTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetCurTaskValue(ctx context.Context, in *GetCurTaskValueRequest, opts ...http.CallOption) (*GetCurTaskValueReply, error) {
	var out GetCurTaskValueReply
	pattern := "/partyaffairs/client/v1/task/value"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetCurTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetTask(ctx context.Context, in *GetTaskRequest, opts ...http.CallOption) (*GetTaskReply, error) {
	var out GetTaskReply
	pattern := "/partyaffairs/api/v1/task"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) GetTaskValue(ctx context.Context, in *GetTaskValueRequest, opts ...http.CallOption) (*GetTaskValueReply, error) {
	var out GetTaskValueReply
	pattern := "/partyaffairs/api/v1/task/value"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskGetTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ListClientTask(ctx context.Context, in *ListClientTaskRequest, opts ...http.CallOption) (*ListClientTaskReply, error) {
	var out ListClientTaskReply
	pattern := "/partyaffairs/client/v1/tasks"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskListClientTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ListTask(ctx context.Context, in *ListTaskRequest, opts ...http.CallOption) (*ListTaskReply, error) {
	var out ListTaskReply
	pattern := "/partyaffairs/api/v1/tasks"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskListTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) ListTaskValue(ctx context.Context, in *ListTaskValueRequest, opts ...http.CallOption) (*ListTaskValueReply, error) {
	var out ListTaskValueReply
	pattern := "/partyaffairs/api/v1/task/values"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTaskListTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*UpdateTaskReply, error) {
	var out UpdateTaskReply
	pattern := "/partyaffairs/api/v1/task"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaskHTTPClientImpl) UpdateTaskValue(ctx context.Context, in *UpdateTaskValueRequest, opts ...http.CallOption) (*UpdateTaskValueReply, error) {
	var out UpdateTaskValueReply
	pattern := "/partyaffairs/client/v1/task/value"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTaskUpdateTaskValue))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
