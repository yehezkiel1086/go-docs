composeup:
	docker compose up -d

composedown:
	docker compose down

postgres:
	docker exec -it postgres -U psql -b go_oauth

run:
	go run cmd/main.go

.PHONY: composeup composedown postgres run
