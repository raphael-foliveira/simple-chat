FROM golang:1.20.2-alpine

WORKDIR /app

COPY . .

EXPOSE 8000

RUN go install github.com/cosmtrek/air@latest && go mod tidy
