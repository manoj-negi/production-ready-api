FROM golang:1.20.2-alpine3.17 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]