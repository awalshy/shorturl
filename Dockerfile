FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o shorturl ./cmd/main.go


FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/shorturl .
EXPOSE 8080

CMD ["./shorturl"]