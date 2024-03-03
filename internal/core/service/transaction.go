package service

import (
	"context"
	"go-playground/internal/core/domain"
	"go-playground/internal/core/port"
	"log/slog"
)

type TransactionService struct {
	transactionRepository port.TransactionRepository
}

func NewTransactionService(transactionRepository port.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepository,
	}
}

func (transactionService *TransactionService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	transaction, err := transactionService.transactionRepository.CreateTransaction(ctx, transaction)

	if err != nil {
		slog.Info("error creating ", err)
		return nil, err
	}
	return transaction, nil
}
