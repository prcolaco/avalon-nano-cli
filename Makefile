# ########################################################## #
# Makefile for Golang Project Cross-compiling
# ########################################################## #

PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64

VERSION=$(shell git describe --long --tags)

PACKAGE=$(shell grep 'module' go.mod |cut -d " " -f 2)

IMAGE="avalon-cron"

default: build

all: clean build_all

build:
	@mkdir -p dist; go build -ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)'" -o dist

build_docker:
	@mkdir -p dist/docker && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)'" -o dist/docker && \
	docker build -t $(IMAGE):$(VERSION) .

build_all:
	$(foreach GOOS, $(PLATFORMS),\
		$(foreach GOARCH, $(ARCHITECTURES),\
			$(shell ./build.sh $(GOOS) $(GOARCH) $(VERSION) $(PACKAGE))\
		)\
	)

clean:
	@rm -rf dist

install:
	@go install

.PHONY: clean build build_docker install build_all all
