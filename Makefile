.DEFAULT_GOAL := build

build:
	go build .

clean:
	rm -f sizeoci

lint:
	golangci-lint run ./...

.PHONY: build clean lint
