
build: es-orders
	@: # sshhh
.PHONY: build

generate: rpc/*.pb.go
	@: # sshhh
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

es-orders: *go
	go build .

%.pb.go: %.proto generate.go
	go generate .