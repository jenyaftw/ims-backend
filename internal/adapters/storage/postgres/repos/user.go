package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (r UserRepository) CreateUser(ctx context.Context, user *domain.User) (domain.User, error) {
	// TODO: Implement create user method
	return domain.User{}, nil
}

func (r UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	// TODO: Implement get user by id method
	return domain.User{}, nil
}

func (r UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	// TODO: Implement get user by email method
	return domain.User{}, nil
}

func (r UserRepository) ListUsers(ctx context.Context, offset, limit uint64) ([]domain.User, error) {
	// TODO: Implement list users method
	return []domain.User{}, nil
}

func (r UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// TODO: Implement delete user method
	return nil
}
