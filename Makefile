
build: es-orders
	@: # sshhh
.PHONY: build

test: es-orders
	@go test ./...
	# @go test -coverprofile=coverage.out .
	# @go tool cover -func=coverage.out
.PHONY: test

bench: es-orders
	@go test -bench=.
.PHONY: bench

generate:
	go generate .
	# sed -i.tmp s@github.com/golang/protobuf/@github.com/gogo/protobuf/@ $(wildcard rpc/*/*.twirp.go)
	sed -i.tmp 's/package github.es.events.v1/package events/' $(wildcard rpc/*/*.es.go)
	rm rpc/*/*.tmp
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

es-orders: $(wildcard *.go **/*.go)
	go build .

