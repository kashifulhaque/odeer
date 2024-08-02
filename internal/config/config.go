package config

import (
	"errors"
)

var (
  accountID string
  authToken string
)

type Config struct {
	AccountID string
	AuthToken string
	ModelName string
}

func LoadConfig() (*Config, error) {
	if accountID == "" || authToken == "" {
		return nil, errors.New("missing required environment variables")
	}

	return &Config{
		AccountID: accountID,
		AuthToken: authToken,
		ModelName: "@cf/meta/llama-3-8b-instruct-awq",
	}, nil
}
