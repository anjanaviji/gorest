FROM golang:1.12.0-alpine3.9

RUN apk add --no-cache git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]