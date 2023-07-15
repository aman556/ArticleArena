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
PROJECT := aman55/articlearena

.PHONY: build server
build-server: ## Build binary.
	cd backend && $(GO) build -a -o ./bin/articleArena-server

.PHONY: run server
run-server: ## Build and run server
	cd backend && $(GO) build -a -o ./bin/articleArena-server && bin/articleArena-server

.PHONY: build server image & pushes to registry
build-server-image: # Build docker image for backend
	docker build --platform=linux/amd64 --no-cache --tag ${PROJECT}:$(release-tag) backend && docker push ${PROJECT}:$(release-tag)
