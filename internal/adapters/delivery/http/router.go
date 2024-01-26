package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/handlers"
)

type Router struct {
	router *chi.Mux
}

func NewRouter(
	userHandler *handlers.UserHandler,
) *Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	userRouter := chi.NewRouter()
	userRouter.Post("/", userHandler.Register)

	r.Mount("/users", userRouter)

	return &Router{
		router: r,
	}
}

func (r *Router) ListenAndServe(host string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r.router)
}
