package orderserver

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/order"
)

func (s server) AddItem(ctx context.Context, req *order.OrderItemAddRequest) (*order.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &order.CommandAddItem{
		CommandModel: eventsource.CommandModel{ID: req.OrderId},
		Type:         req.Item.Type,
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &order.GetOrderRequest{Id: req.OrderId})
}
