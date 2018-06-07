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

SQL_FILE=db/final.sql
not-containing = $(foreach v,$2,$(if $(findstring $1,$v),,$v))
LIST=$(wildcard db/queries/*.sql)
QUERIES=$(call not-containing,final,$(LIST))
$(SQL_FILE): $(QUERIES)
	cat $^ > $@

INTERNAL_FILE=internal/files.go
$(INTERNAL_FILE): $(SQL_FILE) deploy/ca-certificates.crt $(TEMPLATES)
	gotic -package internal $^ > $@

$(OUTPUT): $(INTERNAL_FILE) $(SOURCES)
	$(GO_BUILD_ENV) go build -a ${LDFLAGSWITCH} -o ${OUTPUT} -installsuffix cgo .

install:
	go get ./...

clean:

generate: $(INTERNAL_FILE) $(SQL_FILE)
	go generate

test:
	go test -v  .

vet:
	go vet ./...

build: $(OUTPUT)

build-clean: generate vet test build

deploy: clean build-clean
	gcloud app deploy ./deploy

clean:
	if [ -f ${OUTPUT} ] ; then rm ${OUTPUT} ; fi
	if [ -f $(SQL_FILE) ] ; then rm $(SQL_FILE) ; fi
	if [ -f $(INTERNAL_FILE) ] ; then rm $(INTERNAL_FILE) ; fi
