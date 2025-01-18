ifneq ($(wildcard .env),)
	include .env
endif

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
PROJECT_TMP = $(PROJECT_DIR)/tmp

PLATFORMS=linux/amd64

install-deps:
	GOBIN=$(PROJECT_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


gen-proto:
	rm -f ./api/proto/*.pb.go

	protoc --proto_path=./api/proto/ --go_out=./ ./api/proto/*.proto