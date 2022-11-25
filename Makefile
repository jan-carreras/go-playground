#!/usr/bin/env make

help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test-all: ## Run all the tests on all the projects
	make --directory books/robert-sedewick test