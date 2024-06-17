// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/protocols/reset-refresh-token-repository.go

// Package mock_protocols is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockResetRefreshTokenRepository is a mock of ResetRefreshTokenRepository interface.
type MockResetRefreshTokenRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResetRefreshTokenRepositoryMockRecorder
}

// MockResetRefreshTokenRepositoryMockRecorder is the mock recorder for MockResetRefreshTokenRepository.
type MockResetRefreshTokenRepositoryMockRecorder struct {
	mock *MockResetRefreshTokenRepository
}

// NewMockResetRefreshTokenRepository creates a new mock instance.
func NewMockResetRefreshTokenRepository(ctrl *gomock.Controller) *MockResetRefreshTokenRepository {
	mock := &MockResetRefreshTokenRepository{ctrl: ctrl}
	mock.recorder = &MockResetRefreshTokenRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResetRefreshTokenRepository) EXPECT() *MockResetRefreshTokenRepositoryMockRecorder {
	return m.recorder
}

// Reset mocks base method.
func (m *MockResetRefreshTokenRepository) Reset(userId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reset indicates an expected call of Reset.
func (mr *MockResetRefreshTokenRepositoryMockRecorder) Reset(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockResetRefreshTokenRepository)(nil).Reset), userId)
}