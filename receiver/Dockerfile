FROM golang:1.14-alpine

WORKDIR /code

COPY . ./

RUN go mod download

RUN go build

EXPOSE 8080

CMD ["./receiver"]
