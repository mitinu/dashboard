# 📊 Dashboard System

Система для визуализации данных в виде интерактивных дашбордов с графиками и метриками.

![alt text](image.png)
![alt text](image-1.png)


## 🚀 Основные возможности

- **Автоматический импорт Excel** — система по расписанию проверяет хэши файлов и загружает новые данные
- **Гибкое расписание** — настройка интервала в днях и точного времени запуска (формат `ЧЧ:ММ`)
- **Визуализация данных** — интерактивные диаграммы и графики
- **PostgreSQL в качестве хранилища** — надежное хранение импортированных данных

## 🛠 Технологический стек

| Компонент   | Технология                |
|-------------|---------------------------|
| **Backend** | Go + Gin фреймворк        |
| **Frontend**| Vue 3                      |
| **Database**| PostgreSQL                 |
| **Архитектура** | Луковая (Clean Architecture) |


## 📁 Структура проекта
```markdown
project-root/
├── backend/
│ └── src/
│ ├── cmd/api/ # Точка входа (main.go)
│ ├── internal/
│ │ ├── application/ # Бизнес-логика
│ │ │ └── excel/ # Логика импорта Excel
│ │ ├── config/ # Загрузка и парсинг конфигов
│ │ ├── domain/ # Сущности и интерфейсы репозиториев
│ │ ├── DTO/ # Объекты запросов/ответов и SQL-запросы
│ │ ├── repository/
│ │ │ └── postgres/ # Реализация репозиториев для PostgreSQL
│ │ └── router/ # Маршрутизация HTTP-запросов
├── frontend/
│ └── src/
│ ├── App.vue # Главный компонент
│ ├── router/ # Настройка маршрутизации
│ ├── page/ # Страницы приложения
│ │ └── [name]/
│ │ └── component/ # Компоненты конкретной страницы
│ └── components/ # Переиспользуемые компоненты
│ ├── UI/ # Кнопки, инпуты, базовые элементы
│ ├── icon/ # SVG-иконки
│ └── diagram/ # Модуль диаграмм
│ └── page/
│ └── component/ # Компоненты страниц диаграмм
```
## 🗄 Схема базы данных
![alt text](image-2.png)


### Требования

- Go 1.21+
- Node.js 18+
- PostgreSQL 14+


#### 1. Клонирование репозитория

```bash
git clone https://github.com/mitinu/dashboard.git
cd backend
go mod download
go run src/cmd/api/main.go  #запуск backend


cd ../frontend
npm install
npm run dev                 #запуск frontend
npm run bild                #билдинг frontend
```
для изменение пути билдинга frontend
frontend/vite.config.js
                outputDir: '../backend/dist',  




```.env
# Gin Mode
# debug, release, test
GIN_MODE=           debug

# App Configuration
APP_PORT=           8080

# Server Timeouts (формат: 15s, 30s, 1m, 2m и т.д.)
SERVER_READ_TIMEOUT=    15s
SERVER_WRITE_TIMEOUT=   15s
SERVER_IDLE_TIMEOUT=    60s
SERVER_SHUTDOWN_TIMEOUT=10s

# Excel Files
PATH_EXCEL=         resource/excel
INTERVAL_DAYS_READS=1
CRON_TIME=          08:00           #hh:mm

# Database
DB_HOST=            localhost
DB_PORT=            5432
DB_USER=            postgres
DB_PASSWORD=        1011
DB_NAME=            dashboard
DB_SSL_MODE=        disable
```


⏰ Автоматическое чтение Excel

Система работает по следующему принципу:
    По расписанию (задается через CRON_TIME и INTERVAL_DAYS_READS) запускается задача
    Проверяется хэш Excel-файла по указанному пути PATH_EXCEL
    Если хэш изменился — данные импортируются в PostgreSQL
    Дашборд автоматически отображает актуальные данные

Пример настройки расписания:
Конфигурация	Результат
CRON_TIME=09:00 INTERVAL_DAYS_READS=1	каждый день в 9:00
CRON_TIME=18:30 INTERVAL_DAYS_READS=2	каждые 2 дня в 18:30


```markdown
## 🧪 Режимы работы Gin

| Режим | Описание |
|-------|----------|
| `debug` | Подробные логи, hot-reload (для разработки) |
| `release` | Оптимизированный режим (для продакшена) |
| `test` | Режим для интеграционных тестов |

---

## 🐛 Частые проблемы и решения

| Проблема | Решение |
|----------|---------|
| Порт 8080 уже занят | Измените `APP_PORT` в `.env` на другой, например 8081 |
| Не удается подключиться к БД | Проверьте что PostgreSQL запущен: `sudo systemctl start postgresql` |
| Excel файл не читается | Убедитесь что путь `PATH_EXCEL` существует и файл имеет права на чтение |
```

