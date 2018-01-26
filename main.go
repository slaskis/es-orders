package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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

	store, err := dynamodbstore.New(opts.table,
		dynamodbstore.WithRegion(opts.region),
		dynamodbstore.WithDebug(os.Stderr),
	)
	if err != nil {
		log.Fatalln(err)
	}

	orders := eventsource.New(&rpc.Order{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)

	customers := eventsource.New(&rpc.Order{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)

	mux := http.NewServeMux()
	mux.Handle("/", rpc.NewOrderServiceServer(NewOrderService(orders, customers), nil))
	mux.Handle("/", rpc.NewCustomerServiceServer(NewCustomerService(customers), nil))

	addr := fmt.Sprintf("%s:%d", opts.host, opts.port)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
