package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ksindhwani/pismo/database"
	"github.com/ksindhwani/pismo/internal/pkg/service"
)

func NewRouter(Db database.TransactionDatabase) *gin.Engine {

	r := gin.Default()

	transationService := service.NewTransactionService(Db)
	transactionHandler := NewTransactionHandler(transationService)

	r.POST("/accounts", transactionHandler.CreateAccount)
	r.GET("/accounts/:accountId", transactionHandler.GetAccount)
	r.POST("/transactions", transactionHandler.CreateTransaction)

	return r
}
