# Use the official Go image from the DockerHub
FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
