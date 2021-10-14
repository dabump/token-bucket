GO_FLAGS	?=
NAME	 	:= tokenbukcet
PACKAGE		:= github.com/dabump/$(NAME)

default: help

test: ## Run all tests
	@go clean --testcache && go test ./...

cover: ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

build: ## Builds the token bucket
	@go build ${GO_FLAGS}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'