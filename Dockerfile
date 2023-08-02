# Use the official Go image as the base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container's working directory
COPY . .

# Build the Go application inside the container
RUN go build -o myapp

# Expose the port that the application listens on
EXPOSE 8081

# Copy the config.json file into the container
COPY config.json /app/config.json

# Set the entry point for the container
ENTRYPOINT ["./myapp"]
