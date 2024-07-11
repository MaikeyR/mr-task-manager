# Build stage
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go build -o main ./backend/cmd/server

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
EXPOSE 8080
CMD ["./main"]
