.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/Rue-Foundation/go-rue/cmd/swarm
	go build -o bin/geth     github.com/Rue-Foundation/go-rue/cmd/geth
	go build -o bin/bootnode github.com/Rue-Foundation/go-rue/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/ethereum/go-ethereum/swarm/...
