include .env

init:
	cp .env.example .env
migrate-up:
	migrate -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable -path ./migrations up

migrate-down:
	migrate -database postgres://${DB_USERNAME}:${DB_PASSWORD}@localhost:12000/${DB_DATABASE}?sslmode=disable -path ./migrations down