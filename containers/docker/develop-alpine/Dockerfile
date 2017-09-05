FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 github.com/xenioplatform/go-xenio && \
  (cd go-xenio && make xenio-cli) && \
  cp go-xenio/build/bin/xenio-cli /xenio-cli && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-xenio && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30666

ENTRYPOINT ["/xenio-cli"]
