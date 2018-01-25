//go:generate protoc -I .:vendor --gofast_out=. --twirp_out=. --eventsource_out=. rpc/order_service.proto rpc/events.proto

package main
