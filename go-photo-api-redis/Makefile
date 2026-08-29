run:
	go run cmd/main.go

composeup:
	docker compose up -d

composedown:
	docker compose down

postgres:
	docker exec -it postgres psql -U postgres -b go_photo_app

.PHONY: run composeup composedown postgres
