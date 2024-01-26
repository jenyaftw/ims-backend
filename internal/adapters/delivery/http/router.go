package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	router *chi.Mux

	Host string
	Port int
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) WithPort(port int) *Router {
	r.Port = port
	return r
}

func (r *Router) WithHost(host string) *Router {
	r.Host = host
	return r
}

func (r *Router) Build() *Router {
	r.router = chi.NewRouter()

	return r.BuildRoutes()
}

func (r *Router) BuildRoutes() *Router {
	r.router.Use(middleware.Logger)
	r.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	return r
}

func (r *Router) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", r.Host, r.Port), r.router)
}
