FROM golang:1.21.10

WORKDIR /app
COPY . .
RUN go mod tidy