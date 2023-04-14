FROM golang:1.20.2-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go install github.com/cosmtrek/air@latest && go mod tidy

COPY . .

EXPOSE 8000

