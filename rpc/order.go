package rpc

import (
	fmt "fmt"
	"log"

	context "golang.org/x/net/context"

	"github.com/altairsix/eventsource"
)

func (order *Order) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *OrderCreated:
		order.ID = v.AggregateID()
		order.Name = v.Name
		order.CreatedAt = v.EventAt()
		order.UpdatedAt = v.EventAt()
		log.Printf("order created: %+v", order)

	case *OrderRenamed:
		order.Name = v.Name
		order.UpdatedAt = v.EventAt()

	case *OrderItemAdded:
		order.ItemIDs = append(order.ItemIDs, v.Item)
		order.UpdatedAt = v.EventAt()

	case *OrderSigned:
		order.SignedBy = v.By
		*order.SignedAt = v.EventAt()
		order.UpdatedAt = v.EventAt()

	case *OrderFulfilled:
		order.FulfilledBy = v.By
		*order.FulfilledAt = v.EventAt()
		order.UpdatedAt = v.EventAt()

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}

type CreateOrder struct {
	eventsource.CommandModel
	Name string
}

type AddItem struct {
	eventsource.CommandModel
	ItemID string
}

func (order *Order) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := NewBuilder(command.AggregateID(), int(order.Version))
	switch cmd := command.(type) {
	case *CreateOrder:
		builder.OrderCreated(cmd.Name)
	case *AddItem:
		builder.OrderItemAdded(cmd.ItemID, command.AggregateID())
	default:
		log.Printf("unknown command: %+v", cmd)
	}
	return builder.Events, nil
}
