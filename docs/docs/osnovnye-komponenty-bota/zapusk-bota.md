# Запуск бота

Документация для фичи "Запуск бота" в проекте `trading-news-tgbot`.

## 🚀 Запуск бота

Для запуска Telegram-бота выполните следующие шаги:

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
2. Склонируйте репозиторий проекта:
   ```bash
   git clone https://github.com/yourusername/trading-news-tgbot.git
   cd trading-news-tgbot
   ```
3. Установите зависимости:
   ```bash
   go mod download
   ```
4. Создайте файл `.env` в корне проекта и добавьте в него токен вашего Telegram-бота:
   ```env
   TELEGRAM_TOKEN=your_telegram_bot_token
   ```
5. Запустите бота:
   ```bash
   go run cmd/bot/main.go
   ```

## ⚙️ Конфигурация

Конфигурация бота загружается из файла `.env`. Обязательный параметр:
- `TELEGRAM_TOKEN` — токен вашего Telegram-бота.

Пример файла `.env`:
```env
TELEGRAM_TOKEN=1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ
```

## 🤖 Функционал бота

Бот поддерживает следующие команды:
- `/news` — получает сводку последних новостей для трейдеров.

### Пример использования
```
Пользователь: /news
Бот: 📰 Trader News Summary

- Bitcoin hits new all-time high!
- Ethereum 2.0 launch date announced.
- Major exchange suffers security breach.
- Top trader shares his portfolio secrets.
- Market volatility expected to increase this week.
```

## 📂 Структура проекта

- `cmd/bot/main.go` — точка входа для запуска бота.
- `internal/config/config.go` — загрузка конфигурации.
- `internal/bot/bot.go` — логика работы Telegram-бота.
- `internal/news/news.go` — генерация новостной сводки.

## 🔧 Зависимости

Для работы бота используются следующие библиотеки:
- `github.com/go-telegram-bot-api/telegram-bot-api/v5` — для взаимодействия с Telegram API.
- `github.com/joho/godotenv` — для загрузки переменных окружения из `.env`.

Установите их с помощью:
```bash
go get github.com/go-telegram-bot-api/telegram-bot-api/v5
go get github.com/joho/godotenv
```

## 🛠️ Дополнительно

Если у вас возникли проблемы с запуском, убедитесь, что:
1. Токен Telegram-бота указан корректно.
2. Файл `.env` находится в корне проекта.
3. Все зависимости установлены.

Для остановки бота нажмите `Ctrl+C` в терминале.