
build: *.go
	go build .

generate: rpc/*.pb.go
.PHONY: generate

vendor:
	dep ensure
.PHONY: vendor

%.pb.go: %.proto
	go generate .