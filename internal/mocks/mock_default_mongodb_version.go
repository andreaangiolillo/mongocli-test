// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreangiolillo/mongocli-test/internal/store (interfaces: DefaultVersionGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDefaultVersionGetter is a mock of DefaultVersionGetter interface.
type MockDefaultVersionGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDefaultVersionGetterMockRecorder
}

// MockDefaultVersionGetterMockRecorder is the mock recorder for MockDefaultVersionGetter.
type MockDefaultVersionGetterMockRecorder struct {
	mock *MockDefaultVersionGetter
}

// NewMockDefaultVersionGetter creates a new mock instance.
func NewMockDefaultVersionGetter(ctrl *gomock.Controller) *MockDefaultVersionGetter {
	mock := &MockDefaultVersionGetter{ctrl: ctrl}
	mock.recorder = &MockDefaultVersionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDefaultVersionGetter) EXPECT() *MockDefaultVersionGetterMockRecorder {
	return m.recorder
}

// DefaultMongoDBVersion mocks base method.
func (m *MockDefaultVersionGetter) DefaultMongoDBVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultMongoDBVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultMongoDBVersion indicates an expected call of DefaultMongoDBVersion.
func (mr *MockDefaultVersionGetterMockRecorder) DefaultMongoDBVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultMongoDBVersion", reflect.TypeOf((*MockDefaultVersionGetter)(nil).DefaultMongoDBVersion))
}
