package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/altairsix/eventsource"
	"github.com/oklog/ulid"
	"github.com/slaskis/es-orders/rpc"
)

type service struct {
	orders *eventsource.Repository
	items  *eventsource.Repository
}

var entropy = rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano()))

func (s service) CreateOrder(ctx context.Context, req *rpc.OrderNewRequest) (*rpc.OrderResponse, error) {
	id, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		return nil, err
	}
	b := rpc.NewBuilder(id.String(), 0)
	b.OrderNew(req.Name)
	err = s.orders.Save(ctx, b.Events...)
	if err != nil {
		return nil, err
	}
	agg, err := s.orders.Load(ctx, id.String())
	if err != nil {
		return nil, err
	}
	order := agg.(rpc.Order)
	return &rpc.OrderResponse{Order: &order}, nil
}
func (s service) SignOrder(ctx context.Context, req *rpc.OrderSignRequest) (*rpc.OrderResponse, error) {
	return nil, nil
}
func (s service) FulfillOrder(ctx context.Context, req *rpc.OrderFulfillRequest) (*rpc.OrderResponse, error) {
	return nil, nil
}
func (s service) GetOrder(ctx context.Context, req *rpc.IDRequest) (*rpc.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	order := agg.(rpc.Order)
	return &rpc.OrderResponse{Order: &order}, nil
}
