//go:generate protoc -I .:rpc:vendor:vendor/github.com/gogo/protobuf/protobuf --gogofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. --twirp_out=. --eventsource_out=. rpc/customer_service.proto rpc/order_service.proto rpc/events.proto
//go:generate sed -itmp s@github.com/golang/protobuf/@github.com/gogo/protobuf/@ rpc/customer_service.twirp.go rpc/order_service.twirp.go

package main
