# syntax=docker/dockerfile:1

# ====== Build stage ======
FROM golang:1.25-alpine AS builder
WORKDIR /src

# Инструменты
RUN apk add --no-cache git

# Кеш модулей
COPY go.mod go.sum ./
RUN go mod download

# Исходники
COPY . .

# Сборка бинарей
ENV CGO_ENABLED=0 GOOS=linux
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/sso ./cmd/app && \
    go build -trimpath -ldflags="-s -w" -o /out/migrator ./cmd/migrator

# ====== Runtime stage ======
FROM gcr.io/distroless/base-debian12 AS runtime
WORKDIR /app

# Бинарники
COPY --from=builder /out/sso /app/sso
COPY --from=builder /out/migrator /app/migrator

# Конфиги/серты можно пробрасывать томами; при необходимости раскомментировать:
# COPY ./config.yaml /app/config.yaml
# COPY ./certs/rootCA.crt /app/certs/rootCA.crt

# Некорневой пользователь
USER nonroot:nonroot

# gRPC порт
EXPOSE 2441

# По умолчанию запускаем сервер; мигратор можно вызвать отдельно:
ENTRYPOINT ["/app/sso"]
