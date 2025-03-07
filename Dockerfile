# Use the official Golang image for building the app
FROM golang:1.24 AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o server .

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./server"]

