# Start with the official Golang image
FROM golang:latest AS build

# Take in CGO_ENABLED argument
ARG CGO_ENABLED=0

# Set CGO_ENABLED
ENV CGO_ENABLED=$CGO_ENABLED

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o bin/main .

# Start a new stage from scratch
FROM alpine:latest  AS run
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/bin/main .

# Expose port 8080 to the outside world
EXPOSE 8090

# Command to run the executable
CMD ["./main"]