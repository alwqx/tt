GO=$(shell which go)
CUR_DIR=$(shell pwd)

compile:
	$(GO) build

test: compile
	./tt --help
