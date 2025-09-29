# Smart Transit System Backend - Dockerfile
# Multi-stage build for efficient container image

FROM golang:1.24-alpine AS builder

# Set build metadata
LABEL stage=builder
LABEL service="smart-transit-backend"

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies with caching
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o smart-transit-backend .

# Production stage
FROM alpine:latest

# Install security updates and certificates
RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

# Set metadata
LABEL service="smart-transit-backend"
LABEL version="1.0.0"
LABEL description="Smart Transit System Backend API Service"

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/smart-transit-backend .

# Create non-root user for security compliance (Choreo requirement)
RUN adduser -D -s /bin/sh -u 10001 smarttransit
USER 10001

# Expose port
EXPOSE 8080

# Health check for container orchestration
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/healthz || exit 1

# Command to run the Smart Transit System backend
CMD ["./smart-transit-backend"]