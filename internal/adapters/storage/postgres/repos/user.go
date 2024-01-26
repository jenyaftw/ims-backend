package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// TODO: Implement create user method
	return nil, nil
}

func (ur *UserRepository) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	// TODO: Implement get user by id method
	return nil, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	// TODO: Implement get user by email method
	return nil, nil
}

func (ur *UserRepository) ListUsers(ctx context.Context, offset, limit uint64) ([]*domain.User, error) {
	// TODO: Implement list users method
	return nil, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// TODO: Implement delete user method
	return nil
}
