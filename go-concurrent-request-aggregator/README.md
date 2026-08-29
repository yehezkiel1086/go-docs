# Go Concurrent Request Aggregator

A demonstration of Go's concurrency patterns using **fan-in fan-out pipeline** for aggregating data from external APIs.

## Overview

This project implements a concurrent request aggregation service that fetches user data and their posts from [JSONPlaceholder](https://jsonplaceholder.typicode.com/) API simultaneously, combining the results into a unified dashboard response.

## Architecture

### Concurrency Pattern: Fan-Out Fan-In

```
                    ┌─────────────┐
                    │   Client    │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  Dashboard  │
                    │   Service   │
                    └──────┬──────┘
                           │
           ┌───────────────┴───────────────┐
           │                               │
     ┌─────▼─────┐                 ┌───────▼────────┐
     │ Fan-Out   │                 │    Fan-In      │
     │           │                 │                  │
┌────┴────┐ ┌────┴────┐      ┌────┴───┐      ┌──────┴───┐
│  User   │ │  Posts  │      │  User   │      │  Posts   │
│  Repo   │ │  Repo   │      │ Result  │      │ Result   │
└────┬────┘ └────┬────┘      └────┬────┘      └─────┬────┘
     │           │                │                │
     └─────┬─────┘                └────────┬───────┘
           │                               │
     ┌─────▼─────┐                 ┌──────▼──────┐
     │  External │                 │  Dashboard  │
     │    API    │                 │   Response  │
     └───────────┘                 └─────────────┘
```

### How It Works

1. **Fan-Out**: When a dashboard request is received, the service spawns two concurrent goroutines:
   - One fetches user data from `/users/{id}`
   - One fetches posts from `/posts?userId={id}`

2. **Concurrent Execution**: Both requests execute in parallel, reducing total latency from `t1 + t2` to `max(t1, t2)`.

3. **Fan-In**: Results are collected via channels using a `select` statement that waits for both goroutines to complete.

4. **Aggregation**: User and posts data are combined into a single `Dashboard` domain object.

## Project Structure

```
.
├── internal/
│   ├── adapter/
│   │   ├── api/              # External API client
│   │   │   ├── api.go        # HTTP client wrapper
│   │   │   └── repository/   # Repository implementations
│   │   │       ├── user.go
│   │   │       └── post.go
│   │   └── config/           # Configuration management
│   └── core/
│       ├── domain/           # Domain models
│       │   ├── user.go
│       │   ├── post.go
│       │   └── dashboard.go
│       ├── port/             # Interfaces (ports)
│       └── service/          # Business logic
│           └── dashboard.go  # Fan-in fan-out implementation
```

## Key Implementation

### Dashboard Service (`internal/core/service/dashboard.go`)

```go
func (s *DashboardService) GetDashboard(userId string) (*domain.Dashboard, error) {
    // fan-out: start concurrent requests
    userCh := make(chan userResult, 1)
    postsCh := make(chan postsResult, 1)

    go func() {
        user, err := s.userRepo.GetUser(userId)
        userCh <- userResult{user, err}
    }()

    go func() {
        posts, err := s.postRepo.GetPostsByUserId(userId)
        postsCh <- postsResult{posts, err}
    }()

    // fan-in: collect results
    var dashboard domain.Dashboard
    var dashboardErr error

    for range 2 {
        select {
        case res := <-userCh:
            if res.err != nil {
                dashboardErr = res.err
            } else {
                dashboard.User = res.user
            }
        case res := <-postsCh:
            if res.err != nil {
                dashboardErr = res.err
            } else {
                dashboard.Posts = res.posts
            }
        }
    }

    return &dashboard, dashboardErr
}
```

## Configuration

Create a `.env` file:

```env
APP_NAME=go-concurrent-request-aggregator
APP_ENV=development

HTTP_HOST=127.0.0.1
HTTP_PORT=8080

API_BASE=https://jsonplaceholder.typicode.com/
```

## Running the Project

```bash
# Install dependencies
go mod download

# Run the application
go run cmd/app/main.go
```

## API Endpoints

### Get Dashboard

```
GET /dashboard/{userId}
```

Response:

```json
{
  "user": {
    "id": "1",
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz"
  },
  "posts": [
    {
      "id": "1",
      "userId": "1",
      "title": "sunt aut facere...",
      "content": "quia et suscipit..."
    }
  ]
}
```

## Technologies

- **Go 1.23+** - Concurrency primitives (goroutines, channels, `select`)
- **JSONPlaceholder** - Mock REST API for testing
- **Clean Architecture** - Ports and Adapters pattern

## Concurrency Benefits

| Approach | Latency |
|----------|---------|
| Sequential | ~400ms (user) + ~600ms (posts) = **1000ms** |
| Concurrent | max(400ms, 600ms) = **~600ms** |

The fan-in fan-out pattern provides **~40% latency reduction** by parallelizing independent I/O operations.

## License

MIT
