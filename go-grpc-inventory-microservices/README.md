# Go gRPC Inventory Microservices

A distributed inventory management system built with Go, featuring microservices architecture using **gRPC** for synchronous inter-service communication and **RabbitMQ** for asynchronous event-driven messaging.

![Inventory Microservices High Level Diagram](/assets/high-level-diagram.png)

## Architecture Overview

This project demonstrates a hexagonal/clean architecture implementation of an e-commerce inventory system with the following services:

- [x] **User Service** - Authentication, authorization, and user management
- [x] **Product Service** - Product catalog management (gRPC server)
- [x] **Inventory Service** - Stock/inventory tracking (gRPC server)
- [x] **Order Service** - Order processing with cross-service integration
- [x] **Notification Service** - Event-driven notifications via RabbitMQ

## Tech Stack

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.25 |
| **HTTP Framework** | Gin Gonic |
| **RPC Framework** | gRPC with Protocol Buffers |
| **Message Broker** | RabbitMQ |
| **Database** | PostgreSQL |
| **ORM** | GORM |
| **Authentication** | JWT (JSON Web Tokens) |
| **Task Runner** | Taskfile |
| **Containerization** | Docker & Docker Compose |

## Service Communication Patterns

### Synchronous Communication (gRPC)
- **Product Service** ↔ **Order Service**: Product price lookup
- **Inventory Service** ↔ **Order Service**: Stock verification and deduction

### Asynchronous Communication (RabbitMQ)
- **Order Service** → **Notification Service**: Order placement notifications via `notification` queue

## Project Structure

```
.
├── services/
│   ├── user-service/          # HTTP API (Port 8081)
│   ├── product-service/       # gRPC Server (Port 50022) + HTTP (Port 8082)
│   ├── inventory-service/     # gRPC Server (Port 50021)
│   ├── order-service/         # HTTP API (Port 8083)
│   └── notif-service/         # HTTP API + RabbitMQ Consumer (Port 8084)
├── protobuf/
│   ├── product.proto          # Product service definitions
│   └── inventory.proto        # Inventory service definitions
├── docker-compose.yml         # PostgreSQL & RabbitMQ
├── Taskfile.yml              # Development tasks
└── go.mod                    # Go module definition
```

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose
- Task (taskfile.dev)

### Environment Setup

1. Copy the environment template:
   ```bash
   cp .env.example .env
   ```

2. Fill in your [.env](cci:7://file:///c:/Users/KAKA/repos/golang/microservices/go-grpc-inventory-microservices/.env:0:0-0:0) file:
   ```env
   APP_NAME=go-grpc-inventory-microservices
   APP_ENV=development

   HTTP_HOST=127.0.0.1
   HTTP_PORT_USER=8081
   HTTP_PORT_PRODUCT=8082
   HTTP_PORT_ORDER=8083
   RABBITMQ_PORT_NOTIFICATION=8084

   RPC_PORT_INVENTORY=50021
   RPC_PORT_PRODUCT=50022

   DB_NAME=inventory
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_HOST=127.0.0.1
   DB_PORT=5432

   RABBITMQ_HOST=127.0.0.1
   RABBITMQ_PORT=5672
   RABBITMQ_USER=rabbitmq
   RABBITMQ_PASSWORD=your_password

   JWT_SECRET=your_jwt_secret
   JWT_DURATION=15
   ```

### Infrastructure Startup

Start PostgreSQL and RabbitMQ:
```bash
task compose:up
```

Access RabbitMQ Management UI at: `http://localhost:15672`

### Running Services

Each service can be started independently:

```bash
# User Service (HTTP :8081)
task dev:user:http

# Product Service
task dev:product:http   # HTTP :8082
task dev:product:rpc    # gRPC :50022

# Inventory Service (gRPC :50021)
task dev:inventory:rpc

# Order Service (HTTP :8083)
task dev:order:http

# Notification Service
task dev:notif:rabbitmq  # RabbitMQ Consumer
# HTTP server runs on :8084 (start separately via go run)
```

## API Endpoints

### User Service (`:8081`)
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/auth/register` | User registration |
| POST | `/api/v1/auth/login` | User login |
| GET | `/api/v1/users` | List users (admin) |
| GET | `/api/v1/users/:id` | Get user by ID |

### Product Service (`:8082`)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/products` | List products |
| GET | `/api/v1/products/:id` | Get product details |
| POST | `/api/v1/products` | Create product (admin) |
| PUT | `/api/v1/products/:id` | Update product (admin) |
| DELETE | `/api/v1/products/:id` | Delete product (admin) |

### Order Service (`:8083`)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/orders` | Get user's orders |
| GET | `/api/v1/orders/:id` | Get order details |
| POST | `/api/v1/orders` | Create order |
| PATCH | `/api/v1/orders/:id` | Update order status |
| DELETE | `/api/v1/orders/:id` | Cancel order |

### Notification Service (`:8084`)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/notifications` | Get my notifications |
| GET | `/api/v1/notifications/unread` | Get unread notifications |
| GET | `/api/v1/notifications/:id` | Get notification details |
| PATCH | `/api/v1/notifications/:id/read` | Mark as read |
| PATCH | `/api/v1/notifications/read-all` | Mark all as read |
| DELETE | `/api/v1/notifications/:id` | Delete notification |
| POST | `/api/v1/notifications` | Create notification (admin) |

## gRPC Services

### Product Service (`:50022`)
```protobuf
service ProductService {
  rpc GetProduct(GetProductRequest) returns (GetProductResponse);
  rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse);
  rpc UpdateProduct(UpdateProductRequest) returns (UpdateProductResponse);
  rpc DeleteProduct(DeleteProductRequest) returns (DeleteProductResponse);
  rpc GetAllProducts(GetAllProductsRequest) returns (GetAllProductsResponse);
}
```

### Inventory Service (`:50021`)
```protobuf
service InventoryService {
  rpc GetInventory(GetInventoryRequest) returns (GetInventoryResponse);
  rpc GetInventoryByProductID(GetInventoryByProductIDRequest) returns (GetInventoryByProductIDResponse);
  rpc CreateInventory(CreateInventoryRequest) returns (CreateInventoryResponse);
  rpc UpdateInventory(UpdateInventoryRequest) returns (UpdateInventoryResponse);
  rpc DeleteInventory(DeleteInventoryRequest) returns (DeleteInventoryResponse);
}
```

## Order Flow

1. **User** creates an order via Order Service HTTP API
2. **Order Service** calls **Product Service** (gRPC) to get product price
3. **Order Service** calls **Inventory Service** (gRPC) to check stock availability
4. **Order Service** calls **Inventory Service** (gRPC) to deduct inventory
5. **Order Service** publishes notification message to **RabbitMQ** (`notification` queue)
6. **Notification Service** (RabbitMQ consumer) receives message and stores notification
7. **User** can fetch notifications via Notification Service HTTP API

## Protocol Buffer Generation

```bash
# Generate Product protobuf
task gen:product

# Generate Inventory protobuf
task gen:inventory
```

## Key Features

- **Hexagonal Architecture**: Clean separation between domain, application, and infrastructure layers
- **JWT Authentication**: Secure API access with role-based permissions (user/admin)
- **Inter-service Communication**: gRPC for fast, type-safe synchronous calls
- **Event-driven Messaging**: RabbitMQ for reliable asynchronous notifications
- **Database Per Service**: Each service owns its PostgreSQL schema
- **Auto-migration**: GORM automates database schema migrations

## Development Commands

```bash
# Start infrastructure
task compose:up

# Stop infrastructure
task compose:down

# Access PostgreSQL CLI
task db:cli

# Run any service
task dev:<service>:<type>
```

## License

MIT
