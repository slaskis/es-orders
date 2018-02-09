package userserver

import (
	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/user"
)

type server struct {
	users *eventsource.Repository
}

func NewServer(users *eventsource.Repository) user.UserService {
	return server{users: users}
}
