FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api.go

USER 1000

EXPOSE 8080

CMD ["./bin/api"]