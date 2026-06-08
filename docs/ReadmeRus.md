# Base Go API Engine

Минимальный starter-проект для Go REST API с PostgreSQL, явной конфигурацией, корректной обработкой подключения к БД, graceful shutdown и тестируемым HTTP-слоем.

Проект намеренно небольшой. Это база для backend-сервиса, а не готовая платформа с авторизацией, ролями, метриками, очередями и фронтендом.

## Возможности

- HTTP server и routing на стандартной библиотеке Go.
- PostgreSQL через `database/sql` и `lib/pq`.
- Конфигурация через переменные окружения.
- Graceful shutdown по `SIGINT` и `SIGTERM`.
- JSON-ответы и безопасные сообщения об ошибках.
- Store layer для пользователей и постов.
- SQL-миграции.
- Unit-тесты для config и `/v1/status`.
- GitHub Actions для `gofmt`, `go vet` и `go test`.

## Быстрый запуск

```bash
git clone https://github.com/dimbo1324/Base-Go-API-Engine.git
cd Base-Go-API-Engine
cp .env.example .env
docker compose up -d
make migrate-up
go run ./cmd/api
```

Проверка:

```bash
curl http://localhost:8080/v1/status
```

Ответ:

```json
{"status":"ok"}
```

## Структура

```text
cmd/api                    # точка входа
cmd/migrate/migrations     # SQL-миграции
internal/config            # загрузка и проверка конфигурации
internal/db                # подключение к PostgreSQL
internal/httpapi           # HTTP server, routes, middleware, responses
internal/store             # модели и SQL persistence
```

## Основной принцип

Проект должен оставаться простым: новая логика добавляется в предсказуемые места, зависимости не усложняются без причины, бизнес-логика не размазывается по entry point и HTTP-слою.
