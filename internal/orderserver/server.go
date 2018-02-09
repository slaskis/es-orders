package orderserver

import (
	"math/rand"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/order"
)

type server struct {
	orders    *eventsource.Repository
	customers *eventsource.Repository
	users     *eventsource.Repository
	entropy   *rand.Rand
}

func NewServer(orders, customers, users *eventsource.Repository, entropy *rand.Rand) order.OrderService {
	return server{
		orders:    orders,
		customers: customers,
		users:     users,
		entropy:   entropy,
	}
}
