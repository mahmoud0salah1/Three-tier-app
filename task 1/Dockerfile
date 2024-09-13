# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Create a small image to run the Go binary
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main /app/
EXPOSE 8080
CMD ["./main"]
