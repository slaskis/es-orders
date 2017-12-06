package rpc

import (
	fmt "fmt"

	"github.com/altairsix/eventsource"
)

func (item Item) On(event eventsource.Event) error {
	switch v := event.(type) {
	case *ItemNew:
		item.ID = v.AggregateID()
		item.SKU = v.Sku
		item.Description = v.Description
		item.CreatedAt = v.EventAt()
		item.UpdatedAt = v.EventAt()

	default:
		return fmt.Errorf("unable to handle event, %v", v)
	}

	return nil
}
