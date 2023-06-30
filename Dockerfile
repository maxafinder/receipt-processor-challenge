# Start with golang version 1.20 base image
FROM golang:1.20

# Maintainer info
LABEL maintainer="Max Finder maxafinder@gmail.com"

# Set the working directory inside container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source into container's working directory
COPY . .

# Build Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run binary
CMD ["./main"]