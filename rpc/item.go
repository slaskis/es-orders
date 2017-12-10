package rpc

import (
	fmt "fmt"

	context "golang.org/x/net/context"

	"github.com/altairsix/eventsource"
)

func (item *Item) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *ItemCreated:
		item.ID = v.AggregateID()
		item.SKU = v.Sku
		item.Description = v.Description
		item.CreatedAt = v.EventAt()
		item.UpdatedAt = v.EventAt()

	case *OrderItemAdded:
		item.OrderID = v.Id

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}

type CreateItem struct {
	eventsource.CommandModel
	SKU         string
	Description string
	OrderID     string
}

func (item *Item) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := NewBuilder(command.AggregateID(), int(item.Version))
	switch cmd := command.(type) {
	case *CreateItem:
		builder.ItemCreated(cmd.SKU, cmd.Description)
		builder.OrderItemAdded(item.ID, cmd.OrderID)
	}
	return builder.Events, nil
}
