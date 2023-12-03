# syntax=docker/dockerfile:1

# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Download Go modules
COPY ./ ./

WORKDIR /app/src
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./loginpage"]
