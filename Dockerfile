# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the working directory to /app
WORKDIR /app

# Install Air for live reload
RUN go install github.com/cosmtrek/air@latest

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the Go source code
COPY . .

# Build the Go application
RUN go build -o main .

# Set the command to run the Go application
CMD ["./main"]