# syntax=docker/dockerfile:1

FROM golang:1.24-bookworm AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/fast-stream-bot ./cmd/fsb

FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/fast-stream-bot /usr/local/bin/fast-stream-bot
COPY frontend ./frontend
COPY config.toml ./config.toml

ENV HTTP_PORT=8000
EXPOSE 8000

CMD ["fast-stream-bot"]
