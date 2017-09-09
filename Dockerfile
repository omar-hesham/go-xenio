# Build xenio-cli in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-xenio
RUN cd /go-xenio && make xenio-cli

# Pull xenio-cli into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-xenio/build/bin/xenio-cli /usr/local/bin/

EXPOSE 8545 8546 30666 30666/udp
ENTRYPOINT ["xenio-cli"]
