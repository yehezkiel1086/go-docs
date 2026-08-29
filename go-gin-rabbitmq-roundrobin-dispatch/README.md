# RabbitMQ Work Queues in Go

This project is a practical implementation of the **Work Queues** (also known as **Task Queues**) pattern using RabbitMQ and Go. It demonstrates how to distribute time-consuming tasks among multiple workers using a round-robin dispatching strategy.

The system consists of:
*   A **master service** (producer) that exposes an API to receive tasks and publish them as messages to a RabbitMQ queue.
*   Two **worker services** (consumers) that concurrently pull tasks from the queue and process them.

## How It Works

1.  **Producer (`master-service`)**:
    *   An HTTP server built with Gin listens for POST requests on `/api/v1/publish`.
    *   When a request with a message is received, it publishes the message to a durable RabbitMQ queue named `message`.
    *   Messages are marked as `persistent` to ensure they survive a RabbitMQ broker restart.

2.  **Consumers (`worker-service-1`, `worker-service-2`)**:
    *   Both worker services connect to the same `message` queue.
    *   RabbitMQ dispatches messages to the workers in a **round-robin** fashion. For example, if you send three messages, the first goes to worker 1, the second to worker 2, and the third back to worker 1.
    *   Each worker receives a message, simulates work (by sleeping for a duration proportional to the number of dots `.` in the message), and then sends a manual acknowledgment (`ack`) back to RabbitMQ.
    *   Manual acknowledgment ensures that if a worker dies while processing a task, the message is not lost and will be redelivered to another worker.
    *   The `prefetch count` is set to `1`, meaning RabbitMQ will only give one message at a time to each worker. It won't dispatch a new message to a worker until that worker has processed and acknowledged the previous one.

!Work Queues Diagram
*Image from rabbitmq.com tutorials*

## Prerequisites

*   Go (version 1.24 or higher)
*   Docker and Docker Compose (for running RabbitMQ)
*   curl or an API client like Postman to send requests.

## Getting Started

### 1. Set up RabbitMQ

A `docker-compose.yml` file is provided to easily run a RabbitMQ instance.

```yaml
# docker-compose.yml
version: '3.8'
services:
  rabbitmq:
    image: rabbitmq:3-management
    container_name: rabbitmq
    ports:
      - "5672:5672"  # For AMQP protocol
      - "15672:15672" # For management UI
    environment:
      - RABBITMQ_DEFAULT_USER=guest
      - RABBITMQ_DEFAULT_PASS=guest
```

Run the following command in the project root to start the container:
```bash
docker-compose up -d
```
You can access the RabbitMQ management UI at `http://localhost:15672` (user: `guest`, pass: `guest`).

### 2. Configure Environment Variables

Each service requires a `.env` file for configuration. Create a `.env` file inside each of the three service directories (`master-service`, `worker-service-1`, `worker-service-2`).

**For `master-service/.env`:**
```env
# HTTP Server
HTTP_HOST=0.0.0.0
HTTP_PORT=8080

# RabbitMQ
RABBITMQ_USER=guest
RABBITMQ_PASS=guest
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
```

**For `worker-service-1/.env` and `worker-service-2/.env`:**
```env
# RabbitMQ
RABBITMQ_USER=guest
RABBITMQ_PASS=guest
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
```

### 3. Run the Services

Open three separate terminal windows.

**Terminal 1: Run the Master Service**
```bash
cd master-service
go run main.go
```

**Terminal 2: Run Worker 1**
```bash
cd worker-service-1
go run main.go
```

**Terminal 3: Run Worker 2**
```bash
cd worker-service-2
go run main.go
```

### 4. Publish Messages

Use `curl` to send messages to the master service. The number of dots (`.`) determines how long the "task" will take.

```bash
# This task will take 1 second
curl -X POST http://localhost:8080/api/v1/publish \
  -H "Content-Type: application/json" \
  -d '{"message": "First message."}'

# This task will take 5 seconds
curl -X POST http://localhost:8080/api/v1/publish \
  -H "Content-Type: application/json" \
  -d '{"message": "Second message....."}'

# This task will take 2 seconds
curl -X POST http://localhost:8080/api/v1/publish \
  -H "Content-Type: application/json" \
  -d '{"message": "Third message.."}'
```

Observe the logs in the worker terminals. You will see the messages being distributed between `worker 1` and `worker 2` in a round-robin fashion.

*   Worker 1 will receive "First message."
*   Worker 2 will receive "Second message....."
*   Worker 1 will receive "Third message.."

This demonstrates that even though the second task is long, the workers continue to pick up new tasks as they become available.