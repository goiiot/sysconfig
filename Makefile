.PHONY: build test ensure

build:
	go build -o server cmd/server/main.go

test: build
	./server -c ./config.yaml

ensure:
	go mod verify
