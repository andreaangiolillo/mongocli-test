// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreangiolillo/mongocli-test/internal/cli/auth (interfaces: LoginConfig)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLoginConfig is a mock of LoginConfig interface.
type MockLoginConfig struct {
	ctrl     *gomock.Controller
	recorder *MockLoginConfigMockRecorder
}

// MockLoginConfigMockRecorder is the mock recorder for MockLoginConfig.
type MockLoginConfigMockRecorder struct {
	mock *MockLoginConfig
}

// NewMockLoginConfig creates a new mock instance.
func NewMockLoginConfig(ctrl *gomock.Controller) *MockLoginConfig {
	mock := &MockLoginConfig{ctrl: ctrl}
	mock.recorder = &MockLoginConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginConfig) EXPECT() *MockLoginConfigMockRecorder {
	return m.recorder
}

// AccessTokenSubject mocks base method.
func (m *MockLoginConfig) AccessTokenSubject() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenSubject")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessTokenSubject indicates an expected call of AccessTokenSubject.
func (mr *MockLoginConfigMockRecorder) AccessTokenSubject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenSubject", reflect.TypeOf((*MockLoginConfig)(nil).AccessTokenSubject))
}

// Save mocks base method.
func (m *MockLoginConfig) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockLoginConfigMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockLoginConfig)(nil).Save))
}

// Set mocks base method.
func (m *MockLoginConfig) Set(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockLoginConfigMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLoginConfig)(nil).Set), arg0, arg1)
}

// SetGlobal mocks base method.
func (m *MockLoginConfig) SetGlobal(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGlobal", arg0, arg1)
}

// SetGlobal indicates an expected call of SetGlobal.
func (mr *MockLoginConfigMockRecorder) SetGlobal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGlobal", reflect.TypeOf((*MockLoginConfig)(nil).SetGlobal), arg0, arg1)
}
