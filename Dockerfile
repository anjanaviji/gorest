FROM golang:latest


RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN go build -o patientapi .

EXPOSE 1000

CMD ["./patientapi"]
