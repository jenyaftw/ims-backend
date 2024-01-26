package handlers

import (
	"net/http"

	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type UserHandler struct {
	svc ports.UserService
}

func NewUserHandler(
	svc ports.UserService,
) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
