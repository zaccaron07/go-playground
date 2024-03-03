package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go-playground/internal/adapter/storage/postgres"
	"go-playground/internal/core/domain"
	"log/slog"
)

type TransactionRepository struct {
	database *postgres.Database
}

func NewTransactionRepository(database *postgres.Database) *TransactionRepository {
	return &TransactionRepository{
		database,
	}
}

func (transactionRepository *TransactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	transactionQuery := `INSERT INTO transactions (name, amount, transactionAt) VALUES($1, $2, $3)`

	err := pgx.BeginFunc(ctx, transactionRepository.database, func(tx pgx.Tx) error {
		err := tx.QueryRow(ctx, transactionQuery, transaction.Name, transaction.Amount, transaction.TransactionAt).Scan(
			&transaction.Id,
			&transaction.Name,
			&transaction.Amount,
			&transaction.TransactionAt)
		if err != nil {
			return err
		}
		if transaction.Id == uuid.Nil {
			slog.Info("null id")
			return tx.Rollback(ctx)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
