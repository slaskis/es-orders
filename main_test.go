package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/altairsix/eventsource"
	"github.com/gogo/protobuf/proto"
	"github.com/slaskis/es-orders/rpc"
)

func TestServer(t *testing.T) {
	orders := eventsource.New(&rpc.Order{},
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)
	customers := eventsource.New(&rpc.Customer{},
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)
	users := eventsource.New(&rpc.User{},
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stderr),
	)
	srv := createServiceHTTPHandler(orders, customers, users)

	ores1 := rpc.OrderResponse{}
	request(t, srv, rpc.OrderServicePathPrefix+"CreateOrder", &rpc.OrderNewRequest{}, &ores1)

	ores2 := rpc.OrderResponse{}
	request(t, srv, rpc.OrderServicePathPrefix+"GetOrder", &rpc.GetOrderRequest{ID: ores1.Order.ID}, &ores2)
	if ores2.Order.CustomerID != "" {
		t.Fatalf("expected no customer id")
	}

	ores3 := rpc.OrderResponse{}
	request(t, srv, rpc.OrderServicePathPrefix+"ApproveOrder", &rpc.OrderApproveRequest{ID: ores2.Order.ID}, &ores3)
	if ores3.Order.CustomerID == "" {
		t.Fatalf("expected customer id")
	}

	cres1 := rpc.CustomerResponse{}
	request(t, srv, rpc.CustomerServicePathPrefix+"GetCustomer", &rpc.GetCustomerRequest{ID: ores3.Order.CustomerID}, &cres1)
	if cres1.Customer.ID != ores3.Order.CustomerID {
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
