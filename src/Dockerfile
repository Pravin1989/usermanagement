FROM golang:1.16-alpine

# Set the Current Working Directory inside the container
WORKDIR /usermanagement-service

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

EXPOSE 8080

# Build the Go app
RUN go build -o usermanagement ./src/

# Command to run the executable
ENTRYPOINT [ "/usermanagement-service/usermanagement" ]