package redis

import (
	"fmt"

	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	red "github.com/redis/go-redis/v9"
)

func InitDb(cfg *config.RedisConfig) *red.Client {
	return red.NewClient(&red.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
	})
}
