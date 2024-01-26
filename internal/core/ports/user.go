package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context, offset, limit uint64) ([]*domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id string) (*domain.User, error)
}
