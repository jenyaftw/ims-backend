package services

import (
	"context"

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

/*

type InventoryService interface {
	ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error)
}

*/

func (s InventoryService) ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error) {
	return s.repo.ListInventories(ctx, offset, limit)
}
