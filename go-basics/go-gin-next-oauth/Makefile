composeup:
	docker compose up -d

composedown:
	docker compose down

postgres:
	docker exec -it postgres psql -U postgres -b go_auth

run:
	go run cmd/main.go

.PHONY: composeup composedown postgres run
