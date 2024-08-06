package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo  ports.UserRepository
	email ports.EmailService
	cache ports.CacheRepository
}

func NewUserService(
	repo ports.UserRepository,
	email ports.EmailService,
	cache ports.CacheRepository,
) UserService {
	return UserService{
		repo:  repo,
		email: email,
		cache: cache,
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

	newUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return newUser, err
	}

	code := uuid.New().String()
	cacheIndex := "verification_code_" + newUser.ID.String()
	if err := s.cache.Set(ctx, cacheIndex, code, 0); err != nil {
		return newUser, err
	}

	if err := s.email.SendEmailConfirmation(ctx, newUser.ID.String(), newUser.Name, newUser.Email, code); err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s UserService) Verify(ctx context.Context, id uuid.UUID, code string) error {
	user, err := s.GetUser(ctx, id)
	if err != nil {
		return err
	}

	cacheIndex := "verification_code_" + user.ID.String()
	cachedCode, err := s.cache.Get(ctx, cacheIndex)
	if err != nil {
		return err
	}

	if cachedCode != code {
		return domain.ErrInvalidVerificationCode
	}

	if err := s.cache.Delete(ctx, cacheIndex); err != nil {
		return err
	}

	user.IsVerified = true
	_, err = s.repo.UpdateUser(ctx, user)
	return err
}

func (s UserService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.GetUserById(ctx, id)
}
