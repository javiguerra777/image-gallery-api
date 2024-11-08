# Stage 1: Build the Go application
FROM golang:1.23.2 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Install necessary dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o myapp main.go

# Stage 2: Prepare the runtime environment
FROM golang:1.23.2

# Set the working directory
WORKDIR /app

# Copy the built Go application from the build stage
COPY --from=build /app/myapp .

# Expose the port the app will run on
EXPOSE 4000

# Set the default command to run the application
CMD ["./myapp"]

