# syntax=docker/dockerfile:1
FROM golang:1.17
LABEL maintainer="Daniel Sender"
WORKDIR /app

COPY . .
COPY go.mod go.sum ./
RUN go mod download

RUN go build -o desafio
CMD ./desafio
