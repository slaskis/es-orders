package orderserver

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/order"
)

func (s server) RejectOrder(ctx context.Context, req *order.OrderRejectRequest) (*order.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &order.CommandFulfillOrder{
		CommandModel: eventsource.CommandModel{ID: req.Id},
		By:           req.FulfilledBy,
		Approved:     false,
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &order.GetOrderRequest{Id: req.Id})
}
