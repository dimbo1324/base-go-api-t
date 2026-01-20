# Base Go API Engine

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.3-316192?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-Free_to_Use-green.svg)](LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/dimbo1324)

> **Una base robusta, modular y escalable para construir APIs RESTful con Go (Golang) y PostgreSQL.**

---

### 🌐 Documentación en otros idiomas

* 🇺🇸 [**English (Inglés)**](../Readme.md)
* 🇷🇺 [**Russian (Ruso)**](ReadmeRus.md)

---

## 📖 Resumen

**Base Go API Engine** es una plantilla lista para producción diseñada para acelerar tu desarrollo backend. Implementa un patrón de arquitectura limpia (Clean Architecture), separando las preocupaciones entre la configuración, la gestión de la base de datos y la lógica de negocio.

Ya sea que estés aprendiendo Go o construyendo un microservicio complejo, este motor te proporciona toda la "tubería" esencial —conexiones a base de datos, gestión de configuración y enrutamiento— para que vos solo te preocupes por desarrollar las funcionalidades.

### ✨ Características Clave

* [cite_start]**Arquitectura Modular:** Clara separación entre `cmd`, `internal` y `components`[cite: 3].
* [cite_start]**Enrutamiento de Alto Rendimiento:** Construido sobre `chi` v5 para un enrutamiento ligero e idiomático[cite: 3].
* [cite_start]**Integración con PostgreSQL:** Preconfigurado con pool de conexiones y soporte para la extensión `citext`[cite: 1, 5].
* [cite_start]**Listo para Docker:** Incluye un `docker-compose.yml` para levantar la base de datos al instante[cite: 1].
* [cite_start]**Configuración Inteligente:** Gestión robusta de variables de entorno con valores predeterminados sensatos[cite: 9].
* [cite_start]**Patrón de Almacenamiento Escalable:** Interfaces listas para usar para la gestión de Usuarios (Users) y Publicaciones (Posts)[cite: 11, 12].

---

## 🛠️ Tecnologías

* **Lenguaje:** [Go (Golang)](https://go.dev/)
* **Base de Datos:** [PostgreSQL](https://www.postgresql.org/)
* **Router:** [go-chi/chi](https://github.com/go-chi/chi)
* **Driver:** [lib/pq](https://github.com/lib/pq)
* **Contenedores:** Docker y Docker Compose

---

## 🚀 Cómo Empezar

Seguí estos pasos para tener una copia local funcionando en tu máquina.

### Requisitos Previos

* **Go**: Versión 1.22 o superior instalada.
* **Docker Desktop**: Para correr el contenedor de la base de datos.
* **Git**: Para clonar el repositorio.

### 1. Cloná el Repositorio

```bash
git clone [https://github.com/dimbo1324/Base-Go-API-Engine.git](https://github.com/dimbo1324/Base-Go-API-Engine.git)
cd Base-Go-API-Engine

```

### 2. Configuración del Entorno

La aplicación está diseñada para funcionar "out of the box" con valores por defecto, pero podés personalizarla usando Variables de Entorno.

| Variable                | Descripción                | Valor por Defecto                                 |
| ----------------------- | -------------------------- | ------------------------------------------------- |
| `ADDR`                  | Dirección del Servidor     | `:8080`                                           |
| `DB_ADDR`               | Cadena de Conexión BD      | `postgres://postgres:password@localhost/appdb...` |
| `DB_MAX_OPEN_CONNS`     | Máx. Conexiones Abiertas   | `30`                                              |
| `DB_MAX_IDLE_CONNS`     | Máx. Conexiones Inactivas  | `30`                                              |
| `DB_MAX_IDLE_TIME_MINS` | Tiempo de Vida de Conexión | `15m`                                             |



### 3. Iniciá la Base de Datos

Usamos Docker Compose para levantar una instancia de PostgreSQL con la configuración correcta.

```bash
docker-compose up -d

```

Esto inicia un contenedor PostgreSQL llamado `postgres-db` en el puerto `5432`.

### 4. Migración de la Base de Datos

El proyecto incluye archivos de migración SQL en `cmd/migrate/migrations`. Vas a necesitar aplicarlos para crear las tablas `users` y `posts`.

Podés ejecutar los archivos SQL usando una herramienta de base de datos (como DBeaver o pgAdmin) o vía línea de comandos:

```bash
# Ejemplo usando psql dentro del contenedor
docker exec -it postgres-db psql -U postgres -d appdb -f /path/to/000001_create_users.up.sql

```



### 5. Corré la Aplicación

```bash
go run cmd/api/main.go

```

Deberías ver la salida:

```text
Сервер запущен на :8080

```

*(El log por defecto está en ruso según el código fuente, podés cambiarlo en `cmd/api/components/methods.go`)*.

---

## 📂 Estructura del Proyecto

El proyecto sigue el diseño estándar de proyectos en Go (Standard Go Project Layout):

```text
Base-Go-API-Engine/
├── cmd/
[cite_start]│   ├── api/            # Punto de entrada principal de la aplicación [cite: 2]
[cite_start]│   └── migrate/        # Scripts SQL de migración [cite: 4]
├── internal/
[cite_start]│   ├── config/         # Constantes de configuración y defaults [cite: 9]
[cite_start]│   ├── db/             # Lógica de conexión a la base de datos [cite: 10]
[cite_start]│   ├── env/            # Ayudantes para variables de entorno [cite: 11]
[cite_start]│   └── store/          # Capa de Acceso a Datos (Repository Pattern) [cite: 12]
[cite_start]├── docker-compose.yml  # Definición de servicios Docker [cite: 1]
└── go.mod              # Dependencias del módulo Go

```

---

## 🔌 Endpoints de la API

### Sistema

| Método | Endpoint     | Descripción                                                         |
| ------ | ------------ | ------------------------------------------------------------------- |
| `GET`  | `/v1/status` | Chequeo de salud (Health check) para verificar que la API funciona. |

Nota: La lógica de Usuarios y Publicaciones está implementada en el paquete `internal/store` y lista para ser conectada a nuevos handlers HTTP.

---

## 🤝 Cómo Colaborar

Las contribuciones son lo que hacen a la comunidad de código abierto un lugar increíble para aprender, inspirar y crear. Cualquier contribución que hagas es **muy agradecida**.

1. Hacé un Fork del Proyecto.
2. Creá tu Rama de Funcionalidad (`git checkout -b feature/FeatureIncreible`).
3. Hacé Commit de tus cambios (`git commit -m 'Agrego una FeatureIncreible'`).
4. Hacé Push a la Rama (`git push origin feature/FeatureIncreible`).
5. Abrí un Pull Request.

---

## 📜 Licencia

Este proyecto es libre de uso. Sentite libre de utilizarlo para tus propios desarrollos.

---

## 📬 Contacto

Si tenés preguntas, sugerencias o simplemente querés saludar, ¡no dudes en escribirme!

* **Autor:** dimbo1324
* **Telegram:** [@dimbo1324](https://t.me/dimbo1324)
* **Email:** dimaprihodko180@gmail.com
* **GitHub:** [github.com/dimbo1324](https://github.com/dimbo1324)

---

*Desarrollado con ❤️ por dimbo1324*