package config

import (
	"os"
	"github.com/joho/godotenv"
	"errors"
)

type Config struct {
	TelegramToken string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}
	if cfg.TelegramToken == "" {
		return nil, ErrNoToken
	}
	return cfg, nil
}

var ErrNoToken =  errors.New("TELEGRAM_TOKEN is required in .env") 