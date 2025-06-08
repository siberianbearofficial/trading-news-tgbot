# Обработка новостей

Модуль `news` предоставляет функционал для работы с новостями в проекте `trading-news-tgbot`. В текущей реализации модуль содержит функцию, которая возвращает сводку новостей для трейдеров.

## 📦 Структура модуля

### `GetSummary() string`

Функция возвращает строку с краткой сводкой последних новостей для трейдеров. Новости включают в себя ключевые события, такие как изменения цен на криптовалюты, анонсы обновлений и другие важные события.

#### Возвращаемое значение:
- Строка, содержащая сводку новостей в формате:
  ```plaintext
  📰 Trader News Summary

  - Bitcoin hits new all-time high!
  - Ethereum 2.0 launch date announced.
  - Major exchange suffers security breach.
  - Top trader shares his portfolio secrets.
  - Market volatility expected to increase this week.
  ```

#### Пример использования:
```go
package main

import (
	"fmt"
	"github.com/yourusername/trading-news-tgbot/internal/news"
)

func main() {
	summary := news.GetSummary()
	fmt.Println(summary)
}
```

#### Вывод:
```plaintext
📰 Trader News Summary

- Bitcoin hits new all-time high!
- Ethereum 2.0 launch date announced.
- Major exchange suffers security breach.
- Top trader shares his portfolio secrets.
- Market volatility expected to increase this week.
```

## 🔄 Возможные улучшения
1. **Динамическая загрузка новостей**: В текущей реализации новости захардкожены. Можно добавить поддержку загрузки новостей из внешних источников (например, RSS-лент или API новостных сервисов).
2. **Фильтрация новостей**: Добавить возможность фильтрации новостей по категориям (криптовалюты, акции, форекс и т.д.).
3. **Форматирование**: Поддержка Markdown или HTML для более красивого отображения новостей в Telegram.

---

Документация будет обновляться по мере развития функционала модуля. Если у вас есть предложения по улучшению, пожалуйста, создайте issue в репозитории проекта.