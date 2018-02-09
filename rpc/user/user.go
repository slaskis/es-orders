package user

import (
	"context"
	"log"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/events"
)

func (user *User) On(event eventsource.Event) error {
	switch v := event.(type) {
	default:
		log.Printf("unknown event %T(%+v)", v, v)
	}

	return nil
}

func (user *User) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := events.NewBuilder(command.AggregateID(), int(user.Version))
	switch cmd := command.(type) {
	default:
		log.Printf("unknown command: %T(%+v)", cmd, cmd)
	}
	return builder.Events, nil
}
