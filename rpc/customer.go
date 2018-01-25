package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/altairsix/eventsource"
)

func (customer *Customer) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *CustomerCreated:
		customer.ID = v.AggregateID()
		customer.CreatedAt = v.EventAt()
		customer.UpdatedAt = v.EventAt()

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}

type CreateCustomer struct {
	eventsource.CommandModel
	Name string
}

func (customer *Customer) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := NewBuilder(command.AggregateID(), int(customer.Version))
	switch cmd := command.(type) {
	case *CreateCustomer:
		builder.CustomerCreated(cmd.Name)
	default:
		log.Printf("unknown command: %T", cmd)
	}
	return builder.Events, nil
}
