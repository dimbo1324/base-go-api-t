# Base Go API Engine

[![Go Version](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat&logo=go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.3-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Minimal Go REST API starter with PostgreSQL, explicit configuration, safe database connection handling, graceful shutdown, and a small testable HTTP layer.

This project is intentionally small. It is designed as a clean backend foundation, not as a full production platform with authentication, authorization, observability, background workers, or generated clients.

## Features

- Standard-library HTTP server and routing.
- PostgreSQL access through `database/sql` and `lib/pq`.
- Explicit environment-based configuration with safe defaults.
- Graceful shutdown on `SIGINT` and `SIGTERM`.
- Request logging without query strings or secrets.
- Central JSON response helpers.
- Simple store layer for users and posts.
- SQL migrations for PostgreSQL.
- Unit tests for configuration and the status endpoint.
- GitHub Actions workflow for formatting, vetting, and tests.

## Requirements

- Go 1.23+
- PostgreSQL 16+
- Docker Compose, optional, for local PostgreSQL
- `golang-migrate`, optional, for running the provided migrations through `make`

## Quick start

```bash
git clone https://github.com/dimbo1324/Base-Go-API-Engine.git
cd Base-Go-API-Engine
cp .env.example .env
docker compose up -d
make migrate-up
go run ./cmd/api
```

Check the API:

```bash
curl http://localhost:8080/v1/status
```

Expected response:

```json
{"status":"ok"}
```

## Configuration

The application works with defaults, but local configuration should be placed in `.env`. Keep real secrets out of Git.

| Variable | Description | Default |
| --- | --- | --- |
| `ADDR` | HTTP server address | `:8080` |
| `DB_ADDR` | PostgreSQL connection string | `postgres://postgres:password@localhost/appdb?sslmode=disable` |
| `DB_MAX_OPEN_CONNS` | Maximum open DB connections | `30` |
| `DB_MAX_IDLE_CONNS` | Maximum idle DB connections | `30` |
| `DB_MAX_IDLE_TIME` | Maximum idle connection lifetime | `15m` |
| `DB_MIGRATOR_ADDR` | Migration connection string used by `make migrate-up/down` | same local PostgreSQL DSN |

`DB_MAX_IDLE_TIME_MINS` is still accepted as a legacy fallback, but new configuration should use `DB_MAX_IDLE_TIME`.

## Project structure

```text
.
├── cmd/
│   ├── api/                    # Application entry point
│   └── migrate/migrations/     # SQL migrations
├── internal/
│   ├── config/                 # Environment config loading and validation
│   ├── db/                     # PostgreSQL connection setup
│   ├── httpapi/                # HTTP server, routes, middleware, responses
│   └── store/                  # Database models and persistence logic
├── docs/                       # Additional project documentation
├── .github/workflows/          # CI checks
├── docker-compose.yml          # Local PostgreSQL
├── Makefile                    # Common development commands
└── go.mod
```

## API

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/v1/status` | API status check |

## Development commands

```bash
go test ./...
go vet ./...
make fmt
make fmt-check
make migrate-up
make migrate-down
```

## Design boundaries

This starter deliberately does not include authentication, authorization, ORM, code generation, Docker images for the API, Kubernetes, background jobs, metrics, tracing, or a frontend. Add those only when the real application needs them.

## Documentation

- [Architecture](docs/ARCHITECTURE.md)
- [Development guide](docs/DEVELOPMENT.md)
- [Russian README](docs/ReadmeRus.md)
- [Spanish README](docs/ReadmeSp.md)

## License

MIT. See [LICENSE](LICENSE).
