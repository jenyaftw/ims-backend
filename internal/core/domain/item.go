package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/util"
)

type Item struct {
	ID          uuid.UUID
	InventoryID uuid.UUID
	SectionID   uuid.UUID
	Name        string
	Description string
	Quantity    int
	SKU         string

	util.Timestamps
}

func (u Item) Validate() error {
	if len(u.Name) < 2 || len(u.Name) > 32 {
		return errors.New("name length must be between 2 and 32 characters")
	}

	return nil
}
