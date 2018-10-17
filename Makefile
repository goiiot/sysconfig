.PHONY: build test ensure

BIN=sysconfig-test

build:
	go build -o ${BIN} cmd/sysconfig/main.go

test: build
	./${BIN} -c ./config.yaml

ensure:
	go mod verify

github-action:
	./x-build.sh
	docker build -t sysconfig:github-action -f deploy.dockerfile .
