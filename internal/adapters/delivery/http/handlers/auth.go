package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type AuthHandler struct {
	svc ports.AuthService
}

func NewAuthHandler(
	svc ports.AuthService,
) AuthHandler {
	return AuthHandler{
		svc: svc,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.svc.LoginWithPassword(ctx, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(token)
}
