FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY server.go .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o nget-server server.go

FROM alpine:latest

# Install minimal dependencies
# - curl: for fetching public IP
# - nano: text editor (also provides pico-like interface)
# - vim: text editor (also provides vi)
# - emacs: text editor
RUN apk add --no-cache curl nano vim emacs

# Copy binaries
COPY --from=builder /build/nget-server /usr/local/bin/nget-server
COPY nget /usr/local/bin/nget
RUN chmod +x /usr/local/bin/nget /usr/local/bin/nget-server

# Create working directories
RUN mkdir -p /tmp/nget/state /tmp/nget/www

# Default volume mount point
VOLUME ["/data"]

# Expose port range for nget HTTP servers
EXPOSE 33000-33100

# Set working directory to mounted volume
WORKDIR /data

# Keep container running with shell
CMD ["/bin/sh"]
