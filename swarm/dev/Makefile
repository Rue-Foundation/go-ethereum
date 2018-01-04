.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/Rue-Foundation/go-rue/cmd/swarm
	go build -o bin/grue     github.com/Rue-Foundation/go-rue/cmd/grue
	go build -o bin/bootnode github.com/Rue-Foundation/go-rue/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/Rue-Foundation/go-rue/swarm/...
