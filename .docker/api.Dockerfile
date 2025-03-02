FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

COPY stonewall-api .

RUN --mount=type=cache,target=/gomod-cache \
    go mod tidy -x

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
    go build -v -o main

FROM busybox:latest

WORKDIR /srv

COPY --from=builder --chown=1000 /app/main .

EXPOSE 4000

CMD ["./main"]