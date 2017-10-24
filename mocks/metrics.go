// Code generated by MockGen. DO NOT EDIT.
// Source: ./metrics/metrics.go

// Package mock_metrics is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/uber-go/dosa/metrics"
)

// MockScope is a mock of Scope interface
type MockScope struct {
	ctrl     *gomock.Controller
	recorder *MockScopeMockRecorder
}

// MockScopeMockRecorder is the mock recorder for MockScope
type MockScopeMockRecorder struct {
	mock *MockScope
}

// NewMockScope creates a new mock instance
func NewMockScope(ctrl *gomock.Controller) *MockScope {
	mock := &MockScope{ctrl: ctrl}
	mock.recorder = &MockScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScope) EXPECT() *MockScopeMockRecorder {
	return m.recorder
}

// Counter mocks base method
func (m *MockScope) Counter(name string) metrics.Counter {
	ret := m.ctrl.Call(m, "Counter", name)
	ret0, _ := ret[0].(metrics.Counter)
	return ret0
}

// Counter indicates an expected call of Counter
func (mr *MockScopeMockRecorder) Counter(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Counter", reflect.TypeOf((*MockScope)(nil).Counter), name)
}

// Tagged mocks base method
func (m *MockScope) Tagged(tags map[string]string) metrics.Scope {
	ret := m.ctrl.Call(m, "Tagged", tags)
	ret0, _ := ret[0].(metrics.Scope)
	return ret0
}

// Tagged indicates an expected call of Tagged
func (mr *MockScopeMockRecorder) Tagged(tags interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tagged", reflect.TypeOf((*MockScope)(nil).Tagged), tags)
}

// SubScope mocks base method
func (m *MockScope) SubScope(name string) metrics.Scope {
	ret := m.ctrl.Call(m, "SubScope", name)
	ret0, _ := ret[0].(metrics.Scope)
	return ret0
}

// SubScope indicates an expected call of SubScope
func (mr *MockScopeMockRecorder) SubScope(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubScope", reflect.TypeOf((*MockScope)(nil).SubScope), name)
}

// MockCounter is a mock of Counter interface
type MockCounter struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMockRecorder
}

// MockCounterMockRecorder is the mock recorder for MockCounter
type MockCounterMockRecorder struct {
	mock *MockCounter
}

// NewMockCounter creates a new mock instance
func NewMockCounter(ctrl *gomock.Controller) *MockCounter {
	mock := &MockCounter{ctrl: ctrl}
	mock.recorder = &MockCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCounter) EXPECT() *MockCounterMockRecorder {
	return m.recorder
}

// Inc mocks base method
func (m *MockCounter) Inc(delta int64) {
	m.ctrl.Call(m, "Inc", delta)
}

// Inc indicates an expected call of Inc
func (mr *MockCounterMockRecorder) Inc(delta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inc", reflect.TypeOf((*MockCounter)(nil).Inc), delta)
}

// MockTimer is a mock of Timer interface
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMockRecorder
}

// MockTimerMockRecorder is the mock recorder for MockTimer
type MockTimerMockRecorder struct {
	mock *MockTimer
}
