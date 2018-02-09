package customerserver

import (
	"context"

	"github.com/slaskis/es-orders/rpc/customer"
)

func (s server) GetCustomer(ctx context.Context, req *customer.GetCustomerRequest) (*customer.CustomerResponse, error) {
	agg, err := s.customers.Load(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &customer.CustomerResponse{Customer: agg.(*customer.Customer)}, nil
}
