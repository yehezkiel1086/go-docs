# Go Photo API with Redis Caching

This is a sample Go application that demonstrates how to build a REST API using Redis as Caching with Gin framework. It fetches a list of photos from the [JSONPlaceholder](https://jsonplaceholder.typicode.com/) API and implements a caching layer with Redis to optimize response times and reduce external API calls.

## Features

- **RESTful API**: A simple endpoint to retrieve a list of photos.
- **Caching**: Uses Redis to cache responses from the external API, significantly improving performance on subsequent requests.
- **Containerized**: Comes with a `docker-compose.yml` file to easily spin up the application and its Redis dependency.
- **Configuration Driven**: Manages application settings through environment variables.
- **Clean Architecture**: The project is structured with a clear separation of concerns (router, controllers, storage).

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go (version 1.18 or higher)
- Docker
- Docker Compose

## Getting Started

Follow these steps to get the application up and running.

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-photo-api-redis.git
cd go-photo-api-redis
```

### 2. Configure Environment Variables

Create a `.env` file by copying the example file.

```bash
cp .env.example .env
```

The default values in `.env.example` are configured to work with the `docker-compose` setup. You can customize them if needed.

```ini
# .env
APP_NAME=go-photo-api-redis
APP_ENV=development

HTTP_HOST=127.0.0.1
HTTP_PORT=3500

REDIS_URL=127.0.0.1:6379 # Note: When running with Docker, the Go app connects to 'redis:6379'
REDIS_PASS=your-strong-password
REDIS_DB=0
```

> **Note**: Make sure to set a secure `REDIS_PASS` in your `.env` file.

### 3. Run with Docker Compose

The easiest way to run the application and its dependencies is with Docker Compose.

```bash
# Build and start the services in the background
docker-compose up --build -d
```

The API will be available at `http://127.0.0.1:3500`.

### 4. (Optional) Run Locally

If you prefer to run the Go application directly on your host machine, ensure you have a Redis instance running and accessible.

```bash
# Install dependencies
go mod tidy

# Run the application
go run cmd/main.go
```

## API Endpoints

### Get All Photos

- **Endpoint**: `GET /photos`
- **Description**: Retrieves a list of all photos. On the first request, it fetches data from the JSONPlaceholder API and caches the result in Redis. Subsequent requests will be served directly from the cache.
- **Success Response (200 OK)**:
  ```json
  [
      {
          "albumId": 1,
          "id": 1,
          "title": "accusamus beatae ad facilis cum similique qui sunt",
          "url": "https://via.placeholder.com/600/92c952",
          "thumbnailUrl": "https://via.placeholder.com/150/92c952"
      },
      // ... more photos
  ]
  ```

### Get All Photos in an Album

- **Endpoint**: `GET /photos/:albumId`
- **Description**: Retrieves a list of all photos based on `albumId`.
- **Success Response (200 OK)**:
  ```json
  [
      {
          "albumId": 1,
          "id": 1,
          "title": "accusamus beatae ad facilis cum similique qui sunt",
          "url": "https://via.placeholder.com/600/92c952",
          "thumbnailUrl": "https://via.placeholder.com/150/92c952"
      },
      // ... more photos
  ]
  ```