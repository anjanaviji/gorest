FROM golang:latest

RUN apk add --no-cache git

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN go build -o patientapi .

EXPOSE 8080

CMD ["./patientapi"]