package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ksindhwani/pismo/internal/pkg/model"
	"github.com/ksindhwani/pismo/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		request       CreateAccountRequest
		expectedId    uint
		expectedError error
	}{
		{
			name: "All Good , shall pass",
			request: CreateAccountRequest{
				DocumentNumber: "12345689332",
			},
			expectedId:    1,
			expectedError: nil,
		},
	}

	mockTransactionDb := mocks.NewMockTransactionDatabase(ctrl)
	service := NewTransactionService(mockTransactionDb)
	any := gomock.Any()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTransactionDb.EXPECT().
				CreateAccount(any).
				Return(test.expectedId, test.expectedError).
				Times(1)

			id, _ := service.CreateAccount(test.request)
			assert.Equal(t, test.expectedId, id)
			assert.NoError(t, test.expectedError)

		})
	}

}

func TestCreateAccountError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		request       CreateAccountRequest
		expectedId    uint
		expectedError error
	}{
		{
			name: "Database Return Error, Assert Error",
			request: CreateAccountRequest{
				DocumentNumber: "12345689332",
			},
			expectedId:    0,
			expectedError: errors.New("unable to create account"),
		},
	}

	mockTransactionDb := mocks.NewMockTransactionDatabase(ctrl)
	service := NewTransactionService(mockTransactionDb)
	any := gomock.Any()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTransactionDb.EXPECT().
				CreateAccount(any).
				Return(test.expectedId, test.expectedError).
				Times(1)

			id, err := service.CreateAccount(test.request)
			assert.Equal(t, test.expectedId, id)
			assert.EqualError(t, test.expectedError, err.Error())

		})
	}

}

func TestCreateTransaction(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		request       CreateTransactionRequest
		expectedId    uint
		expectedError error
	}{
		{
			name: "All Good Credit Transaction, shall pass",
			request: CreateTransactionRequest{
				AccountID:       1,
				OperationTypeId: 4,
				Amount:          120,
			},
			expectedId:    1,
			expectedError: nil,
		},
		{
			name: "All Good Purchase Transaction, shall pass",
			request: CreateTransactionRequest{
				AccountID:       1,
				OperationTypeId: 3,
				Amount:          120,
			},
			expectedId:    1,
			expectedError: nil,
		},
	}

	mockTransactionDb := mocks.NewMockTransactionDatabase(ctrl)
	service := NewTransactionService(mockTransactionDb)
	any := gomock.Any()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTransactionDb.EXPECT().
				CreateTransaction(any).
				Return(test.expectedId, test.expectedError).
				Times(1)

			id, _ := service.CreateTransaction(test.request)
			assert.Equal(t, test.expectedId, id)
			assert.NoError(t, test.expectedError)

		})
	}

}

func TestCreateTransactionError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		request       CreateTransactionRequest
		expectedId    uint
		expectedError error
	}{
		{
			name: "Database Return Error, Assert Error",
			request: CreateTransactionRequest{
				AccountID:       1,
				OperationTypeId: 4,
				Amount:          120,
			},
			expectedId:    0,
			expectedError: errors.New("unable to create transaction"),
		},
	}

	mockTransactionDb := mocks.NewMockTransactionDatabase(ctrl)
	service := NewTransactionService(mockTransactionDb)
	any := gomock.Any()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTransactionDb.EXPECT().
				CreateTransaction(any).
				Return(test.expectedId, test.expectedError).
				Times(1)

			id, err := service.CreateTransaction(test.request)
			assert.Equal(t, test.expectedId, id)
			assert.EqualError(t, test.expectedError, err.Error())

		})
	}

}

func TestGetAccount(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		request         int
		expectedAccount model.Account
		expectedError   error
	}{
		{
			name:    "All Good , shall pass",
			request: 1,
			expectedAccount: model.Account{
				AccountId:      1,
				DocumentNumber: "12345689332",
			},
			expectedError: nil,
		},
	}

	mockTransactionDb := mocks.NewMockTransactionDatabase(ctrl)
	service := NewTransactionService(mockTransactionDb)
	any := gomock.Any()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTransactionDb.EXPECT().
				GetAccount(any).
				Return(test.expectedAccount, test.expectedError).
				Times(1)

			account, _ := service.GetAccount(test.request)
			assert.Equal(t, test.expectedAccount, account)
			assert.NoError(t, test.expectedError)

		})
	}

}

func TestGetAccountError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name            string
		request         int
		expectedAccount model.Account
		expectedError   error
	}{
		{
			name:            "Database Return Error, Assert Error",
			request:         10,
			expectedAccount: model.Account{},
			expectedError:   errors.New("account doesn't exist"),
		},
	}

	mockTransactionDb := mocks.NewMockTransactionDatabase(ctrl)
	service := NewTransactionService(mockTransactionDb)
	any := gomock.Any()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTransactionDb.EXPECT().
				GetAccount(any).
				Return(test.expectedAccount, test.expectedError).
				Times(1)

			account, err := service.GetAccount(test.request)
			assert.Equal(t, test.expectedAccount, account)
			assert.EqualError(t, test.expectedError, err.Error())

		})
	}

}
