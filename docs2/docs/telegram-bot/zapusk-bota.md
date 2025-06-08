# Запуск бота

Документация описывает процесс запуска Telegram-бота для получения торговых новостей. Бот использует библиотеку `go-telegram-bot-api` для взаимодействия с Telegram API.

## 🚀 Точка входа

Основной код запуска бота находится в файле `cmd/bot/main.go`. Здесь происходит загрузка конфигурации и запуск бота.

```go
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
```

### Шаги запуска:
1. Загрузка конфигурации (`config.Load()`).
2. Запуск бота (`bot.Run(cfg)`).

---

## ⚙️ Конфигурация

Конфигурация бота загружается из переменных окружения (файл `.env`). Основной параметр — `TELEGRAM_TOKEN`, который необходим для авторизации бота в Telegram.

```go
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

var ErrNoToken = errors.New("TELEGRAM_TOKEN is required in .env")
```

### Требования:
1. Создайте файл `.env` в корне проекта.
2. Добавьте в него переменную `TELEGRAM_TOKEN` с токеном вашего бота.

Пример `.env`:
```env
TELEGRAM_TOKEN=ваш_токен
```

---

## 🤖 Логика бота

Основная логика бота реализована в файле `internal/bot/bot.go`. Бот поддерживает команду `/news`, которая возвращает сводку торговых новостей.

```go
package bot

import (
	"log"
	"trading-news-tgbot/internal/config"
	"trading-news-tgbot/internal/news"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return err
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "news":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, news.GetSummary())
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Try /news")
				bot.Send(msg)
			}
		}
	}
	return nil
}
```

### Команды бота:
- `/news` — возвращает сводку торговых новостей.

---

## 🛠 Установка и запуск

1. Установите зависимости:
   ```bash
   go mod tidy
   ```

2. Создайте файл `.env` и добавьте токен:
   ```env
   TELEGRAM_TOKEN=ваш_токен
   ```

3. Запустите бота:
   ```bash
   go run cmd/bot/main.go
   ```

После запуска бот будет доступен в Telegram и начнет обрабатывать команды.