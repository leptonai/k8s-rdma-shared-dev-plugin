// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

import v1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"

// ResourceServer is an autogenerated mock type for the ResourceServer type
type ResourceServer struct {
	mock.Mock
}

// Allocate provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) Allocate(_a0 context.Context, _a1 *v1beta1.AllocateRequest) (*v1beta1.AllocateResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.AllocateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.AllocateRequest) *v1beta1.AllocateResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.AllocateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.AllocateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDevicePluginOptions provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) GetDevicePluginOptions(_a0 context.Context, _a1 *v1beta1.Empty) (*v1beta1.DevicePluginOptions, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.DevicePluginOptions
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Empty) *v1beta1.DevicePluginOptions); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.DevicePluginOptions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAndWatch provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) ListAndWatch(_a0 *v1beta1.Empty, _a1 v1beta1.DevicePlugin_ListAndWatchServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1beta1.Empty, v1beta1.DevicePlugin_ListAndWatchServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreStartContainer provides a mock function with given fields: _a0, _a1
func (_m *ResourceServer) PreStartContainer(_a0 context.Context, _a1 *v1beta1.PreStartContainerRequest) (*v1beta1.PreStartContainerResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta1.PreStartContainerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.PreStartContainerRequest) *v1beta1.PreStartContainerResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.PreStartContainerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.PreStartContainerRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Restart provides a mock function with given fields:
func (_m *ResourceServer) Restart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *ResourceServer) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *ResourceServer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields:
func (_m *ResourceServer) Watch() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
