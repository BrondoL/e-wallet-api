package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTransactions(c *gin.Context) {
	query := &dto.TransactionRequestQuery{}
	err := c.ShouldBindQuery(query)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("get transaction failed", http.StatusUnprocessableEntity, errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}
	query = dto.FormatQuery(query)

	user := c.MustGet("user").(*model.User)
	transactions, err := h.transactionService.GetTransactions(int(user.ID), query)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("get transactions failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}
	totalTransactions, err := h.transactionService.CountTransaction(int(user.ID))
	if err != nil {
		response := utils.ErrorResponse("get transactions failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedTransaction := dto.FormatTransactions(transactions)
	metadata := utils.Metadata{Resource: "transactions", TotalAll: int(totalTransactions), TotalNow: len(transactions), Page: query.Page, Limit: query.Limit, Sort: query.Sort}
	response := utils.ResponseWithPagination("get transaction success", http.StatusOK, formattedTransaction, metadata)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) TopUp(c *gin.Context) {
	input := &dto.TopUpRequestBody{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("top up failed", http.StatusUnprocessableEntity, errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	user := c.MustGet("user").(*model.User)
	input.User = user
	transaction, err := h.transactionService.TopUp(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("top up failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedTransaction := dto.FormatTopUp(transaction)
	response := utils.SuccessResponse("top up success", http.StatusOK, formattedTransaction)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Transfer(c *gin.Context) {
	input := &dto.TransferRequestBody{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		if err != nil {
			errors := utils.FormatValidationError(err)
			response := utils.ErrorResponse("transfer failed", http.StatusUnprocessableEntity, errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	user := c.MustGet("user").(*model.User)
	input.User = user
	transaction, err := h.transactionService.Transfer(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("transfer failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedTransaction := dto.FormatTransfer(transaction)
	response := utils.SuccessResponse("transfer success", http.StatusOK, formattedTransaction)
	c.JSON(http.StatusOK, response)
}
