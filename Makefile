GO ?= go

# Run all tests.
test:
	@$(GO) test	-cover ./...
.PHONY: test

# Install all depenences.
install.dev:
	@$(GO) get -u github.com/golang/dep/cmd/dep
	@dep ensure
.PHONY: install.dev

# Run cmd for single build.
build:
	@echo	"===> run build"
	@$(GO) build -o bin/gost main.go
.PHONY: build
