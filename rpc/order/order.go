package order

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/altairsix/eventsource"
	"github.com/slaskis/es-orders/rpc/events"
)

func (order *Order) On(event eventsource.Event) error {
	// log.Printf("%T(%+v)", event, event)
	switch v := event.(type) {
	case *events.OrderCreated:
		if order.Id != "" {
			return errors.New("order already exists")
		}
		order.Id = v.AggregateID()
		order.Status = OrderStatus_EMPTY
		order.CreatedAt = v.EventAt().Format(time.RFC3339)
		order.UpdatedAt = v.EventAt().Format(time.RFC3339)

	case *events.OrderItemAAdded:
		item := &OrderItem{
			Quantity:  1,
			UpdatedAt: v.EventAt().Format(time.RFC3339),
			Item:      &Item{Type: ItemType_ITEM_A, Id: v.ItemA},
		}
		order.Status = OrderStatus_PENDING
		order.Items = append(order.Items, item)
		order.UpdatedAt = v.EventAt().Format(time.RFC3339)

	case *events.OrderItemBAdded:
		item := &OrderItem{
			Quantity:  1,
			UpdatedAt: v.EventAt().Format(time.RFC3339),
			Item:      &Item{Type: ItemType_ITEM_B, Id: v.ItemB},
		}
		order.Status = OrderStatus_PENDING
		order.Items = append(order.Items, item)
		order.UpdatedAt = v.EventAt().Format(time.RFC3339)

	case *events.OrderItemRemoved:
		items := order.Items[:0]
		for _, i := range order.Items {
			if i.Id != v.ItemId {
				items = append(items, i)
			}
		}
		if len(items) == 0 {
			order.Status = OrderStatus_EMPTY
		}
		order.Items = items
		order.UpdatedAt = v.EventAt().Format(time.RFC3339)

	case *events.OrderFulfilled:
		order.FulfilledBy = v.By
		order.FulfilledAt = v.EventAt().Format(time.RFC3339)
		order.UpdatedAt = v.EventAt().Format(time.RFC3339)
		if v.Approved {
			order.Status = OrderStatus_APPROVED
		} else {
			order.Status = OrderStatus_REJECTED
		}

	case *events.OrderAssignCustomer:
		order.CustomerId = v.CustomerId
		order.UpdatedAt = v.EventAt().Format(time.RFC3339)

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}

type CommandCreateOrder struct {
	eventsource.CommandModel
	Name string
}

type CommandAddItem struct {
	eventsource.CommandModel
	Type ItemType
}

type CommandRemoveItem struct {
	eventsource.CommandModel
	ItemID string
}

type CommandFulfillOrder struct {
	eventsource.CommandModel
	Approved bool
	By       string
}

type CommandAssignCustomer struct {
	eventsource.CommandModel
	CustomerID string
}

func (order *Order) Apply(ctx context.Context, command eventsource.Command) ([]eventsource.Event, error) {
	builder := events.NewBuilder(command.AggregateID(), int(order.Version))
	switch cmd := command.(type) {
	case *CommandCreateOrder:
		builder.OrderCreated()
	case *CommandAddItem:
		switch cmd.Type {
		case ItemType_ITEM_A:
			builder.OrderItemAAdded("1")
		case ItemType_ITEM_B:
			builder.OrderItemBAdded("2")
		}
	case *CommandRemoveItem:
		builder.OrderItemRemoved(cmd.ItemID)
	case *CommandFulfillOrder:
		if order.FulfilledAt != "" {
			return builder.Events, errors.New("already fulfilled")
		}
		builder.OrderFulfilled(cmd.By, cmd.Approved)
	case *CommandAssignCustomer:
		builder.OrderAssignCustomer(cmd.CustomerID)
	default:
		log.Printf("unknown command: %T", cmd)
	}
	return builder.Events, nil
}
