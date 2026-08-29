# Go Gin Todo Graphql

Simple Graphql todo app built with Go and Gin

## Installation and Run

Install: `go mod tidy` \
Build: `go build cmd/main.go -o bin/main` \
Run: `go run cmd/main.go` 

## Graphql Requests

URI: `http://127.0.0.1:3500/graphql`

Get all todos:
```json
{
  "query": "{ todos { id title done } }"
}
```

Get todo by id:
```json
{
  "query": "{ todo(id: 1) { id title done } }"
}
```

Create todo:
```json
{
  "query": "mutation { createTodo(title: \"Learn Go GraphQL\") { id title done } }"
}
```


Toggle Todo:
```
{
  "query": "mutation { toggleTodo(id: 1) { id title done } }"
}
```
