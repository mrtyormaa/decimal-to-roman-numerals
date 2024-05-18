# Stage 1: Builder
FROM golang:1.21.10 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init
RUN CGO_ENABLED=1 go build -o bin/main main.go

# Stage 2: Tester
FROM builder AS tester
WORKDIR /app
COPY tests/ tests/
CMD ["go", "test", "./tests/..."]

# Stage 3: Final Image
FROM golang:1.21.10
WORKDIR /app
COPY --from=builder /app/bin/main ./bin/main
EXPOSE 8001
CMD ["./bin/main"]
