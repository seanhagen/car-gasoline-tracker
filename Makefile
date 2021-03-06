-include .env
export

BUILD_DIR ?= $(CURDIR)
CACHE_DIR ?= $(BUILD_DIR)/.tmp

GOVERSION = 1.5.3

ifeq ($(STACK),cedar-14)
export GOROOT := $(CACHE_DIR)/go/$(GOVERSION)
export PATH := $(GOROOT)/bin:$(PATH)
endif

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -type f -name '*.go')

BINARY=gasweb

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

REPO=github.com/seanhagen/car-gasoline-tracker
LDFLAGS=-ldflags "-X ${REPO}/core.Version=${VERSION} -X ${REPO}/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)
.PHONY: clean generate test vet deps all

$(BINARY): $(SOURCES) clean test vet
	godep go build ${LDFLAGS} -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

generate:
	go generate

test:
	go test -v  .

vet:
	go vet ./...

deps:
	godep save
	godep go install

all: deps generate $(BINARY) test vet

heroku: deps generate $(BINARY)
