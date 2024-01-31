package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(
	repo ports.UserRepository,
) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) Register(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return user, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hash)
	user.ID = uuid.New()
	user.InitTimestamps()

	return s.repo.CreateUser(ctx, user)
}

func (s UserService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.GetUserById(ctx, id)
}
