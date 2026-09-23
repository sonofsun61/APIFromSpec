package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MustConnectToDatabase(ctx context.Context, connString string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(fmt.Sprintf("could not create pgxpool: %v", err))
	}
	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Sprintf("could not ping pgxpool: %v", err))
	}
	return pool
}
