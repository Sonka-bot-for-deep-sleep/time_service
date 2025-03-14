FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o time_service ./cmd/time_service/main.go

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /app/time_service .
ENTRYPOINT ["./time_service"]
