FROM golang:1.21.0-alpine
EXPOSE 80

WORKDIR /app
COPY go.mod cmd/*.go ./

RUN go mod tidy
RUN go build -v -o main

CMD ["./main"]