# syntax=docker/dockerfile:1.2

# Stage 1: Build the binary
FROM golang:1.22 AS builder

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /request-counter ./cmd/request-counter

# Stage 2: Create a minimal image
FROM gcr.io/distroless/static-debian11

COPY --from=builder /request-counter /request-counter

EXPOSE 1378

ENTRYPOINT ["/request-counter"]
