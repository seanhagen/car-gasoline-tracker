SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=gasweb

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

REPO=github.com/seanhagen/car-gasoline-tracker
LDFLAGS=-ldflags "-X ${REPO}/core.Version=${VERSION} -X ${REPO}/core.BuildTime=${BUILD_TIME}"

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES) clean
	go build ${LDFLAGS} -o ${BINARY}

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: generate
generate:
	go generate

.PHONY: test
test:
	go test -v -coverprofile=out.cov -covermode atomic -cover ./... | go2xunit -output tests.xml
	gocover-cobertura < out.cov > coverage.xml

.PHONY: vet
vet:
	go vet

.PHONY: deps
deps:
	go get golang.org/x/tools/cmd/vet
	go get golang.org/x/tools/cmd/cover
	go get github.com/golang/lint/golint
	go get bitbucket.org/tebeka/go2xunit
	go get github.com/gchaincl/dotsql
	go get github.com/gorilla/context
	go get github.com/julienschmidt/httprouter
	go get github.com/justinas/alice
	go get github.com/lib/pq
	go get github.com/rs/cors
	go get github.com/satori/go.uuid
	go get github.com/seanhagen/ldap
	go get github.com/unrolled/render
	go get github.com/vharitonsky/iniflags
	go get github.com/go-gomail/gomail
	go get github.com/t-yuki/gocover-cobertura

.PHONY: all
all: deps generate $(BINARY) test vet
