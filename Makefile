VERSION:=1.0.1

TIMESTAMP=$(shell date +'%Y%m%d%H%M%S')
SERVICE_NAME=meshtastic-go
BIN?=$(pwd)/bin
BINARY?=$(BIN)/meshtastic-go
BUILD_ARGS?=

CONFIG_PATH?=$(pwd)/etc/config.yaml

OSFLAG :=
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSFLAG = linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OSFLAG = darwin
	endif

.PHONY: help
help:
	@echo "clean                remove artifacts"
	@echo "test                 run unit-tests for all service packages"
	@echo "coverage             run tests with coverage"
	@echo "fmt                  format application sources"
	@echo "vet                  check code with go vet command"
	@echo "lint                 check for linter errors"
	@echo "mod                  install dependencies modules to global GOPATH"
	@echo "modupdate            update all dependencies (it's not recommended, to update one package use go get)"
	@echo "build                build application from sources"
	@echo "build-release        build application for all supported platforms"
	@echo "run                  build and run application"

.PHONY: clean
clean: env
	cd src && go clean $(allpackages)
	rm -rf $(BIN)/*

.PHONY: test
test: clean env code-generate
	cd src && CONFIG_PATH=$(CONFIG_PATH_TEST) SERVICE_NAME=$(SERVICE_NAME) go test -v $(allpackages)

.PHONY: coverage
coverage: clean env code-generate
	cd src && CONFIG_PATH=$(CONFIG_PATH_TEST) SERVICE_NAME=$(SERVICE_NAME) go test -v -cover $(allpackages)

.PHONY: fmt
fmt: env
	cd src && go fmt $(allpackages)

.PHONY: vet
vet: env
	cd src && go vet $(allpackages)

.PHONY: lint
lint: fmt vet
	cd src && GOROOT=$(shell go env GOROOT) golint -set_exit_status $(allpackages)

.PHONY: mod
mod: env
	cd src && go mod tidy

.PHONY: modupdate
modupdate: env
	cd src && go list -m -u all

.PHONY: build
build: env fmt
	cd src && \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	go build $(BUILD_ARGS) -o $(BINARY) service/cmd

.PHONY: build-release
build-release: build-linux-amd64 build-darwin-arm64 build-darwin-amd64 build-windows-amd64

.PHONY: build-linux-amd64
build-linux-amd64: env fmt
	cd src && \
	GOOS=linux \
	GOARCH=amd64 \
	go build $(BUILD_ARGS) -o $(BINARY)-$(VERSION)-linux-amd64 service/cmd

.PHONY: build-darwin-arm64
build-darwin-arm64: env fmt
	cd src && \
	GOOS=darwin \
	GOARCH=arm64 \
	go build $(BUILD_ARGS) -o $(BINARY)-$(VERSION)-darwin-arm64 service/cmd

.PHONY: build-darwin-amd64
build-darwin-amd64: env fmt
	cd src && \
	GOOS=darwin \
	GOARCH=amd64 \
	go build $(BUILD_ARGS) -o $(BINARY)-$(VERSION)-darwin-amd64 service/cmd

.PHONY: build-windows-amd64
build-windows-amd64: env fmt
	cd src && \
	GOOS=windows \
	GOARCH=amd64 \
	go build $(BUILD_ARGS) -o $(BINARY)-$(VERSION)-windows-amd64.exe service/cmd

.PHONY: run
run: build
	CONFIG_PATH=$(CONFIG_PATH) SERVICE_NAME=$(SERVICE_NAME) $(BINARY)

.PHONY: env
env:
	go env -w GOPROXY=proxy.golang.org,$(PRIVATE_PROXY),direct
	go env -w GOSUMDB=off

_allpackages=$(shell find src -name '*.go' -exec dirname {} \; | sed -e 's/src/service/g' | uniq )
allpackages=$(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)

_pwd=$(shell pwd)
pwd=$(if $(__pwd),,$(eval __pwd := $$(_pwd)))$(__pwd)