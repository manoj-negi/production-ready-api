FROM golang:1.20.2-alpine3.17 AS builder
WORKDIR /app
COPY . .

RUN go build -o main cmd/server/main.go

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./app/main"]