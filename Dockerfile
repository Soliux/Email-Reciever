# Use an official Go runtime as a parent image
FROM golang:1.20-alpine AS builder

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the Go app
RUN go build -o /go/bin/app

# Use a smaller base image
FROM alpine:3.14

# Copy the built binary from the previous stage
COPY --from=builder /go/bin/app /usr/local/bin/app

# Set the working directory to /app
WORKDIR /app

# Run the app when the container starts
CMD ["/usr/local/bin/app"]
