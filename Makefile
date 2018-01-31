
build: generate es-orders
	@: # sshhh
.PHONY: build

test: generate es-orders
	@go test -coverprofile=coverage.out .
	@go tool cover -func=coverage.out
.PHONY: test

bench: generate es-orders
	@go test -bench=.
.PHONY: bench

generate: $(wildcard rpc/*.pb.go)
	@: # sshhh
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

es-orders: $(wildcard *.go **/*.go)
	go build .

%.pb.go: %.proto generate.go
	go generate .
	sed -i.tmp s@github.com/golang/protobuf/@github.com/gogo/protobuf/@ $(wildcard rpc/*.twirp.go)
	rm rpc/*.tmp
