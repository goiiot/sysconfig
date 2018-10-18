.PHONY: build test ensure

BIN=sysconfig-test

build:
	go build -o ${BIN} cmd/sysconfig/main.go

test: build
	./${BIN} -c ./config.yaml

ensure:
	go mod download
	go mod verify

demo-image:
	docker build -t sysconfig:demo -f demo.dockerfile .
