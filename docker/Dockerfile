FROM golang:1.21.0-alpine as builder

WORKDIR /app
COPY go.mod cmd/*.go ./

RUN go mod tidy
RUN go build -v -o main

FROM golang:1.21.0-alpine

COPY --from=builder /app/main .

EXPOSE 80

ENTRYPOINT ["./main"]