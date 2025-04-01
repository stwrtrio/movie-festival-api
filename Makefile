# Variables
APP_CMD = cmd/main.go
APP_NAME = movie-festival-api
TEST_CMD = go test -v ./tests/...
LINT_CMD = go vet ./... && go fmt ./...
MOCK_SOURCE_DIR = ./internal/repositories/*.go
MOCK_DEST_DIR = ./internal/mocks

# Tasks
.PHONY: all run lint test start clean-test build mocks

# Default target
run: lint mocks test start

# Target to check syntax (lint)
lint:
	@echo "Checking code syntax..."
	@$(LINT_CMD)
	@echo "Syntax check passed."

# Target to run tests
test:
	@echo "Running tests..."
	@$(TEST_CMD)
	@echo "All tests passed."

# Target to run the application
start:
	@echo "Starting the application..."
	@go run $(APP_CMD)

# Build the application
build:
	@echo "Building the application..."
	go build -o $(APP_NAME) $(APP_CMD)

# Help target to display usage
help:
	@echo "Usage:"
	@echo "  make run   - Run the application with syntax check and tests"
	@echo "  make lint  - Check code syntax"
	@echo "  make test  - Run all tests"
	@echo "  make start - Run the application directly"
	@echo "  make build - Build the application"
