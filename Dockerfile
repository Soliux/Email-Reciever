# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 25 and 8080 to the outside world
EXPOSE 25
EXPOSE 8080

# Command to run the app
CMD ["./main"]
