FROM golang:alpine

WORKDIR /src/

COPY go.mod src/*.go ./

RUN go build -o main

CMD ./main