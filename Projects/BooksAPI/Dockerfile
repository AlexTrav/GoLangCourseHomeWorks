# build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/app

# runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/docs ./docs

EXPOSE 3000

CMD ["./app"]