package rpc

import (
	fmt "fmt"

	"github.com/altairsix/eventsource"
)

func (order Order) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *OrderNew:
		order.ID = v.AggregateID()
		order.Name = v.Name
		order.CreatedAt = v.EventAt()
		order.UpdatedAt = v.EventAt()

	case *OrderRename:
		order.Name = v.Name
		order.UpdatedAt = v.EventAt()

	case *OrderAddItem:
		// order.Items = append(order.Items, v.Item)
		order.UpdatedAt = v.EventAt()

	case *OrderSign:
		order.SignedBy = v.By
		*order.SignedAt = v.EventAt()
		order.UpdatedAt = v.EventAt()

	case *OrderFulfill:
		order.FulfilledBy = v.By
		*order.FulfilledAt = v.EventAt()
		order.UpdatedAt = v.EventAt()

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}
