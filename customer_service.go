package main

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc"
)

type customerService struct {
	customers *eventsource.Repository
}

func NewCustomerService(customers *eventsource.Repository) rpc.CustomerService {
	return customerService{customers: customers}
}

func (s customerService) GetCustomer(ctx context.Context, req *rpc.GetCustomerRequest) (*rpc.CustomerResponse, error) {
	agg, err := s.customers.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	customer := agg.(*rpc.Customer)

	return &rpc.CustomerResponse{Customer: customer}, nil
}
