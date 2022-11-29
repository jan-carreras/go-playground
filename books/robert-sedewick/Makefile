#!/usr/bin/env make

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test: ## Run all the tests
	@go test ./...

exercise:
	go build tools/exercise.go
	mv exercise ${HOME}/go/bin
