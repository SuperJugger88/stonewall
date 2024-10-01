FROM golang:1.23.1-alpine AS builder

WORKDIR /app

RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

COPY stonewall-api/ .

RUN --mount=type=cache,target=/gomod-cache \
    go mod tidy

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
    go build -o main

FROM busybox:latest

USER gopher

WORKDIR /srv

COPY --from=builder --chown=gopher /app/main .

EXPOSE 4000

ENTRYPOINT ["./main"]