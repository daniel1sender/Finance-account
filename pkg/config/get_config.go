package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

func GetConfig() (Config, error) {
	os.Setenv("DB_URL", "postgres://postgres:1234@localhost:5432/desafio")
	os.Setenv("API_PORT", ":3000")
	os.Setenv("TOKEN_SECRET", "123")
	var apiConfig Config
	err := envconfig.Process("", &apiConfig)
	if err != nil {
		return Config{}, err
	}
	return apiConfig, nil
}
