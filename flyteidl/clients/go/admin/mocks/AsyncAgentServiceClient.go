// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AsyncAgentServiceClient is an autogenerated mock type for the AsyncAgentServiceClient type
type AsyncAgentServiceClient struct {
	mock.Mock
}

type AsyncAgentServiceClient_CreateTask struct {
	*mock.Call
}

func (_m AsyncAgentServiceClient_CreateTask) Return(_a0 *admin.CreateTaskResponse, _a1 error) *AsyncAgentServiceClient_CreateTask {
	return &AsyncAgentServiceClient_CreateTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentServiceClient) OnCreateTask(ctx context.Context, in *admin.CreateTaskRequest, opts ...grpc.CallOption) *AsyncAgentServiceClient_CreateTask {
	c_call := _m.On("CreateTask", ctx, in, opts)
	return &AsyncAgentServiceClient_CreateTask{Call: c_call}
}

func (_m *AsyncAgentServiceClient) OnCreateTaskMatch(matchers ...interface{}) *AsyncAgentServiceClient_CreateTask {
	c_call := _m.On("CreateTask", matchers...)
	return &AsyncAgentServiceClient_CreateTask{Call: c_call}
}

// CreateTask provides a mock function with given fields: ctx, in, opts
func (_m *AsyncAgentServiceClient) CreateTask(ctx context.Context, in *admin.CreateTaskRequest, opts ...grpc.CallOption) (*admin.CreateTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.CreateTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateTaskRequest, ...grpc.CallOption) *admin.CreateTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.CreateTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.CreateTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AsyncAgentServiceClient_DeleteTask struct {
	*mock.Call
}

func (_m AsyncAgentServiceClient_DeleteTask) Return(_a0 *admin.DeleteTaskResponse, _a1 error) *AsyncAgentServiceClient_DeleteTask {
	return &AsyncAgentServiceClient_DeleteTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentServiceClient) OnDeleteTask(ctx context.Context, in *admin.DeleteTaskRequest, opts ...grpc.CallOption) *AsyncAgentServiceClient_DeleteTask {
	c_call := _m.On("DeleteTask", ctx, in, opts)
	return &AsyncAgentServiceClient_DeleteTask{Call: c_call}
}

func (_m *AsyncAgentServiceClient) OnDeleteTaskMatch(matchers ...interface{}) *AsyncAgentServiceClient_DeleteTask {
	c_call := _m.On("DeleteTask", matchers...)
	return &AsyncAgentServiceClient_DeleteTask{Call: c_call}
}

// DeleteTask provides a mock function with given fields: ctx, in, opts
func (_m *AsyncAgentServiceClient) DeleteTask(ctx context.Context, in *admin.DeleteTaskRequest, opts ...grpc.CallOption) (*admin.DeleteTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.DeleteTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteTaskRequest, ...grpc.CallOption) *admin.DeleteTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.DeleteTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.DeleteTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AsyncAgentServiceClient_GetTask struct {
	*mock.Call
}

func (_m AsyncAgentServiceClient_GetTask) Return(_a0 *admin.GetTaskResponse, _a1 error) *AsyncAgentServiceClient_GetTask {
	return &AsyncAgentServiceClient_GetTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentServiceClient) OnGetTask(ctx context.Context, in *admin.GetTaskRequest, opts ...grpc.CallOption) *AsyncAgentServiceClient_GetTask {
	c_call := _m.On("GetTask", ctx, in, opts)
	return &AsyncAgentServiceClient_GetTask{Call: c_call}
}

func (_m *AsyncAgentServiceClient) OnGetTaskMatch(matchers ...interface{}) *AsyncAgentServiceClient_GetTask {
	c_call := _m.On("GetTask", matchers...)
	return &AsyncAgentServiceClient_GetTask{Call: c_call}
}

// GetTask provides a mock function with given fields: ctx, in, opts
func (_m *AsyncAgentServiceClient) GetTask(ctx context.Context, in *admin.GetTaskRequest, opts ...grpc.CallOption) (*admin.GetTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.GetTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetTaskRequest, ...grpc.CallOption) *admin.GetTaskResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GetTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.GetTaskRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AsyncAgentServiceClient_GetTaskLogs struct {
	*mock.Call
}

func (_m AsyncAgentServiceClient_GetTaskLogs) Return(_a0 *admin.GetTaskLogsResponse, _a1 error) *AsyncAgentServiceClient_GetTaskLogs {
	return &AsyncAgentServiceClient_GetTaskLogs{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentServiceClient) OnGetTaskLogs(ctx context.Context, in *admin.GetTaskLogsRequest, opts ...grpc.CallOption) *AsyncAgentServiceClient_GetTaskLogs {
	c_call := _m.On("GetTaskLogs", ctx, in, opts)
	return &AsyncAgentServiceClient_GetTaskLogs{Call: c_call}
}

func (_m *AsyncAgentServiceClient) OnGetTaskLogsMatch(matchers ...interface{}) *AsyncAgentServiceClient_GetTaskLogs {
	c_call := _m.On("GetTaskLogs", matchers...)
	return &AsyncAgentServiceClient_GetTaskLogs{Call: c_call}
}

// GetTaskLogs provides a mock function with given fields: ctx, in, opts
func (_m *AsyncAgentServiceClient) GetTaskLogs(ctx context.Context, in *admin.GetTaskLogsRequest, opts ...grpc.CallOption) (*admin.GetTaskLogsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.GetTaskLogsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetTaskLogsRequest, ...grpc.CallOption) *admin.GetTaskLogsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GetTaskLogsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.GetTaskLogsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AsyncAgentServiceClient_GetTaskMetrics struct {
	*mock.Call
}

func (_m AsyncAgentServiceClient_GetTaskMetrics) Return(_a0 *admin.GetTaskMetricsResponse, _a1 error) *AsyncAgentServiceClient_GetTaskMetrics {
	return &AsyncAgentServiceClient_GetTaskMetrics{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AsyncAgentServiceClient) OnGetTaskMetrics(ctx context.Context, in *admin.GetTaskMetricsRequest, opts ...grpc.CallOption) *AsyncAgentServiceClient_GetTaskMetrics {
	c_call := _m.On("GetTaskMetrics", ctx, in, opts)
	return &AsyncAgentServiceClient_GetTaskMetrics{Call: c_call}
}

func (_m *AsyncAgentServiceClient) OnGetTaskMetricsMatch(matchers ...interface{}) *AsyncAgentServiceClient_GetTaskMetrics {
	c_call := _m.On("GetTaskMetrics", matchers...)
	return &AsyncAgentServiceClient_GetTaskMetrics{Call: c_call}
}

// GetTaskMetrics provides a mock function with given fields: ctx, in, opts
func (_m *AsyncAgentServiceClient) GetTaskMetrics(ctx context.Context, in *admin.GetTaskMetricsRequest, opts ...grpc.CallOption) (*admin.GetTaskMetricsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.GetTaskMetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetTaskMetricsRequest, ...grpc.CallOption) *admin.GetTaskMetricsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GetTaskMetricsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.GetTaskMetricsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
