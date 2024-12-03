FROM golang:1.23

WORKDIR /app
COPY . .

# Download and install any dependencies
RUN go mod tidy
# Build the Go application
RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
