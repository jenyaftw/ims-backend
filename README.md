# Inventory Management System

This repository holds the code for the backend of my inventory management system. It's built on Golang with a bunch of other external libraries.

## Requirements

- Have [Golang](https://go.dev/dl/) compiler installed on your machine
- Docker Engine
- Docker Compose

If you have PostgreSQL and Redis already installed, no need to have Docker installed.

## How To

1) Clone the repository: `git clone git@github.com:jenyaftw/ims-backend.git`
2) Copy the following config file into `config.toml` in the root directory:
```
[http]
host = "0.0.0.0" # Host the server should listen on
port = 8000 # Port you want the server to run on

[jwt]
secret = "<SECRET>" # Auth tokens will be encrypted with this key, keep it hush

# These are all default from the docker-compose.yml
[db]
driver = "postgres"
host = "127.0.0.1"
port = 5432
username = "postgres"
password = "password"
name = "postgres"
secure = false
migrations = "./internal/adapters/storage/postgres/migrations"

# This isn't required, email verification was temporarily disabled for this project.
[email]
api_key = "<API_KEY>"
from = "Scaffold <scaffold@resend.dev>"

# These are all default from the docker-compose.yml
[redis]
host = "127.0.0.1"
port = 6379
password = "password"
db = 0
```
3) Launch database and cache with the following command: `docker-compose up -d`
4) You're almost there, just have to make sure your database is synced: `go run cmd/goose/main.go up`
5) Now just launch the backend itself: `go run cmd/http/main.go`

Woohoo! Congrats, you got it up and running!

Now onto the [frontend](https://github.com/jenyaftw/ims-app)
