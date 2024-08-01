// Code generated by MockGen.
// Source: /Users/kunalsindhwani/take-home-assingment/Pismo/database/transactionDb.go

// Package mock_database is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ksindhwani/pismo/internal/pkg/model"
)

// MockTransactionDatabase is a mock of TransactionDatabase interface.
type MockTransactionDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionDatabaseMockRecorder
}

// MockTransactionDatabaseMockRecorder is the mock recorder for MockTransactionDatabase.
type MockTransactionDatabaseMockRecorder struct {
	mock *MockTransactionDatabase
}

// NewMockTransactionDatabase creates a new mock instance.
func NewMockTransactionDatabase(ctrl *gomock.Controller) *MockTransactionDatabase {
	mock := &MockTransactionDatabase{ctrl: ctrl}
	mock.recorder = &MockTransactionDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionDatabase) EXPECT() *MockTransactionDatabaseMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockTransactionDatabase) CreateAccount(account *model.Account) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", account)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockTransactionDatabaseMockRecorder) CreateAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockTransactionDatabase)(nil).CreateAccount), account)
}

// CreateTransaction mocks base method.
func (m *MockTransactionDatabase) CreateTransaction(arg0 *model.Transaction) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionDatabaseMockRecorder) CreateTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionDatabase)(nil).CreateTransaction), arg0)
}

// GetAccount mocks base method.
func (m *MockTransactionDatabase) GetAccount(accountId int) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", accountId)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockTransactionDatabaseMockRecorder) GetAccount(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockTransactionDatabase)(nil).GetAccount), accountId)
}
