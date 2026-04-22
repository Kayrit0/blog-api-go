package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	db *pgxpool.Pool
}

func Setup(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
