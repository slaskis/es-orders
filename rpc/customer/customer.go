package customer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/events"
)

func (customer *Customer) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *events.CustomerCreated:
		if customer.Id != "" {
			return errors.New("customer already exist")
		}
		customer.Id = v.AggregateID()
		customer.CreatedAt = v.EventAt().Format(time.RFC3339)
		customer.UpdatedAt = v.EventAt().Format(time.RFC3339)

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}

type CommandCreateCustomer struct {
	eventsource.CommandModel
	Name string
}

func (customer *Customer) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := events.NewBuilder(command.AggregateID(), int(customer.Version))
	switch cmd := command.(type) {
	case *CommandCreateCustomer:
		builder.CustomerCreated(cmd.Name)
	default:
		log.Printf("unknown command: %T", cmd)
	}
	return builder.Events, nil
}
