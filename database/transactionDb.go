package database

import (
	"fmt"

	"github.com/ksindhwani/pismo/config"
	"github.com/ksindhwani/pismo/internal/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TransactionDatabase interface {
	CreateTransaction(*model.Transaction) (uint, error)
	CreateAccount(account *model.Account) (uint, error)
	GetAccount(accountId int) (model.Account, error)
}

type transactionDatabase struct {
	TranactionDB *gorm.DB
}

func NewTransactionDb(config *config.Config) (TransactionDatabase, error) {

	dsn := getTransactionDbConnectionString(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &transactionDatabase{
		TranactionDB: db,
	}, nil
}

func (td *transactionDatabase) CreateTransaction(transaction *model.Transaction) (uint, error) {
	return transaction.TransactionId, td.TranactionDB.Create(&transaction).Error
}

func (td *transactionDatabase) CreateAccount(account *model.Account) (uint, error) {
	return account.AccountId, td.TranactionDB.Create(&account).Error
}

func (td *transactionDatabase) GetAccount(accountId int) (model.Account, error) {
	var account model.Account

	err := td.TranactionDB.First(&account, accountId).Error
	return account, err
}

func getTransactionDbConnectionString(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DbHost,
		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbName,
		cfg.DbPort)
	return dsn
}
