package main

import (
	"log"

	"trading-news-tgbot/internal/bot"
	"trading-news-tgbot/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := bot.Run(cfg); err != nil {
		log.Fatalf("bot stopped with error: %v", err)
	}
} 