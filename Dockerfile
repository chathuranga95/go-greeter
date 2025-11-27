# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY main.go .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o greeter main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS support (if needed)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/greeter .

# Expose the port the app runs on
EXPOSE 9090

# Run the binary
CMD ["./greeter"]

