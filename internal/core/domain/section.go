package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/util"
)

type Section struct {
	ID          uuid.UUID
	InventoryID uuid.UUID
	Name        string
	Description string

	util.Timestamps
}

func (u Section) Validate() error {
	if len(u.Name) < 2 || len(u.Name) > 32 {
		return errors.New("name length must be between 2 and 32 characters")
	}

	return nil
}
