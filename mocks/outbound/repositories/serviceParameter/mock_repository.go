// Code generated by MockGen. DO NOT EDIT.
// Source: internal/outbound/repositories/serviceParameter/repository.go

// Package mock_repository_serviceParameter is a generated GoMock package.
package mock_repository_serviceParameter

import (
	context "context"
	reflect "reflect"

	model "github.com/enricodg/leaf-example/internal/outbound/model"
	gomock "github.com/golang/mock/gomock"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) (model.ServiceParameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, orm, serviceParameter)
	ret0, _ := ret[0].(model.ServiceParameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, orm, serviceParameter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, orm, serviceParameter)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, orm, serviceParameter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, orm, serviceParameter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, orm, serviceParameter)
}

// FindAllPaginated mocks base method.
func (m *MockRepository) FindAllPaginated(ctx context.Context, orm leafSql.ORM, request model.ServiceParameterPagingRequest) (model.ServiceParameterPagingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllPaginated", ctx, orm, request)
	ret0, _ := ret[0].(model.ServiceParameterPagingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllPaginated indicates an expected call of FindAllPaginated.
func (mr *MockRepositoryMockRecorder) FindAllPaginated(ctx, orm, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllPaginated", reflect.TypeOf((*MockRepository)(nil).FindAllPaginated), ctx, orm, request)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, orm leafSql.ORM, id uint64) (model.ServiceParameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, orm, id)
	ret0, _ := ret[0].(model.ServiceParameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, orm, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, orm, id)
}

// FindByVariable mocks base method.
func (m *MockRepository) FindByVariable(ctx context.Context, orm leafSql.ORM, variable string) (model.ServiceParameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByVariable", ctx, orm, variable)
	ret0, _ := ret[0].(model.ServiceParameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByVariable indicates an expected call of FindByVariable.
func (mr *MockRepositoryMockRecorder) FindByVariable(ctx, orm, variable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByVariable", reflect.TypeOf((*MockRepository)(nil).FindByVariable), ctx, orm, variable)
}
