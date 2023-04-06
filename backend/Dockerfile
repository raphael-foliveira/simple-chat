FROM golang:1.20.2-alpine

WORKDIR /app

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy
