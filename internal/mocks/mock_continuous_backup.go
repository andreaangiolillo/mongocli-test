// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andreaangiolillo/mongocli-test/internal/store (interfaces: CheckpointsLister,ContinuousJobLister,ContinuousJobCreator,ContinuousSnapshotsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockCheckpointsLister is a mock of CheckpointsLister interface.
type MockCheckpointsLister struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointsListerMockRecorder
}

// MockCheckpointsListerMockRecorder is the mock recorder for MockCheckpointsLister.
type MockCheckpointsListerMockRecorder struct {
	mock *MockCheckpointsLister
}

// NewMockCheckpointsLister creates a new mock instance.
func NewMockCheckpointsLister(ctrl *gomock.Controller) *MockCheckpointsLister {
	mock := &MockCheckpointsLister{ctrl: ctrl}
	mock.recorder = &MockCheckpointsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointsLister) EXPECT() *MockCheckpointsListerMockRecorder {
	return m.recorder
}

// Checkpoints mocks base method.
func (m *MockCheckpointsLister) Checkpoints(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*mongodbatlas.Checkpoints, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkpoints", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Checkpoints)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkpoints indicates an expected call of Checkpoints.
func (mr *MockCheckpointsListerMockRecorder) Checkpoints(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkpoints", reflect.TypeOf((*MockCheckpointsLister)(nil).Checkpoints), arg0, arg1, arg2)
}

// MockContinuousJobLister is a mock of ContinuousJobLister interface.
type MockContinuousJobLister struct {
	ctrl     *gomock.Controller
	recorder *MockContinuousJobListerMockRecorder
}

// MockContinuousJobListerMockRecorder is the mock recorder for MockContinuousJobLister.
type MockContinuousJobListerMockRecorder struct {
	mock *MockContinuousJobLister
}

// NewMockContinuousJobLister creates a new mock instance.
func NewMockContinuousJobLister(ctrl *gomock.Controller) *MockContinuousJobLister {
	mock := &MockContinuousJobLister{ctrl: ctrl}
	mock.recorder = &MockContinuousJobListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContinuousJobLister) EXPECT() *MockContinuousJobListerMockRecorder {
	return m.recorder
}

// ContinuousRestoreJobs mocks base method.
func (m *MockContinuousJobLister) ContinuousRestoreJobs(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*mongodbatlas.ContinuousJobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContinuousRestoreJobs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.ContinuousJobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContinuousRestoreJobs indicates an expected call of ContinuousRestoreJobs.
func (mr *MockContinuousJobListerMockRecorder) ContinuousRestoreJobs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContinuousRestoreJobs", reflect.TypeOf((*MockContinuousJobLister)(nil).ContinuousRestoreJobs), arg0, arg1, arg2)
}

// MockContinuousJobCreator is a mock of ContinuousJobCreator interface.
type MockContinuousJobCreator struct {
	ctrl     *gomock.Controller
	recorder *MockContinuousJobCreatorMockRecorder
}

// MockContinuousJobCreatorMockRecorder is the mock recorder for MockContinuousJobCreator.
type MockContinuousJobCreatorMockRecorder struct {
	mock *MockContinuousJobCreator
}

// NewMockContinuousJobCreator creates a new mock instance.
func NewMockContinuousJobCreator(ctrl *gomock.Controller) *MockContinuousJobCreator {
	mock := &MockContinuousJobCreator{ctrl: ctrl}
	mock.recorder = &MockContinuousJobCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContinuousJobCreator) EXPECT() *MockContinuousJobCreatorMockRecorder {
	return m.recorder
}

// CreateContinuousRestoreJob mocks base method.
func (m *MockContinuousJobCreator) CreateContinuousRestoreJob(arg0, arg1 string, arg2 *mongodbatlas.ContinuousJobRequest) (*mongodbatlas.ContinuousJobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContinuousRestoreJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.ContinuousJobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContinuousRestoreJob indicates an expected call of CreateContinuousRestoreJob.
func (mr *MockContinuousJobCreatorMockRecorder) CreateContinuousRestoreJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContinuousRestoreJob", reflect.TypeOf((*MockContinuousJobCreator)(nil).CreateContinuousRestoreJob), arg0, arg1, arg2)
}

// MockContinuousSnapshotsLister is a mock of ContinuousSnapshotsLister interface.
type MockContinuousSnapshotsLister struct {
	ctrl     *gomock.Controller
	recorder *MockContinuousSnapshotsListerMockRecorder
}

// MockContinuousSnapshotsListerMockRecorder is the mock recorder for MockContinuousSnapshotsLister.
type MockContinuousSnapshotsListerMockRecorder struct {
	mock *MockContinuousSnapshotsLister
}

// NewMockContinuousSnapshotsLister creates a new mock instance.
func NewMockContinuousSnapshotsLister(ctrl *gomock.Controller) *MockContinuousSnapshotsLister {
	mock := &MockContinuousSnapshotsLister{ctrl: ctrl}
	mock.recorder = &MockContinuousSnapshotsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContinuousSnapshotsLister) EXPECT() *MockContinuousSnapshotsListerMockRecorder {
	return m.recorder
}

// ContinuousSnapshots mocks base method.
func (m *MockContinuousSnapshotsLister) ContinuousSnapshots(arg0, arg1 string, arg2 *mongodbatlas.ListOptions) (*mongodbatlas.ContinuousSnapshots, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContinuousSnapshots", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.ContinuousSnapshots)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContinuousSnapshots indicates an expected call of ContinuousSnapshots.
func (mr *MockContinuousSnapshotsListerMockRecorder) ContinuousSnapshots(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContinuousSnapshots", reflect.TypeOf((*MockContinuousSnapshotsLister)(nil).ContinuousSnapshots), arg0, arg1, arg2)
}
