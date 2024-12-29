package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type TransferRepository struct {
	db *pgx.Conn
}

func NewTransferRepository(db *pgx.Conn) *TransferRepository {
	return &TransferRepository{
		db: db,
	}
}

func (r TransferRepository) CreateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error) {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO transfer_requests (id, from_inventory_id, from_section_id, to_inventory_id, to_section_id, item_id, quantity, transfer_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		transfer.ID.String(),
		transfer.FromInventoryID.String(),
		transfer.FromSectionID.String(),
		transfer.ToInventoryID.String(),
		transfer.ToSectionID.String(),
		transfer.Item.ID.String(),
		transfer.Quantity,
		transfer.Status,
		transfer.CreatedAt,
		transfer.UpdatedAt,
	)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	return transfer, nil
}

func (r TransferRepository) GetTransferById(ctx context.Context, id uuid.UUID) (domain.TransferRequest, error) {
	var transfer domain.TransferRequest

	row := r.db.QueryRow(
		ctx,
		`SELECT * FROM transfer_requests WHERE id = $1`,
		id.String(),
	)
	err := row.Scan(
		&transfer.ID,
		&transfer.Item.ID,
		&transfer.Quantity,
		&transfer.FromInventoryID,
		&transfer.ToInventoryID,
		&transfer.FromSectionID,
		&transfer.ToSectionID,
		&transfer.Status,
		&transfer.UpdatedAt,
		&transfer.CreatedAt,
	)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	return transfer, nil
}

func (r TransferRepository) GetTransfers(ctx context.Context, offset, limit uint64) ([]domain.TransferRequest, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM transfer_requests LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}

	var transfers []domain.TransferRequest
	for rows.Next() {
		var transfer domain.TransferRequest
		err = rows.Scan(
			&transfer.ID,
			&transfer.Item.ID,
			&transfer.Quantity,
			&transfer.FromInventoryID,
			&transfer.ToInventoryID,
			&transfer.FromSectionID,
			&transfer.ToSectionID,
			&transfer.Status,
			&transfer.UpdatedAt,
			&transfer.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (r TransferRepository) UpdateTransfer(ctx context.Context, transfer domain.TransferRequest) (domain.TransferRequest, error) {
	_, err := r.db.Exec(
		ctx,
		`UPDATE transfer_requests SET from_inventory_id = $2, to_inventory_id = $3, item_id = $4, quantity = $5, transfer_status = $6, updated_at = $7 WHERE id = $1`,
		transfer.ID.String(),
		transfer.FromInventoryID.String(),
		transfer.ToInventoryID.String(),
		transfer.Item.ID.String(),
		transfer.Quantity,
		transfer.Status,
		transfer.UpdatedAt,
	)
	if err != nil {
		return domain.TransferRequest{}, err
	}

	return transfer, nil
}

func (r TransferRepository) DeleteTransfer(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM transfer_requests WHERE id = $1`,
		id.String(),
	)
	if err != nil {
		return err
	}

	return nil
}
