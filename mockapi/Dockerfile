# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (if you have them)
# COPY go.mod go.sum ./
COPY go.mod ./

# Download Go modules (if you have them)
RUN go mod download

# Copy the Go source code
COPY . .

# Build the Go application
RUN go build -o mockapi

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./mockapi"]
