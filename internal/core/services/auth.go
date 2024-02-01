package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg  *config.JwtConfig
	repo ports.UserRepository
}

func NewAuthService(
	cfg *config.JwtConfig,
	repo ports.UserRepository,
) AuthService {
	return AuthService{
		cfg:  cfg,
		repo: repo,
	}
}

func (s AuthService) LoginWithPassword(ctx context.Context, email, password string) (domain.Token, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.Token{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return domain.Token{}, domain.ErrInvalidPassword
	}

	now := time.Now().Unix()
	exp := time.Now().AddDate(0, 0, 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"nbf": now,
		"exp": exp,
	})

	tokenString, err := token.SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return domain.Token{}, domain.ErrInternal
	}

	return domain.Token{
		Text:      tokenString,
		ExpiresAt: exp,
	}, nil
}
