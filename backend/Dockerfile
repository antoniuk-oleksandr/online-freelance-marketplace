# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app source code, including keys
COPY . .

# Build the app
RUN go build -o app ./cmd/ofm_backend

# Stage 2: Run the Go app in a minimal image
FROM alpine:latest

WORKDIR /app

# Install CA certificates (required for HTTPS)
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /app/app .

# 🔑 Copy keys from builder stage
COPY --from=builder /app/keys ./keys
COPY --from=builder /app/internal ./internal
COPY --from=builder /app/cmd ./cmd

# Expose the port your Fiber app runs on (default is 3000 or 8000)
EXPOSE 8000

# Run the app
CMD ["./app"]

