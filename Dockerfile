# Build stage
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate swagger files
RUN go generate

# Build the Go app
RUN go build -o main .

# Run stage
FROM debian:stable-slim

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]