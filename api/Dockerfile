# Используем многоэтапную сборку
FROM golang:1.24.0-alpine3.21 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod tidy

COPY db/migrations /app/db/migrations

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main main.go

# Второй этап: создаем финальный образ
FROM alpine:3.21

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из builder
COPY --from=builder /app/main .

# Открываем порт
EXPOSE 8000

# Запускаем приложение
ENTRYPOINT [ "./main" ]