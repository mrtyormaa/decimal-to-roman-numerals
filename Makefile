.PHONY: setup build test cover up down restart clean push

# Variables
APP_NAME := decimal-to-roman-numerals
DOCKER_IMAGE := $(APP_NAME):latest
DOCKERFILE := Dockerfile
COMPOSE_FILE := docker-compose.yml
COVERAGE_CONTAINER := $(APP_NAME)-coverage

# Detect OS
ifdef ComSpec
    RM = del /F /Q
    RMDIR = rmdir /S /Q
    SEP = ;
else
    RM = rm -f
    RMDIR = rm -rf
    SEP = :
endif

setup:
	@echo "Installing Swag..."
	go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Downloading Go modules..."
	go mod download
	@echo "Initializing Swag documentation..."
	swag init
	@echo "Building Go project..."
	go build -o bin/main main.go


build:
	@echo "Building Docker images without cache..."
	docker-compose -f $(COMPOSE_FILE) build --no-cache

test:
	@echo "Testing the application..."
	docker-compose -f $(COMPOSE_FILE) up --build roman-numerals-tests
	@echo "Tests completed."

cover:
	@echo "Running tests and generating coverage report in Docker..."
	docker build -t $(COVERAGE_CONTAINER) -f $(DOCKERFILE) --target coverage .
	docker run --rm -v $(CURDIR)/coverage:/coverage $(COVERAGE_CONTAINER)
	@echo "Coverage report generated at coverage/coverage.html"

up:
	@echo "Starting Docker containers..."
	docker-compose -f $(COMPOSE_FILE) up -d roman-numerals prometheus grafana

down:
	@echo "Stopping Docker containers..."
	docker-compose -f $(COMPOSE_FILE) down

restart:
	@echo "Restarting Docker containers..."
	docker-compose -f $(COMPOSE_FILE) restart

clean:
	@echo "Stopping and removing Docker containers and images..."
	-docker-compose -f $(COMPOSE_FILE) down --rmi all
	-docker volume prune -f
	@echo "Removing build artifacts..."
	-$(RM) bin$(SEP)main
	-$(RMDIR) coverage
