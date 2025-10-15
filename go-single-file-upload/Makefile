composeup:
	docker compose up -d

composedown:
	docker composedown

postgres:
	docker exec -it postgres psql -U postgres -b file_uploads

run:
	go run cmd/main.go

.PHONY: composeup composedown postgres run
