// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	types "github.com/Mellanox/k8s-rdma-shared-dev-plugin/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// CDI is an autogenerated mock type for the CDI type
type CDI struct {
	mock.Mock
}

// CreateCDISpec provides a mock function with given fields: resourcePrefix, resourceName, poolName, devices
func (_m *CDI) CreateCDISpec(resourcePrefix string, resourceName string, poolName string, devices []types.PciNetDevice) error {
	ret := _m.Called(resourcePrefix, resourceName, poolName, devices)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []types.PciNetDevice) error); ok {
		r0 = rf(resourcePrefix, resourceName, poolName, devices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateContainerAnnotations provides a mock function with given fields: devices, resourcePrefix, resourceKind
func (_m *CDI) CreateContainerAnnotations(devices []types.PciNetDevice, resourcePrefix string, resourceKind string) (map[string]string, error) {
	ret := _m.Called(devices, resourcePrefix, resourceKind)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func([]types.PciNetDevice, string, string) (map[string]string, error)); ok {
		return rf(devices, resourcePrefix, resourceKind)
	}
	if rf, ok := ret.Get(0).(func([]types.PciNetDevice, string, string) map[string]string); ok {
		r0 = rf(devices, resourcePrefix, resourceKind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]types.PciNetDevice, string, string) error); ok {
		r1 = rf(devices, resourcePrefix, resourceKind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCDI interface {
	mock.TestingT
	Cleanup(func())
}

// NewCDI creates a new instance of CDI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCDI(t mockConstructorTestingTNewCDI) *CDI {
	mock := &CDI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
