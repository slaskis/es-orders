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
	users     *eventsource.Repository
}

func NewOrderService(orders, customers, users *eventsource.Repository) rpc.OrderService {
	return orderService{orders: orders, customers: customers, users: users}
}

func (s orderService) CreateOrder(ctx context.Context, req *rpc.OrderNewRequest) (*rpc.OrderResponse, error) {
	orderID, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		return nil, err
	}

	_, err = s.orders.Apply(ctx, &rpc.CommandCreateOrder{
		CommandModel: eventsource.CommandModel{ID: orderID.String()},
	})
	if err != nil {
		return nil, err
	}

	for _, i := range req.Items {
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &rpc.CommandAddItem{
			CommandModel: eventsource.CommandModel{ID: orderID.String()},
			Type:         i.Type,
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
	_, err = s.orders.Apply(ctx, &rpc.CommandFulfillOrder{
		CommandModel: eventsource.CommandModel{ID: req.ID},
		By:           req.FulfilledBy,
		Approved:     true,
	})
	if err != nil {
		return nil, err
	}

	// create customer (if not exists)
	order := agg.(*rpc.Order)
	if order.CustomerID == "" {
		customerID, err := ulid.New(ulid.Now(), entropy)
		if err != nil {
			return nil, err
		}
		_, err = s.customers.Apply(ctx, &rpc.CommandCreateCustomer{
			CommandModel: eventsource.CommandModel{ID: customerID.String()},
			Name:         "new customer",
		})
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &rpc.CommandAssignCustomer{
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
	_, err = s.orders.Apply(ctx, &rpc.CommandFulfillOrder{
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
	_, err = s.orders.Apply(ctx, &rpc.CommandAddItem{
		CommandModel: eventsource.CommandModel{ID: req.OrderID},
		Type:         req.Item.Type,
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
	_, err = s.orders.Apply(ctx, &rpc.CommandRemoveItem{
		CommandModel: eventsource.CommandModel{ID: req.OrderID},
		ItemID:       req.OrderItemID,
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
