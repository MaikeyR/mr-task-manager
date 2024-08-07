FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o main ./backend/cmd/server
EXPOSE 8080
CMD ["./main"]
