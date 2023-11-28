APP_NAME=$(shell basename $(PWD))

VERSION=$(shell git describe)
BUILD_TIME=$(shell date +%FT%T%z)
GO_MODULE_PATH=$(shell go run ./internal/cmd/getmodulepath)

all: build

build: $(APP_NAME)

$(APP_NAME): $(wildcard *.go) $(wildcard cmd/*.go)
	CGO_ENABLED=0 GOFLAGS=-trimpath \
	go build -mod=readonly -ldflags='-s -w -X $(GO_MODULE_PATH)/internal.Release=$(VERSION) -X $(GO_MODULE_PATH)/internal.BuildTime=$(BUILD_TIME)' \
		-o ./$@ ./cmd

clean:
	rm -f $(APP_NAME)

.PHONY: all build clean
