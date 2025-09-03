postgresinit:
	docker-compose up -d

postgresdown:
	docker-compose down

postgres:
	docker exec -it postgres psql -U postgres go_jwt

.PHONY: postgresinit postgresdown postgres
