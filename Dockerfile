# Start with the official Go image for building
FROM docker.arvancloud.ir/golang:1.18 AS builder

WORKDIR /app

# Set the GOPROXY environment variable
ENV GOPROXY=https://goproxy.io,direct

# Copy go mod and sum files first, leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application for the production environment
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage with a minimal image
FROM docker.arvancloud.ir/alpine:latest

RUN apk add --no-cache ca-certificates curl

WORKDIR /root/

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 CMD curl --fail http://localhost/health || exit 1

# Copy the pre-built binary from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"]
