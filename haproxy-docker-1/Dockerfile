# syntax=docker/dockerfile:1
# Create build stage based on buster image
FROM golang:1.16-alpine
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY goapp/go.* ./
# Install any required modules
ENV GO111MODULE=on
RUN go mod download
# Copy over Go source code
COPY goapp/*.go ./
# Run the Go build and output binary under hello_go_http
RUN go build -o /goapp
# Run the app binary when we run the container
ENTRYPOINT ["/goapp"]