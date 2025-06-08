# Запуск бота

Документация по запуску Telegram-бота для проекта `trading-news-tgbot`.

## 🚀 Точка входа

Основной файл для запуска бота находится в `cmd/bot/main.go`. Он выполняет следующие действия:
1. Загружает конфигурацию бота.
2. Запускает бота с использованием загруженной конфигурации.

### Пример кода:
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

---

## ⚙️ Конфигурация бота

Конфигурация бота загружается из файла `.env` с помощью библиотеки `godotenv`. Основной параметр — `TELEGRAM_TOKEN`, который должен быть указан в `.env`.

### Пример `.env`:
```env
TELEGRAM_TOKEN=your_telegram_bot_token
```

### Код конфигурации (`internal/config/config.go`):
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

---

## 🤖 Логика бота

Логика работы бота реализована в файле `internal/bot/bot.go`. Бот поддерживает команду `/news`, которая возвращает сводку новостей.

### Основные функции:
1. **`Run(cfg *config.Config) error`**:
   - Инициализирует бота с использованием токена из конфигурации.
   - Настраивает канал для получения обновлений.
   - Обрабатывает команды пользователей.

### Пример кода:
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

---

## 🛠 Установка и запуск

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
2. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/yourusername/trading-news-tgbot.git
   cd trading-news-tgbot
   ```
3. Создайте файл `.env` и добавьте токен бота:
   ```env
   TELEGRAM_TOKEN=your_telegram_bot_token
   ```
4. Запустите бота:
   ```bash
   go run cmd/bot/main.go
   ```

---

## 📝 Примечания

- Убедитесь, что токен бота действителен и указан в `.env`.
- Для работы бота требуется подключение к интернету.