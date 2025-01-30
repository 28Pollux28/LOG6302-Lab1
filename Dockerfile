# Stage 1: Build the Go program
FROM golang:1.23 AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Enable CGO and build the binary
ENV CGO_ENABLED=1
RUN go build -o /bin/go-php-parser .

# Stage 2: Create a minimal runtime image
FROM debian:bookworm-slim

WORKDIR /app

# Install required runtime dependencies for CGO
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Create a volume for passing files
VOLUME /data

# Copy the built binary
COPY --from=builder /bin/go-php-parser /usr/local/bin/go-php-parser

# Default command
CMD ["bash"]
