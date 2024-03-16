# Use the official Golang image
FROM golang:latest AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o CC-Compiler-System .

# Use a minimal base image to reduce size
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/CC-Compiler-System .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./CC-Compiler-System"]
