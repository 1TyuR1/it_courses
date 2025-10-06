# Makefile
.PHONY: help infra-up infra-down proto migrate test

help: ## Показать помощь
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

infra-up: ## Запустить всю инфраструктуру
	docker-compose up -d

infra-down: ## Остановить инфраструктуру
	docker-compose down

infra-logs: ## Показать логи инфраструктуры
	docker-compose logs -f

infra-clean: ## Очистить все данные
	docker-compose down -v

proto: ## Генерировать protobuf файлы
	@echo "Generating protobuf files..."
	@./scripts/generate-proto.sh

migrate-up: ## Применить миграции
	@echo "Running migrations..."
	@./infrastructure/scripts/migrate.sh up

migrate-down: ## Откатить миграции
	@echo "Rolling back migrations..."
	@./infrastructure/scripts/migrate.sh down

# Auth Service
auth-run: ## Запустить auth-service
	cd services/auth-service && go run cmd/auth/main.go

auth-build: ## Собрать auth-service
	cd services/auth-service && go build -o ../../bin/auth-service cmd/auth/main.go

# Тесты
test: ## Запустить все тесты
	go test -v ./...

test-coverage: ## Тесты с покрытием
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Утилиты
fmt: ## Форматировать код
	go fmt ./...

lint: ## Линтинг
	golangci-lint run ./...

tidy: ## Очистить зависимости
	go mod tidy
	cd shared && go mod tidy
