//go:generate protoc -I .:rpc:vendor:vendor/github.com/gogo/protobuf/protobuf --gogofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. --twirp_out=. rpc/customer_service.proto rpc/order_service.proto rpc/user_service.proto
//go:generate protoc -I .:rpc:vendor --gogofast_out=. --eventsource_out=. rpc/events.proto

package main
