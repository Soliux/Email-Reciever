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

# Expose ports 25, 8080, and 5432
EXPOSE 25
EXPOSE 8080
EXPOSE 5432

# Install PostgreSQL client and server
RUN apt-get update && apt-get install -y postgresql postgresql-contrib

# Create a new PostgreSQL user and database
USER postgres
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER myuser WITH PASSWORD 'mypassword';" &&\
    createdb -O myuser mydb
USER root

# Command to run the app and PostgreSQL
CMD service postgresql start && ./main