// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/protocols/add_account_repository.go

// Package mock_protocols is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	protocols "github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	gomock "github.com/golang/mock/gomock"
)

// MockAddAccountRepository is a mock of AddAccountRepository interface.
type MockAddAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAddAccountRepositoryMockRecorder
}

// MockAddAccountRepositoryMockRecorder is the mock recorder for MockAddAccountRepository.
type MockAddAccountRepositoryMockRecorder struct {
	mock *MockAddAccountRepository
}

// NewMockAddAccountRepository creates a new mock instance.
func NewMockAddAccountRepository(ctrl *gomock.Controller) *MockAddAccountRepository {
	mock := &MockAddAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAddAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddAccountRepository) EXPECT() *MockAddAccountRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAddAccountRepository) Add(account *protocols.AddAccountRepositoryInput) (*protocols.AddAccountRepositoryOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", account)
	ret0, _ := ret[0].(*protocols.AddAccountRepositoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockAddAccountRepositoryMockRecorder) Add(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAddAccountRepository)(nil).Add), account)
}
