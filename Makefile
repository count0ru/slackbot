# Переменные
BIN_DIR := bin
APP_NAME := slackbot

# Цели
.PHONY: all build clean lint test generate-mocks regenerate-mocks fmt

all: build

fmt:
	@echo "Formatting Go files (excluding mocks)..."
	@find . -type f -name '*.go' ! -name '*mocks.go' -exec gofmt -s -w {} \;
	@find . -type f -name '*.go' ! -name '*mocks.go' -exec goimports -w {} \;
	@echo "Formatting complete."

# Обновление зависимостей и генерация go.mod
deps:
	go mod tidy
	go mod vendor

# Сборка приложения
build: deps
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/slackbot

# Очистка собранных файлов
clean:
	rm -rf $(BIN_DIR)

# Линтинг всех Go-файлов
lint:
	golangci-lint run ./...

# Запуск тестов (кроме functional_tests)
test:
	go test ./... -v -coverprofile=coverage.out -covermode=atomic

# Генерация моков
generate-mocks:
	mockgen -source=internal/db/interfaces.go -destination=internal/db/mocks.go -package=db
	mockgen -source=internal/httpclient/interfaces.go -destination=internal/httpclient/mocks.go -package=httpclient
	mockgen -source=pkg/tracker/interfaces.go -destination=pkg/tracker/mocks.go -package=tracker

# Перегенерация моков
regenerate-mocks: clean generate-mocks

# Запуск всех целей по умолчанию
default: build
