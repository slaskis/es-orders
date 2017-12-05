package main

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc"
)

type service struct {
	repo    *eventsource.Repository
	builder *rpc.Builder
}

func (s service) CreateOrder(context.Context, *rpc.OrderNewRequest) (*rpc.OrderResponse, error) {
	s.repo.Save(ctx)
	return nil, nil
}
func (s service) SignOrder(context.Context, *rpc.OrderSignRequest) (*rpc.OrderResponse, error) {
	return nil, nil
}
func (s service) FulfillOrder(context.Context, *rpc.OrderFulfillRequest) (*rpc.OrderResponse, error) {
	return nil, nil
}
func (s service) GetOrder(context.Context, *rpc.IDRequest) (*rpc.OrderResponse, error) {
	return nil, nil
}
