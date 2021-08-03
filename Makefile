MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
ifeq ($(word 1,$(subst ., ,$(MAKE_VERSION))),4)
.SHELLFLAGS := -eu -o pipefail -c
endif
.DEFAULT_GOAL := help
.ONESHELL:

PACKAGE := $(shell go list)

.PHONY: help
help:  ## Show this help
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.* ##/:/g' | column -t -s:

.PHONY: test
test:  ## Run go test
	go test -v ./...

bin/go-dependency-list.darwin:  ## Build the application binary for current OS

bin/go-dependency-list.%:  ## Build the application binary for target OS, for example bin/go-dependency-list.linux
	GOOS=$* go build -o $@ -ldflags "-X $(PACKAGE)/version=$(GIT_TAG)+$(GIT_REF)" main.go

.PHONY: install
install: bin/go-dependency-list.darwin ## Install the binary
	cp $< ~/bin/go-dependency-list

