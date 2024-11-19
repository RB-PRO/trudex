#include .env

# Go related variables.
GOBASE=$(shell pwd)
GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

DOCKER_IMAGE_BASE=1.0

## clean: Clean build files. Runs `go clean` internally.
clean:
	@(MAKEFILE) go-clean

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean


.PHONY: docker

build: build-base push-base

build-base:
	@echo "  >  Build base docker image"
	@docker build --file base.Dockerfile --tag rbproxy/trudex_base:$(DOCKER_IMAGE_BASE) --force-rm .

push-base:
	@echo "  >  Push base docker image"
	@docker push rbproxy/trudex_base:$(DOCKER_IMAGE_BASE)

pull-base:
	@echo "  >  Push base docker image"
	@docker pull rbproxy/trudex_base:$(DOCKER_IMAGE_BASE)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo