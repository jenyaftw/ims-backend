package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type TransferService struct {
	repo    ports.TransferRepository
	invRepo ports.InventoryRepository
}

func NewTransferService(
	repo ports.TransferRepository,
	invRepo ports.InventoryRepository,
) TransferService {
	return TransferService{
		repo:    repo,
		invRepo: invRepo,
	}
}

func (s TransferService) CreateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error) {
	transfer.ID = uuid.New()
	transfer.InitTimestamps()

	item, err := s.invRepo.GetInventoryItemByID(ctx, transfer.Item.ID)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	transfer.Item = item
	transfer.FromInventoryID = item.InventoryID
	transfer.FromSectionID = item.SectionID
	transfer.Status = "pending"

	section, err := s.invRepo.GetInventorySection(ctx, transfer.ToSectionID)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	transfer.ToInventoryID = section.InventoryID

	transfer, err = s.repo.CreateTransfer(ctx, transfer)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	return transfer, nil
}

func (s TransferService) DeleteTransfer(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTransfer(ctx, id)
}

func (s TransferService) GetTransferById(ctx context.Context, id uuid.UUID) (domain.TransferRequest, error) {
	transfer, err := s.repo.GetTransferById(ctx, id)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	item, err := s.invRepo.GetInventoryItemByID(ctx, transfer.Item.ID)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	transfer.Item = item
	return transfer, nil
}

func (s TransferService) GetTransfers(ctx context.Context, offset, limit uint64) ([]domain.TransferRequest, error) {
	transfers, err := s.repo.GetTransfers(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	for i, transfer := range transfers {
		item, err := s.invRepo.GetInventoryItemByID(ctx, transfer.Item.ID)
		if err != nil {
			return nil, err
		}

		transfer.Item = item
		transfers[i] = transfer
	}

	return transfers, nil
}

func (s TransferService) UpdateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error) {
	return s.repo.UpdateTransfer(ctx, transfer)
}

func (s TransferService) ProcessTransfer(ctx context.Context, id uuid.UUID) error {
	transfer, err := s.GetTransferById(ctx, id)
	if err != nil {
		return err
	}

	if transfer.Status != "pending" {
		return domain.ErrTransferAlreadyProcessed
	}

	fromItem, err := s.invRepo.GetInventoryItemByID(ctx, transfer.Item.ID)
	if err != nil {
		return err
	}

	if fromItem.Quantity < int(transfer.Quantity) {
		return domain.ErrNotEnoughItems
	}

	fromItem.Quantity -= int(transfer.Quantity)

	toItem := domain.Item{
		ID:          uuid.New(),
		InventoryID: transfer.ToInventoryID,
		SectionID:   transfer.ToSectionID,
		Name:        fromItem.Name,
		Description: fromItem.Description,
		Quantity:    int(transfer.Quantity),
		SKU:         fromItem.SKU,
	}

	_, err = s.invRepo.UpdateInventoryItem(ctx, fromItem)
	if err != nil {
		return err
	}

	_, err = s.invRepo.CreateInventoryItem(ctx, toItem)
	if err != nil {
		return err
	}

	transfer.Status = "processed"
	_, err = s.UpdateTransfer(ctx, transfer)
	if err != nil {
		return err
	}

	return nil
}
