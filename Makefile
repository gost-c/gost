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

# Run cmd for up tool build.
build.up:
	@echo	"===> run build up tools"
	@$(GO) build -o bin/up-gen cmd/up/main.go
.PHONY: build.up

init.up:
	./bin/up-gen
.PHONY: init.up

install.up:
	@curl -sf https://up.apex.sh/install | sh
.PHONY: install.up
