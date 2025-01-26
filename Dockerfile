# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app

# Install required packages
RUN apk --no-cache add ca-certificates

# Copy the binary and any other required files from the builder stage
COPY --from=builder /app/main .
COPY env/ env/

# Expose the required port
EXPOSE 8080

# Set default environment variables (can be overridden)
ENV ACTIVE_PROFILE=prod

# Set default command to include "server" subcommand
ENTRYPOINT ["./main", "server"]
