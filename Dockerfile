# Используем официальный образ Golang в качестве базового
FROM golang:1.21-alpine

RUN apk update \
    && apk add --no-cache \ 
    make \
    gcc \
    musl-dev

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum внутрь контейнера
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все файлы проекта внутрь контейнера
COPY . .

# Устанавливаем CGO_ENABLED=1
ENV CGO_ENABLED=1

# Собираем приложение
RUN go build -o .bin/app ./cmd

EXPOSE 80

# Указываем команду, которая будет выполняться при запуске контейнера
CMD ["/app/.bin/app"]