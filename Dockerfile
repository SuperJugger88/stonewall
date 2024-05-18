FROM golang:1.21.0-alpine
EXPOSE 80

WORKDIR /app
COPY go.mod cmd/*.go ./

RUN go mod vendor
RUN go build -v -o app
RUN mv ./app /go/bin/

CMD ["app"]