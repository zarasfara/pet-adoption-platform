# Этап 1
FROM golang:1.21-alpine AS build-stage

RUN apk update \
    && apk add --no-cache \
    make \
    gcc \
    musl-dev \
    postgresql-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go

# Этап 2
FROM alpine as run-stage

RUN apk update && apk add --no-cache postgresql-client
WORKDIR /

# Убедитесь, что бинарный файл исполняемый и находится в правильном месте
COPY --from=build-stage /app/bin/app /app
COPY --from=build-stage /app/configs /configs
COPY --from=build-stage /app/wait-for-postgres.sh /
COPY --from=build-stage /app/.env /

# Сделайте бинарный файл исполняемым
RUN chmod +x /app

EXPOSE 80

# Укажите правильный путь к бинарному файлу в команде CMD
CMD ["/app"]
