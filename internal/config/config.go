package config

import (
	"errors"
	"os"
)

type Config struct {
	AccountID string
	AuthToken string
	ModelName string
}

func LoadConfig() (*Config, error) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	authToken := os.Getenv("CLOUDFLARE_WORKERS_AI_API_KEY")

	if accountID == "" || authToken == "" {
		return nil, errors.New("missing required environment variables")
	}

	return &Config{
		AccountID: accountID,
		AuthToken: authToken,
		ModelName: "@cf/meta/llama-3-8b-instruct-awq",
	}, nil
}
