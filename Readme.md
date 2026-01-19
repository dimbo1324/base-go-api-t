
# Base Go API Engine

A robust, modular, and scalable foundation for building RESTful APIs using **Go (Golang)** and **PostgreSQL**. This project implements a clean architecture pattern, separating configuration, database storage, and HTTP transport layers.

---

### 🌐 Documentation Translations

* [🇷🇺 Русский (Russian)](./docs/ReadmeRus.md)
* [🇪🇸 Español (Spanish)](./docs/ReadmeSp.md)

---

## ✨ Features

* **Modular Architecture**: Clean separation of concerns (CMD, Internal, Components, Store).
* 
**High-Performance Routing**: Built on top of [chi v5](https://github.com/go-chi/chi), a lightweight and idiomatic router.


* **Database Management**:
* PostgreSQL 16.3 integration via Docker.


* Connection pooling configuration (Max Open/Idle connections, Idle time).


* 
`citext` extension support for case-insensitive email handling.




* **Middleware Stack**:
* Request ID tagging.
* Real IP resolution.
* Structured Logging.
* Panic Recovery.
* Timeouts (60s global).




* 
**Environment Configuration**: Flexible config management with default fallback values.


* **Data Models**:
* **Users**: Storage logic for user management (Email, Username, Password).
* **Posts**: Storage logic for blog-style posts with tagging support.



## 🛠 Tech Stack

* **Language**: Go 1.22+
* **Database**: PostgreSQL 16.3
* **Router**: go-chi/chi/v5
* **Driver**: lib/pq
* **Containerization**: Docker & Docker Compose

## 🚀 Getting Started

### Prerequisites

* **Go** (version 1.21 or higher installed)
* **Docker** and **Docker Compose**

### Installation

1. **Clone the repository:**
```bash
git clone https://github.com/dimbo1324/base-go-api-t.git
cd base-go-api-t

```


2. **Start the Database:**
Use Docker Compose to spin up the PostgreSQL container.
```bash
docker-compose up -d

```


This will start a Postgres container named `postgres-db` on port `5432` with a persistent volume.


3. **Run Migrations:**
*Ensure the database is initialized with the schemas provided in `cmd/migrate/migrations`.*
The project includes SQL scripts to create the `users` and `posts` tables.


4. **Run the Application:**
```bash
go run cmd/api/main.go

```


The server will start on the address defined in your configuration (default: `:8080`).



## ⚙️ Configuration

The application uses environment variables for configuration. If variables are not set, it falls back to hardcoded defaults defined in `internal/config/config.go`.

| Variable | Default Value | Description              |
| -------- | ------------- | ------------------------ |
| `ADDR`   | `:8080`       | Server listening address |

 |
| `DB_ADDR` | `postgres://...` | Database connection string 

 |
| `DB_MAX_OPEN_CONNS` | `30` | Max open DB connections |
| `DB_MAX_IDLE_CONNS` | `30` | Max idle DB connections |
| `DB_MAX_IDLE_TIME_MINS` | `15m` | Max connection idle time 

 |

## 📂 Project Structure

```text
base-go-api-t/
├── cmd/
│   ├── api/
│   │   ├── components/     # App struct, router mounting, handlers
│   │   └── main.go         # Entry point, dependency injection
│   └── migrate/            # SQL migration files
├── internal/
│   ├── config/             # Configuration constants and keys
│   ├── db/                 # Low-level database connection logic
│   ├── env/                # Helper functions for reading env vars
│   └── store/              # Repository pattern implementations (Users, Posts)
├── scripts/
│   └── initDb.sql          # Database initialization script
├── docker-compose.yml      # Docker infrastructure definition
└── go.mod                  # Go module definitions

```

## 📡 API Endpoints

### System

* `GET /v1/status`: Health check endpoint. Returns "OK: it works".



*(Note: While the storage layer supports User and Post creation, specific HTTP handlers for these entities should be registered in `cmd/api/components/methods.go`).*

## 📄 License

This project is open source.

