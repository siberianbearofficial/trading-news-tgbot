# CI/CD для проекта trading-news-tgbot

В проекте **trading-news-tgbot** используется **GitHub Actions** для автоматизации процессов непрерывной интеграции и доставки (CI/CD). Основная цель — автоматическое развертывание документации проекта, созданной с помощью **Docusaurus**, при каждом изменении в ветке `main`.

## 🚀 GitHub Actions Workflow

Конфигурация CI/CD находится в файле `.github/workflows/main.yaml` и включает следующие шаги:

### 1. **Триггер**
Workflow запускается автоматически при каждом `push` в ветку `main`.

### 2. **Среда выполнения**
- Используется последняя версия Ubuntu (`ubuntu-latest`).

### 3. **Шаги выполнения**

#### **1. Checkout кода**
```yaml
- name: Checkout code
  uses: actions/checkout@v4
```
Клонирует репозиторий для дальнейшей работы.

---

#### **2. Установка Node.js**
```yaml
- name: Set up Node.js
  uses: actions/setup-node@v3
  with:
    node-version: '20'
```
Устанавливает Node.js версии 20, необходимую для сборки Docusaurus.

---

#### **3. Установка зависимостей**
```yaml
- name: Navigate to /docs and install dependencies
  run: |
    cd docs
    npm install
```
Переходит в директорию `docs` и устанавливает зависимости проекта.

---

#### **4. Сборка Docusaurus**
```yaml
- name: Build Docusaurus site
  run: |
    cd docs
    npm run build
```
Собирает статический сайт документации.

---

#### **5. Развертывание на GitHub Pages**
```yaml
- name: Deploy to GitHub Pages
  uses: peaceiris/actions-gh-pages@v3
  with:
    github_token: ${{ secrets.GITHUB_TOKEN }}
    publish_dir: ./docs/build
```
Развертывает собранный сайт на GitHub Pages, используя токен доступа.

---

## 🔧 Настройка
Для корректной работы workflow убедитесь, что:
1. Ветка `main` является основной веткой репозитория.
2. В настройках репозитория включен GitHub Pages и указана ветка `gh-pages` (если требуется).

---

## 📌 Пример использования
После каждого `push` в ветку `main`:
1. GitHub Actions автоматически запускает workflow.
2. Собирается документация.
3. Развертывается на GitHub Pages.

Пример лога выполнения можно посмотреть в разделе **Actions** репозитория.

---

## 🔗 Полезные ссылки
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Docusaurus Deployment Guide](https://docusaurus.io/docs/deployment)