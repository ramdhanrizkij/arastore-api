# Arastore API

A RESTful API for an e-commerce platform built with Go. This API provides authentication features (JWT + Refresh Token), user management with RBAC (Role-Based Access Control), product categories, products, and shipping addresses.

## Architecture & Code Style

This project adopts **Clean Architecture** with a **feature-based structure**. Each feature is organized independently with clear separation of concerns:

```
internal/
├── core/                 # Infrastructure & cross-cutting concerns
│   ├── config/           # Configuration from environment variables
│   ├── database/         # PostgreSQL connection via GORM
│   ├── middleware/        # JWT auth, RBAC, logging, recovery
│   ├── server/           # Fiber HTTP server & route wiring
│   ├── storage/          # File storage (local / MinIO)
│   └── worker/           # Background worker pool & scheduler
│
├── features/             # Domain-driven feature modules
│   ├── auth/             # Authentication (register, login, refresh, logout)
│   ├── user/             # User management
│   ├── role/             # Role management
│   └── permission/       # Permission management
│
├── model/                # GORM models (database entities)
├── seeder/               # Database seeders
└── shared/               # Shared utilities (errors, pagination, response, validator)
```

Each feature module follows a consistent structure:

```
feature/
├── domain/
│   ├── interface.go      # Service & Repository interfaces
│   └── dto.go            # Request/Response DTOs
├── handler/
│   ├── http_handler.go   # HTTP request handlers
│   └── route.go          # Route registration
├── repository/            # Database operations
└── service/               # Business logic
```

### Design Principles

- **Interface-based contracts** — The domain layer defines interfaces, implementations are separate
- **Manual Dependency Injection** — Wiring is done in `server.SetupRoutes()` without a DI framework
- **Standard response envelope** — All responses follow the format `{ meta, data, pagination }`
- **Pluggable storage** — File storage can be switched between local and MinIO via config
- **Graceful shutdown** — Server, worker, and database connections are properly closed on OS signal

## Tech Stack

| Technology | Purpose |
|---|---|
| [Go](https://go.dev/) 1.26.1 | Programming language |
| [Fiber](https://gofiber.io/) v3 | HTTP framework |
| [GORM](https://gorm.io/) | ORM for PostgreSQL |
| [PostgreSQL](https://www.postgresql.org/) 17 | Database |
| [Goose](https://github.com/pressly/goose) | Database migration |
| [JWT](https://github.com/golang-jwt/jwt) | Authentication (HS256) |
| [go-playground/validator](https://github.com/go-playground/validator) | Request validation |
| [Redis](https://redis.io/) | Caching (optional) |
| [MinIO](https://min.io/) | S3-compatible object storage (optional) |
| [Zap](https://uber.github.io/zap/) | Structured logging |
| [Docker](https://www.docker.com/) | Containerization |

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.26.1 or later
- [PostgreSQL](https://www.postgresql.org/download/) 17
- [Docker](https://docs.docker.com/get-docker/) & Docker Compose (optional, for dependency services)
- [Make](https://www.gnu.org/software/make/) (optional, for running commands)

### Clone Repository

```bash
git clone https://github.com/ramdhanrizkij/arastore-api.git
cd arastore-api
```

### Setup Environment Variables

```bash
cp .env.example .env
```

Edit the `.env` file according to your local configuration. At minimum, update the database section:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=bytecode_api

JWT_SECRET=your-super-secret-key
```

### Install Dependencies

```bash
go mod download
```

### Run Services with Docker Compose

Start PostgreSQL (and additional services like Redis/MinIO if needed):

```bash
# PostgreSQL only
docker compose up -d

# PostgreSQL + Redis + MinIO
COMPOSE_PROFILES=redis,minio docker compose up -d
```

### Run Migrations

```bash
make migrate-up
```

Migrations will create all necessary tables, including default seed data for roles and permissions.

To check migration status:

```bash
make migrate-status
```

### Run Seeders

```bash
make seeder-run
```

Seeders will create the following initial data:

| Data | Details |
|---|---|
| Roles | `superadmin`, `admin`, `user` |
| Users | 3 accounts (password: `password`) — `superadmin@arastore.id`, `admin@arastore.id`, `budi@arastore.id` |
| Addresses | 2 addresses for Budi (Jakarta) |
| Categories | Electronics, Fashion, Food & Beverage, Home & Living, Sports |
| Products | 5 sample products |

### Run the Project

```bash
# Run API server
make run

# Run background worker (optional, for scheduled tasks)
make run-worker
```

The API server will run at `http://localhost:8080`.

## Make Commands

| Command | Description |
|---|---|
| `make build` | Build all binaries (`api`, `worker`, `migrate`) |
| `make run` | Run the API server |
| `make run-worker` | Run the background worker |
| `make test` | Run all tests |
| `make test-unit` | Run unit tests only |
| `make test-integration` | Run integration tests (requires Docker) |
| `make tidy` | Run `go mod tidy` |
| `make migrate-up` | Run all pending migrations |
| `make migrate-down` | Rollback the last migration |
| `make migrate-redo` | Rollback and re-apply the last migration |
| `make migrate-status` | Show migration status |
| `make migrate-create name=xxx` | Create a new migration file |
| `make seeder-run` | Run database seeders |

## API Endpoints

All endpoints are under `/api/v1`.

### Public

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/health` | Health check |
| `POST` | `/api/v1/auth/register` | Register a new user |
| `POST` | `/api/v1/auth/login` | Login (returns JWT + refresh token) |
| `POST` | `/api/v1/auth/refresh` | Refresh access token |
| `POST` | `/api/v1/auth/logout` | Logout (revoke refresh token) |

### Protected (JWT + Permission)

**Users** `/api/v1/users`:

| Method | Endpoint | Permission |
|---|---|---|
| `GET` | `/me` | JWT only |
| `PUT` | `/me/profile` | JWT only |
| `GET` | `/me/permissions` | JWT only |
| `GET` | `/` | `users.view` |
| `GET` | `/:id` | `users.view` |
| `POST` | `/` | `users.create` |
| `PUT` | `/:id` | `users.edit` |
| `DELETE` | `/:id` | `users.delete` |

**Roles** `/api/v1/roles`:

| Method | Endpoint | Permission |
|---|---|---|
| `GET` | `/` | `roles.view` |
| `GET` | `/:id` | `roles.view` |
| `POST` | `/` | `roles.create` |
| `PUT` | `/:id` | `roles.edit` |
| `DELETE` | `/:id` | `roles.delete` |
| `POST` | `/:id/permissions` | `roles.assign-permission` |
| `DELETE` | `/:id/permissions` | `roles.remove-permission` |

**Permissions** `/api/v1/permissions`:

| Method | Endpoint | Permission |
|---|---|---|
| `GET` | `/` | `permissions.view` |
| `GET` | `/:id` | `permissions.view` |
| `POST` | `/` | `permissions.create` |
| `PUT` | `/:id` | `permissions.edit` |
| `DELETE` | `/:id` | `permissions.delete` |

> **Note:** The `superadmin` role bypasses all permission checks.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
