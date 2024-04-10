# Start with a base image containing Go runtime
FROM golang:alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .




# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
