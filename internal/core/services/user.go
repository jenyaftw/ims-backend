package services

import (
	"context"

	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(
	repo ports.UserRepository,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	// TODO: Implement register method
	return nil, nil
}

func (us *UserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	// TODO: Implement get user method
	return nil, nil
}
