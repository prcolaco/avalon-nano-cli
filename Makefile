# ########################################################## #
# Makefile for Golang Project Cross-compiling
# ########################################################## #

PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64

VERSION=$(shell git describe --long --tags)

PACKAGE=$(shell grep 'module' go.mod |cut -d " " -f 2)

default: build

all: clean build_all

build:
	@mkdir -p dist; go build -ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)'" -o dist

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

.PHONY: clean build install build_all all
