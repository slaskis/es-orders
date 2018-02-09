package userserver

import (
	"context"

	"github.com/slaskis/es-orders/rpc/user"
)

func (s server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
	agg, err := s.users.Load(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{User: agg.(*user.User)}, nil
}
