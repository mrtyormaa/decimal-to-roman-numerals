setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/server/main.go -o docs/
	go build -o bin/server cmd/server/main.go

build:
	docker compose build --no-cache

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose restart

clean:
	docker stop decimal-to-roman-numerals
	docker rm decimal-to-roman-numerals
	docker image rm decimal-to-roman-numerals-backend