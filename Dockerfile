FROM golang:1.15 AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN mkdir -p main && go build -o main ./...

# Build a small image
FROM alpine:latest

RUN apk add --no-cache bash
RUN mkdir fixtures

ENV PORT 8082

COPY --from=builder /build/main/tiny-url-service /
EXPOSE 8082

# Command to run
ENTRYPOINT ["/tiny-url-service"]