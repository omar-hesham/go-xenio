.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    https://github.com/xenioplatform/go-xenio/cmd/swarm
	go build -o bin/geth     https://github.com/xenioplatform/go-xenio/cmd/geth
	go build -o bin/bootnode https://github.com/xenioplatform/go-xenio/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v https://github.com/xenioplatform/go-xenio/swarm/...
