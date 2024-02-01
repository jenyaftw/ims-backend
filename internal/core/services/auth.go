package services

import (
	"context"

	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type AuthService struct {
	repo ports.UserRepository
}

func NewAuthService(
	repo ports.UserRepository,
) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s AuthService) LoginWithPassword(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
