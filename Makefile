.PHONY: all
all: format test lint

.PHONY: format
format:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
