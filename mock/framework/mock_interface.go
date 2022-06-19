// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_framework is a generated GoMock package.
package mock_framework

import (
	context "context"
	reflect "reflect"

	framework "github.com/BinacsLee/escheduler/framework"
	gomock "github.com/golang/mock/gomock"
)

// MockStrategy is a mock of Strategy interface.
type MockStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockStrategyMockRecorder
}

// MockStrategyMockRecorder is the mock recorder for MockStrategy.
type MockStrategyMockRecorder struct {
	mock *MockStrategy
}

// NewMockStrategy creates a new mock instance.
func NewMockStrategy(ctrl *gomock.Controller) *MockStrategy {
	mock := &MockStrategy{ctrl: ctrl}
	mock.recorder = &MockStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStrategy) EXPECT() *MockStrategyMockRecorder {
	return m.recorder
}

// Schedule mocks base method.
func (m *MockStrategy) Schedule(arg0 context.Context, arg1 []framework.Relation) ([]framework.Relations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schedule", arg0, arg1)
	ret0, _ := ret[0].([]framework.Relations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Schedule indicates an expected call of Schedule.
func (mr *MockStrategyMockRecorder) Schedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockStrategy)(nil).Schedule), arg0, arg1)
}

// MockPrepare is a mock of Prepare interface.
type MockPrepare struct {
	ctrl     *gomock.Controller
	recorder *MockPrepareMockRecorder
}

// MockPrepareMockRecorder is the mock recorder for MockPrepare.
type MockPrepareMockRecorder struct {
	mock *MockPrepare
}

// NewMockPrepare creates a new mock instance.
func NewMockPrepare(ctrl *gomock.Controller) *MockPrepare {
	mock := &MockPrepare{ctrl: ctrl}
	mock.recorder = &MockPrepareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrepare) EXPECT() *MockPrepareMockRecorder {
	return m.recorder
}

// GenerateGraph mocks base method.
func (m *MockPrepare) GenerateGraph(arg0 context.Context, arg1 []framework.Node, arg2 []framework.Edge) (framework.Graph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateGraph", arg0, arg1, arg2)
	ret0, _ := ret[0].(framework.Graph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateGraph indicates an expected call of GenerateGraph.
func (mr *MockPrepareMockRecorder) GenerateGraph(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateGraph", reflect.TypeOf((*MockPrepare)(nil).GenerateGraph), arg0, arg1, arg2)
}

// MockProcess is a mock of Process interface.
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess.
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance.
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return m.recorder
}

// ProcessGraph mocks base method.
func (m *MockProcess) ProcessGraph(arg0 context.Context, arg1 framework.Graph) (framework.DepthChart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessGraph", arg0, arg1)
	ret0, _ := ret[0].(framework.DepthChart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessGraph indicates an expected call of ProcessGraph.
func (mr *MockProcessMockRecorder) ProcessGraph(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessGraph", reflect.TypeOf((*MockProcess)(nil).ProcessGraph), arg0, arg1)
}

// MockDecision is a mock of Decision interface.
type MockDecision struct {
	ctrl     *gomock.Controller
	recorder *MockDecisionMockRecorder
}

// MockDecisionMockRecorder is the mock recorder for MockDecision.
type MockDecisionMockRecorder struct {
	mock *MockDecision
}

// NewMockDecision creates a new mock instance.
func NewMockDecision(ctrl *gomock.Controller) *MockDecision {
	mock := &MockDecision{ctrl: ctrl}
	mock.recorder = &MockDecisionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecision) EXPECT() *MockDecisionMockRecorder {
	return m.recorder
}

// SelectEdges mocks base method.
func (m *MockDecision) SelectEdges(arg0 context.Context, arg1 framework.Graph, arg2 framework.DepthChart, arg3 framework.CheckFunc) (framework.EdgeSet, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectEdges", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(framework.EdgeSet)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// SelectEdges indicates an expected call of SelectEdges.
func (mr *MockDecisionMockRecorder) SelectEdges(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectEdges", reflect.TypeOf((*MockDecision)(nil).SelectEdges), arg0, arg1, arg2, arg3)
}
