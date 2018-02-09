package orderserver

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/oklog/ulid"
	"github.com/slaskis/es-orders/rpc/customer"
	"github.com/slaskis/es-orders/rpc/order"
)

func (s server) ApproveOrder(ctx context.Context, req *order.OrderApproveRequest) (*order.OrderResponse, error) {
	agg, err := s.orders.Load(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.orders.Apply(ctx, &order.CommandFulfillOrder{
		CommandModel: eventsource.CommandModel{ID: req.Id},
		By:           req.FulfilledBy,
		Approved:     true,
	})
	if err != nil {
		return nil, err
	}

	ord := agg.(*order.Order)
	if ord.CustomerId == "" {
		customerID, err := ulid.New(ulid.Now(), s.entropy)
		if err != nil {
			return nil, err
		}
		_, err = s.customers.Apply(ctx, &customer.CommandCreateCustomer{
			CommandModel: eventsource.CommandModel{ID: customerID.String()},
			Name:         "new customer",
		})
		if err != nil {
			return nil, err
		}
		_, err = s.orders.Apply(ctx, &order.CommandAssignCustomer{
			CommandModel: eventsource.CommandModel{ID: ord.Id},
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

	return s.GetOrder(ctx, &order.GetOrderRequest{Id: req.Id})
}
