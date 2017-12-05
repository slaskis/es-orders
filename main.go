//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --gofast_out=. rpc/events.proto
//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --eventsource_out=. rpc/events.proto
//go:generate protoc -I .:vendor:vendor/github.com/gogo/protobuf/protobuf --gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc:. rpc/service.proto

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/altairsix/eventsource"
	"github.com/altairsix/eventsource/dynamodbstore"
	"github.com/slaskis/es-orders/rpc"
)

type options struct {
	host   string
	port   int
	table  string
	region string
}

func main() {
	opts := options{}
	flag.StringVar(&opts.host, "host", "", "listening host")
	flag.IntVar(&opts.port, "port", 7070, "listening port")
	flag.StringVar(&opts.table, "table", "orders", "dynamodb table name")
	flag.StringVar(&opts.region, "region", "eu-central-1", "dynamodb region")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", opts.host, opts.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("listening on %s\n", lis.Addr())

	store, err := dynamodbstore.New(opts.table,
		dynamodbstore.WithRegion(opts.region),
		dynamodbstore.WithDebug(os.Stderr),
	)
	if err != nil {
		log.Fatalln(err)
	}

	repo := eventsource.New(&rpc.Order{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)

	svc := service{repo: repo}

	srv := grpc.NewServer()
	rpc.RegisterCoreServiceServer(srv, svc)
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
