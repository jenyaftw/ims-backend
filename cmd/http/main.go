package main

import (
	"fmt"

	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/handlers"
	"github.com/jenyaftw/scaffold-go/internal/adapters/storage/postgres/repos"
	"github.com/jenyaftw/scaffold-go/internal/core/services"
)

func main() {
	userRepo := repos.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := http.NewRouter(userHandler)

	host := "localhost"
	port := 3333
	fmt.Printf("Listening on http://%s:%d\n", host, port)
	if err := r.ListenAndServe(host, port); err != nil {
		panic(err)
	}
}
