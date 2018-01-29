package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/twitchtv/twirp"

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

	customers := eventsource.New(&rpc.Customer{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)

	mux := createServiceHTTPHandler(orders, customers)
	addr := fmt.Sprintf("%s:%d", opts.host, opts.port)
	log.Printf("listening at %s", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}

func logger() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			p, _ := twirp.PackageName(ctx)
			s, _ := twirp.ServiceName(ctx)
			m, _ := twirp.MethodName(ctx)
			log.Printf("RequestReceived %s.%s/%s", p, s, m)
			return ctx, nil
		},
	}
}

func createServiceHTTPHandler(orders, customers *eventsource.Repository) http.Handler {
	return handleServers(
		rpc.NewOrderServiceServer(NewOrderService(orders, customers), logger()),
		rpc.NewCustomerServiceServer(NewCustomerService(customers), logger()),
	)
}

func handleServers(orders rpc.TwirpServer, customers rpc.TwirpServer) http.HandlerFunc {
	notFound := http.NotFoundHandler()
	return func(res http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, rpc.OrderServicePathPrefix) {
			orders.ServeHTTP(res, req)
		} else if strings.HasPrefix(req.URL.Path, rpc.CustomerServicePathPrefix) {
			customers.ServeHTTP(res, req)
		} else {
			notFound.ServeHTTP(res, req)
		}
	}
}
