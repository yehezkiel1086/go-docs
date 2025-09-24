postgresup:
	docker compose up -d

postgresdown:
	docker compose down

postgres:
	docker exec -it postgres psql -U postgres -d socket_chat

build:
	go build cmd/ -o bin/server

run:
	go run cmd/main.go

.PHONY: postgresup postgresdown postgres build
