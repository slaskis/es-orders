package orderserver

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/oklog/ulid"
	"github.com/slaskis/es-orders/rpc/order"
)

func (s server) CreateOrder(ctx context.Context, req *order.OrderNewRequest) (*order.OrderResponse, error) {
	orderID, err := ulid.New(ulid.Now(), s.entropy)
	if err != nil {
		return nil, err
	}

	_, err = s.orders.Apply(ctx, &order.CommandCreateOrder{
		CommandModel: eventsource.CommandModel{ID: orderID.String()},
	})
	if err != nil {
		return nil, err
	}

	for _, i := range req.Items {
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &order.CommandAddItem{
			CommandModel: eventsource.CommandModel{ID: orderID.String()},
			Type:         i.Type,
		})
		if err != nil {
			return nil, err
		}
	}

	return s.GetOrder(ctx, &order.GetOrderRequest{Id: orderID.String()})
}
