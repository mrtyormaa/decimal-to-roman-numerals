setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init
	go build -o bin/ main.go

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