FROM golang:1.23.0-alpine AS builder

RUN apk add

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd

FROM alpine:latest

RUN adduser -D appuser

COPY --from=builder /app/app /app/app

USER appuser

WORKDIR /app

WORKDIR /app

ENTRYPOINT ["./app"]
