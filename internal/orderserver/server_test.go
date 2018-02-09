package orderserver

import (
	"context"
	"math/rand"
	"os"
	"testing"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/customer"
	"github.com/slaskis/es-orders/rpc/events"
	"github.com/slaskis/es-orders/rpc/order"
	"github.com/slaskis/es-orders/rpc/user"
)

func TestOrderServer(t *testing.T) {
	obs := func(event eventsource.Event) {
		t.Logf("observer event: %s", event)
	}

	orders := eventsource.New(&order.Order{},
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stdout),
		eventsource.WithObservers(obs),
	)

	users := eventsource.New(&user.User{},
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stdout),
		eventsource.WithObservers(obs),
	)

	customers := eventsource.New(&customer.Customer{},
		eventsource.WithSerializer(events.NewSerializer()),
		eventsource.WithDebug(os.Stdout),
		eventsource.WithObservers(obs),
	)

	osrv := NewServer(orders, customers, users, rand.New(rand.NewSource(1)))

	ctx := context.Background()
	res, err := osrv.CreateOrder(ctx, &order.OrderNewRequest{
		Items: []*order.NewItem{
			&order.NewItem{Type: order.ItemType_ITEM_B},
		},
	})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Logf("order: %+v", res.Order)
	if res.Order == nil {
		t.Fatalf("order must not be nil")
	}
	if res.Order.Status != order.OrderStatus_PENDING {
		t.Fatalf("order must be pending. was %s", res.Order.Status)
	}
	if res.Order.CreatedAt == "" {
		t.Fatalf("order must have created at date")
	}
	if res.Order.UpdatedAt == "" {
		t.Fatalf("order must have updated at date")
	}
	if res.Order.DeletedAt != "" {
		t.Fatalf("order must not have deleted at date")
	}
	if len(res.Order.Items) != 1 {
		t.Fatalf("order must have 1 items. had %d", len(res.Order.Items))
	}

	item := res.Order.Items[0]
	res2, err := osrv.RemoveItem(ctx, &order.OrderItemRemoveRequest{OrderId: res.Order.Id, OrderItemId: item.Id})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res2.Order.Status != order.OrderStatus_EMPTY {
		t.Fatalf("order must be empty. was %s", res2.Order.Status)
	}
	if len(res2.Order.Items) != 0 {
		t.Fatalf("order must have 0 items")
	}
	if res2.Order.CustomerId != "" {
		t.Fatalf("order must not have customer id")
	}

	res, err = osrv.AddItem(ctx, &order.OrderItemAddRequest{
		OrderId: res.Order.Id,
		Item: &order.NewItem{
			Type: order.ItemType_ITEM_A,
		},
	})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res.Order == nil {
		t.Fatalf("order must not be nil")
	}
	if res.Order.Status != order.OrderStatus_PENDING {
		t.Fatalf("order must be pending. was %s (%d items)", res.Order.Status, len(res.Order.Items))
	}
	if len(res.Order.Items) != 1 {
		t.Fatalf("order must have 1 items. had %d", len(res.Order.Items))
	}

	res3, err := osrv.ApproveOrder(ctx, &order.OrderApproveRequest{Id: res.Order.Id, FulfilledBy: "Alice"})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res3.Order.Status != order.OrderStatus_APPROVED {
		t.Fatalf("order must be approved")
	}
	if res3.Order.CustomerId == "" {
		t.Fatalf("order must have customer id")
	}

	_, err = osrv.RejectOrder(ctx, &order.OrderRejectRequest{Id: res.Order.Id, FulfilledBy: "Alice"})
	if err.Error() != "already fulfilled" {
		t.Fatalf("%+v", err)
	}

	res4, err := osrv.GetOrder(ctx, &order.GetOrderRequest{Id: res.Order.Id})
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if res4.Order.Status != order.OrderStatus_APPROVED {
		t.Fatalf("order must be approved")
	}
}
