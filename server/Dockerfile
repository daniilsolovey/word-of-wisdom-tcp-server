# Use Golang base image
FROM golang:1.22-alpine

# Set work directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the server
RUN go build -o server .

# Expose the port
EXPOSE 8080

# Run the server
CMD ["./server"]
