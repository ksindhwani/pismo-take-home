package service

import (
	"fmt"
	"math"
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
	CreateTransaction(req CreateTransactionRequest, operationType model.OperationType) (uint, error)
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

func (ts *TService) CreateTransaction(req CreateTransactionRequest, operationType model.OperationType) (uint, error) {

	transaction := model.Transaction{
		AccountId:       req.AccountID,
		OperationTypeId: req.OperationTypeId,
		Amount:          req.Amount,
		Balance:         req.Amount,
		EventDate:       time.Now(),
	}

	if !operationType.IsPurchaseType() {
		return ts.dischargeTransaction(transaction)
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

func (ts *TService) dischargeTransaction(transaction model.Transaction) (uint, error) {

	unPaidTranactions, err := ts.TransactionDb.GetUnPaidTransactions()

	if err != nil {
		return 0, fmt.Errorf("unable to fetch unpaid transactions %v", err)
	}

	for index := 0; index < len(unPaidTranactions); index++ {

		if transaction.Balance <= 0 {
			break
		}

		balance := math.Abs(unPaidTranactions[index].Balance)
		if transaction.Balance >= balance {
			unPaidTranactions[index].Balance = 0
			transaction.Balance -= balance
		} else {
			unPaidTranactions[index].Balance += transaction.Balance
			transaction.Balance = 0
		}
	}

	return ts.TransactionDb.UpdateBalance(unPaidTranactions, transaction)

}
