# Go JWT Auth

This project demonstrates a production-ready authentication system using Go, Gin, PostgreSQL, and Redis. It covers the full auth lifecycle — registration, login, token refresh, and logout — with security hardening applied at every layer.

## Features

### Authentication
- **JWT access & refresh tokens** — short-lived access tokens (minutes) paired with long-lived refresh tokens (days)
- **Refresh token rotation** — every refresh issues a new refresh token and invalidates the old one, preventing token reuse attacks
- **Secure cookie storage** — tokens are stored in `HttpOnly`, `Secure` cookies, never exposed to JavaScript
- **Password hashing** — passwords are hashed using bcrypt before storage

### Authorization
- **Role-based access control (RBAC)** — `UserRole` and `AdminRole` with middleware-enforced route protection
- **Self-or-admin middleware** — users can only access their own resources; admins can access any

### Security Hardening
- **Secure HTTP headers** — `X-Frame-Options`, `X-Content-Type-Options`, `Strict-Transport-Security`, `Content-Security-Policy`, `Referrer-Policy`, `Permissions-Policy`
- **CORS** — origin allowlist with credentials support, configurable via environment variables
- **Rate limiting** — global rate limiter (60 req/min) and strict limiter on auth endpoints (10 req/min) backed by Redis
- **Server fingerprint removal** — `Server` header is stripped from all responses

### Architecture
- **Hexagonal architecture** — domain, port, service, repository, and handler layers are fully separated
- **Interface-driven** — all dependencies are injected via interfaces for testability
- **Redis-backed state** — refresh tokens and rate limit counters are stored in Redis with automatic TTL expiry

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go |
| Web framework | Gin |
| Database | PostgreSQL + GORM |
| Cache / token store | Redis |
| Authentication | JWT (`golang-jwt/jwt`) |
| Config | `godotenv` |

## Project Structure

```
.
├── cmd/
│   └── main.go
└── internal/
    ├── adapter/
    │   ├── config/
    │   ├── handler/         # gin handlers, middleware, router
    │   └── storage/
    │       ├── postgres/
    │       │   └── repository/
    │       └── redis/
    └── core/
        ├── domain/          # user, jwt claims
        ├── port/            # repository and service interfaces
        ├── service/         # auth, user business logic
        └── util/            # jwt, password, rate limiter
```

## Environment Variables

```env
APP_NAME=go-jwt-auth
APP_ENV=development

HTTP_HOST=localhost
HTTP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_jwt_auth

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

ACCESS_TOKEN_SECRET=your-access-secret
REFRESH_TOKEN_SECRET=your-refresh-secret
ACCESS_TOKEN_DURATION=15
REFRESH_TOKEN_DURATION=7

CORS_ALLOWED_ORIGINS=http://localhost:3000
```

## API Endpoints

### Public
| Method | Path | Description |
|---|---|---|
| `POST` | `/api/v1/register` | register a new user |
| `POST` | `/api/v1/auth/login` | login and receive tokens |
| `POST` | `/api/v1/auth/refresh` | rotate refresh token and get new access token |
| `POST` | `/api/v1/auth/logout` | revoke refresh token |

### Authenticated users
| Method | Path | Description |
|---|---|---|
| `GET` | `/api/v1/users/:id` | get own profile |
| `PUT` | `/api/v1/users/:id` | update own profile |
| `DELETE` | `/api/v1/users/:id` | delete own account |

### Admin only
| Method | Path | Description |
|---|---|---|
| `GET` | `/api/v1/admin/users` | list all users |
| `GET` | `/api/v1/admin/users/:id` | get any user by id |
| `DELETE` | `/api/v1/admin/users/:id` | delete any user |

## Getting Started

```bash
# clone the repo
git clone https://github.com/yehezkiel1086/go-jwt-auth
cd go-jwt-auth

# copy and fill in environment variables
cp .env.example .env

# run the server
go run cmd/main.go
```