# Go Workerpool File Processing

A high-performance concurrent CSV importer for PostgreSQL, written in Go. Designed to bulk-insert millions of rows fast using a worker-pool concurrency pattern with graceful shutdown support.

---

## Features

- Concurrent insertion via a configurable worker pool
- Graceful shutdown on `Ctrl+C` / `SIGTERM` — in-flight inserts finish, pending rows are drained cleanly
- Auto-migration — creates the `domain` table if it doesn't exist
- Environment-based configuration via `.env`
- PostgreSQL-native placeholder syntax (`$1, $2, ...`)
- Graceful connection pool management

---

## Project Structure

```
.
├── main.go         # Entry point, orchestrates the pipeline
├── config.go       # Loads and exposes config from .env
├── db.go           # Opens and configures the DB connection pool
├── migration.go    # Auto-creates the domain table if missing
├── worker.go       # Worker pool, CSV reader, and insert logic
├── csv.go          # Opens and returns the CSV reader
├── .env            # Environment variables (not committed)
└── .env.example    # Example env file to copy from
```

---

## Prerequisites

- Go 1.21+
- Docker & Docker Compose
- A copy of [`majestic_million.csv`](https://downloads.majestic.com/majestic_million.csv) in the project root

---

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/csv-to-pg.git
cd csv-to-pg
```

### 2. Set up environment variables

```bash
cp .env.example .env
```

Edit `.env` to match your setup:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=test
DB_SSL_MODE=disable
DB_MAX_CONNS=100
DB_MAX_IDLE_CONNS=4
TOTAL_WORKER=50
CSV_FILE=majestic_million.csv
```

### 3. Start PostgreSQL

```bash
docker-compose up -d
```

### 4. Install dependencies

```bash
go mod tidy
```

### 5. Run the importer

```bash
go run .
```

---

## Configuration Reference

| Variable           | Default                | Description                                      |
|--------------------|------------------------|--------------------------------------------------|
| `DB_HOST`          | `localhost`            | PostgreSQL host                                  |
| `DB_PORT`          | `5432`                 | PostgreSQL port                                  |
| `DB_USER`          | `postgres`             | Database user                                    |
| `DB_PASSWORD`      | `postgres`             | Database password                                |
| `DB_NAME`          | `test`                 | Target database name                             |
| `DB_SSL_MODE`      | `disable`              | SSL mode (`disable`, `require`, `verify-full`)   |
| `DB_MAX_CONNS`     | `100`                  | Max open DB connections                          |
| `DB_MAX_IDLE_CONNS`| `4`                    | Max idle DB connections                          |
| `TOTAL_WORKER`     | `50`                   | Number of concurrent insert workers              |
| `CSV_FILE`         | `majestic_million.csv` | Path to the CSV file to import                   |

> **Note:** Keep `TOTAL_WORKER` and `DB_MAX_CONNS` well below PostgreSQL's `max_connections` limit. The docker-compose config sets this to `200`.

---

## Database Schema

The `domain` table is created automatically on startup:

```sql
CREATE TABLE IF NOT EXISTS domain (
    GlobalRank      INT,
    TldRank         INT,
    Domain          VARCHAR(255),
    TLD             VARCHAR(255),
    RefSubNets      INT,
    RefIPs          INT,
    IDN_Domain      VARCHAR(255),
    IDN_TLD         VARCHAR(255),
    PrevGlobalRank  INT,
    PrevTldRank     INT,
    PrevRefSubNets  INT,
    PrevRefIPs      INT
);
```

---

## Docker Compose

```yaml
services:
  db:
    image: postgres:17-alpine
    restart: always
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME}
    ports:
      - "5432:5432"
    volumes:
      - ./postgres_data:/var/lib/postgresql/data
    command: postgres -c max_connections=200
    container_name: postgres

volumes:
  postgres_data:
```

---

## Dependencies

| Package                    | Purpose                        |
|----------------------------|--------------------------------|
| `github.com/lib/pq`        | PostgreSQL driver for `database/sql` |
| `github.com/joho/godotenv` | `.env` file loader             |

Install with:

```bash
go get github.com/lib/pq
go get github.com/joho/godotenv
```

---

## How It Works

1. `.env` is loaded and validated into a `Config` struct
2. A PostgreSQL connection pool is opened with the configured limits
3. The `domain` table is created if it doesn't already exist
4. A `context.WithCancel` root context and OS signal listener are set up for shutdown handling
5. The CSV file is opened and read line by line
6. Each row is dispatched as a job to a channel via the worker-pool concurrency pattern
7. A pool of goroutine workers picks up jobs and inserts rows concurrently using the shared context
8. A `WaitGroup` ensures the program waits for all in-flight work to complete before exiting

---

## Graceful Shutdown

The importer handles `Ctrl+C` (SIGINT) and `SIGTERM` (e.g. Docker/Kubernetes stop) without data corruption or hanging goroutines.

When a signal is received:

- The root `context` is cancelled and a shutdown message is logged
- The CSV reader stops dispatching new rows immediately
- Any row blocked on the `jobs` channel is released via a `select/ctx.Done()` arm and its `WaitGroup` counter is decremented
- Active workers check `ctx.Err()` before each insert — already-started DB calls are also cancelled via `ExecContext(ctx, ...)`
- Workers continue draining the `jobs` channel, calling `wg.Done()` for each remaining job without inserting, so the `WaitGroup` always reaches zero
- The program exits cleanly and reports total elapsed time

```
^C
=> received signal: interrupt, shutting down gracefully...
=> stopping csv dispatch due to shutdown
=> import interrupted before completion
done in 4 seconds
```

---

## License

MIT