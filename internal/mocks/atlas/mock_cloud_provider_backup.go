// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreangiolillo/mongocli-test/internal/store/atlas (interfaces: ScheduleDescriber)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

// MockScheduleDescriber is a mock of ScheduleDescriber interface.
type MockScheduleDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleDescriberMockRecorder
}

// MockScheduleDescriberMockRecorder is the mock recorder for MockScheduleDescriber.
type MockScheduleDescriberMockRecorder struct {
	mock *MockScheduleDescriber
}

// NewMockScheduleDescriber creates a new mock instance.
func NewMockScheduleDescriber(ctrl *gomock.Controller) *MockScheduleDescriber {
	mock := &MockScheduleDescriber{ctrl: ctrl}
	mock.recorder = &MockScheduleDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleDescriber) EXPECT() *MockScheduleDescriberMockRecorder {
	return m.recorder
}

// DescribeSchedule mocks base method.
func (m *MockScheduleDescriber) DescribeSchedule(arg0, arg1 string) (*admin.DiskBackupSnapshotSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSchedule", arg0, arg1)
	ret0, _ := ret[0].(*admin.DiskBackupSnapshotSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSchedule indicates an expected call of DescribeSchedule.
func (mr *MockScheduleDescriberMockRecorder) DescribeSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSchedule", reflect.TypeOf((*MockScheduleDescriber)(nil).DescribeSchedule), arg0, arg1)
}
