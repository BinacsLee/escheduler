// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_core is a generated GoMock package.
package mock_core

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEScheduler is a mock of EScheduler interface.
type MockEScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockESchedulerMockRecorder
}

// MockESchedulerMockRecorder is the mock recorder for MockEScheduler.
type MockESchedulerMockRecorder struct {
	mock *MockEScheduler
}

// NewMockEScheduler creates a new mock instance.
func NewMockEScheduler(ctrl *gomock.Controller) *MockEScheduler {
	mock := &MockEScheduler{ctrl: ctrl}
	mock.recorder = &MockESchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEScheduler) EXPECT() *MockESchedulerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockEScheduler) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockESchedulerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEScheduler)(nil).Run))
}
