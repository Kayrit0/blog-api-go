package database

import (
	"context"

	"github.com/Kayrit0/blog-api-go/internal/libs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool(cfg *libs.Config) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), cfg.DB_URL)
	if err != nil {
		panic(err)
	}
	if err := dbpool.Ping(context.Background()); err != nil {
		panic(err)
	}
	return dbpool
}

func HealthCheck(pool *pgxpool.Pool) error {
	return pool.Ping(context.Background())
}
