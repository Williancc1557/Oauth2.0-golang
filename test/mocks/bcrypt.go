// Code generated by MockGen. DO NOT EDIT.
// Source: internal/utils/encrypter.go

// Package mock_utils is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCryptoUtil is a mock of CryptoUtil interface.
type MockCryptoUtil struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoUtilMockRecorder
}

// MockCryptoUtilMockRecorder is the mock recorder for MockCryptoUtil.
type MockCryptoUtilMockRecorder struct {
	mock *MockCryptoUtil
}

// NewMockCryptoUtil creates a new mock instance.
func NewMockCryptoUtil(ctrl *gomock.Controller) *MockCryptoUtil {
	mock := &MockCryptoUtil{ctrl: ctrl}
	mock.recorder = &MockCryptoUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoUtil) EXPECT() *MockCryptoUtilMockRecorder {
	return m.recorder
}

// GenerateFromPassword mocks base method.
func (m *MockCryptoUtil) GenerateFromPassword(arg0 []byte, arg1 int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateFromPassword", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateFromPassword indicates an expected call of GenerateFromPassword.
func (mr *MockCryptoUtilMockRecorder) GenerateFromPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateFromPassword", reflect.TypeOf((*MockCryptoUtil)(nil).GenerateFromPassword), arg0, arg1)
}
