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
}

var entropy = rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano()))

func (s service) CreateOrder(ctx context.Context, req *rpc.OrderNewRequest) (*rpc.OrderResponse, error) {
	orderID, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		return nil, err
	}

	_, err = s.orders.Apply(ctx, &rpc.CreateOrder{
		CommandModel: eventsource.CommandModel{ID: orderID.String()},
	})
	if err != nil {
		return nil, err
	}

	for _, i := range req.Items {
		itemID, err := ulid.New(ulid.Now(), entropy)
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &rpc.AddItem{
			CommandModel: eventsource.CommandModel{ID: orderID.String()},
			Item: &rpc.Item{
				ID:          itemID.String(),
				SKU:         i.SKU,
				Description: i.Description,
			},
		})
		if err != nil {
			return nil, err
		}
	}

	return s.GetOrder(ctx, &rpc.IDRequest{ID: orderID.String()})
}
func (s service) ApproveOrder(ctx context.Context, req *rpc.OrderApproveRequest) (*rpc.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &rpc.FulfillOrder{
		CommandModel: eventsource.CommandModel{ID: req.ID},
		By:           req.FulfilledBy,
		Approved:     true,
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &rpc.IDRequest{ID: req.ID})
}
func (s service) RejectOrder(ctx context.Context, req *rpc.OrderRejectRequest) (*rpc.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &rpc.FulfillOrder{
		CommandModel: eventsource.CommandModel{ID: req.ID},
		By:           req.FulfilledBy,
		Approved:     false,
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &rpc.IDRequest{ID: req.ID})
}
func (s service) AddItem(ctx context.Context, req *rpc.OrderItemAddRequest) (*rpc.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}
	itemID, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &rpc.AddItem{
		CommandModel: eventsource.CommandModel{ID: req.OrderID},
		Item: &rpc.Item{
			ID:          itemID.String(),
			SKU:         req.Item.SKU,
			Description: req.Item.Description,
		},
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &rpc.IDRequest{ID: req.OrderID})
}
func (s service) RemoveItem(ctx context.Context, req *rpc.OrderItemRemoveRequest) (*rpc.OrderResponse, error) {
	_, err := s.orders.Load(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &rpc.RemoveItem{
		CommandModel: eventsource.CommandModel{ID: req.OrderID},
		ItemID:       req.ItemID,
	})
	if err != nil {
		return nil, err
	}
	return s.GetOrder(ctx, &rpc.IDRequest{ID: req.OrderID})
}
func (s service) GetOrder(ctx context.Context, req *rpc.IDRequest) (*rpc.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &rpc.OrderResponse{Order: agg.(*rpc.Order)}, nil
}
