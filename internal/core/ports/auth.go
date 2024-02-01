package ports

import (
	"context"

	"github.com/jenyaftw/scaffold-go/internal/core/domain"
)

type AuthService interface {
	LoginWithPassword(ctx context.Context, email, password string) (domain.Token, error)
}
