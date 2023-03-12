FROM golang:1.20.0-alpine3.17


# Set the working directory to /code
WORKDIR /code

# Copy the current directory contents into the container at /code
COPY . .

# Install C compiler
RUN apk add build-base

# Build the go binaries
RUN mkdir -p ./gobin
RUN go build -o ./gobin .

# Use the alpine:3.17 image as the final image
FROM alpine:3.17

# Set the working directory to /code
WORKDIR /code

# Copy the built binaries from the builder image to the final image
COPY --from=builder /code/gobin/. .

# Make the binary executable
RUN chmod +x ./*
