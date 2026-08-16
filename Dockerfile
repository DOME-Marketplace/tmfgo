# -------------------------------------------------------------------------
# STAGE 1: Build sqlite3_rsync
# -------------------------------------------------------------------------
FROM hesusruiz/sqlite3-rsync-alpine:3.51.2 AS sqlite3_rsync

# Stage 2: Build TMForum API server
FROM golang:1.25-alpine AS tmfbuilder

# Install build tools for CGO and sqlite tools
RUN apk update && \
    apk add --no-cache \
    build-base \
    musl-dev \
    linux-headers \
    git \
    gcc \
    sqlite-tools && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary with CGO enabled
# -ldflags="-w -s" strips debug information and symbols, reducing the binary size
RUN go build -ldflags="-w -s" -o /isbetmf .

# Final stage
FROM alpine/curl:latest

WORKDIR /
COPY --from=tmfbuilder /isbetmf /isbetmf
COPY www /www
COPY ./auth_policies.star /auth_policies.star
COPY --from=sqlite3_rsync --chmod=755 /usr/local/bin/sqlite3_rsync /usr/local/bin/sqlite3_rsync

HEALTHCHECK \
    --interval=300s \
    --timeout=5s \
    --start-period=10s \
    --start-interval=3s \
    --retries=3 \
    CMD curl -f http://localhost:9991/health || exit 1

# Expose the port the server runs on
EXPOSE 9991

# Run the binary
ENTRYPOINT ["/isbetmf"]