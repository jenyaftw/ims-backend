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

	GetInventorySections(ctx context.Context, inventoryID uuid.UUID) ([]domain.Section, error)
	CreateInventorySection(ctx context.Context, section domain.Section) (domain.Section, error)
	DeleteInventorySection(ctx context.Context, id uuid.UUID) error
	UpdateInventorySection(ctx context.Context, section domain.Section) (domain.Section, error)

	CreateInventoryItem(ctx context.Context, item domain.Item) (domain.Item, error)
	GetInventoryItems(ctx context.Context, inventoryId uuid.UUID, sectionID *uuid.UUID, offset, limit uint64) ([]domain.Item, error)
}

type InventoryService interface {
	ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error)
	GetInventoryById(ctx context.Context, id uuid.UUID) (domain.Inventory, error)
	CreateInventory(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error)
	CreateInventorySection(ctx context.Context, inventory domain.Inventory, section domain.Section) (domain.Section, error)
	GetInventoryItems(ctx context.Context, inventoryId uuid.UUID, offset, limit uint64) ([]domain.Item, error)
	GetInventoryItemsBySection(ctx context.Context, inventoryId, sectionID uuid.UUID, offset, limit uint64) ([]domain.Item, error)
	CreateInventoryItem(ctx context.Context, inventory domain.Inventory, section domain.Section, item domain.Item) (domain.Item, error)
}
