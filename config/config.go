package config

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql"
	"os"
)

func InitConfigPostgresql() postgresql.Config {
	return postgresql.Config{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}
}
