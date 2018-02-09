package orderserver

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/order"
)

func (s server) RemoveItem(ctx context.Context, req *order.OrderItemRemoveRequest) (*order.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &order.CommandRemoveItem{
		CommandModel: eventsource.CommandModel{ID: req.OrderId},
		ItemID:       req.OrderItemId,
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &order.GetOrderRequest{Id: req.OrderId})
}
