GO ?= $(shell command -v go 2> /dev/null)
GOPROXY := $(shell go env GOPROXY)
ifeq ($(GOPROXY),)
GOPROXY := https://proxy.golang.org
endif
export GOPROXY

GOPATH := $(shell go env GOPATH)

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

# Harbor registry
PROJECT := aman55/articlearenadb

.PHONY: build database image & pushes to registry
build-database-image: # Build docker image for backend
	docker build --platform=linux/amd64 --tag ${PROJECT}:$(release-tag) . && docker push ${PROJECT}:$(release-tag)
