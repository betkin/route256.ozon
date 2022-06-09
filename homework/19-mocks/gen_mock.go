// Code generated by MockGen. DO NOT EDIT.
// Source: ./homework/19-mocks/gen.go

// Package mocks_learning is a generated GoMock package.
package mocks_learning

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockgenRand is a mock of genRand interface.
type MockgenRand struct {
	ctrl     *gomock.Controller
	recorder *MockgenRandMockRecorder
}

// MockgenRandMockRecorder is the mock recorder for MockgenRand.
type MockgenRandMockRecorder struct {
	mock *MockgenRand
}

// NewMockgenRand creates a new mock instance.
func NewMockgenRand(ctrl *gomock.Controller) *MockgenRand {
	mock := &MockgenRand{ctrl: ctrl}
	mock.recorder = &MockgenRandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgenRand) EXPECT() *MockgenRandMockRecorder {
	return m.recorder
}

// Doer mocks base method.
func (m *MockgenRand) Doer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Doer")
}

// Doer indicates an expected call of Doer.
func (mr *MockgenRandMockRecorder) Doer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Doer", reflect.TypeOf((*MockgenRand)(nil).Doer))
}

// Random mocks base method.
func (m *MockgenRand) Random(n int) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Random", n)
	ret0, _ := ret[0].(int64)
	return ret0
}

// Random indicates an expected call of Random.
func (mr *MockgenRandMockRecorder) Random(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Random", reflect.TypeOf((*MockgenRand)(nil).Random), n)
}
