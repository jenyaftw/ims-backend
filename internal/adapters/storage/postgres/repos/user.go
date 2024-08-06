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

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	rows, err := r.db.Query(
		ctx,
		`INSERT INTO users (id, name, email, password, updated_at, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		user.ID.String(), user.Name, user.Email, user.Password, user.UpdatedAt, user.CreatedAt,
	)
	if err != nil {
		return domain.User{}, err
	}

	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // Error code for when insert violates unique key constraint
				return user, domain.ErrDataConflict
			}
		}
	}
	return user, err
}

func (r UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM users WHERE id = $1 LIMIT 1`,
		id.String(),
	)
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, domain.ErrUserNotFound
		}
	}
	return user, err
}

func (r UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM users WHERE email = $1 LIMIT 1`,
		email,
	)
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, domain.ErrUserNotFound
		}
	}
	return user, err
}

func (r UserRepository) ListUsers(ctx context.Context, offset, limit uint64) ([]domain.User, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT * FROM users LIMIT $1 OFFSET $2`,
		offset, limit,
	)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = pgxscan.ScanAll(users, rows)
	return users, err
}

func (r UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Query(
		ctx,
		`DELETE FROM users WHERE id = $1`,
		id.String(),
	)
	if err != nil {
		return err
	}

	return nil
}
