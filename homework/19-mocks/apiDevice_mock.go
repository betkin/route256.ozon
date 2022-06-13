// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/repo/device.go

// Package mocks_learning is a generated GoMock package.
package mocks_learning

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "gitlab.ozon.dev/betkin/device-api/internal/model"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockRepo) CreateDevice(ctx context.Context, device *model.Device) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", ctx, device)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockRepoMockRecorder) CreateDevice(ctx, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockRepo)(nil).CreateDevice), ctx, device)
}

// DescribeDevice mocks base method.
func (m *MockRepo) DescribeDevice(ctx context.Context, deviceID uint64) (*model.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDevice", ctx, deviceID)
	ret0, _ := ret[0].(*model.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDevice indicates an expected call of DescribeDevice.
func (mr *MockRepoMockRecorder) DescribeDevice(ctx, deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDevice", reflect.TypeOf((*MockRepo)(nil).DescribeDevice), ctx, deviceID)
}

// DescribeLastDevice mocks base method.
func (m *MockRepo) DescribeLastDevice(ctx context.Context) (*model.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLastDevice", ctx)
	ret0, _ := ret[0].(*model.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLastDevice indicates an expected call of DescribeLastDevice.
func (mr *MockRepoMockRecorder) DescribeLastDevice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLastDevice", reflect.TypeOf((*MockRepo)(nil).DescribeLastDevice), ctx)
}

// ListDevices mocks base method.
func (m *MockRepo) ListDevices(ctx context.Context, page, perPage uint64) ([]*model.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDevices", ctx, page, perPage)
	ret0, _ := ret[0].([]*model.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDevices indicates an expected call of ListDevices.
func (mr *MockRepoMockRecorder) ListDevices(ctx, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDevices", reflect.TypeOf((*MockRepo)(nil).ListDevices), ctx, page, perPage)
}

// LogDevice mocks base method.
func (m *MockRepo) LogDevice(ctx context.Context, deviceID uint64) ([]*model.DeviceEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogDevice", ctx, deviceID)
	ret0, _ := ret[0].([]*model.DeviceEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogDevice indicates an expected call of LogDevice.
func (mr *MockRepoMockRecorder) LogDevice(ctx, deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogDevice", reflect.TypeOf((*MockRepo)(nil).LogDevice), ctx, deviceID)
}

// RemoveDevice mocks base method.
func (m *MockRepo) RemoveDevice(ctx context.Context, deviceID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDevice", ctx, deviceID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveDevice indicates an expected call of RemoveDevice.
func (mr *MockRepoMockRecorder) RemoveDevice(ctx, deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDevice", reflect.TypeOf((*MockRepo)(nil).RemoveDevice), ctx, deviceID)
}

// RemoveLastDevice mocks base method.
func (m *MockRepo) RemoveLastDevice(ctx context.Context, deviceID *uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLastDevice", ctx, deviceID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLastDevice indicates an expected call of RemoveLastDevice.
func (mr *MockRepoMockRecorder) RemoveLastDevice(ctx, deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLastDevice", reflect.TypeOf((*MockRepo)(nil).RemoveLastDevice), ctx, deviceID)
}

// UpdateDevice mocks base method.
func (m *MockRepo) UpdateDevice(ctx context.Context, device *model.Device) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDevice", ctx, device)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDevice indicates an expected call of UpdateDevice.
func (mr *MockRepoMockRecorder) UpdateDevice(ctx, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDevice", reflect.TypeOf((*MockRepo)(nil).UpdateDevice), ctx, device)
}

// UpdateLastDevice mocks base method.
func (m *MockRepo) UpdateLastDevice(ctx context.Context, device *model.Device) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastDevice", ctx, device)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLastDevice indicates an expected call of UpdateLastDevice.
func (mr *MockRepoMockRecorder) UpdateLastDevice(ctx, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastDevice", reflect.TypeOf((*MockRepo)(nil).UpdateLastDevice), ctx, device)
}
