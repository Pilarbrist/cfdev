// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/start (interfaces: CFDevD)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCFDevD is a mock of CFDevD interface
type MockCFDevD struct {
	ctrl     *gomock.Controller
	recorder *MockCFDevDMockRecorder
}

// MockCFDevDMockRecorder is the mock recorder for MockCFDevD
type MockCFDevDMockRecorder struct {
	mock *MockCFDevD
}

// NewMockCFDevD creates a new mock instance
func NewMockCFDevD(ctrl *gomock.Controller) *MockCFDevD {
	mock := &MockCFDevD{ctrl: ctrl}
	mock.recorder = &MockCFDevDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCFDevD) EXPECT() *MockCFDevDMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockCFDevD) Install() error {
	ret := m.ctrl.Call(m, "Install")
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockCFDevDMockRecorder) Install() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockCFDevD)(nil).Install))
}
