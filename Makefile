#include .env

# Go related variables.
GOBASE=$(shell pwd)
GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
CLO_PORT=8080
CLI_TOKEN=$(CLI_TOKEN)

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
	@docker build --file trud_base.Dockerfile --tag rbproxy/trud_base:$(DOCKER_IMAGE_BASE) --force-rm .

push-base:
	@echo "  >  Push base docker image"
	@docker push rbproxy/trud_base:$(DOCKER_IMAGE_BASE)

pull-base:
	@echo "  >  Push base docker image"
	@docker pull rbproxy/trud_base:$(DOCKER_IMAGE_BASE)

clo:
	@clo publish http $(CLO_PORT)

clo-install:
	@wget https://cloudpub.ru/download/stable/clo-1.2.21-stable-linux-x86_64.tar.gz
	@sudo dpkg -i cloudpub-1.1.21-stable-linux-x86_64.deb
	@clo set token $(CLI_TOKEN)

build-start-compose:
	@docker-compose up -d --build

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo