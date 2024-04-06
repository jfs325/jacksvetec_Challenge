# Start from the official Go image to build our application
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go files and static files into the container
COPY . .

# Build the Go app
RUN go env -w GO111MODULE=off
 
RUN go build -o webserver

# Use a small Alpine image to run the application
FROM alpine:latest  

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/webserver .
COPY --from=builder /app/static ./static

# Expose port 8080
EXPOSE 8080

# Run the web server
CMD ["./webserver"]
