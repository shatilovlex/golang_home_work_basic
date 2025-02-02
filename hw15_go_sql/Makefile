include .env

LOCAL_BIN:=$(CURDIR)/bin

ENV_DIR = .env

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


env-prepare: ## copy .env from example
	cp -n .env.example .env

audit: ## Run linter checks
	golangci-lint run ./...

tidy: ## go  mod tidy run and reformat code
	go mod tidy
	go fmt ./...
	golangci-lint run --fix ./...

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.19.1

migration-status:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-add:
	goose -dir ${MIGRATION_DIR} create $(name) sql

migration-up:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v