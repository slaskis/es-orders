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
	orderID, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		return nil, err
	}

	_, err = s.orders.Apply(ctx, &rpc.CreateOrder{
		CommandModel: eventsource.CommandModel{ID: orderID.String()},
		Name:         req.Name,
	})
	if err != nil {
		return nil, err
	}

	for _, i := range req.Items {
		itemID, err := ulid.New(ulid.Now(), entropy)
		if err != nil {
			return nil, err
		}
		_, err = s.items.Apply(ctx, &rpc.CreateItem{
			CommandModel: eventsource.CommandModel{ID: itemID.String()},
			SKU:          i.SKU,
			Description:  i.Description,
			OrderID:      orderID.String(),
		})
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &rpc.AddItem{
			CommandModel: eventsource.CommandModel{ID: orderID.String()},
			ItemID:       itemID.String(),
		})
		if err != nil {
			return nil, err
		}
	}

	return s.GetOrder(ctx, &rpc.IDRequest{ID: orderID.String()})
}
func (s service) SignOrder(ctx context.Context, req *rpc.OrderSignRequest) (*rpc.OrderResponse, error) {
	// TODO sign order...
	return s.GetOrder(ctx, &rpc.IDRequest{ID: req.ID})
}
func (s service) FulfillOrder(ctx context.Context, req *rpc.OrderFulfillRequest) (*rpc.OrderResponse, error) {
	return nil, nil
}
func (s service) GetOrder(ctx context.Context, req *rpc.IDRequest) (*rpc.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &rpc.OrderResponse{Order: agg.(*rpc.Order)}, nil
}

func (s service) ListItemsOfOrder(ctx context.Context, req *rpc.IDRequest) (*rpc.ItemsResponse, error) {
	var items []*rpc.Item
	res, err := s.GetOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, i := range res.Order.ItemIDs {
		agg, err := s.items.Load(ctx, i)
		if err != nil {
			return nil, err
		}
		items = append(items, agg.(*rpc.Item))
	}
	return &rpc.ItemsResponse{
		Items: items,
	}, nil
}
