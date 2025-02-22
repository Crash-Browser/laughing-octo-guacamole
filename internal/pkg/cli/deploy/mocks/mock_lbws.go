// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/cli/deploy/lbws.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	elbv2 "github.com/aws/copilot-cli/internal/pkg/aws/elbv2"
	gomock "github.com/golang/mock/gomock"
)

// MockelbGetter is a mock of elbGetter interface.
type MockelbGetter struct {
	ctrl     *gomock.Controller
	recorder *MockelbGetterMockRecorder
}

// MockelbGetterMockRecorder is the mock recorder for MockelbGetter.
type MockelbGetterMockRecorder struct {
	mock *MockelbGetter
}

// NewMockelbGetter creates a new mock instance.
func NewMockelbGetter(ctrl *gomock.Controller) *MockelbGetter {
	mock := &MockelbGetter{ctrl: ctrl}
	mock.recorder = &MockelbGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockelbGetter) EXPECT() *MockelbGetterMockRecorder {
	return m.recorder
}

// LoadBalancer mocks base method.
func (m *MockelbGetter) LoadBalancer(nameOrARN string) (*elbv2.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBalancer", nameOrARN)
	ret0, _ := ret[0].(*elbv2.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadBalancer indicates an expected call of LoadBalancer.
func (mr *MockelbGetterMockRecorder) LoadBalancer(nameOrARN interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBalancer", reflect.TypeOf((*MockelbGetter)(nil).LoadBalancer), nameOrARN)
}
