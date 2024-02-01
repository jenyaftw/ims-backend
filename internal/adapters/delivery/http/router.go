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
	userHandler handlers.UserHandler,
	authHandler handlers.AuthHandler,
) Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	userRouter := chi.NewRouter()
	userRouter.Post("/", userHandler.Register)

	authRouter := chi.NewRouter()
	authRouter.Post("/login", authHandler.Login)

	r.Mount("/users", userRouter)
	r.Mount("/auth", authRouter)

	return Router{
		router: r,
	}
}

func (r Router) ListenAndServe(host string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r.router)
}
