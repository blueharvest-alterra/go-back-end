package config

import (
	// "github.com/blueharvest-alterra/go-back-end/drivers/redis"
	"os"

	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql"
	"github.com/blueharvest-alterra/go-back-end/utils"
)

func InitConfigPostgresql() postgresql.Config {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = utils.GetConfig("DB_NAME")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = utils.GetConfig("DB_USER")
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = utils.GetConfig("DB_PASS")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = utils.GetConfig("DB_HOST")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = utils.GetConfig("DB_PORT")
	}

	return postgresql.Config{
		Name: dbName,
		User: dbUser,
		Pass: dbPass,
		Host: dbHost,
		Port: dbPort,
	}
}

// func InitConfigRedis() redis.Config {
// 	connectionURL := os.Getenv("REDIS_URI")
// 	if connectionURL == "" {
// 		connectionURL = utils.GetConfig("REDIS_URI")
// 	}
// 	return redis.Config{
// 		ConnectionURL: connectionURL,
// 	}
// }
