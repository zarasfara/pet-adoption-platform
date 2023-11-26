include .env

init:
	cp .env.example .env
migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable up

migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable down -all

swag:
	docker run -v $(shell pwd):/app -w /app ghcr.io/swaggo/swag:latest /root/swag init -g cmd/main.go

swag-fmt:
	docker run -v $(shell pwd):/app -w /app ghcr.io/swaggo/swag:latest /root/swag fmt
