FROM golang:1.18-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

ENV GO111MODULE=on

WORKDIR /app

ADD . .

RUN make build

FROM alpine:latest

RUN set -x \
    && apk add --update --no-cache \
       ca-certificates \
    && rm -rf /var/cache/apk/*

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /app/polygon-edge /usr/bin/

RUN mkdir -p /opt/logs

ENV GIN_MODE=release

WORKDIR /opt

EXPOSE 8545 9632 1478
ENTRYPOINT ["polygon-edge"]