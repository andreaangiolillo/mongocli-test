// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreangiolillo/mongocli-test/internal/store/atlas (interfaces: IntegrationLister)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// MockIntegrationLister is a mock of IntegrationLister interface.
type MockIntegrationLister struct {
	ctrl     *gomock.Controller
	recorder *MockIntegrationListerMockRecorder
}

// MockIntegrationListerMockRecorder is the mock recorder for MockIntegrationLister.
type MockIntegrationListerMockRecorder struct {
	mock *MockIntegrationLister
}

// NewMockIntegrationLister creates a new mock instance.
func NewMockIntegrationLister(ctrl *gomock.Controller) *MockIntegrationLister {
	mock := &MockIntegrationLister{ctrl: ctrl}
	mock.recorder = &MockIntegrationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntegrationLister) EXPECT() *MockIntegrationListerMockRecorder {
	return m.recorder
}

// Integrations mocks base method.
func (m *MockIntegrationLister) Integrations(arg0 string) (*admin.PaginatedIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Integrations", arg0)
	ret0, _ := ret[0].(*admin.PaginatedIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Integrations indicates an expected call of Integrations.
func (mr *MockIntegrationListerMockRecorder) Integrations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Integrations", reflect.TypeOf((*MockIntegrationLister)(nil).Integrations), arg0)
}
