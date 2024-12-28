package repos

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type InventoryRepository struct {
	db *pgx.Conn
}

func NewInventoryRepository(db *pgx.Conn) *InventoryRepository {
	return &InventoryRepository{
		db: db,
	}
}

func (r InventoryRepository) CreateInventory(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	rows, err := r.db.Query(
		ctx,
		`INSERT INTO inventories (id, name, description, updated_at, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		inventory.ID.String(), inventory.Name, inventory.Description, inventory.UpdatedAt, inventory.CreatedAt,
	)
	if err != nil {
		return domain.Inventory{}, err
	}

	err = pgxscan.ScanOne(&inventory, rows)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return inventory, domain.ErrDataConflict
			}
		}
	}

	return inventory, err
}

func (r InventoryRepository) UpdateInventory(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	rows, err := r.db.Query(
		ctx,
		`UPDATE inventories SET name = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING *`,
		inventory.Name, inventory.Description, inventory.UpdatedAt, inventory.ID.String(),
	)
	if err != nil {
		return domain.Inventory{}, err
	}

	err = pgxscan.ScanOne(&inventory, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return inventory, domain.ErrInventoryNotFound
		}
	}
	return inventory, err
}

func (r InventoryRepository) GetInventoryById(ctx context.Context, id uuid.UUID) (domain.Inventory, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM inventories WHERE id = $1 LIMIT 1`,
		id.String(),
	)
	if err != nil {
		return domain.Inventory{}, err
	}

	var inventory domain.Inventory
	err = pgxscan.ScanOne(&inventory, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return inventory, domain.ErrInventoryNotFound
		}
	}
	return inventory, err
}

func (r InventoryRepository) GetInventoryByName(ctx context.Context, name string) (domain.Inventory, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM inventories WHERE name = $1 LIMIT 1`,
		name,
	)
	if err != nil {
		return domain.Inventory{}, err
	}

	var inventory domain.Inventory
	err = pgxscan.ScanOne(&inventory, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return inventory, domain.ErrInventoryNotFound
		}
	}
	return inventory, err
}

func (r InventoryRepository) ListInventories(ctx context.Context, offset, limit uint64) ([]domain.Inventory, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM inventories LIMIT $1 OFFSET $2`,
		offset, limit,
	)
	if err != nil {
		return nil, err
	}

	var inventories []domain.Inventory
	err = pgxscan.ScanAll(inventories, rows)
	return inventories, err
}

func (r InventoryRepository) DeleteInventory(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Query(
		ctx,
		`DELETE FROM inventories WHERE id = $1`,
		id.String(),
	)
	if err != nil {
		return err
	}

	return nil
}
