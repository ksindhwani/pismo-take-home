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
	GetUnPaidTransactions() ([]model.Transaction, error)
	UpdateBalance(unPaidTranactions []model.Transaction, transaction model.Transaction) (uint, error)
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

func (td *transactionDatabase) GetUnPaidTransactions() ([]model.Transaction, error) {

	var unPaidTransactions []model.Transaction

	td.TranactionDB.Raw("SELECT * from transactions WHERE balance < 0  order by event_date").Scan(&unPaidTransactions)

	return unPaidTransactions, nil
}

func (td *transactionDatabase) UpdateBalance(unPaidTranactions []model.Transaction, transaction model.Transaction) (uint, error) {

	td.TranactionDB.Transaction(func(tx *gorm.DB) error {

		for _, unPaid := range unPaidTranactions {
			if err := tx.Save(&model.Transaction{
				TransactionId:   unPaid.TransactionId,
				Balance:         unPaid.Balance,
				AccountId:       unPaid.AccountId,
				Amount:          unPaid.Amount,
				OperationTypeId: unPaid.OperationTypeId}).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil

	})
	return transaction.TransactionId, nil

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
