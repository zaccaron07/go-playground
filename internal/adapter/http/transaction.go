package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-playground/internal/core/domain"
	"go-playground/internal/core/port"
	"net/http"
	"time"
)

type TransactionHandler struct {
	service port.TransactionService
}

func NewTransactionHandler(service port.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service,
	}
}

type createTransactionRequest struct {
	Name          string    `json:"name"`
	Amount        string    `json:"amount"`
	TransactionAt time.Time `json:"transaction_at1"`
}

type createTransactionResponse struct {
	Id uuid.UUID `json:"id"`
}

func (transactionHandler *TransactionHandler) CreateTransaction(ctx echo.Context) error {
	var createTransactionRequest createTransactionRequest
	err := ctx.Bind(&createTransactionRequest)
	if err != nil {
		err := ctx.JSON(http.StatusBadRequest, err)
		if err != nil {
			return err
		}
		return nil
	}
	transaction := domain.Transaction{
		Name:          createTransactionRequest.Name,
		Amount:        createTransactionRequest.Amount,
		TransactionAt: createTransactionRequest.TransactionAt,
	}
	newTransaction, err := transactionHandler.service.CreateTransaction(ctx.Request().Context(), &transaction)

	if err != nil {
		err = ctx.JSON(http.StatusInternalServerError, err.Error())
		if err != nil {
			return err
		}
	}
	response := createTransactionResponse{Id: newTransaction.Id}
	ctx.JSON(http.StatusOK, response)
	return nil
}
