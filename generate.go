//go:generate protoc -I .:rpc:vendor:vendor/github.com/gogo/protobuf/protobuf --lint_out=. --gogofast_out=. --twirp_out=. rpc/user/service.proto
//go:generate protoc -I .:rpc:vendor:vendor/github.com/gogo/protobuf/protobuf --lint_out=. --gogofast_out=. --twirp_out=. rpc/order/service.proto
//go:generate protoc -I .:rpc:vendor:vendor/github.com/gogo/protobuf/protobuf --lint_out=. --gogofast_out=. --twirp_out=. rpc/customer/service.proto
//go:generate protoc -I .:rpc:vendor --gogofast_out=. --eventsource_out=.  rpc/events/events.proto

package main
