# CI/CD для проекта `trading-news-tgbot`

Документация описывает настройку и работу CI/CD (Continuous Integration/Continuous Deployment) для проекта `trading-news-tgbot` с использованием GitHub Actions.

## 📌 Обзор

GitHub Actions используется для автоматизации процессов сборки, тестирования и деплоя проекта. В данном случае, workflow настроен для деплоя сайта, созданного с помощью Docusaurus, на GitHub Pages.

## 🛠️ Workflow

Workflow активируется при пуше в ветку `main` и выполняет следующие шаги:

### 1. **Checkout code**
   - Используется действие `actions/checkout@v4` для получения кода из репозитория.

### 2. **Set up Node.js**
   - Устанавливается Node.js версии 20 с помощью действия `actions/setup-node@v3`.

### 3. **Navigate to /docs and install dependencies**
   - Переход в директорию `/docs` и установка зависимостей с помощью команды `npm install`.

### 4. **Build Docusaurus site**
   - Сборка сайта Docusaurus с помощью команды `npm run build`.

### 5. **Deploy to GitHub Pages**
   - Деплой собранного сайта на GitHub Pages с использованием действия `peaceiris/actions-gh-pages@v3`.
   - Для авторизации используется токен `GITHUB_TOKEN`.
   - Публикуется содержимое директории `./docs/build`.

## 🔄 Триггеры
Workflow запускается автоматически при:
- Пуше в ветку `main`.

## 🚀 Пример файла `.github/workflows/main.yaml`
```yaml
name: Deploy Docusaurus Site

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Node.js
        uses: actions/setup-node@v3
        with:
          node-version: '20'

      - name: Navigate to /docs and install dependencies
        run: |
          cd docs
          npm install

      - name: Build Docusaurus site
        run: |
          cd docs
          npm run build

      - name: Deploy to GitHub Pages
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./docs/build
```

## 📌 Примечания
- Убедитесь, что в репозитории включены GitHub Pages и указана ветка для публикации (обычно `gh-pages`).
- Для работы с приватными репозиториями может потребоваться дополнительная настройка токена.