package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ksindhwani/pismo/internal/pkg/model"
	"github.com/ksindhwani/pismo/internal/pkg/service"
	"github.com/ksindhwani/pismo/internal/pkg/utils"
)

type TransactionHandler struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: transactionService,
	}
}

func (th *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req service.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, struct{}{}, err)
		return
	}

	OperationType, err := model.GetOperationType(int(req.OperationTypeId))
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, req.OperationTypeId, err)
		return
	}

	if OperationType.IsPurchaseType() {
		req.Amount = -req.Amount
	}

	// create record in database
	id, err := th.TransactionService.CreateTransaction(req, OperationType)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, req, fmt.Errorf("unable to create transaction - %v", err))
		return
	}

	utils.WriteResponse(c, http.StatusCreated, id)

}

func (th *TransactionHandler) CreateAccount(c *gin.Context) {

	var req service.CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, struct{}{}, err)
		return
	}
	// create record in database
	id, err := th.TransactionService.CreateAccount(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, req, fmt.Errorf("unable to create account - %v", err))
		return
	}

	utils.WriteResponse(c, http.StatusCreated, id)

}
func (th *TransactionHandler) GetAccount(c *gin.Context) {

	accountId, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "unable to parse accountId from url", err)
		return
	}

	account, err := th.TransactionService.GetAccount(accountId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "unable to fetch account from database", err)
		return
	}

	utils.WriteResponse(c, http.StatusOK, account)
}
