package domain

import (
	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/util"
)

type TransferRequest struct {
	ID       uuid.UUID
	Item     Item
	Quantity uint64
	Status   string

	FromInventoryID uuid.UUID
	ToInventoryID   uuid.UUID
	FromSectionID   uuid.UUID
	ToSectionID     uuid.UUID

	util.Timestamps
}
