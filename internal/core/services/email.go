package services

import (
	"context"
	"strings"

	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	"github.com/jenyaftw/scaffold-go/internal/core/services/templates"
	"github.com/resend/resend-go/v2"
)

type EmailService struct {
	cfg    *config.EmailConfig
	client *resend.Client
}

func NewEmailService(cfg *config.EmailConfig, client *resend.Client) EmailService {
	return EmailService{
		cfg:    cfg,
		client: client,
	}
}

func (s EmailService) SendEmailConfirmation(ctx context.Context, id, name, email, code string) error {
	b := new(strings.Builder)
	component := templates.VerifyUser(name, id, code)
	component.Render(context.Background(), b)
	body := b.String()

	params := &resend.SendEmailRequest{
		From:    s.cfg.From,
		To:      []string{email},
		Subject: "Scaffold - Confirm your email adddress",
		Html:    body,
	}

	_, err := s.client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}
