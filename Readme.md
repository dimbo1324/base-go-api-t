# Base Go API Engine

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.3-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-Free_to_Use-green.svg)](LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/dimbo1324)

> **A robust, modular, and scalable foundation for building RESTful APIs with Go (Golang) and PostgreSQL.**

---

### 🌐 Documentation / Документация / Documentación

* 🇷🇺 [**Russian (Русский)**](docs/ReadmeRus.md)
* 🇪🇸 [**Spanish (Español)**](docs/ReadmeSp.md)

---

## 📖 Overview

**Base Go API Engine** is a production-ready template designed to jumpstart your backend development. It implements a clean architecture pattern, separating concerns between configuration, database management, and business logic.

Whether you are learning Go or building a complex microservice, this engine provides the essential plumbing—database connections, configuration management, and routing—so you can focus on building features.

### ✨ Key Features

* **Modular Architecture:** Clean separation of `cmd`, `internal`, and `components`.
* **High-Performance Routing:** Built on top of `chi` v5 for lightweight and idiomatic routing.
* **PostgreSQL Integration:** Pre-configured with connection pooling and `citext` extension support.
* **Docker Ready:** Includes a `docker-compose.yml` for instant database setup.
* **Smart Configuration:** robust environment variable management with sensible defaults.
* **Scalable Store Pattern:** Ready-to-use interfaces for User and Post management.

---

## 🛠️ Tech Stack

* **Language:** [Go (Golang)](https://go.dev/)
* **Database:** [PostgreSQL](https://www.postgresql.org/)
* **Router:** [go-chi/chi](https://github.com/go-chi/chi)
* **Driver:** [lib/pq](https://github.com/lib/pq)
* **Containerization:** Docker & Docker Compose

---

## 🚀 Getting Started

Follow these steps to get a local copy up and running.

### Prerequisites

* **Go**: Version 1.22 or higher installed on your machine.
* **Docker Desktop**: For running the database container.
* **Git**: To clone the repository.

### 1. Clone the Repository

```bash
git clone [https://github.com/dimbo1324/Base-Go-API-Engine.git](https://github.com/dimbo1324/Base-Go-API-Engine.git)
cd Base-Go-API-Engine

```

### 2. Environment Configuration

The application is designed to run out of the box with defaults, but you can customize it using Environment Variables.

| Variable                | Description                | Default Value                                     |
| ----------------------- | -------------------------- | ------------------------------------------------- |
| `ADDR`                  | Server Address             | `:8080`                                           |
| `DB_ADDR`               | Database Connection String | `postgres://postgres:password@localhost/appdb...` |
| `DB_MAX_OPEN_CONNS`     | Max Open DB Connections    | `30`                                              |
| `DB_MAX_IDLE_CONNS`     | Max Idle DB Connections    | `30`                                              |
| `DB_MAX_IDLE_TIME_MINS` | Connection Max Lifetime    | `15m`                                             |

### 3. Start the Database

We use Docker Compose to spin up a PostgreSQL instance with the correct settings.

```bash
docker-compose up -d

```

*This starts a PostgreSQL container named `postgres-db` on port `5432`.*

### 4. Database Migration

The project includes SQL migration files in `cmd/migrate/migrations`. You will need to apply these to create the `users` and `posts` tables.

You can execute the SQL files using a database tool (like DBeaver or pgAdmin) or via command line:

```bash
# Example using psql inside the container
docker exec -it postgres-db psql -U postgres -d appdb -f /path/to/000001_create_users.up.sql

```

### 5. Run the Application

```bash
go run cmd/api/main.go

```

You should see the output:

```text
Server started on :8080

```

---

## 📂 Project Structure

The project follows the Standard Go Project Layout:

```text
Base-Go-API-Engine/
├── cmd/
│   ├── api/            # Main application entry point
│   └── migrate/        # Database migration SQL scripts
├── internal/
│   ├── config/         # Configuration constants and defaults
│   ├── db/             # Database connection logic
│   ├── env/            # Environment variable helpers
│   └── store/          # Data Access Layer (Repository Pattern)
├── docker-compose.yml  # Docker services definition
└── go.mod              # Go module dependencies

```

---

## 🔌 API Endpoints

### System

| Method | Endpoint     | Description                                |
| ------ | ------------ | ------------------------------------------ |
| `GET`  | `/v1/status` | Health check to verify the API is running. |

*Note: The User and Post logic is implemented in the `internal/store` package and ready to be connected to new HTTP handlers.*

---

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📜 License

This project is free to use.

---

## 📬 Contact

If you have questions, suggestions, or just want to say hi, feel free to reach out!

* **Author:** dimbo1324
* **Telegram:** [@dimbo1324](https://t.me/dimbo1324)
* **Email:** dimaprihodko180@gmail.com
* **GitHub:** [github.com/dimbo1324](https://github.com/dimbo1324)

---

*Developed with ❤️ by dimbo1324*

