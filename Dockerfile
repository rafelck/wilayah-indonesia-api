# Gunakan image base Golang versi terbaru
FROM golang:1.24-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum (jika ada)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode aplikasi ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["./main"]