// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecase/create_access_token.go

// Package mock_usecase is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	usecase "github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
	gomock "github.com/golang/mock/gomock"
)

// MockCreateAccessToken is a mock of CreateAccessToken interface.
type MockCreateAccessToken struct {
	ctrl     *gomock.Controller
	recorder *MockCreateAccessTokenMockRecorder
}

// MockCreateAccessTokenMockRecorder is the mock recorder for MockCreateAccessToken.
type MockCreateAccessTokenMockRecorder struct {
	mock *MockCreateAccessToken
}

// NewMockCreateAccessToken creates a new mock instance.
func NewMockCreateAccessToken(ctrl *gomock.Controller) *MockCreateAccessToken {
	mock := &MockCreateAccessToken{ctrl: ctrl}
	mock.recorder = &MockCreateAccessTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateAccessToken) EXPECT() *MockCreateAccessTokenMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreateAccessToken) Create(userId string) (*usecase.CreateAccessTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId)
	ret0, _ := ret[0].(*usecase.CreateAccessTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCreateAccessTokenMockRecorder) Create(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreateAccessToken)(nil).Create), userId)
}