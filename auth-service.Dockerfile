# auth-service/Dockerfile
# Start from a base image with Go pre-installed
FROM golang:1.22-alpine

# Set the current working directory inside the container to /app
WORKDIR /app

# Copy the Go module files
COPY auth-service/go.mod auth-service/go.sum .env ./
COPY ./shared /shared

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY ./auth-service .

# Build the application
RUN go build -o auth-service

# Set the environment variables
ENV MYSQL_DSN="user:password@tcp(mysql:3306)/dbname"
ENV JWT_SECRET="your_secret_key"
ENV AUTH_SERVICE_ADDRESS=":8000"

# Expose the port the service runs on
EXPOSE 8000

# Run the binary
CMD ["./auth-service"]