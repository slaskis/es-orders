package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/altairsix/eventsource"
	"github.com/gogo/protobuf/proto"
	"github.com/slaskis/es-orders/rpc/customer"
	"github.com/slaskis/es-orders/rpc/events"
	"github.com/slaskis/es-orders/rpc/order"
	"github.com/slaskis/es-orders/rpc/user"
)

func TestServer(t *testing.T) {
	orders := eventsource.New(&order.Order{},
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)
	customers := eventsource.New(&customer.Customer{},
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)
	users := eventsource.New(&user.User{},
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)
	srv := createServiceHTTPHandler(orders, customers, users)

	ores1 := order.OrderResponse{}
	request(t, srv, order.OrderServicePathPrefix+"CreateOrder", &order.OrderNewRequest{}, &ores1)

	ores2 := order.OrderResponse{}
	request(t, srv, order.OrderServicePathPrefix+"GetOrder", &order.GetOrderRequest{Id: ores1.Order.Id}, &ores2)
	if ores2.Order.CustomerId != "" {
		t.Fatalf("expected no customer id")
	}

	ores3 := order.OrderResponse{}
	request(t, srv, order.OrderServicePathPrefix+"ApproveOrder", &order.OrderApproveRequest{Id: ores2.Order.Id}, &ores3)
	if ores3.Order.CustomerId == "" {
		t.Fatalf("expected customer id")
	}

	cres1 := customer.CustomerResponse{}
	request(t, srv, customer.CustomerServicePathPrefix+"GetCustomer", &customer.GetCustomerRequest{Id: ores3.Order.CustomerId}, &cres1)
	if cres1.Customer.Id != ores3.Order.CustomerId {
		t.Fatalf("expected customer id")
	}
}

func request(t *testing.T, srv http.Handler, path string, in proto.Message, out proto.Message) {
	bod, err := proto.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("POST", path, bytes.NewReader(bod))
	req.Header.Set("content-type", "application/protobuf")
	req.Header.Set("accept", "application/protobuf")
	srv.ServeHTTP(rec, req)
	t.Logf("%+v", rec)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected ok status code. got %d", rec.Code)
	}
	err = proto.Unmarshal(rec.Body.Bytes(), out)
	if err != nil {
		t.Fatal(err)
	}
}
