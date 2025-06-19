# Use Golang base image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o app .

# Expose the port (you can modify as per need)
EXPOSE 8080

# Run the Go application
CMD ["./app"]
