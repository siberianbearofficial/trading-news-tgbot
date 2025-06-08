На основе анализа файлов `internal/bot/bot.go`, `internal/news/news.go` и `cmd/bot/main.go` я составил документацию для фичи "Команда /news". Вот её содержимое:

---

# Команда `/news`

Команда `/news` предоставляет пользователю краткую сводку последних новостей из мира трейдинга. Она реализована в Telegram-боте и возвращает предопределённый список новостей.

## 🚀 Использование

1. Отправьте команду `/news` в чат с ботом.
2. Бот ответит сообщением с краткой сводкой новостей.

**Пример ответа бота:**
```
📰 Trader News Summary

- Bitcoin hits new all-time high!
- Ethereum 2.0 launch date announced.
- Major exchange suffers security breach.
- Top trader shares his portfolio secrets.
- Market volatility expected to increase this week.
```

## 🛠 Реализация

### `internal/bot/bot.go`
Файл содержит логику обработки команд Telegram-бота. Команда `/news` обрабатывается следующим образом:
- При получении команды `/news` бот вызывает функцию `news.GetSummary()`.
- Результат функции отправляется пользователю в виде сообщения.

```go
switch update.Message.Command() {
case "news":
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, news.GetSummary())
    bot.Send(msg)
default:
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Try /news")
    bot.Send(msg)
}
```

### `internal/news/news.go`
Файл содержит функцию `GetSummary()`, которая возвращает строку с предопределённой сводкой новостей.

```go
func GetSummary() string {
    return `📰 Trader News Summary

- Bitcoin hits new all-time high!
- Ethereum 2.0 launch date announced.
- Major exchange suffers security breach.
- Top trader shares his portfolio secrets.
- Market volatility expected to increase this week.`
}
```

### `cmd/bot/main.go`
Файл является точкой входа для запуска бота. Он загружает конфигурацию и запускает бота с помощью функции `bot.Run()`.

```go
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

## 🔄 Дальнейшее развитие
- Добавить возможность получать новости из внешних источников (например, через API).
- Реализовать фильтрацию новостей по категориям (криптовалюты, акции и т. д.).
- Добавить поддержку мультиязычности.

--- 

Если у вас есть дополнительные вопросы или предложения по улучшению документации, дайте знать!