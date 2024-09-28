FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o ./bin/consumer ./cmd/consumer.go

USER 1000

CMD ["./bin/consumer"]