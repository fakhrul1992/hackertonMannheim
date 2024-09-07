# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Install curl
RUN apk update && apk add --no-cache curl

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
# COPY go.mod go.sum ./
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o mywebsite

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./mywebsite"]
