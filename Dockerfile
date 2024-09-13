FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/main ./cmd

EXPOSE 50051

CMD ["./bin/main"]