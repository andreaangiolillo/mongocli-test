// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: LDAPConfigurationVerifier,LDAPConfigurationDescriber,LDAPConfigurationSaver,LDAPConfigurationDeleter,LDAPConfigurationGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/admin"
)

// MockLDAPConfigurationVerifier is a mock of LDAPConfigurationVerifier interface.
type MockLDAPConfigurationVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConfigurationVerifierMockRecorder
}

// MockLDAPConfigurationVerifierMockRecorder is the mock recorder for MockLDAPConfigurationVerifier.
type MockLDAPConfigurationVerifierMockRecorder struct {
	mock *MockLDAPConfigurationVerifier
}

// NewMockLDAPConfigurationVerifier creates a new mock instance.
func NewMockLDAPConfigurationVerifier(ctrl *gomock.Controller) *MockLDAPConfigurationVerifier {
	mock := &MockLDAPConfigurationVerifier{ctrl: ctrl}
	mock.recorder = &MockLDAPConfigurationVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPConfigurationVerifier) EXPECT() *MockLDAPConfigurationVerifierMockRecorder {
	return m.recorder
}

// VerifyLDAPConfiguration mocks base method.
func (m *MockLDAPConfigurationVerifier) VerifyLDAPConfiguration(arg0 string, arg1 *admin.NDSLDAPVerifyConnectivityJobRequestParams) (*admin.NDSLDAPVerifyConnectivityJobRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLDAPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*admin.NDSLDAPVerifyConnectivityJobRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyLDAPConfiguration indicates an expected call of VerifyLDAPConfiguration.
func (mr *MockLDAPConfigurationVerifierMockRecorder) VerifyLDAPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLDAPConfiguration", reflect.TypeOf((*MockLDAPConfigurationVerifier)(nil).VerifyLDAPConfiguration), arg0, arg1)
}

// MockLDAPConfigurationDescriber is a mock of LDAPConfigurationDescriber interface.
type MockLDAPConfigurationDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConfigurationDescriberMockRecorder
}

// MockLDAPConfigurationDescriberMockRecorder is the mock recorder for MockLDAPConfigurationDescriber.
type MockLDAPConfigurationDescriberMockRecorder struct {
	mock *MockLDAPConfigurationDescriber
}

// NewMockLDAPConfigurationDescriber creates a new mock instance.
func NewMockLDAPConfigurationDescriber(ctrl *gomock.Controller) *MockLDAPConfigurationDescriber {
	mock := &MockLDAPConfigurationDescriber{ctrl: ctrl}
	mock.recorder = &MockLDAPConfigurationDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPConfigurationDescriber) EXPECT() *MockLDAPConfigurationDescriberMockRecorder {
	return m.recorder
}

// GetStatusLDAPConfiguration mocks base method.
func (m *MockLDAPConfigurationDescriber) GetStatusLDAPConfiguration(arg0, arg1 string) (*admin.NDSLDAPVerifyConnectivityJobRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatusLDAPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*admin.NDSLDAPVerifyConnectivityJobRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatusLDAPConfiguration indicates an expected call of GetStatusLDAPConfiguration.
func (mr *MockLDAPConfigurationDescriberMockRecorder) GetStatusLDAPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatusLDAPConfiguration", reflect.TypeOf((*MockLDAPConfigurationDescriber)(nil).GetStatusLDAPConfiguration), arg0, arg1)
}

// MockLDAPConfigurationSaver is a mock of LDAPConfigurationSaver interface.
type MockLDAPConfigurationSaver struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConfigurationSaverMockRecorder
}

// MockLDAPConfigurationSaverMockRecorder is the mock recorder for MockLDAPConfigurationSaver.
type MockLDAPConfigurationSaverMockRecorder struct {
	mock *MockLDAPConfigurationSaver
}

// NewMockLDAPConfigurationSaver creates a new mock instance.
func NewMockLDAPConfigurationSaver(ctrl *gomock.Controller) *MockLDAPConfigurationSaver {
	mock := &MockLDAPConfigurationSaver{ctrl: ctrl}
	mock.recorder = &MockLDAPConfigurationSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPConfigurationSaver) EXPECT() *MockLDAPConfigurationSaverMockRecorder {
	return m.recorder
}

// SaveLDAPConfiguration mocks base method.
func (m *MockLDAPConfigurationSaver) SaveLDAPConfiguration(arg0 string, arg1 *admin.UserSecurity) (*admin.UserSecurity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLDAPConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*admin.UserSecurity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveLDAPConfiguration indicates an expected call of SaveLDAPConfiguration.
func (mr *MockLDAPConfigurationSaverMockRecorder) SaveLDAPConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLDAPConfiguration", reflect.TypeOf((*MockLDAPConfigurationSaver)(nil).SaveLDAPConfiguration), arg0, arg1)
}

// MockLDAPConfigurationDeleter is a mock of LDAPConfigurationDeleter interface.
type MockLDAPConfigurationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConfigurationDeleterMockRecorder
}

// MockLDAPConfigurationDeleterMockRecorder is the mock recorder for MockLDAPConfigurationDeleter.
type MockLDAPConfigurationDeleterMockRecorder struct {
	mock *MockLDAPConfigurationDeleter
}

// NewMockLDAPConfigurationDeleter creates a new mock instance.
func NewMockLDAPConfigurationDeleter(ctrl *gomock.Controller) *MockLDAPConfigurationDeleter {
	mock := &MockLDAPConfigurationDeleter{ctrl: ctrl}
	mock.recorder = &MockLDAPConfigurationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPConfigurationDeleter) EXPECT() *MockLDAPConfigurationDeleterMockRecorder {
	return m.recorder
}

// DeleteLDAPConfiguration mocks base method.
func (m *MockLDAPConfigurationDeleter) DeleteLDAPConfiguration(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLDAPConfiguration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLDAPConfiguration indicates an expected call of DeleteLDAPConfiguration.
func (mr *MockLDAPConfigurationDeleterMockRecorder) DeleteLDAPConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLDAPConfiguration", reflect.TypeOf((*MockLDAPConfigurationDeleter)(nil).DeleteLDAPConfiguration), arg0)
}

// MockLDAPConfigurationGetter is a mock of LDAPConfigurationGetter interface.
type MockLDAPConfigurationGetter struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConfigurationGetterMockRecorder
}

// MockLDAPConfigurationGetterMockRecorder is the mock recorder for MockLDAPConfigurationGetter.
type MockLDAPConfigurationGetterMockRecorder struct {
	mock *MockLDAPConfigurationGetter
}

// NewMockLDAPConfigurationGetter creates a new mock instance.
func NewMockLDAPConfigurationGetter(ctrl *gomock.Controller) *MockLDAPConfigurationGetter {
	mock := &MockLDAPConfigurationGetter{ctrl: ctrl}
	mock.recorder = &MockLDAPConfigurationGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPConfigurationGetter) EXPECT() *MockLDAPConfigurationGetterMockRecorder {
	return m.recorder
}

// GetLDAPConfiguration mocks base method.
func (m *MockLDAPConfigurationGetter) GetLDAPConfiguration(arg0 string) (*admin.UserSecurity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLDAPConfiguration", arg0)
	ret0, _ := ret[0].(*admin.UserSecurity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLDAPConfiguration indicates an expected call of GetLDAPConfiguration.
func (mr *MockLDAPConfigurationGetterMockRecorder) GetLDAPConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLDAPConfiguration", reflect.TypeOf((*MockLDAPConfigurationGetter)(nil).GetLDAPConfiguration), arg0)
}
