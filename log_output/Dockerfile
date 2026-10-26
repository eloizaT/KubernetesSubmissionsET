# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o log_output .

# Runtime stage
FROM scratch

COPY --from=builder /app/log_output /log_output

USER 65532:65532

ENTRYPOINT ["/log_output"]