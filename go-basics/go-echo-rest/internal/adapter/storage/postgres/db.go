package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/adapter/config"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, conf *config.DB) (*DB, error) {
	dsn := "postgres://" + conf.User + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.Name
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{pool: pool}, nil
}

func (d *DB) Close() {
	d.pool.Close()
}
