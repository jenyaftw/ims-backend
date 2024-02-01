package handlers

import "github.com/jenyaftw/scaffold-go/internal/core/domain"

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newUserResponse(user domain.User) userResponse {
	return userResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

type tokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   int64  `json:"expiresAt"`
}

func newTokenResponse(token domain.Token) tokenResponse {
	return tokenResponse{
		AccessToken: token.Text,
		ExpiresAt:   token.ExpiresAt,
	}
}
