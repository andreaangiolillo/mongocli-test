// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store (interfaces: AuditingDescriber,AuditingUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// MockAuditingDescriber is a mock of AuditingDescriber interface.
type MockAuditingDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockAuditingDescriberMockRecorder
}

// MockAuditingDescriberMockRecorder is the mock recorder for MockAuditingDescriber.
type MockAuditingDescriberMockRecorder struct {
	mock *MockAuditingDescriber
}

// NewMockAuditingDescriber creates a new mock instance.
func NewMockAuditingDescriber(ctrl *gomock.Controller) *MockAuditingDescriber {
	mock := &MockAuditingDescriber{ctrl: ctrl}
	mock.recorder = &MockAuditingDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditingDescriber) EXPECT() *MockAuditingDescriberMockRecorder {
	return m.recorder
}

// Auditing mocks base method.
func (m *MockAuditingDescriber) Auditing(arg0 string) (*admin.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auditing", arg0)
	ret0, _ := ret[0].(*admin.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auditing indicates an expected call of Auditing.
func (mr *MockAuditingDescriberMockRecorder) Auditing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auditing", reflect.TypeOf((*MockAuditingDescriber)(nil).Auditing), arg0)
}

// MockAuditingUpdater is a mock of AuditingUpdater interface.
type MockAuditingUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockAuditingUpdaterMockRecorder
}

// MockAuditingUpdaterMockRecorder is the mock recorder for MockAuditingUpdater.
type MockAuditingUpdaterMockRecorder struct {
	mock *MockAuditingUpdater
}

// NewMockAuditingUpdater creates a new mock instance.
func NewMockAuditingUpdater(ctrl *gomock.Controller) *MockAuditingUpdater {
	mock := &MockAuditingUpdater{ctrl: ctrl}
	mock.recorder = &MockAuditingUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditingUpdater) EXPECT() *MockAuditingUpdaterMockRecorder {
	return m.recorder
}

// UpdateAuditingConfig mocks base method.
func (m *MockAuditingUpdater) UpdateAuditingConfig(arg0 string, arg1 *admin.AuditLog) (*admin.AuditLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuditingConfig", arg0, arg1)
	ret0, _ := ret[0].(*admin.AuditLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuditingConfig indicates an expected call of UpdateAuditingConfig.
func (mr *MockAuditingUpdaterMockRecorder) UpdateAuditingConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuditingConfig", reflect.TypeOf((*MockAuditingUpdater)(nil).UpdateAuditingConfig), arg0, arg1)
}
