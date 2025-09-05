FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o app ./cmd

FROM golang:1.23.0-alpine

RUN adduser -D appuser

COPY --from=builder /app/app /app/app

USER appuser

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./app"]
