package orderserver

import (
	"context"

	"github.com/slaskis/es-orders/rpc/order"
)

func (s server) GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &order.OrderResponse{Order: agg.(*order.Order)}, nil
}
