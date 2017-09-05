.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/xenioplatform/go-xenio/cmd/swarm
	go build -o bin/xenio-cli     github.com/xenioplatform/go-xenio/cmd/xenio-cli
	go build -o bin/bootnode github.com/xenioplatform/go-xenio/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/xenioplatform/go-xenio/swarm/...
