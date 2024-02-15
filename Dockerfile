# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files to the container
COPY . .

# Build the Go application
RUN go build -o server .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./server"]
