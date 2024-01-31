package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jenyaftw/scaffold-go/internal/core/domain"
	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type UserHandler struct {
	svc ports.UserService
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(
	svc ports.UserService,
) UserHandler {
	return UserHandler{
		svc: svc,
	}
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.svc.Register(ctx, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(newUserResponse(user))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(res)
}
