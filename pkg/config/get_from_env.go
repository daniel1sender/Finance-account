package config

import (
	"github.com/kelseyhightower/envconfig"
)

func GetFromEnv() (Config, error) {
	var apiEnv Config
	err := envconfig.Process("", &apiEnv)
	if err != nil {
		return Config{}, err
	}
	return apiEnv, nil
}
