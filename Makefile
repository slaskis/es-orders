
build: es-orders
	@: # sshhh
.PHONY: build

generate: rpc/*.pb.go
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

es-orders: *go
	go build .

%.pb.go: %.proto
	go generate .