package service

import (
	"time"

	"github.com/ksindhwani/pismo/database"
	"github.com/ksindhwani/pismo/internal/pkg/model"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type CreateTransactionRequest struct {
	AccountID       uint    `json:"account_id"`
	OperationTypeId uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type TransactionService interface {
	CreateTransaction(request CreateTransactionRequest) (uint, error)
	CreateAccount(request CreateAccountRequest) (uint, error)
	GetAccount(accountId int) (model.Account, error)
}

type TService struct {
	TransactionDb database.TransactionDatabase
}

func NewTransactionService(tDb database.TransactionDatabase) TransactionService {
	return &TService{
		TransactionDb: tDb,
	}

}

func (ts *TService) CreateTransaction(req CreateTransactionRequest) (uint, error) {

	transaction := model.Transaction{
		AccountId:       req.AccountID,
		OperationTypeId: req.OperationTypeId,
		Amount:          req.Amount,
		EventDate:       time.Now(),
	}

	return ts.TransactionDb.CreateTransaction(&transaction)
}

func (ts *TService) CreateAccount(request CreateAccountRequest) (uint, error) {
	// create new record in database
	account := model.Account{DocumentNumber: request.DocumentNumber}

	return ts.TransactionDb.CreateAccount(&account)
}

func (ts *TService) GetAccount(accountId int) (model.Account, error) {
	return ts.TransactionDb.GetAccount(accountId)

}
