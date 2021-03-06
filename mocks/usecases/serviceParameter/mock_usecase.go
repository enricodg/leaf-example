// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecases/serviceParameter/usecase.go

// Package mock_usecase_serviceParameter is a generated GoMock package.
package mock_usecase_serviceParameter

import (
	context "context"
	reflect "reflect"

	model "github.com/enricodg/leaf-example/internal/usecases/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(ctx context.Context, request model.ServiceParameterUpsertRequest) (model.ServiceParameterUpsertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(model.ServiceParameterUpsertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, id)
}

// FindAllPaginated mocks base method.
func (m *MockUseCase) FindAllPaginated(ctx context.Context, request model.ServiceParameterPagingRequest) (model.ServiceParameterPagingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllPaginated", ctx, request)
	ret0, _ := ret[0].(model.ServiceParameterPagingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllPaginated indicates an expected call of FindAllPaginated.
func (mr *MockUseCaseMockRecorder) FindAllPaginated(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllPaginated", reflect.TypeOf((*MockUseCase)(nil).FindAllPaginated), ctx, request)
}

// FindByVariable mocks base method.
func (m *MockUseCase) FindByVariable(ctx context.Context, variable string) (model.ServiceParameterUpsertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVariable", ctx, variable)
	ret0, _ := ret[0].(model.ServiceParameterUpsertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVariable indicates an expected call of FindByVariable.
func (mr *MockUseCaseMockRecorder) FindByVariable(ctx, variable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVariable", reflect.TypeOf((*MockUseCase)(nil).FindByVariable), ctx, variable)
}
