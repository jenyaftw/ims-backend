package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type TransferRepository interface {
	CreateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error)
	DeleteTransfer(ctx context.Context, id uuid.UUID) error
	GetTransferById(ctx context.Context, id uuid.UUID) (domain.TransferRequest, error)
	GetTransfers(ctx context.Context, offset, limit uint64) ([]domain.TransferRequest, error)
	UpdateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error)
}

type TransferService interface {
	CreateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error)
	DeleteTransfer(ctx context.Context, id uuid.UUID) error
	GetTransferById(ctx context.Context, id uuid.UUID) (domain.TransferRequest, error)
	GetTransfers(ctx context.Context, offset, limit uint64) ([]domain.TransferRequest, error)
	UpdateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error)
	ProcessTransfer(ctx context.Context, id uuid.UUID) error
}
