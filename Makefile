include .env

init:
	cp .env.example .env

migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable up

migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable down -all

migrate-create:
	docker run -v ./migrations:/migrations migrate/migrate create -dir=/migrations -ext sql -seq create_$(name)_table

migrate-fix:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable force $(version)

swag:
	docker run --rm -v $(shell pwd):/app -w /app ghcr.io/swaggo/swag:latest init -g cmd/main.go

swag-fmt:
	docker run --rm -v $(shell pwd):/app -w /app ghcr.io/swaggo/swag:latest fmt

rebuild:
	docker-compose up -d --no-deps --build $(service)