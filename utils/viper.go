package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	fileInfo, err := os.Stat(".env")
	if err == nil && fileInfo.Size() > 0 {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error when reading configuration file: %s\n", err)
		}
	}

	return viper.GetString(key)
}
