FROM golang:1.18 as builder

ENV GO111MODULE=on

WORKDIR /app

ADD . .

RUN go build -o polygon-edge main.go

FROM alpine:latest

RUN set -x \
    && apk add --update --no-cache \
       ca-certificates \
    && rm -rf /var/cache/apk/*

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /app/polygon-edge /usr/local/bin/

COPY --from=builder /app/spartan/config.json /opt/

COPY --from=builder /app/spartan/genesis.json /opt/

RUN mkdir -p /opt/logs

ENV GIN_MODE=release

WORKDIR /opt

EXPOSE 8545 9632 1478
ENTRYPOINT ["polygon-edge","server" ,"--config","config.json"]