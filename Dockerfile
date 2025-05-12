# ------------------ STAGE 1: Build the Go Application ------------------
# Use the official Golang image as the builder
FROM golang:1.24.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only go.mod and go.sum for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o /cryptoapi main.go

# ------------------ STAGE 2: Create a Small and Secure Image ------------------
# Use a minimal base image (Distroless for security)
FROM gcr.io/distroless/base-debian12

# Set the working directory
WORKDIR /app

# Copy the compiled Go application from the builder stage
COPY --from=builder /cryptoapi /app/cryptoapi

# Expose the application port (change this to your app port)
EXPOSE 8080

# Run the compiled Go binary
CMD ["/app/cryptoapi"]