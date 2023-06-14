FROM golang:1.18-alpine

WORKDIR /

COPY . /app

RUN go mod init github.com/Nelwhix/docker-go

RUN go mod download

RUN go build -o main .

EXPOSE 9000

CMD ["./main"]