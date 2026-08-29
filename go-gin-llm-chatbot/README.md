# Go Gin LLM Chatbot with Gemini API

This project is a backend service for a chatbot built with Go, the Gin web framework, and Google's Gemini Large Language Model (LLM) API. It provides a simple and extensible foundation for creating conversational AI applications.

## Features

- **Text-to-Text Generation**: Send a text prompt and receive a generated text response from the Gemini model.
- **Image-to-Text (Multimodal)**: Send an image along with a text prompt to get a contextual response based on the visual information.
- **Clean Architecture**: Organized into layers (adapter, core, service) for better separation of concerns, testability, and maintainability.
- **Configuration-driven**: Easily configure the application using environment variables.

## Tech Stack

- **Language**: [Go](https://golang.org/)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **LLM**: [Google Gemini API](https://ai.google.dev/docs/gemini_api_overview)
- **Configuration**: [godotenv](https://github.com/joho/godotenv)

## Project Structure

The project follows a clean architecture-inspired structure:

```
go-gemini-llm/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── adapter/
│   │   ├── config/       # Environment configuration
│   │   ├── handler/      # HTTP handlers (Gin)
│   │   └── llm/          # Adapters for LLM clients (Gemini)
│   └── core/
│       ├── port/         # Interfaces for services (ports)
│       └── service/      # Core business logic
├── .env.example          # Example environment variables
└── go.mod
```

## Getting Started

Follow these steps to get the project up and running on your local machine.

### Prerequisites

- [Go](https://go.dev/doc/install) (version 1.21 or later)
- A Google Gemini API Key. You can get one from [Google AI Studio](https://makersuite.google.com/app/apikey).

### Installation & Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/your-username/go-gemini-llm.git
    cd go-gemini-llm
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Set up environment variables:**
    Create a `.env` file in the root of the project by copying the example file:
    ```sh
    cp .env.example .env
    ```
    Now, open the `.env` file and add your configuration details:
    ```env
    # LLM Configuration
    API_KEY="YOUR_GEMINI_API_KEY"
    AI_MODEL="gemini-pro-vision" # or "gemini-pro" for text-only

    # HTTP Server Configuration
    HTTP_HOST="localhost"
    HTTP_PORT="8080"
    ```

### Running the Application

Execute the following command to start the server:

```sh
go run cmd/main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### 1. Text-to-Text Prompt

- **Endpoint**: `POST /api/v1/prompt/text-to-text`
- **Request**: `application/json`
  ```json
  {
    "prompt": "Write a short story about a robot who discovers music."
  }
  ```
- **cURL Example**:
  ```sh
  curl -X POST http://localhost:8080/api/v1/prompt/text-to-text \
  -H "Content-Type: application/json" \
  -d '{"prompt": "Write a short story about a robot who discovers music."}'
  ```

### 2. Image-to-Text Prompt

- **Endpoint**: `POST /api/v1/prompt/image-to-text`
- **Request**: `multipart/form-data`
  - `prompt` (form field): The text prompt.
  - `file` (file field): The image file.
- **cURL Example**:
  ```sh
  curl -X POST http://localhost:8080/api/v1/prompt/image-to-text \
  -F "prompt=What is in this image?" \
  -F "file=@/path/to/your/image.jpg"
  ```