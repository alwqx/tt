GO=$(shell which go)
CUR_DIR=$(shell pwd)

compile:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build

test: compile
	./tt --help
