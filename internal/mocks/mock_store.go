// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreangiolillo/mongocli-test/internal/store (interfaces: CredentialsGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/andreangiolillo/mongocli-test/internal/config"
	auth "go.mongodb.org/atlas/auth"
)

// MockCredentialsGetter is a mock of CredentialsGetter interface.
type MockCredentialsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialsGetterMockRecorder
}

// MockCredentialsGetterMockRecorder is the mock recorder for MockCredentialsGetter.
type MockCredentialsGetterMockRecorder struct {
	mock *MockCredentialsGetter
}

// NewMockCredentialsGetter creates a new mock instance.
func NewMockCredentialsGetter(ctrl *gomock.Controller) *MockCredentialsGetter {
	mock := &MockCredentialsGetter{ctrl: ctrl}
	mock.recorder = &MockCredentialsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialsGetter) EXPECT() *MockCredentialsGetterMockRecorder {
	return m.recorder
}

// AuthType mocks base method.
func (m *MockCredentialsGetter) AuthType() config.AuthMechanism {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthType")
	ret0, _ := ret[0].(config.AuthMechanism)
	return ret0
}

// AuthType indicates an expected call of AuthType.
func (mr *MockCredentialsGetterMockRecorder) AuthType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthType", reflect.TypeOf((*MockCredentialsGetter)(nil).AuthType))
}

// PrivateAPIKey mocks base method.
func (m *MockCredentialsGetter) PrivateAPIKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateAPIKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// PrivateAPIKey indicates an expected call of PrivateAPIKey.
func (mr *MockCredentialsGetterMockRecorder) PrivateAPIKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateAPIKey", reflect.TypeOf((*MockCredentialsGetter)(nil).PrivateAPIKey))
}

// PublicAPIKey mocks base method.
func (m *MockCredentialsGetter) PublicAPIKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicAPIKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// PublicAPIKey indicates an expected call of PublicAPIKey.
func (mr *MockCredentialsGetterMockRecorder) PublicAPIKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicAPIKey", reflect.TypeOf((*MockCredentialsGetter)(nil).PublicAPIKey))
}

// Token mocks base method.
func (m *MockCredentialsGetter) Token() (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token.
func (mr *MockCredentialsGetterMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockCredentialsGetter)(nil).Token))
}
