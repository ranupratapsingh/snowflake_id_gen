FROM golang:1.26.8-alpine AS builder
RUN apk upgrade --no-cache
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o main .
 
FROM alpine:latest
RUN apk upgrade --no-cache
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
 
CMD ["./main"]
