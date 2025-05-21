# Step 1: Build the Go binary
FROM golang:1.23.6-alpine3.20 AS builder
WORKDIR /app
COPY . .
COPY .env .
RUN go mod download
RUN go build -o main

# Step 2: Run it in a lightweight container
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

# Make sure Dockerfile is in the same directory as main.go, go.mod, etc. For example:
# /to_do_list
# ├── main.go
# ├── go.mod
# ├── Dockerfile ✅ should be here