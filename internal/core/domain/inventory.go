package domain

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/util"
)

type Inventory struct {
	ID          uuid.UUID
	Name        string
	Description string

	Sections []Section

	util.Timestamps
}

func (u Inventory) Validate() error {
	if len(u.Name) < 2 || len(u.Name) > 32 {
		return errors.New("name length must be between 2 and 32 characters")
	}

	return nil
}
