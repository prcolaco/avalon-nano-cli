# ########################################################## #
# Makefile for Golang Project Cross-compiling
# ########################################################## #

PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64

DOCKER_PLATFORM=linux
DOCKER_ARCHITECTURE=amd64
DOCKER_REGISTRY?=no-registry
DOCKER_IMAGE=avalon-cron

VERSION=$(shell git describe --long --tags)

PACKAGE=$(shell grep 'module' go.mod |cut -d " " -f 2)


default: build

all: clean build_release

build:
	@mkdir -p dist; go build -ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)'" -o dist

build_release:
	$(foreach GOOS, $(PLATFORMS),\
		$(foreach GOARCH, $(ARCHITECTURES),\
			$(shell ./build.sh $(GOOS) $(GOARCH) $(VERSION) $(PACKAGE))\
		)\
	)

clean:
	@rm -rf dist

install:
	@go install

docker: docker_build docker_publish

docker_build: check_env
	@mkdir -p dist/docker && \
	CGO_ENABLED=0 GOOS=$(DOCKER_PLATFORM) GOARCH=$(DOCKER_ARCHITECTURE) go build -ldflags="-X '$(PACKAGE)/cmd.Version=$(VERSION)'" -o dist/docker && \
	docker build --platform $(DOCKER_PLATFORM)/$(DOCKER_ARCHITECTURE) -t $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(VERSION) -t $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest .

docker_publish: check_env
	@docker push --all-tags $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)

check_env:
ifeq ($(DOCKER_REGISTRY),no-registry)
	$(error please set the docker registry in DOCKER_REGISTRY variable)
endif

.PHONY: clean build install build_release docker_build docker_publish
