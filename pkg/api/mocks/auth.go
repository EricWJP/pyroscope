// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pyroscope-io/pyroscope/pkg/api (interfaces: AuthService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/pyroscope-io/pyroscope/pkg/model"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// APIKeyFromJWTToken mocks base method.
func (m *MockAuthService) APIKeyTokenFromJWTToken(arg0 context.Context, arg1 string) (model.APIKeyToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIKeyTokenFromJWTToken", arg0, arg1)
	ret0, _ := ret[0].(model.APIKeyToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIKeyFromJWTToken indicates an expected call of APIKeyFromJWTToken.
func (mr *MockAuthServiceMockRecorder) APIKeyFromJWTToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIKeyTokenFromJWTToken", reflect.TypeOf((*MockAuthService)(nil).APIKeyTokenFromJWTToken), arg0, arg1)
}

// AuthenticateUser mocks base method.
func (m *MockAuthService) AuthenticateUser(arg0 context.Context, arg1, arg2 string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockAuthServiceMockRecorder) AuthenticateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockAuthService)(nil).AuthenticateUser), arg0, arg1, arg2)
}

// UserFromJWTToken mocks base method.
func (m *MockAuthService) UserFromJWTToken(arg0 context.Context, arg1 string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserFromJWTToken", arg0, arg1)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserFromJWTToken indicates an expected call of UserFromJWTToken.
func (mr *MockAuthServiceMockRecorder) UserFromJWTToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserFromJWTToken", reflect.TypeOf((*MockAuthService)(nil).UserFromJWTToken), arg0, arg1)
}
