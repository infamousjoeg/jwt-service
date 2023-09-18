# Use the official Golang image to create a build artifact.
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app without CGO
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jwt-service .

# Start from a scratch image
FROM scratch

# Set a non-root user for the application (though this won't have an actual system user representation in the scratch image)
USER 1001

WORKDIR /app

# Copy the pre-built binary
COPY --from=builder /app/jwt-service .

# Command to run the executable
CMD ["./jwt-service"]
