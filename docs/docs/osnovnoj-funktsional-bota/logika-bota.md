На основе анализа файлов `internal/bot/bot.go`, `internal/news/news.go` и `cmd/bot/main.go` я составил документацию для фичи "Логика бота". Вот её содержимое:

---

# Логика бота

Логика бота реализована в модуле `internal/bot/bot.go` и включает обработку команд пользователя, а также взаимодействие с модулем новостей (`internal/news/news.go`).

## 🏗️ Структура логики

### Основной цикл работы бота
1. **Инициализация бота**:
   - Бот авторизуется с использованием токена Telegram.
   - Настраивается канал для получения обновлений (сообщений от пользователей).

2. **Обработка сообщений**:
   - Бот проверяет, является ли сообщение командой (начинается с `/`).
   - Поддерживается команда `/news`, которая возвращает сводку новостей.

3. **Взаимодействие с модулем новостей**:
   - Команда `/news` вызывает функцию `GetSummary()` из модуля `internal/news/news.go`, которая возвращает строку с краткой сводкой новостей.

### Пример кода
```go
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

## 📌 Команды бота

### `/news`
- **Описание**: Возвращает краткую сводку последних новостей.
- **Пример ответа**:
  ```
  📰 Trader News Summary

  - Bitcoin hits new all-time high!
  - Ethereum 2.0 launch date announced.
  - Major exchange suffers security breach.
  - Top trader shares his portfolio secrets.
  - Market volatility expected to increase this week.
  ```

## 🚀 Запуск бота
Бот запускается через функцию `main()` в файле `cmd/bot/main.go`, которая загружает конфигурацию и вызывает функцию `Run()` из модуля `bot`.

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

## 🔄 Взаимодействие с другими модулями
- **Конфигурация**: Используется модуль `internal/config` для загрузки настроек (например, токена Telegram).
- **Новости**: Используется модуль `internal/news` для получения данных о новостях.

---

Эта документация охватывает основную логику работы бота, его команды и взаимодействие с другими модулями. Если потребуется дополнить её или изменить, дайте знать!