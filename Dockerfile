FROM golang:1.24.6

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd

CMD ["./main"]

EXPOSE 8080