.PHONY: setup build up down restart cover clean

setup:
	@echo "Installing Swag..."
	go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Initializing Swag documentation..."
	swag init
	@echo "Building Go project..."
	go build -o bin/ main.go

build:
	@echo "Building Docker images without cache..."
	docker compose build --no-cache

up:
	@echo "Starting Docker containers..."
	docker compose up

down:
	@echo "Stopping Docker containers..."
	docker compose down

restart:
	@echo "Restarting Docker containers..."
	docker compose restart

cover:
	@echo "Running tests and generating coverage report..."
	go test ./... -coverprofile="docs/coverage.out"
	go tool cover -html="docs/coverage.out"

clean:
	@echo "Stopping and removing Docker containers and images..."
	-docker stop decimal-to-roman-numerals
	-docker rm decimal-to-roman-numerals
	-docker image rm decimal-to-roman-numerals-backend
