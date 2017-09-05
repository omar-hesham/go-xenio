FROM alpine:3.6

ADD . /go-xenio
RUN \
  apk add --no-cache git go make gcc musl-dev linux-headers && \
  (cd go-xenio && make xenio-cli)                           && \
  cp go-xenio/build/bin/xenio-cli /usr/local/bin/           && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /go-xenio && rm -rf /var/cache/apk/*

EXPOSE 8545 30666 30666/udp

ENTRYPOINT ["xenio-cli"]
