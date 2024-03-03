package main

import (
	"context"
	"fmt"
	"go-playground/internal/adapter/config"
	"go-playground/internal/adapter/storage/postgres"
	"go-playground/internal/adapter/storage/postgres/repository"
	"go-playground/internal/core/service"
	"log/slog"
	"os"
)

func main() {
	configData, err := config.New()
	if err != nil {
		fmt.Printf("Failed to load config, %s\n", err)
	}
	ctx := context.Background()
	database, err := postgres.New(ctx, configData.Database)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer database.Close()
	slog.Info("Successfully connected to the database")

	err = database.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully migrated the database")

	transactionRepository := repository.NewTransactionRepository(database)
	transactionService := service.NewTransactionService(transactionRepository)

}
