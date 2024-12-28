package main

import (
	"fmt"

	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/handlers"
	"github.com/jenyaftw/scaffold-go/internal/adapters/storage/postgres"
	"github.com/jenyaftw/scaffold-go/internal/adapters/storage/postgres/repos"
	"github.com/jenyaftw/scaffold-go/internal/adapters/storage/redis"
	redRepos "github.com/jenyaftw/scaffold-go/internal/adapters/storage/redis/repos"
	"github.com/jenyaftw/scaffold-go/internal/core/services"
	"github.com/resend/resend-go/v2"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := postgres.InitDb(cfg.Db)
	if err != nil {
		panic(err)
	}

	rdb := redis.InitDb(cfg.Redis)
	cacheRepo := redRepos.NewCacheRepository(rdb)

	rs := resend.NewClient(cfg.Email.ApiKey)
	emailService := services.NewEmailService(cfg.Email, rs)

	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo, emailService, cacheRepo)
	userHandler := handlers.NewUserHandler(userService)

	authService := services.NewAuthService(cfg.Jwt, userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	inventoryRepo := repos.NewInventoryRepository(db)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)

	protectedHandler := handlers.NewProtectedHandler(userService)

	r := http.NewRouter(userHandler, authHandler, inventoryHandler, protectedHandler)

	fmt.Printf("Listening on http://%s:%d\n", cfg.Http.Host, cfg.Http.Port)
	if err := r.ListenAndServe(cfg.Http.Host, cfg.Http.Port); err != nil {
		panic(err)
	}
}
