# Use Golang base image
FROM golang:1.22-alpine

# Set work directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the client
RUN go build -o client .

# Run the client
CMD ["./client"]
