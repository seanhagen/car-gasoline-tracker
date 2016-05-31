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

$(BINARY): $(SOURCES) clean
	godep go build ${LDFLAGS} -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

generate:
	go generate

test:
	go test -v -coverprofile=out.cov -covermode atomic -cover ./... | go2xunit -output tests.xml
	gocover-cobertura < out.cov > coverage.xml

vet:
	go vet

deps:
	godep go install
	# go get golang.org/x/tools/cmd/vet
	# go get golang.org/x/tools/cmd/cover
	# go get github.com/golang/lint/golint
	# go get bitbucket.org/tebeka/go2xunit
	# go get github.com/gchaincl/dotsql
	# go get github.com/gorilla/context
	# go get github.com/julienschmidt/httprouter
	# go get github.com/justinas/alice
	# go get github.com/lib/pq
	# go get github.com/rs/cors
	# go get github.com/satori/go.uuid
	# go get github.com/seanhagen/ldap
	# go get github.com/unrolled/render
	# go get github.com/vharitonsky/iniflags
	# go get github.com/go-gomail/gomail
	# go get github.com/t-yuki/gocover-cobertura

all: deps generate $(BINARY) test vet

heroku: deps generate $(BINARY)
