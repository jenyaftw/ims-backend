package domain

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"github.com/jenyaftw/scaffold-go/internal/core/util"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string

	util.Timestamps
}

func (u User) Validate() error {
	if len(u.Name) < 2 || len(u.Name) > 32 {
		return errors.New("name length must be between 2 and 32 characters")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email address")
	}

	return nil
}
