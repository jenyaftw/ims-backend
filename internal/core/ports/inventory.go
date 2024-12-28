package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type InventoryRepository interface {
	CreateInventory(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error)
	UpdateInventory(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error)
	GetInventoryById(ctx context.Context, id uuid.UUID) (domain.Inventory, error)
	GetInventoryByName(ctx context.Context, name string) (domain.Inventory, error)
	ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error)
	DeleteInventory(ctx context.Context, id uuid.UUID) error
}

type InventoryService interface {
	ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error)
}
