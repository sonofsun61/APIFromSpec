package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB() (*pgxpool.Pool) {
	dns := os.Getenv("DATABASE_URL")
	if dns == "" {
		panic("No database environment database provided")
	}
	pool, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		panic("Can not create pgxpool")
	}
	defer pool.Close()

	return pool
}