#!make
include .env

ifeq ($(CONTAINER_ENV),)
CONTAINER_ENV=production
endif

ifeq ($(BINARY),)
BINARY=$(CMDNAME)
endif

ifeq ($(VERSION),)
export VERSION=$(shell cat VERSION)
export BUILD=$(shell ./version.sh)
export REPOBASE=github.com/seanhagen/gas-web
export LDFLAGSBASE=-X main.Version=${VERSION} -X main.Build=${BUILD}
endif

GO_BUILD_ENV:=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

TEMPLATE_DIR=./templates
TEMPLATES=$(shell find $(TEMPLATE_DIR) -name '*.tmpl')

REPO=$(shell git remote get-url origin)/$(BUILD)
OUTPUT=deploy/server

BUILD_TIME=$(shell date +%s)

LDFLAGS=$(LDFLAGSBASE) -X main.AppName=${BINARY} -X main.Repo=${REPO} -X main.BuildTime=${BUILD_TIME}
LDFLAGSWITCH=-ldflags "$(LDFLAGS)"

CONTAINER_TAG=gcr.io/biba-services/$(BINARY):$(VERSION)
CONTAINER_LATEST_TAG=gcr.io/biba-services/$(BINARY):latest


.DEFAULT_GOAL: $(OUTPUT)
.PHONY: clean generate test vet deps all

not-containing = $(foreach v,$2,$(if $(findstring $1,$v),,$v))
LIST=$(wildcard db/queries/*.sql)
QUERIES=$(call not-containing,final,$(LIST))
db/final.sql: $(QUERIES)
	cat $^ > db/final.sql

internal/files.go: config/ca-certificates.crt db/final.sql $(TEMPLATES)
	gotic -package internal $^ > $@

$(OUTPUT): internal/files.go $(SOURCES)
	$(GO_BUILD_ENV) go build -a ${LDFLAGSWITCH} -o ${OUTPUT} -installsuffix cgo .

install:
	go get ./...

clean:

generate: internal/files.go db/final.sql
	go generate

test:
	go test -v  .

vet:
	go vet ./...

build: $(OUTPUT)

build-clean: generate vet test build

build-container: build
	@echo "Building container using Google Container Builer"
	@gcloud container builds submit --tag $(CONTAINER_TAG) .
	@gcloud container images add-tag $(CONTAINER_TAG) $(CONTAINER_LATEST_TAG) --quiet
	@echo "Done and tagged"

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f db/final.sql ] ; then rm db/final.sql ; fi
	if [ -f internal/files.go ] ; then rm internal/files.go ; fi
