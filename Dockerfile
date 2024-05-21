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
FROM golang:1.21.10 AS tester
WORKDIR /app
COPY --from=builder /app /app
CMD ["sh", "-c", "go test ./..."]

# Stage 3: Tester and Coverage Generator
FROM golang:1.21.10 AS coverage
WORKDIR /app
COPY --from=builder /app /app
RUN mkdir -p /coverage
CMD ["sh", "-c", "go test ./... -coverprofile=/coverage/coverage.out && go tool cover -html=/coverage/coverage.out -o /coverage/coverage.html"]


# Stage 4: Final Image
FROM golang:1.21.10
WORKDIR /app
COPY --from=builder /app/bin/main ./bin/main
# Set environment variable for the port with a default value of 8001
ENV PORT 8001
# Make port configurable at runtime with a default fallback
EXPOSE ${PORT}
CMD ["./bin/main"]
