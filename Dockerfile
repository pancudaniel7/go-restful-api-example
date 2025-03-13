FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download && \
    go build -o main cmd/main.go

FROM alpine:3.21.3

WORKDIR /
COPY --from=builder /app/configs/ /configs
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]