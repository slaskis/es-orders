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
		items: eventsource.New(&rpc.Item{},
			eventsource.WithSerializer(rpc.NewSerializer()),
			eventsource.WithDebug(os.Stdout),
			eventsource.WithObservers(obs),
		),
	}

	ctx := context.Background()
	res, err := svc.CreateOrder(ctx, &rpc.OrderNewRequest{
		Name: "First!",
		Items: []*rpc.OrderNewRequest_NewItem{
			&rpc.OrderNewRequest_NewItem{SKU: "1", Description: "One"},
			&rpc.OrderNewRequest_NewItem{SKU: "2", Description: "Two"},
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
	if res.Order.Name != "First!" {
		t.Fatalf("order must have name")
	}
	if len(res.Order.ItemIDs) != 2 {
		t.Fatalf("order must have 2 items")
	}

	ires, err := svc.ListItemsOfOrder(ctx, &rpc.IDRequest{ID: res.Order.ID})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if len(ires.Items) != 2 {
		t.Fatalf("order must have 2 items")
	}
	for idx, i := range ires.Items {
		if i.CreatedAt.IsZero() {
			t.Fatalf("item must have created at date. was: %s", res.Order.CreatedAt.Format(time.RFC3339))
		}
		if i.UpdatedAt.IsZero() {
			t.Fatalf("item must have updated at date")
		}
		if i.DeletedAt != nil {
			t.Fatalf("item must not have deleted at date")
		}
		if idx == 0 && i.SKU != "1" {
			t.Fatalf("item 0 expected to have sku 1")
		} else if idx == 1 && i.SKU != "2" {
			t.Fatalf("item 1 expected to have sku 2")
		}
	}
}
