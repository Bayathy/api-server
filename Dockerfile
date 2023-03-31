FROM golang:1.20.2-alpine3.17

WORKDIR /app

RUN go mod download