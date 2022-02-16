# syntax=docker/dockerfile:1
FROM golang:1.17
LABEL Daniel Sender
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o desafio