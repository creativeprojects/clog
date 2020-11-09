# 
# Makefile for clog
# 
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOGET=$(GOCMD) get
GOPATH?=`$(GOCMD) env GOPATH`

TESTS=./...
COVERAGE_FILE=coverage.out

.PHONY: all test coverage clean staticcheck

all: test

test:
		$(GOTEST) -race -v $(TESTS)

coverage:
		$(GOTEST) -coverprofile=$(COVERAGE_FILE) $(TESTS)
		$(GOTOOL) cover -html=$(COVERAGE_FILE)

clean:
		$(GOCLEAN)
		rm -rf $(COVERAGE_FILE)

staticcheck:
	go get -u honnef.co/go/tools/cmd/staticcheck
	go run honnef.co/go/tools/cmd/staticcheck ./...
	go mod tidy
