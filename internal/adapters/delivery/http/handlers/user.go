package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
		HandleError(w, err)
		return
	}

	newUser := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.svc.Register(ctx, newUser)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newUserResponse(user))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userIdString := r.Context().Value("user").(string)
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	user, err := h.svc.GetUser(ctx, userId)
	if err != nil {
		HandleError(w, err)
		return
	}

	res, err := json.Marshal(newUserResponse(user))
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Write(res)
}

func (h UserHandler) Verify(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userIdString := chi.URLParam(r, "id")
	code := chi.URLParam(r, "code")

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		HandleError(w, err)
		return
	}

	if err := h.svc.Verify(ctx, userId, code); err != nil {
		HandleError(w, err)
		return
	}

	w.Write([]byte("User verified"))
}
