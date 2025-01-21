# Use the official Golang image to build the application
FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o hr-system ./cmd/app

# Start a new stage from scratch
FROM alpine:latest

RUN apk update && apk add --no-cache bash

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/hr-system .
COPY --from=builder /app/configs/config.yml ./configs/config.yml

COPY ./scripts/wait-for-it.sh /usr/local/bin/

# Command to run the executable
CMD ["wait-for-it.sh", "mysql:3306", "--", "./hr-system"]
