# Конфигурация

Модуль `config` предоставляет функционал для загрузки и управления конфигурацией приложения. Он использует переменные окружения для хранения чувствительных данных, таких как токен Telegram бота.

## 📦 Установка

Для использования модуля `config` необходимо добавить его в зависимости проекта:

```bash
go get github.com/joho/godotenv
```

## 🔧 Структура и функции

### `Config`
Структура `Config` содержит поля для хранения конфигурационных данных приложения.

```go
type Config struct {
    TelegramToken string
}
```

- **Поля:**
  - `TelegramToken` — токен Telegram бота, необходимый для работы с API Telegram.

---

### `Load() (*Config, error)`
Функция `Load` загружает конфигурацию из переменных окружения и возвращает экземпляр `Config`.

```go
func Load() (*Config, error)
```

- **Возвращает:**
  - Указатель на `Config`, если загрузка прошла успешно.
  - Ошибку, если не удалось загрузить токен Telegram.

- **Пример использования:**
  ```go
  cfg, err := config.Load()
  if err != nil {
      log.Fatal(err)
  }
  ```

---

### `ErrNoToken`
Ошибка `ErrNoToken` возвращается, если токен Telegram не был найден в переменных окружения.

```go
var ErrNoToken = errors.New("TELEGRAM_TOKEN is required in .env")
```

## 🔄 Переменные окружения

Для работы модуля `config` необходимо создать файл `.env` в корне проекта и добавить в него токен Telegram бота:

```env
TELEGRAM_TOKEN=your_telegram_bot_token
```

## 🚀 Пример использования

```go
package main

import (
    "log"
    "your_project_path/internal/config"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Telegram Token: %s", cfg.TelegramToken)
}
```

Этот пример демонстрирует загрузку конфигурации и вывод токена Telegram в лог.