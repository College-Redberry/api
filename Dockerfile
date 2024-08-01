# Use the official Golang image as a build stage
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM debian:stable-slim

# Install curl to get wait-for-it
RUN apt-get update && apt-get install -y curl

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Download wait-for-it script
RUN curl -o /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && chmod +x /usr/local/bin/wait-for-it.sh

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["wait-for-it.sh", "postgres:5432", "--", "/app/main"]
