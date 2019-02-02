generate:
	@go generate ./...
.PHONY: generate

build: generate
	@echo "====> Build gost cli"
	@go build -o ./bin/gost main.go
.PHONY: build

test:
	@go test ./...
.PHONY: test

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic
.PHONY: test.cov
