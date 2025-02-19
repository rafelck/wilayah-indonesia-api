# Gunakan image dasar dari Golang
FROM golang:1.21 AS builder

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependencies
RUN go mod tidy

# Build aplikasi Go
RUN go build -o app .

# Stage baru untuk image yang lebih ringan
FROM alpine:latest

WORKDIR /root/

# Install dependencies jika dibutuhkan
RUN apk add --no-cache bash curl

# Copy binary dari tahap builder
COPY --from=builder /app/app .

## Copy file .env
#COPY .env .

# Jalankan aplikasi
CMD ["./app"]
