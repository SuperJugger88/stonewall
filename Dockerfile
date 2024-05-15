FROM golang:1.21.0-alpine
ENV PORT 80
EXPOSE 80

WORKDIR /go/src/app
COPY go.mod src/*.go ./

RUN go mod vendor
RUN go build -v -o app
RUN mv ./app /go/bin/

CMD ["app"]