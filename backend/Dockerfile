# Билд стадии
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server /app/cmd/server/main.go

# Финальная стадия
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/config/ ./config/
COPY --from=builder /app/db/postgres/migration/ ./db/postgres/migration/

EXPOSE 8080

CMD ["./server"]
