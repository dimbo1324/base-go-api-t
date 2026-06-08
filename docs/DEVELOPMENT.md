# Development Guide

## Local setup

```bash
cp .env.example .env
docker compose up -d
make migrate-up
go run ./cmd/api
```

## Quality checks

Run these before pushing changes:

```bash
make fmt-check
go vet ./...
go test ./...
```

Use `make fmt` to format Go files.

## Adding a new endpoint

1. Add the handler in `internal/httpapi`.
2. Register the route in `internal/httpapi/routes.go`.
3. Put persistence logic in `internal/store` if the endpoint needs database access.
4. Add tests for success and error cases.
5. Update `Readme.md` if the public API changed.

## Adding database logic

1. Add or update SQL migrations in `cmd/migrate/migrations`.
2. Keep SQL queries close to the store that uses them.
3. Use parameterized queries only.
4. Use transactions when one business operation changes multiple records that must stay consistent.
5. Make models match database columns.

## Removing old functionality

1. Remove the route or public entry point first.
2. Remove unused handlers, store methods, models, and tests.
3. Run `go test ./...` and `go vet ./...`.
4. Update documentation.

## Release checklist

- `make fmt-check`
- `go vet ./...`
- `go test ./...`
- Migrations checked against a clean local database
- README updated if API, config, or setup changed
- No secrets committed
