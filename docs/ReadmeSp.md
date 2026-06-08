# Base Go API Engine

Starter mínimo para una REST API en Go con PostgreSQL, configuración explícita, manejo seguro de conexión a la base de datos, graceful shutdown y una capa HTTP fácil de probar.

El proyecto es pequeño a propósito. Es una base backend, no una plataforma completa con autenticación, autorización, métricas, workers, colas o frontend.

## Características

- HTTP server y routing con la biblioteca estándar de Go.
- PostgreSQL con `database/sql` y `lib/pq`.
- Configuración por variables de entorno.
- Graceful shutdown con `SIGINT` y `SIGTERM`.
- Respuestas JSON y errores públicos seguros.
- Store layer para users y posts.
- Migraciones SQL.
- Unit tests para config y `/v1/status`.
- GitHub Actions para `gofmt`, `go vet` y `go test`.

## Inicio rápido

```bash
git clone https://github.com/dimbo1324/Base-Go-API-Engine.git
cd Base-Go-API-Engine
cp .env.example .env
docker compose up -d
make migrate-up
go run ./cmd/api
```

Verificación:

```bash
curl http://localhost:8080/v1/status
```

Respuesta:

```json
{"status":"ok"}
```

## Estructura

```text
cmd/api                    # punto de entrada
cmd/migrate/migrations     # migraciones SQL
internal/config            # carga y validación de configuración
internal/db                # conexión a PostgreSQL
internal/httpapi           # HTTP server, rutas, middleware, respuestas
internal/store             # modelos y persistencia SQL
```

## Principio principal

El proyecto debe mantenerse simple: la nueva lógica debe ir a lugares predecibles, las dependencias no deben complicarse sin motivo y la lógica de negocio no debe mezclarse con el entry point o la capa HTTP.
