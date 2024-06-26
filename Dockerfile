# Build the binary
FROM golang:1.22.1-alpine as builder
WORKDIR /app
COPY . .
RUN go clean --modcache
RUN go get
RUN go build -o executable

# Run the app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/executable /app/executable

EXPOSE 5000

CMD ["./executable"]
