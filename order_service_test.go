package main

import (
	"context"
	"os"
	"testing"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc"
)

func TestOrderService(t *testing.T) {
	obs := func(event eventsource.Event) {
		t.Logf("observer event: %s", event)
	}

	orders := eventsource.New(&rpc.Order{},
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stdout),
		eventsource.WithObservers(obs),
	)

	customers := eventsource.New(&rpc.Customer{},
		eventsource.WithSerializer(rpc.NewSerializer()),
		eventsource.WithDebug(os.Stdout),
		eventsource.WithObservers(obs),
	)

	osvc := NewOrderService(orders, customers)
	csvc := NewCustomerService(customers)

	ctx := context.Background()
	res, err := osvc.CreateOrder(ctx, &rpc.OrderNewRequest{
		Items: []*rpc.NewItem{
			&rpc.NewItem{SKU: "1", Description: "One"},
			&rpc.NewItem{SKU: "2", Description: "Two"},
		},
	})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Logf("order: %+v", res.Order)
	if res.Order == nil {
		t.Fatalf("order must not be nil")
	}
	if res.Order.Status != rpc.OrderStatus_PENDING {
		t.Fatalf("order must be pending")
	}
	if res.Order.CreatedAt.IsZero() {
		t.Fatalf("order must have created at date")
	}
	if res.Order.UpdatedAt.IsZero() {
		t.Fatalf("order must have updated at date")
	}
	if res.Order.DeletedAt != nil {
		t.Fatalf("order must not have deleted at date")
	}
	if len(res.Order.Items) != 2 {
		t.Fatalf("order must have 2 items")
	}
	item := res.Order.Items[0]
	res2, err := osvc.RemoveItem(ctx, &rpc.OrderItemRemoveRequest{OrderID: res.Order.ID, ItemID: item.ID})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res2.Order.Status != rpc.OrderStatus_PENDING {
		t.Fatalf("order must be pending")
	}
	if len(res2.Order.Items) != 1 {
		t.Fatalf("order must have 1 items")
	}
	if res2.Order.CustomerID != "" {
		t.Fatalf("order must not have customer id")
	}

	res3, err := osvc.ApproveOrder(ctx, &rpc.OrderApproveRequest{ID: res.Order.ID, FulfilledBy: "Alice"})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res3.Order.Status != rpc.OrderStatus_APPROVED {
		t.Fatalf("order must be approved")
	}
	if res3.Order.CustomerID == "" {
		t.Fatalf("order must have customer id")
	}

	ces1, err := csvc.GetCustomer(ctx, &rpc.GetCustomerRequest{ID: res3.Order.CustomerID})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ces1.Customer == nil {
		t.Fatalf("customer must exist")
	}

	_, err = osvc.RejectOrder(ctx, &rpc.OrderRejectRequest{ID: res.Order.ID, FulfilledBy: "Alice"})
	if err.Error() != "already fulfilled" {
		t.Fatalf("%+v", err)
	}

	res4, err := osvc.GetOrder(ctx, &rpc.GetOrderRequest{ID: res.Order.ID})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res4.Order.Status != rpc.OrderStatus_APPROVED {
		t.Fatalf("order must be approved")
	}
}
