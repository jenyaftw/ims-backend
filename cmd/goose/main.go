package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	disableSSL := ""
	if !cfg.Db.Secure {
		disableSSL = "?sslmode=disable"
	}

	url := fmt.Sprintf("%s://%s:%s@%s:%d/%s%s",
		cfg.Db.Driver,
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
		disableSSL,
	)
	db, err := goose.OpenDBWithDriver(cfg.Db.Driver, url)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.RunContext(context.Background(), args[0], db, cfg.Db.Migrations, arguments...); err != nil {
		log.Fatalf("goose %v: %v", args[0], err)
	}
}
