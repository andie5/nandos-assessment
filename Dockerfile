# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

# Create directory to run program
WORKDIR /app

# Copy go.mod file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy all files
COPY ./ ./

# Run the program
RUN go build -o marsrover .

# Run the executable
CMD ["./marsrover"]