# Use the official Golang image to create a build artifact.
# This is a multi-stage build. This stage is named "builder".
FROM golang:1.21-alpine as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies.
# Doing this before copying the entire codebase optimizes re-building the image by caching the dependencies layer.
COPY go.mod go.sum ./

# Download the dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application. This will produce a binary named "server".
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# Use a small Alpine Linux image to run the application.
# This is the second stage of the build, where the binary from the "builder" stage is copied.
FROM alpine:latest  

# Set the working directory to /root/.
WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/server .
COPY .env .env
# Expose the port the app runs on.
EXPOSE 8080

# Run the binary.
CMD ["./server"]
