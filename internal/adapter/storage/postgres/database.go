package postgres

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-playground/internal/adapter/config"
)

//go:embed migrations/*.sql
var migrationsDirectory embed.FS

type Database struct {
	*pgxpool.Pool
	url string
}

func New(context context.Context, config *config.Database) (*Database, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Driver,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	db, err := pgxpool.New(context, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context)
	if err != nil {
		return nil, err
	}

	return &Database{
		db,
		url,
	}, nil
}

func (db *Database) Migrate() error {
	driver, err := iofs.New(migrationsDirectory, "migrations")
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
	if err != nil {
		return err
	}

	err = migrations.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
