FROM alpine:3.6

ADD . /go-xenio
RUN \
  apk add --no-cache git go make gcc musl-dev linux-headers && \
  (cd go-xenio && make geth)                           && \
  cp go-xenio/build/bin/geth /usr/local/bin/           && \
  apk del git go make gcc musl-dev linux-headers          && \
  rm -rf /go-xenio && rm -rf /var/cache/apk/*

EXPOSE 8545 30303 30303/udp

ENTRYPOINT ["geth"]
