// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/protocols/get_account_by_email_repository.go

// Package mock_protocols is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockGetAccountByEmailRepository is a mock of GetAccountByEmailRepository interface.
type MockGetAccountByEmailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGetAccountByEmailRepositoryMockRecorder
}

// MockGetAccountByEmailRepositoryMockRecorder is the mock recorder for MockGetAccountByEmailRepository.
type MockGetAccountByEmailRepositoryMockRecorder struct {
	mock *MockGetAccountByEmailRepository
}

// NewMockGetAccountByEmailRepository creates a new mock instance.
func NewMockGetAccountByEmailRepository(ctrl *gomock.Controller) *MockGetAccountByEmailRepository {
	mock := &MockGetAccountByEmailRepository{ctrl: ctrl}
	mock.recorder = &MockGetAccountByEmailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAccountByEmailRepository) EXPECT() *MockGetAccountByEmailRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetAccountByEmailRepository) Get(email string) (*models.AccountModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*models.AccountModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetAccountByEmailRepositoryMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetAccountByEmailRepository)(nil).Get), email)
}
