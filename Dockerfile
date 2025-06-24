# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY src/ ./
RUN go mod tidy && go build -o api main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api ./
EXPOSE 8080
CMD ["./api"]
