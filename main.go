package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/twitchtv/twirp"

	"github.com/altairsix/eventsource"
	"github.com/altairsix/eventsource/dynamodbstore"
	"github.com/slaskis/es-orders/internal/customerserver"
	"github.com/slaskis/es-orders/internal/orderserver"
	"github.com/slaskis/es-orders/internal/userserver"
	"github.com/slaskis/es-orders/rpc/customer"
	"github.com/slaskis/es-orders/rpc/events"
	"github.com/slaskis/es-orders/rpc/order"
	"github.com/slaskis/es-orders/rpc/user"
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

	orders := eventsource.New(&order.Order{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)

	customers := eventsource.New(&customer.Customer{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)

	users := eventsource.New(&user.User{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stdout),
	)

	mux := createServiceHTTPHandler(orders, customers, users)
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

func createServiceHTTPHandler(orders, customers, users *eventsource.Repository) http.Handler {
	var entropy = rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano()))

	return handleServers(
		order.NewOrderServiceServer(orderserver.NewServer(orders, customers, users, entropy), logger()),
		customer.NewCustomerServiceServer(customerserver.NewServer(customers), logger()),
		user.NewUserServiceServer(userserver.NewServer(users), logger()),
	)
}

func handleServers(orders order.TwirpServer, customers customer.TwirpServer, users user.TwirpServer) http.HandlerFunc {
	notFound := http.NotFoundHandler()
	return func(res http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, order.OrderServicePathPrefix) {
			orders.ServeHTTP(res, req)
		} else if strings.HasPrefix(req.URL.Path, customer.CustomerServicePathPrefix) {
			customers.ServeHTTP(res, req)
		} else if strings.HasPrefix(req.URL.Path, user.UserServicePathPrefix) {
			users.ServeHTTP(res, req)
		} else {
			notFound.ServeHTTP(res, req)
		}
	}
}
