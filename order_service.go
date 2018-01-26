package main

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/oklog/ulid"
	"github.com/slaskis/es-orders/rpc"
)

type orderService struct {
	orders    *eventsource.Repository
	customers *eventsource.Repository
}

func NewOrderService(orders *eventsource.Repository, customers *eventsource.Repository) rpc.OrderService {
	return orderService{orders: orders, customers: customers}
}

func (s orderService) CreateOrder(ctx context.Context, req *rpc.OrderNewRequest) (*rpc.OrderResponse, error) {
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

	return s.GetOrder(ctx, &rpc.GetOrderRequest{ID: orderID.String()})
}
func (s orderService) ApproveOrder(ctx context.Context, req *rpc.OrderApproveRequest) (*rpc.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.ID)
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

	// create customer (if not exists)
	order := agg.(*rpc.Order)
	if order.CustomerID != "" {
		customerID, err := ulid.New(ulid.Now(), entropy)
		if err != nil {
			return nil, err
		}
		_, err = s.customers.Apply(ctx, &rpc.CreateCustomer{
			CommandModel: eventsource.CommandModel{ID: customerID.String()},
			Name:         "new customer",
		})
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &rpc.AssignCustomer{
			CommandModel: eventsource.CommandModel{ID: order.ID},
			CustomerID:   customerID.String(),
		})
		if err != nil {
			return nil, err
		}
	}

	// TODO add oneOf to Item for different kinds of items
	// then switch() on them here and generate s.customer.Apply()
	// for _, item := range order.Items {
	// }

	return s.GetOrder(ctx, &rpc.GetOrderRequest{ID: req.ID})
}
func (s orderService) RejectOrder(ctx context.Context, req *rpc.OrderRejectRequest) (*rpc.OrderResponse, error) {
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
	return s.GetOrder(ctx, &rpc.GetOrderRequest{ID: req.ID})
}
func (s orderService) AddItem(ctx context.Context, req *rpc.OrderItemAddRequest) (*rpc.OrderResponse, error) {
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
	return s.GetOrder(ctx, &rpc.GetOrderRequest{ID: req.OrderID})
}
func (s orderService) RemoveItem(ctx context.Context, req *rpc.OrderItemRemoveRequest) (*rpc.OrderResponse, error) {
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
	return s.GetOrder(ctx, &rpc.GetOrderRequest{ID: req.OrderID})
}
func (s orderService) GetOrder(ctx context.Context, req *rpc.GetOrderRequest) (*rpc.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &rpc.OrderResponse{Order: agg.(*rpc.Order)}, nil
}
