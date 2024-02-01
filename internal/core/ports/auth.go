package ports

import "context"

type AuthService interface {
	LoginWithPassword(ctx context.Context, email, password string) (string, error)
}
