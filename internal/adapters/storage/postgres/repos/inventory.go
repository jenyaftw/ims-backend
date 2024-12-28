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
	err = pgxscan.ScanAll(&inventories, rows)
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

func (r InventoryRepository) GetInventorySections(ctx context.Context, inventoryID uuid.UUID) ([]domain.Section, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM inventory_sections WHERE inventory_id = $1`,
		inventoryID.String(),
	)
	if err != nil {
		return nil, err
	}

	var sections []domain.Section
	err = pgxscan.ScanAll(&sections, rows)
	return sections, err
}

func (r InventoryRepository) CreateInventorySection(ctx context.Context, section domain.Section) (domain.Section, error) {
	rows, err := r.db.Query(
		ctx,
		`INSERT INTO inventory_sections (id, inventory_id, name, description, updated_at, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		section.ID.String(), section.InventoryID.String(), section.Name, section.Description, section.UpdatedAt, section.CreatedAt,
	)
	if err != nil {
		return domain.Section{}, err
	}

	err = pgxscan.ScanOne(&section, rows)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return section, domain.ErrDataConflict
			}
		}
	}

	return section, err
}

func (r InventoryRepository) DeleteInventorySection(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Query(
		ctx,
		`DELETE FROM inventory_sections WHERE id = $1`,
		id.String(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r InventoryRepository) UpdateInventorySection(ctx context.Context, section domain.Section) (domain.Section, error) {
	rows, err := r.db.Query(
		ctx,
		`UPDATE inventory_sections SET name = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING *`,
		section.Name, section.Description, section.UpdatedAt, section.ID.String(),
	)
	if err != nil {
		return domain.Section{}, err
	}

	err = pgxscan.ScanOne(&section, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return section, domain.ErrSectionNotFound
		}
	}
	return section, err
}

func (r InventoryRepository) GetInventoryItems(ctx context.Context, inventoryID uuid.UUID, sectionID *uuid.UUID, offset, limit uint64) ([]domain.Item, error) {
	var rows pgx.Rows
	var err error
	if sectionID != nil {
		rows, err = r.db.Query(
			ctx,
			`SELECT * FROM inventory_items WHERE inventory_id = $1 AND section_id = $2 LIMIT $3 OFFSET $4`,
			inventoryID.String(), sectionID.String(), limit, offset,
		)
	} else {
		rows, err = r.db.Query(
			ctx,
			`SELECT * FROM inventory_items WHERE inventory_id = $1 LIMIT
			$2 OFFSET $3`,
			inventoryID.String(), limit, offset,
		)
	}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrItemNotFound
		}
		return nil, err
	}

	var items []domain.Item
	err = pgxscan.ScanAll(&items, rows)
	return items, err
}

func (r InventoryRepository) CreateInventoryItem(ctx context.Context, item domain.Item) (domain.Item, error) {
	rows, err := r.db.Query(
		ctx,
		`INSERT INTO inventory_items (id, inventory_id, section_id, name, description, quantity, sku, updated_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`,
		item.ID.String(), item.InventoryID.String(), item.SectionID.String(), item.Name, item.Description, item.Quantity, item.SKU, item.UpdatedAt, item.CreatedAt,
	)
	if err != nil {
		return domain.Item{}, err
	}

	err = pgxscan.ScanOne(&item, rows)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return item, domain.ErrDataConflict
			}
		}
	}

	return item, err
}
