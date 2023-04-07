FROM golang:1.20.2-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build ./cmd/main.go
