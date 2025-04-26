# Build stage
FROM golang:1.22.3-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy any necessary config files or static assets
COPY --from=builder /app/data ./data

# Expose the port your application runs on
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Run the application
CMD ["./main"] 