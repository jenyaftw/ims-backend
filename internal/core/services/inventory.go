package services

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type InventoryService struct {
	repo ports.InventoryRepository
}

func NewInventoryService(
	repo ports.InventoryRepository,
) InventoryService {
	return InventoryService{
		repo: repo,
	}
}

func (s InventoryService) ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error) {
	inventories, err := s.repo.ListInventories(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	newInventories := make([]domain.Inventory, 0, len(inventories))
	for _, inventory := range inventories {
		sections, err := s.repo.GetInventorySections(ctx, inventory.ID)
		if err != nil {
			return nil, err
		}

		inventory.Sections = sections
		newInventories = append(newInventories, inventory)
	}

	return newInventories, nil
}

func (s InventoryService) CreateInventory(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	inventory.ID = uuid.New()
	inventory.InitTimestamps()

	return s.repo.CreateInventory(ctx, inventory)
}

func (s InventoryService) CreateInventorySection(ctx context.Context, inventory domain.Inventory, section domain.Section) (domain.Section, error) {
	section.ID = uuid.New()
	section.InventoryID = inventory.ID
	section.InitTimestamps()

	return s.repo.CreateInventorySection(ctx, section)
}

func (s InventoryService) GetInventoryById(ctx context.Context, id uuid.UUID) (domain.Inventory, error) {
	inventory, err := s.repo.GetInventoryById(ctx, id)
	if err != nil {
		return domain.Inventory{}, err
	}

	sections, err := s.repo.GetInventorySections(ctx, inventory.ID)
	if err != nil {
		return domain.Inventory{}, err
	}

	inventory.Sections = sections

	return inventory, nil
}

func (s InventoryService) GetInventoryItems(ctx context.Context, inventoryId uuid.UUID, offset, limit uint64) ([]domain.Item, error) {
	return s.repo.GetInventoryItems(ctx, inventoryId, nil, offset, limit)
}

func (s InventoryService) GetInventoryItemsBySection(ctx context.Context, inventoryId, sectionID uuid.UUID, offset, limit uint64) ([]domain.Item, error) {
	return s.repo.GetInventoryItems(ctx, inventoryId, &sectionID, offset, limit)
}

func (s InventoryService) CreateInventoryItem(ctx context.Context, inventory domain.Inventory, section domain.Section, item domain.Item) (domain.Item, error) {
	item.ID = uuid.New()
	item.InventoryID = inventory.ID
	item.SectionID = section.ID
	item.InitTimestamps()

	item.SKU = fmt.Sprintf("%x", sha256.Sum256([]byte(item.ID.String()+section.ID.String())))[:12]

	return s.repo.CreateInventoryItem(ctx, item)
}
