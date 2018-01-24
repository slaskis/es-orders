package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc"
)

func TestService(t *testing.T) {
	obs := func(event eventsource.Event) {
		t.Logf("observer event: %s", event)
	}

	svc := service{
		orders: eventsource.New(&rpc.Order{},
			eventsource.WithSerializer(rpc.NewSerializer()),
			eventsource.WithDebug(os.Stdout),
			eventsource.WithObservers(obs),
		),
	}

	ctx := context.Background()
	res, err := svc.CreateOrder(ctx, &rpc.OrderNewRequest{
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
	if res.Order.CreatedAt.IsZero() {
		t.Fatalf("order must have created at date. was: %s", res.Order.CreatedAt.Format(time.RFC3339))
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
	res2, err := svc.RemoveItem(ctx, &rpc.OrderItemRemoveRequest{OrderID: res.Order.ID, ItemID: item.ID})
	if len(res2.Order.Items) != 1 {
		t.Fatalf("order must have 1 items")
	}
}
