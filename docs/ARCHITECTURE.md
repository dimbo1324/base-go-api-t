# Architecture

The project is a small Go API starter. The architecture keeps responsibilities explicit and avoids unnecessary abstractions.

## Layers

```text
cmd/api
  starts the application, loads config, opens the database, handles shutdown signals

internal/config
  reads and validates environment variables

internal/db
  opens and verifies the PostgreSQL connection

internal/httpapi
  owns HTTP routing, middleware, response formatting, and server lifecycle

internal/store
  owns database models and SQL persistence logic
```

## Rules

- Keep `cmd/api` thin. It should wire dependencies, not contain business logic.
- Put HTTP-specific code in `internal/httpapi`.
- Put SQL and persistence logic in `internal/store`.
- Put environment parsing and validation in `internal/config`.
- Do not add interfaces until there is a real need for substitution.
- Do not add frameworks, ORMs, or code generation unless they solve a concrete problem.
- Keep dependencies one-way. Lower-level packages must not import `cmd/api` or `internal/httpapi`.

## Current API surface

The project currently exposes only:

```text
GET /v1/status
```

The `users` and `posts` stores are available as persistence examples, but HTTP handlers for them are intentionally not implemented yet.

## Security notes

- Secrets are read from environment variables, not hardcoded into Go code.
- SQL writes use parameterized queries.
- HTTP errors return safe public messages.
- Request logs use path only and do not log query strings.
- Passwords are represented as `password_hash`; raw password handling is outside the scope of this starter.
