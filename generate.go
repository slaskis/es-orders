//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --gofast_out=. rpc/events.proto
//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --eventsource_out=. rpc/events.proto
//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:. rpc/service.proto

package main
