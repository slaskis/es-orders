package main

import (
	"context"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc"
)

type userService struct {
	users *eventsource.Repository
}

func NewUserService(users *eventsource.Repository) rpc.UserService {
	return userService{users: users}
}

func (s userService) GetUser(ctx context.Context, req *rpc.GetUserRequest) (*rpc.UserResponse, error) {
	agg, err := s.users.Load(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	user := agg.(*rpc.User)

	return &rpc.UserResponse{User: user}, nil
}
