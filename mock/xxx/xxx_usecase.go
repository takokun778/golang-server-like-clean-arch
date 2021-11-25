// Code generated by MockGen. DO NOT EDIT.
// Source: xxx_usecase.go

// Package mock_xxx is a generated GoMock package.
package mock_xxx

import (
	context "context"
	reflect "reflect"
	xxx "xxx/app/domain/xxx"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecaseCreate is a mock of UsecaseCreate interface.
type MockUsecaseCreate struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseCreateMockRecorder
}

// MockUsecaseCreateMockRecorder is the mock recorder for MockUsecaseCreate.
type MockUsecaseCreateMockRecorder struct {
	mock *MockUsecaseCreate
}

// NewMockUsecaseCreate creates a new mock instance.
func NewMockUsecaseCreate(ctrl *gomock.Controller) *MockUsecaseCreate {
	mock := &MockUsecaseCreate{ctrl: ctrl}
	mock.recorder = &MockUsecaseCreateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseCreate) EXPECT() *MockUsecaseCreateMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockUsecaseCreate) Handle(ctx context.Context, input *xxx.UsecaseCreateInput) (*xxx.UsecaseCreateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, input)
	ret0, _ := ret[0].(*xxx.UsecaseCreateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockUsecaseCreateMockRecorder) Handle(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUsecaseCreate)(nil).Handle), ctx, input)
}

// MockUsecaseRead is a mock of UsecaseRead interface.
type MockUsecaseRead struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseReadMockRecorder
}

// MockUsecaseReadMockRecorder is the mock recorder for MockUsecaseRead.
type MockUsecaseReadMockRecorder struct {
	mock *MockUsecaseRead
}

// NewMockUsecaseRead creates a new mock instance.
func NewMockUsecaseRead(ctrl *gomock.Controller) *MockUsecaseRead {
	mock := &MockUsecaseRead{ctrl: ctrl}
	mock.recorder = &MockUsecaseReadMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseRead) EXPECT() *MockUsecaseReadMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockUsecaseRead) Handle(ctx context.Context, input *xxx.UsecaseReadInput) (*xxx.UsecaseReadOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, input)
	ret0, _ := ret[0].(*xxx.UsecaseReadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockUsecaseReadMockRecorder) Handle(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUsecaseRead)(nil).Handle), ctx, input)
}

// MockUsecaseReadAll is a mock of UsecaseReadAll interface.
type MockUsecaseReadAll struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseReadAllMockRecorder
}

// MockUsecaseReadAllMockRecorder is the mock recorder for MockUsecaseReadAll.
type MockUsecaseReadAllMockRecorder struct {
	mock *MockUsecaseReadAll
}

// NewMockUsecaseReadAll creates a new mock instance.
func NewMockUsecaseReadAll(ctrl *gomock.Controller) *MockUsecaseReadAll {
	mock := &MockUsecaseReadAll{ctrl: ctrl}
	mock.recorder = &MockUsecaseReadAllMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseReadAll) EXPECT() *MockUsecaseReadAllMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockUsecaseReadAll) Handle(ctx context.Context, input *xxx.UsecaseReadAllInput) (*xxx.UsecaseReadAllOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, input)
	ret0, _ := ret[0].(*xxx.UsecaseReadAllOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockUsecaseReadAllMockRecorder) Handle(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUsecaseReadAll)(nil).Handle), ctx, input)
}

// MockUsecaseUpdate is a mock of UsecaseUpdate interface.
type MockUsecaseUpdate struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseUpdateMockRecorder
}

// MockUsecaseUpdateMockRecorder is the mock recorder for MockUsecaseUpdate.
type MockUsecaseUpdateMockRecorder struct {
	mock *MockUsecaseUpdate
}

// NewMockUsecaseUpdate creates a new mock instance.
func NewMockUsecaseUpdate(ctrl *gomock.Controller) *MockUsecaseUpdate {
	mock := &MockUsecaseUpdate{ctrl: ctrl}
	mock.recorder = &MockUsecaseUpdateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseUpdate) EXPECT() *MockUsecaseUpdateMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockUsecaseUpdate) Handle(ctx context.Context, input *xxx.UsecaseUpdateInput) (*xxx.UsecaseUpdateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, input)
	ret0, _ := ret[0].(*xxx.UsecaseUpdateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockUsecaseUpdateMockRecorder) Handle(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUsecaseUpdate)(nil).Handle), ctx, input)
}

// MockUsecaseDelete is a mock of UsecaseDelete interface.
type MockUsecaseDelete struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseDeleteMockRecorder
}

// MockUsecaseDeleteMockRecorder is the mock recorder for MockUsecaseDelete.
type MockUsecaseDeleteMockRecorder struct {
	mock *MockUsecaseDelete
}

// NewMockUsecaseDelete creates a new mock instance.
func NewMockUsecaseDelete(ctrl *gomock.Controller) *MockUsecaseDelete {
	mock := &MockUsecaseDelete{ctrl: ctrl}
	mock.recorder = &MockUsecaseDeleteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseDelete) EXPECT() *MockUsecaseDeleteMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockUsecaseDelete) Handle(ctx context.Context, input *xxx.UsecaseDeleteInput) (*xxx.UsecaseDeleteOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, input)
	ret0, _ := ret[0].(*xxx.UsecaseDeleteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockUsecaseDeleteMockRecorder) Handle(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUsecaseDelete)(nil).Handle), ctx, input)
}
