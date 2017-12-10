//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --gogofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:. --eventsource_out=. rpc/service.proto rpc/events.proto

package main
