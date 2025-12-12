# Stage 1: Build the Go application
FROM golang:1.25-alpine AS builder

# Install git required for fetching dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# CGO_ENABLED=0 creates a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o cv .

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app/

# Copy the binary from the builder stage
COPY --from=builder /app/cv .

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./cv"]
