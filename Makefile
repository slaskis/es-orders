
build: generate es-orders
	@: # sshhh
.PHONY: build

test: generate es-orders
	@go test -coverprofile=coverage.out .
	@go tool cover -func=coverage.out
.PHONY: test

generate: rpc/*.pb.go
	@: # sshhh
.PHONY: generate

fix: rpc/*.twirp.go
	@: # sshhh
.PHONY: fix

vendor:
	dep ensure
.PHONY: vendor

es-orders: $(wildcard *.go **/*.go)
	go build .

%.pb.go: %.proto generate.go
	go generate .

%.twirp.go: %.proto
	sed -i.tmp s@github.com/golang/protobuf/@github.com/gogo/protobuf/@ $@
	rm $@.tmp