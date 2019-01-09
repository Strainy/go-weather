GIT ?= git
GO_VARS ?=
GO ?= go
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= 1.0.0
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
ROOT := main
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)

.PHONY: clean help test weather

default: clean test weather

clean:
	rm -rf weather

help:
	@echo "Please use \`make <TARGET>' where <TARGET> is one of"
	@echo "  clean		to clean the working directory"
	@echo "  test		to run unittests"
	@echo "  weather	to build the main binary for current platform"

test:
	$(GO_VARS) $(GO) test -v

weather:
	$(GO_VARS) $(GO) build -v -ldflags "$(LD_FLAGS)" ./cmd/weather