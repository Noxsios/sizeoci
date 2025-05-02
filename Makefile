.DEFAULT_GOAL := build

build:
	go build .

clean:
	rm -f sizeoci

test:
	go test ./...

lint:
	golangci-lint run ./...

.PHONY: build clean test lint
