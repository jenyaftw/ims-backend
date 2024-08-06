package handlers

import (
	"net/http"

	"github.com/jenyaftw/scaffold-go/internal/core/ports"
)

type ProtectedHandler struct {
	svc ports.UserService
}

func NewProtectedHandler(
	svc ports.UserService,
) ProtectedHandler {
	return ProtectedHandler{
		svc: svc,
	}
}

func (h ProtectedHandler) TestRoute(w http.ResponseWriter, r *http.Request) {

}
