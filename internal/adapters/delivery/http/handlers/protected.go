package handlers

import "net/http"

type ProtectedHandler struct{}

func NewProtectedHandler() ProtectedHandler {
	return ProtectedHandler{}
}

func (h ProtectedHandler) TestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are authenticated"))
}
