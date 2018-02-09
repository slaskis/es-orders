package customerserver

import (
	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/customer"
)

type server struct {
	customers *eventsource.Repository
}

func NewServer(customers *eventsource.Repository) customer.CustomerService {
	return server{customers: customers}
}
