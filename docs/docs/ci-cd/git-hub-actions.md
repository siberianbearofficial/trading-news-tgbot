# GitHub Actions

GitHub Actions используется в проекте для автоматизации процесса деплоя сайта Docusaurus при каждом пуше в ветку `main`. Ниже приведено описание конфигурации и шагов, выполняемых в рамках workflow.

## 🚀 Workflow: Deploy Docusaurus Site

### Триггеры
Workflow запускается при каждом пуше в ветку `main`.

```yaml
on:
  push:
    branches:
      - main
```

### Задачи (Jobs)
Workflow содержит одну задачу (`deploy`), которая выполняется на виртуальной машине с Ubuntu.

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
```

### Шаги (Steps)
1. **Checkout code**  
   Код проекта клонируется из репозитория.

   ```yaml
   - name: Checkout code
     uses: actions/checkout@v4
   ```

2. **Set up Node.js**  
   Устанавливается Node.js версии 20.

   ```yaml
   - name: Set up Node.js
     uses: actions/setup-node@v3
     with:
       node-version: '20'
   ```

3. **Navigate to /docs and install dependencies**  
   Переход в директорию `/docs` и установка зависимостей проекта.

   ```yaml
   - name: Navigate to /docs and install dependencies
     run: |
       cd docs
       npm install
   ```

4. **Build Docusaurus site**  
   Сборка сайта Docusaurus.

   ```yaml
   - name: Build Docusaurus site
     run: |
       cd docs
       npm run build
   ```

5. **Deploy to GitHub Pages**  
   Деплой собранного сайта на GitHub Pages с использованием токена `GITHUB_TOKEN`.

   ```yaml
   - name: Deploy to GitHub Pages
     uses: peaceiris/actions-gh-pages@v3
     with:
       github_token: ${{ secrets.GITHUB_TOKEN }}
       publish_dir: ./docs/build
   ```

---

## 🔧 Настройка
Для корректной работы workflow убедитесь, что:
- Ветка `main` существует в репозитории.
- В директории `/docs` находится проект Docusaurus с корректными зависимостями.
- Токен `GITHUB_TOKEN` доступен в секретах репозитория.

## 📌 Пример использования
После настройки workflow автоматически запустится при каждом пуше в ветку `main`, выполнит сборку и деплой сайта Docusaurus на GitHub Pages.