.SUFFIXES:

SOURCE := $(wildcard *.go **/*.go)

build: es-orders
	@: # sshhh
.PHONY: build

test: es-orders
	@go test ./...
.PHONY: test

cover: es-orders
	@go test -cover ./...
.PHONY: cover

cover-func: coverage-func.out
	@: # sshhh
.PHONY: cover-func

coverage-%.out: FORCE
	@echo "mode: count" > $@
	@$(foreach pkg,$(shell go list ./...),\
	go test -coverprofile=tmp.out  $(pkg) > /dev/null;\
	tail -n +2 tmp.out >> $@;)
	@go tool cover -$*=$@
	@rm tmp.out
FORCE:

bench: es-orders
	@go test -bench=.
.PHONY: bench

generate:
	go generate .
	sed -i.tmp 's/package github.es.events.v1/package events/' $(wildcard rpc/*/*.es.go)
	rm rpc/*/*.tmp
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

es-orders: ${SOURCE}
	go build .


