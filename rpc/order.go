package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/altairsix/eventsource"
)

func (order *Order) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *OrderCreated:
		order.ID = v.AggregateID()
		order.Status = OrderStatus_EMPTY
		order.CreatedAt = v.EventAt()
		order.UpdatedAt = v.EventAt()

	case *OrderItemAdded:
		item := &Item{
			ID:          v.ItemId,
			SKU:         v.Sku,
			Description: v.Description,
			CreatedAt:   v.EventAt(),
			UpdatedAt:   v.EventAt(),
		}
		order.Status = OrderStatus_PENDING
		order.Items = append(order.Items, item)
		order.UpdatedAt = v.EventAt()

	case *OrderItemRemoved:
		items := order.Items[:0]
		for _, i := range order.Items {
			if i.ID != v.ItemId {
				items = append(items, i)
			}
		}
		if len(items) == 0 {
			order.Status = OrderStatus_EMPTY
		}
		order.Items = items
		order.UpdatedAt = v.EventAt()

	case *OrderFulfilled:
		order.FulfilledBy = v.By
		*order.FulfilledAt = v.EventAt()
		order.UpdatedAt = v.EventAt()
		if v.Approved {
			order.Status = OrderStatus_APPROVED
		} else {
			order.Status = OrderStatus_REJECTED
		}

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
	Item *Item
}

type RemoveItem struct {
	eventsource.CommandModel
	ItemID string
}

type FulfillOrder struct {
	eventsource.CommandModel
	Approved bool
	By       string
}

func (order *Order) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := NewBuilder(command.AggregateID(), int(order.Version))
	switch cmd := command.(type) {
	case *CreateOrder:
		builder.OrderCreated()
	case *AddItem:
		builder.OrderItemAdded(cmd.Item.ID, cmd.Item.SKU, cmd.Item.Description)
	case *RemoveItem:
		builder.OrderItemRemoved(cmd.ItemID)
	case *FulfillOrder:
		builder.OrderFulfilled(cmd.By, cmd.Approved)
	default:
		log.Printf("unknown command: %T", cmd)
	}
	return builder.Events, nil
}
