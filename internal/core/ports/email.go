package ports

import (
	"context"
)

type EmailService interface {
	SendEmailConfirmation(ctx context.Context, id, name, email, code string) error
}
