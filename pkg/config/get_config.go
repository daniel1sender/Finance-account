package config

import (
	"github.com/kelseyhightower/envconfig"
)

func GetFromEnv() (Config, error) {
	var apiConfig Config
	err := envconfig.Process("", &apiConfig)
	if err != nil {
		return Config{}, err
	}
	return apiConfig, nil
}
