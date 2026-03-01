include .env
export

# ==================================================================================
# VARIABLES
# ==================================================================================

APP_NAME = piggybank
MAIN_PATH = ./cmd/server
BINARY = ./bin/$(APP_NAME)
MIGRATIONS_PATH = internal/database/migrations

# ==================================================================================
# BUILD
# ==================================================================================

build:
	@echo "Building..."
	@mkdir -p bin
	go build -o $(BINARY) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	@echo "Cleaning..."
	@rm -rf bin/
	
# ==================================================================================
# DATABASE & MIGRATIONS
# ==================================================================================

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-drop:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" drop -f

migrate-create:
	@read -p "Migration name:" name; \
	migrate create -ext sql $(MIGRATIONS_PATH) -seq $$name

# ==================================================================================
# CODE QUALITY
# ==================================================================================

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

# ==================================================================================
# TESTING
# ==================================================================================

test:
	go test ./... -v

test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# ==================================================================================
# HELP
# ==================================================================================

help:
	@echo "Available commands:"
	@echo "  make build          - Build the binary"
	@echo "  make run            - Run the server"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make migrate-up     - Run all pending migrations"
	@echo "  make migrate-down   - Roll back the last migration"
	@echo "  make migrate-drop   - Drop everything in the database"
	@echo "  make migrate-create - Create a new migration file"
	@echo "  make fmt            - Format all Go code"
	@echo "  make vet            - Run Go vet"
	@echo "  make tidy           - Tidy go modules"
	@echo "  make test           - Run all tests"
	@echo "  make test-cover     - Run tests with coverage report"

.PHONY: build run clean migrate-up migrate-down migrate-drop migrate-create fmt vet tidy test test-cover help
```
