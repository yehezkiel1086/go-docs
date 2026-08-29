package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	// 1. Connect to Postgres using pgxpool
	pool, _ := pgxpool.New(ctx, "postgres://user:pass@localhost:5432/dbname")
	defer pool.Close()

	// 2. Initialize Adapter
	userRepo := pg.NewUserRepository(pool)

	// 3. Inject into Usecase/Service
	// userService := services.NewUserService(userRepo)

	// 4. Start Server...
}
