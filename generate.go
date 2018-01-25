//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --gogofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. --twirp_out=. --eventsource_out=. rpc/order_service.proto rpc/events.proto

package main
