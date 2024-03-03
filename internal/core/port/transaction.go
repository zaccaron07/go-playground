package port

import (
	"context"
	"go-playground/internal/core/domain"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
}
