package config

import (
	"github.com/joho/godotenv"
	"os"
)

type (
	Container struct {
		App      *App
		Database *Database
	}

	App struct {
		Name string
		Env  string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Driver   string
	}
)

func New() (*Container, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}

	if os.Getenv("APP_ENV") != "dev" {
		println("TODO Secret should be loaded from env")
	}
	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}
	database := &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	return &Container{
		App:      app,
		Database: database,
	}, nil
}
